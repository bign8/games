// +build js

package main

import (
	"github.com/bign8/games/io/client/sock"
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	js.Global.Get("document").Call("write", "Hello world!")
	sock, err := sock.New("ws://localhost:8000/ws")
	if err != nil {
		panic(err)
	}

	js.Global.Set("games", map[string]interface{}{
		"sock": sock,
		"test": func() {
			sock.Send(js.Global.Get("Date").Call("now"))
		},
		"rand": func() float64 {
			return js.Global.Get("Math").Call("random").Float()
		},
	})
}
