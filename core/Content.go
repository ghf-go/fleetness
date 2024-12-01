package core

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/smtp"
	"strconv"
	"strings"
	"time"

	"github.com/ghf-go/fleetness/core/conf"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
	uuid "github.com/satori/go.uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type GContent struct {
	confData      *conf.Conf
	ReqID         string //请求id
	r             *http.Request
	w             http.ResponseWriter
	handles       []Handle
	reqIP         string
	isAbort       bool
	currentNext   int
	responseBytes []byte
	userId        uint64
	ctx           context.Context
	clientlang    string
}

// 新建GContent
func newWebGContent(confData *conf.Conf, w http.ResponseWriter, r *http.Request, handles []Handle) *GContent {

	return &GContent{
		confData:    confData,
		r:           r,
		w:           w,
		handles:     handles,
		ReqID:       uuid.NewV4().String(),
		isAbort:     false,
		currentNext: 0,
		ctx:         context.Background(),
	}
}

// 获取客户端IP
func (c *GContent) GetIP() string {
	if c.r == nil {
		return "127.0.0.1"
	}
	if c.reqIP == "" {
		ret := c.r.Header.Get("ipv4")
		if ret != "" {
			c.reqIP = ret
			return c.reqIP
		}
		ret = c.r.Header.Get("X-Forwarded-For")
		if ret != "" {
			rs := strings.Split(ret, ",")
			if rs[0] != "" {
				c.reqIP = rs[0]
				return c.reqIP
			}
		}
		ret = c.r.Header.Get("XForwardedFor")
		if ret != "" {
			rs := strings.Split(ret, ",")
			if rs[0] != "" {
				c.reqIP = rs[0]
				return c.reqIP
			}
		}
		ret = c.r.Header.Get("X-Real-Ip")
		if ret != "" {
			rs := strings.Split(ret, ",")
			if rs[0] != "" {
				c.reqIP = rs[0]
				return c.reqIP
			}
		}
		ret = c.r.Header.Get("X-Real-IP")
		if ret != "" {
			rs := strings.Split(ret, ",")
			if rs[0] != "" {
				c.reqIP = rs[0]
				return c.reqIP
			}
		}
		ret = c.r.RemoteAddr
		if ret != "" {
			ips := strings.Split(ret, ":")
			c.reqIP = ips[0]
			return c.reqIP
		}

		c.reqIP = "unknow"
	}
	return c.reqIP
}

func (c *GContent) Abort() {
	c.isAbort = true
}
func (c *GContent) Next() {
	if c.r == nil {
		return
	}
	if c.currentNext < len(c.handles) {
		ci := c.currentNext
		c.currentNext++
		if !c.isAbort {
			c.handles[ci](c)
		}
	}

}

// 是否开启DEBUG 日志输出模式
func (c *GContent) IsDebug() bool {
	return c.confData.App.Debug
}

// 获取数据库
func (c *GContent) GetDB(dbname ...string) *gorm.DB {
	conName := "default"
	if len(dbname) > 0 {
		conName = dbname[0]
	}
	if r, ok := dbCon[conName]; ok {
		return r
	}
	if dbconf, ok := c.confData.Dbs[conName]; ok {
		db, e := gorm.Open(mysql.Open(dbconf.Write), &gorm.Config{})
		if e != nil {
			panic(e.Error())
		}
		rs := []gorm.Dialector{}
		for _, rc := range dbconf.Reads {
			rs = append(rs, mysql.Open(rc))
		}
		db.Use(dbresolver.Register(dbresolver.Config{
			Sources:  []gorm.Dialector{mysql.Open(dbconf.Write)},
			Replicas: rs,
		}).SetMaxIdleConns(dbconf.ConMaxIdleTime).SetMaxOpenConns(dbconf.MaxOpenCons).SetConnMaxIdleTime(time.Minute * time.Duration(dbconf.ConMaxIdleTime)).SetConnMaxLifetime(time.Minute * time.Duration(dbconf.ConMaxLifeTime)))
		if c.confData.App.Debug {
			db = db.Debug()
		}
		dbCon[conName] = db
		return db
	} else {
		panic("数据配置错误 " + conName)
	}

}

