// +build !js

//go:generate gopherjs -m

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/websocket"
)

// sock is the main websocket handler.
func sock(ws *websocket.Conn) {
	fmt.Println("Somebody connected!", ws.RemoteAddr())
	ws.Write([]byte("server: connected!"))
	io.Copy(ws, ws)
	// TODO: http://www.meetspaceapp.com/2016/03/29/tutorial-getting-started-websockets-go.html
}

var htmlTest = `<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
		<title>Testing...</title>
		<link type="text/css" rel="stylesheet" href="/test/css.css">
		<script type="text/javascript" src="/test/js.js"></script>
	</head>
	<body>

	</body>
</html>`

func test(w http.ResponseWriter, r *http.Request) {
	if strings.HasSuffix(r.URL.Path, ".html") {
		log.Println(r.Proto)
		if p, ok := w.(http.Pusher); ok {
			log.Println("Promoted!")
			p.Push("/test/css.css", nil) // err
		}
		w.Write([]byte(htmlTest)) // err
	} else {
		w.Write([]byte(`/* static */`)) // err
	}
}

func run(f func() error, c chan<- error) { c <- f() }

// http://buildnewgames.com/real-time-multiplayer/
// https://github.com/denji/golang-tls
func main() {
	// TODO: use http 2.0 to push more assets to the consumer
	// TODO: use service workers on the client side as necessary
	fmt.Println("Running server...")
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/test/", test)
	http.Handle("/ws", websocket.Handler(sock))

	errc := make(chan error, 1)
	go run(func() error { return http.ListenAndServe(":8000", nil) }, errc)
	go run(func() error { return http.ListenAndServeTLS(":8443", "server.crt", "server.key", nil) }, errc)
	err := <-errc
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
