package websocket

import (
	"net/http"
	"golang.org/x/net/websocket"
	"log"
	"fmt"
	"connect/conf"
)

func InitWebSocketServer(){
	var listenOption string
	listenOption = fmt.Sprintf("%s:%d",conf.Setting.Websocket.Address, conf.Setting.Websocket.Port)
	//接受websocket的路由地址

	http.Handle("/websocket", websocket.Handler(Echo))

	//html页面

	http.HandleFunc("/web", Web)

	log.Printf("Start WebSocket service on port: %v",listenOption)

	if err := http.ListenAndServe(listenOption, nil); err != nil {

		log.Fatal("Start WebSocket,ListenAndServe:", err)

	}

}