// 获取缓存配置
func (c *GContent) GetCache(conname ...string) *redis.Client {
	conName := "default"
	if len(conname) > 0 {
		conName = conname[0]
	}
	if r, ok := cacheCon[conName]; ok {
		return r
	}
	if rconf, ok := c.confData.Cache[conName]; ok {
		r := redis.NewClient(&redis.Options{
			Addr:            rconf.Host,
			Username:        rconf.UserName,
			Password:        rconf.Passwd,
			MinIdleConns:    rconf.MinIdleConns,
			MaxIdleConns:    rconf.MaxIdleConns,
			MaxActiveConns:  rconf.MaxActiveConns,
			ConnMaxIdleTime: time.Minute * time.Duration(rconf.ConnMaxIdleTime),
			ConnMaxLifetime: time.Minute * time.Duration(rconf.ConnMaxLifetime),
		})
		cacheCon[conName] = r
		return r
	}
	panic("缓存配置不存在" + conName)

}

// 国际化语言
func (c *GContent) Lang(key string) string {
	if c.clientlang == "" {
		kks := strings.Split(c.r.Header.Get("accept-language"), ",")
		for _, item := range kks {
			iks := strings.Split(item, ";")
			for _, v := range iks {
				if strings.Index(v, "-") > 0 {
					c.clientlang = strings.ToLower(v)
					return c.confData.Lang.Lang(key, c.clientlang)
				}
			}
		}
	}
	return c.confData.Lang.Lang(key, c.clientlang)
}

// 绑定数据
func (c *GContent) BindJson(obj any) error {
	if c.r == nil {
		return errors.New("计划任务不能解析参数")
	}
	body := c.r.Body
	defer body.Close()
	data, e := io.ReadAll(body)
	if e != nil {
		return e
	}
	return json.Unmarshal(data, obj)
}

// 接口正常返回
func (c *GContent) SuccessJson(data any) {
	c.json(0, "", data)
}

// 接口保存信息
func (c *GContent) FailJson(code int, msg string) {
	c.json(code, msg, nil)
}

// 输出JSON信息
func (c *GContent) json(code int, msg string, data any) {
	if c.r == nil {
		return
	}
	c.w.Header().Set("content-type", "application/json;charset=utf8")
	ret := map[string]any{
		"code": code,
		"msg":  msg,
		"data": data,
	}
	dd, e := json.Marshal(ret)
	if e != nil {
		panic(e.Error())
	}
	c.responseBytes = dd

}

// 刷新缓存
func (c *GContent) flush() {
	if c.r == nil {
		return
	}
	c.w.(http.Flusher).Flush()
}

// 开启Event事件
func (c *GContent) Sse(call func(s *Sse)) {
	if c.r == nil {
		return
	}
	c.w.Header().Set("Content-Type", "text/event-stream")
	// 这行代码设置 HTTP 响应的 Cache-Control 为 no-cache，告诉浏览器不要缓存此响应。
	c.w.Header().Set("Cache-Control", "no-cache")
	// 这行代码设置 HTTP 响应的 Connection 为 keep-alive，保持长连接，以便服务器可以持续发送事件到客户端。
	c.w.Header().Set("Connection", "keep-alive")
	// 这行代码设置 HTTP 响应的自定义头部 X-Accel-Buffering 为 no，用于禁用某些代理或 Web 服务器（如 Nginx）的缓冲。这有助于确保服务器发送事件在传输过程中不会受到缓冲影响
	c.w.Header().Set("X-Accel-Buffering", "no")
	c.w.Header().Set("Access-Control-Allow-Origin", "*")
	c.flush()
	call(&Sse{c: c, isClose: false, key: c.ReqID})
}

