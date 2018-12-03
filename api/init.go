package api

import (
	"fmt"
	"connect/conf"
	"net/http"
	"github.com/emicklei/go-restful"
	"log"
)

func InitApiServer(){
	//默认路由器的写法
	//model.Server{}.Register()

	Container := restful.NewContainer()
	Container.Router(restful.CurlyRouter{})
	//server := model.Server{}
	//server.Register(Container)

	//add middleware
	for _,ws := range Container.RegisteredWebServices(){
		ws.Filter(JwtAuthMiddleware)
	}

	addr := fmt.Sprintf("%s:%d",conf.Setting.Api.Address,conf.Setting.Api.Port)
	log.Printf("Run api server on port: %v",addr)
	log.Fatalf("Run api server %v error: %v",addr,http.ListenAndServe(addr, Container))
}
