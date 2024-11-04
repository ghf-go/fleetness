package core

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/ghf-go/fleetness/core/conf"
	"github.com/redis/go-redis/v9"
	uuid "github.com/satori/go.uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

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
	}
}

// 获取客户端IP
func (c *GContent) GetIP() string {
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
	if c.currentNext < len(c.handles) {
		ci := c.currentNext
		c.currentNext++
		c.handles[ci](c)

	}

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
			Sources:  []gorm.Dialector{db},
			Replicas: rs,
		}).SetMaxIdleConns(dbconf.ConMaxIdleTime).SetMaxOpenConns(dbconf.MaxOpenCons).SetConnMaxIdleTime(time.Minute * time.Duration(dbconf.ConMaxIdleTime)).SetConnMaxLifetime(time.Minute * time.Duration(dbconf.ConMaxLifeTime)))
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

// 绑定数据
func (c *GContent) BindJson(obj any) error {
	body := c.r.Body
	defer body.Close()
	data, e := io.ReadAll(body)
	if e != nil {
		return e
	}
	return json.Unmarshal(data, obj)
}

// 发送队列
func (c *GContent) SendMq() {}

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

// 显示模版
func (c *GContent) Display() {}

// 显示模版
func (c *GContent) DisplayLayout() {}

// 发送邮件
func (c *GContent) SendMail() {}

// 获取配置信息
func (c *GContent) GetConf() *conf.Conf {
	return c.confData
}
func (c *GContent) Flush() {
	c.w.Write(c.responseBytes)
}
