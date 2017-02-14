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
	ws.Write([]byte("onload: yep"))
	io.Copy(ws, ws)
}

func main() {
	fmt.Println("Running server")
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.Handle("/sock", websocket.Handler(sock))
	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
