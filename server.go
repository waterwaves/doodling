package main

import (
	"github.com/googollee/go-socket.io"
	"log"
	"net/http"
)

func main() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	server.On("connection", func(so socketio.Socket) {
		log.Println("on connection")
		so.Join("chat")
		so.On("draw", func(msg string) {
			//log.Println("emit ", msg, so.Emit("draw", msg))
			so.BroadcastTo("chat", "draw", msg)
		})
		so.On("disconnection", func() {
			log.Println("Disconnectioned")
		})
	})
	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error ", err)
	})

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./view")))
	log.Println("Serving at localhost:5000...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
