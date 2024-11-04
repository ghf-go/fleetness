package core

import (
	"fmt"
	"net/http"

	"github.com/ghf-go/fleetness/core/conf"
	"gopkg.in/yaml.v3"
)

type GEngine struct {
	webRouter *WebRouter //webRouter路径
	confData  *conf.Conf
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
	}
}
func (ge *GEngine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS,GET")
		w.Header().Set("content-type", "application/json;charset=utf8")
		w.Header().Set("Access-Control-Allow-Headers", "Appid,Appver,x-requested-with,Token,content-type,Cookie,Authorization,Sid,Set-Cookie,Access-Control-Allow-Origin")
		w.WriteHeader(204)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write([]byte("sadfsdfsd"))
}

// 是否debug运行
func (ge *GEngine) SetDebug(isdebug bool) {}

// 程序开始运行
func (ge *GEngine) Run() {
	hserver := &http.Server{
		Addr:    fmt.Sprintf(":%d", ge.confData.App.Port),
		Handler: ge,
	}
	func() {
		if e := hserver.ListenAndServe(); e != nil {
			panic(e.Error())
		}
	}()
	fmt.Println("----")
}

// 注册消息队列
func (ge *GEngine) RegisterMq() {

}

// 注册web handle
func (ge *GEngine) RegisterWebHandle() {}

// 注册job
func (ge *GEngine) RegisterJob() {}

// 注册Webvue
func (ge *GEngine) RegisterVue() {}

// 注册模版
func (ge *GEngine) RegisterTemplate() {}
