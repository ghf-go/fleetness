package core

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ghf-go/fleetness/core/conf"
	"gopkg.in/yaml.v3"
)

type GEngine struct {
	webRouter *WebRouter //webRouter路径
	confData  *conf.Conf
	job       *job
}

func NewGengine(confdata []byte) *GEngine {
	conf := &conf.Conf{}
	if e := yaml.Unmarshal(confdata, conf); e != nil {
		panic(e.Error())
	}
	fmt.Println(conf.App.Port)
	return &GEngine{
		webRouter: newRootRouter(),
		confData:  conf,
		job:       newJob(newWebGContent(conf, nil, nil, []Handle{})),
	}
}
func (ge *GEngine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS,GET")
		w.Header().Set("Access-Control-Allow-Headers", "Appid,Appver,x-requested-with,Token,content-type,Cookie,Authorization,Sid,Set-Cookie,Access-Control-Allow-Origin")
		w.WriteHeader(204)
		return
	}
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	AppDebug("请求接口 %s %s", r.URL.Path, r.URL.RawQuery)
	isOk, hands := ge.webRouter.FindHandle(r.Method, r.URL.Path)
	c := newWebGContent(ge.confData, w, r, hands)
	if isOk {
		c.Next()
	} else {
		c.Next()
	}

}

// 是否debug运行
func (ge *GEngine) SetDebug(isdebug bool) {
	isAppDebug = isdebug
}

// 程序开始运行
func (ge *GEngine) Run() {
	hserver := &http.Server{
		Addr:    fmt.Sprintf(":%d", ge.confData.App.Port),
		Handler: ge,
	}
	go ge.job.start()
	go func() {
		if e := hserver.ListenAndServe(); e != nil {
			panic(e.Error())
		}
	}()
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
	<-sigc
	ct, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	hserver.Shutdown(ct)
}

// 注册web handle
func (ge *GEngine) RegisterWebHandle() {}

// 添加计划任务
func (ge *GEngine) AddCronJob(name string, crontab string, handle Handle) {
	ge.job.addCronJob(name, crontab, handle)
}

// 添加时间间隔的任务
func (ge *GEngine) AddAfterJob(name string, second int, handle Handle) {
	ge.job.addAfterJob(name, second, handle)
}

// 添加一直运行的JOB
func (ge *GEngine) AddAlwaysJob(name string, handle Handle) {
	ge.job.addAlwaysJob(name, handle)
}

// 注册Webvue
func (ge *GEngine) RegisterVue() {}

// 注册模版
func (ge *GEngine) RegisterTemplate() {}

func (ge *GEngine) RouterPost(path string, hand Handle, args ...Handle) {
	ge.webRouter.Post(path, hand, args...)
}
func (ge *GEngine) RouterGet(path string, hand Handle, args ...Handle) {
	ge.webRouter.Get(path, hand, args...)
}
func (ge *GEngine) RouterAny(path string, hand Handle, args ...Handle) {
	ge.webRouter.Any(path, hand, args...)
}
func (ge *GEngine) RouterDelete(path string, hand Handle, args ...Handle) {
	ge.webRouter.Delete(path, hand, args...)
}
func (ge *GEngine) RouterPut(path string, hand Handle, args ...Handle) {
	ge.webRouter.Post(path, hand, args...)
}
func (ge *GEngine) RouterGroup(path string, err404 Handle, args ...Handle) *WebRouter {
	return ge.webRouter.Group(path, err404, args...)
}
