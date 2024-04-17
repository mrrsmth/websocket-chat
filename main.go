package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"

)

//	var upgrader = websocket.Upgrader{
//		ReadBufferSize:  1024,
//		WriteBufferSize: 1024,
//	}
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	http.HandleFunc("/ws", wsHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}

		log.Println(string(msg))
		conn.WriteMessage(websocket.TextMessage, msg)
	}
}
