// +build js

package main

import "github.com/gopherjs/gopherjs/js"

func main() {
	js.Global.Get("document").Call("write", "Hello world!")
	var sock = js.Global.Get("WebSocket").New("ws://localhost:8000/sock")

	sock.Set("onmessage", func(e *js.Object) {
		print(e.Get("data"))
		print(e)
	})

	js.Global.Set("games", map[string]interface{}{
		"sock": sock,
	})
}
