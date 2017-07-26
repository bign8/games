package sock

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/websocket/websocketjs"
)

const (
	retryCap  = 60 * 1000 // one-minute max retry
	retryBase = 50        // 50ms base retrys
)

var rand = func(num int) int { return int(js.Global.Get("Math").Call("random").Float() * float64(num)) }

// New constructs a new socket for a given ws url
func New(addr string) (c *Conn, err error) {
	c = &Conn{Object: js.Global.Get("Object").New()}
	return c, c.reconnect(addr)
}

// Conn is a websocket connection
type Conn struct {
	*js.Object
	*websocketjs.WebSocket `js:"sock"` // Core embedded object
	attempt                int         // current retry attempt
}

// retryDelay gives the current retry delay in milliseconds
// TODO: https://www.awsarchitectureblog.com/2015/03/backoff.html
// TODO: benchmark
func retryDelay(attempt int) (delay, next int) {
	delay = 1 << uint(attempt-1) // 2^retryAttp
	delay *= retryBase
	if delay > retryCap {
		delay = retryCap
	}
	return int(rand(delay)), attempt + 1
}

func (c *Conn) reconnect(addr string) (err error) {
	// TODO: if WS is already set, unset listeners and destroy object
	print("reconnecting!")
	if !(c.WebSocket == nil || c.WebSocket.Object == js.Undefined) {
		// TODO: teardown handlers
		addr = c.WebSocket.URL
	}

	// Construct the new websocket and attach the correct handlers
	c.WebSocket, err = websocketjs.New(addr)
	if err == nil {
		c.WebSocket.AddEventListener("error", false, c.onErr)
		c.WebSocket.AddEventListener("close", false, c.onClose)
		c.WebSocket.AddEventListener("open", false, c.onOpen)
		c.WebSocket.AddEventListener("message", false, c.onMsg)
	}
	return err
}

// onClose is what the js socket will call on close
func (c *Conn) onClose(e *js.Object) {
	var delay int
	delay, c.attempt = retryDelay(c.attempt)
	print("setting reconnect timeouts", delay, e)
	js.Global.Call("setTimeout", c.reconnect, delay)
}

// onOpen is what the js socket will call on open
func (c *Conn) onOpen(e *js.Object) {
	c.attempt = 0
	print("open", e)
}

// onErr is what the js socket will call on errors
func (c *Conn) onErr(e *js.Object) {
	print("err", e)
}

// onMsg is what the js socket will call on messages
func (c *Conn) onMsg(e *js.Object) {
	print("TIME", js.Global.Get("Date").Call("now"), "MSG", e.Get("data"))
}

// // See: https://github.com/gopherjs/gopherjs/wiki/JavaScript-Tips-and-Gotchas
// func recov(err *error) {
// 	if e := recover(); e == nil {
// 		return
// 	} else if er, ok := e.(*js.Error); ok {
// 		*err = er
// 	} else {
// 		panic(e)
// 	}
// }
