package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/gorilla/websocket"
)
// 需要定义一个 Upgrader
// 它需要定义 ReadBufferSize 和 WriteBufferSize
var upgrader = websocket.Upgrader{
	ReadBufferSize:1024,
	WriteBufferSize:1024,
	//可以用来检查链接来源
	//允许从React 服务向这里请求
	//暂时，不需要检查并运行任何连接
	CheckOrigin: func(r *http.Request) bool {return true},
}

//监听往ws发送的新消息
func reader(conn *websocket.Conn){

	for{
		//读取消息
		messageType,p,err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		//打印消息
		fmt.Println(string(p))
		if err := conn.WriteMessage(messageType,p); err != nil{
			log.Println(err)
            return
		}
	}

}

func serveWs(w http.ResponseWriter, r*http.Request){
	 fmt.Println(r.Host)
	 //将连接更新为 WebSocket 连接
	 ws , err := upgrader.Upgrade(w,r,nil)
	 if err != nil {
        log.Println(err)
	 }
	 // 一直监听 WebSocket 连接上传来的新消息
	 reader(ws)
}


func setupRoutes(){
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w,"Simple Server")
	})

	// 将 `/ws` 端点交给 `serveWs` 函数处理
    http.HandleFunc("/ws", serveWs)
}

func main() {
  fmt.Println("Chat App v0.01")
  setupRoutes()
  http.ListenAndServe(":8080",nil)
}