// 开始websocket
func (c *GContent) WebSocket(call func(con *websocket.Conn)) {
	if c.r == nil {
		return
	}
	conn, err := upgrader.Upgrade(c.w, c.r, nil)
	if err != nil {
		fmt.Printf("链接失败 %s\n", err.Error())
		log.Println(err)
		return
	}
	call(conn)
	defer fmt.Printf("链接管理了\n")
	defer conn.Close()

	// if c.r.Header.Get("Upgrade") != "websocket" {
	// 	return errors.New("协议错误")
	// }
	// if c.r.Header.Get("Sec-WebSocket-Version") != "13" {
	// 	return errors.New("协议错误")
	// }
	// k := c.r.Header.Get("Sec-WebSocket-Key")
	// if k == "" {
	// 	return errors.New("协议错误")
	// }
	// c.w.WriteHeader(http.StatusSwitchingProtocols)
	// c.w.Header().Set("Upgrade", "websocket")
	// c.w.Header().Set("Sec-WebSocket-Version", "13")
	// c.w.Header().Set("Connection", "Upgrade")
	// dd := sha1.Sum([]byte(k + "258EAFA5-E914-47DA-95CA-C5AB0DC85B11"))
	// c.w.Header().Set("Sec-WebSocket-Accept", base64.StdEncoding.EncodeToString(dd[:]))
	// c.flush()
	// return nil
}

// 显示模版
func (c *GContent) Display() {}

// 显示模版
func (c *GContent) DisplayLayout() {}

// 获取配置信息
func (c *GContent) GetConf() *conf.Conf {
	return c.confData
}
func (c *GContent) Flush() {
	if c.r == nil {
		AppDebug(" %s", string(c.responseBytes))
		return
	}
	c.w.Write(c.responseBytes)
}
func (c *GContent) GetRequest() *http.Request {
	if c.r == nil {
		return nil
	}
	return c.r
}
func (c *GContent) GetResponseWriter() http.ResponseWriter {
	if c.r == nil {
		return nil
	}
	return c.w
}
func (c *GContent) SetUserID(uid string) {
	if c.r == nil {
		return
	}
	if id, e := strconv.ParseUint(uid, 10, 64); e == nil {
		c.userId = id
	}
}

func (c *GContent) GetUserID() uint64 {
	if c.r == nil {
		return 0
	}
	return c.userId
}
func (c *GContent) IsLogin() bool {
	if c.r == nil {
		return false
	}
	return c.userId > 0
}
func (c *GContent) GetContext() context.Context {
	return c.ctx
}

// 处理事务
func (c *GContent) Tx(db *gorm.DB, call func(tx *gorm.DB) (error, any)) (error, any) {
	tx := db.Begin()
	e, ret := call(tx)
	if e != nil {
		tx.Rollback()
		return e, ret
	}
	tx.Commit()
	return e, ret
}

// 通过本机发送邮件
func (c *GContent) SendLocalMail(conname, to, subject string, isHtml bool, msg []byte) error {
	i := strings.Index(to, "@")
	host := to[i+1:]
	if sc, ok := c.confData.Stmp[conname]; ok {
		if dd, e := net.LookupMX(host); e == nil {
			content_type := ""
			if isHtml {
				content_type = "Content-Type: text/html; charset=UTF-8"
			} else {
				content_type = "Content-Type: text/plain" + "; charset=UTF-8"
			}
			msg = []byte("To: " + to + "\r\nFrom: " + sc.UserName + "\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + string(msg))

			return smtp.SendMail(dd[0].Host+":25", nil, sc.UserName, []string{to}, msg)
		}
		return errors.New("获取信息失败")
	}
	fmt.Println(c.confData.Stmp)
	return errors.New("配置不存在" + conname)

}

// 发送邮件
func (c *GContent) SendMail(conname, to, subject string, isHtml bool, msg []byte) error {
	if sc, ok := c.confData.Stmp[conname]; ok {
		var auth smtp.Auth
		switch strings.ToUpper(sc.AuthType) {
		case "CRAMMD5":
			auth = smtp.CRAMMD5Auth(sc.UserName, sc.Passwd)
		case "HOTMAIL":
			auth = conf.NewHotmailStmpAuth(sc.UserName, sc.Passwd)
		default:
			auth = smtp.PlainAuth("", sc.UserName, sc.Passwd, sc.Host)
		}
		content_type := ""
		if isHtml {
			content_type = "Content-Type: text/html; charset=UTF-8"
		} else {
			content_type = "Content-Type: text/plain" + "; charset=UTF-8"
		}
		msg = []byte("To: " + to + "\r\nFrom: " + sc.UserName + "\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + string(msg))

		return smtp.SendMail(sc.Host, auth, sc.UserName, []string{to}, msg)
	}
	return errors.New("配置不存在")
}
