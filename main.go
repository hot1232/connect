package main

import (
	"flag"
	"connect/conf"
	"github.com/golang/glog"
	"connect/audit"
	"connect/sshserver"
	"connect/websocket"
	"connect/api"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var cf = flag.String("config","config.yaml","config path")

func main(){
	var err error
	flag.Parse()
	config := conf.NewSettings()
	conf.Setting,err = config.Load(*cf)
	if err != nil {
		glog.Fatalf("load config: %v err: %v",*cf,err)
	}

	audit.InitDb()

	//启动ssh服务
	go sshserver.InitSsh()

	//启动api服务
	go api.InitApiServer()

	//启动websocket
	go websocket.InitWebSocketServer()

	select {

	}
}