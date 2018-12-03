package api

import (
	"github.com/emicklei/go-restful"
)

var Container *restful.Container

func init(){
	//默认路由器的写法
	//model.Server{}.Register()

	Container = restful.NewContainer()
	Container.Router(restful.CurlyRouter{})
	//server := model.Server{}
	//server.Register(Container)


	//add middleware
	for _,ws := range Container.RegisteredWebServices(){
		ws.Filter(JwtAuthMiddleware)
	}
}