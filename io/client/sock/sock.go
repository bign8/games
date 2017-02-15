package sock

import "github.com/gopherjs/gopherjs/js"

// New constructs a new socket for a given ws url
func New(addr string) *Conn {
	c := &Conn{Object: js.Global.Get("Object").New()}
	c.addr = addr
	c.retryCap = 60 * 1000 // on-minute max retry
	c.retryBase = 50       // 250ms base retrys
	c.reconnect()
	return c
}

// Conn is a websocket connection
type Conn struct {
	*js.Object
	WS   *js.Object `js:"ws"`   // javascript websocket
	addr string     `js:"addr"` // address of websocket

	// retry parameters
	retryCap  int // maximum allowed retry (milliseconds)
	retryBase int // base multiple for retry (milliseconds)
	retryAttp int // current retry attempt
}

// retryDelay gives the current retry delay in milliseconds
// TODO: https://www.awsarchitectureblog.com/2015/03/backoff.html
func (c *Conn) retryDelay() int {
	delay := 1 << uint(c.retryAttp-1) // 2^retryAttp
	delay *= c.retryBase
	if delay > c.retryCap {
		delay = c.retryCap
	}
	c.retryAttp++
	return int(js.Global.Get("Math").Call("random").Float() * float64(delay))
}

func (c *Conn) reconnect() {
	print("reconnecting!")
	c.WS = js.Global.Get("WebSocket").New(c.addr)
	c.WS.Set("onerror", c.onErr)
	c.WS.Set("onclose", c.onClose)
	c.WS.Set("onopen", c.onOpen)
	c.WS.Set("onmessage", c.onMsg)
}

// onClose is what the js socket will call on close
func (c *Conn) onClose(e *js.Object) {
	delay := c.retryDelay()
	print("setting reconnect timeouts", delay, e)
	js.Global.Call("setTimeout", c.reconnect, delay)
}

// onOpen is what the js socket will call on open
func (c *Conn) onOpen(e *js.Object) {
	c.retryAttp = 0
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

// See: https://github.com/gopherjs/gopherjs/wiki/JavaScript-Tips-and-Gotchas
func recov(err *error) {
	if e := recover(); e == nil {
		return
	} else if er, ok := e.(*js.Error); ok {
		*err = er
	} else {
		panic(e)
	}
}

// Send sends a message on the WebSocket.
// See: http://dev.w3.org/html5/websockets/#dom-websocket-send
func (c *Conn) Send(data interface{}) (err error) {
	defer recov(&err)
	c.WS.Call("send", data)
	return err
}

// Close closes the underlying WebSocket.
// See: http://dev.w3.org/html5/websockets/#dom-websocket-close
func (c *Conn) Close() (err error) {
	defer recov(&err)
	c.WS.Call("close")
	return err
}
