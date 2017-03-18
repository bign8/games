// +build !js

//go:generate gopherjs -m

package main

import (
	"fmt"
	"io"
	"net/http"

	"golang.org/x/net/websocket"
)

// sock is the main websocket handler.
func sock(ws *websocket.Conn) {
	ws.Write([]byte("server: connected!"))
	io.Copy(ws, ws)
	// TODO: http://www.meetspaceapp.com/2016/03/29/tutorial-getting-started-websockets-go.html
}

// http://buildnewgames.com/real-time-multiplayer/
func main() {
	// TODO: use http 2.0 to push more assets to the consumer
	// TODO: use service workers on the client side as necessary
	fmt.Println("Running server...")
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.Handle("/ws", websocket.Handler(sock))
	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
