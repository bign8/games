package sock

import "github.com/gopherjs/gopherjs/js"

// New constructs a new socket for a given ws url
func New(addr string) *js.Object {
	c := &Conn{} //Object: js.Global.Get("Object").New()}
	c.addr = addr
	c.retryCap = 60 * 1000 // on-minute max retry
	c.retryBase = 50       // 250ms base retrys
	c.reconnect()
	return js.MakeWrapper(c)
}

// Conn is a websocket connection
type Conn struct {
	// *js.Object
	ws   *js.Object // javascript websocket
	addr string     // address of websocket

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
	return js.Global.Get("Math").Call("random").Int() * delay
	// if delay == 0 {
	// 	return delay
	// }
	// return rand.Intn(delay) // TODO: transpile a better random generator here
}

func (c *Conn) reconnect() {
	print("reconnecting!")
	c.ws = js.Global.Get("WebSocket").New(c.addr)
	c.ws.Set("onerror", c.onErr)
	c.ws.Set("onclose", c.onClose)
	c.ws.Set("onopen", c.onOpen)
	c.ws.Set("onmessage", c.onMsg)
}

// onClose is what the js socket will call on close
func (c *Conn) onClose(e *js.Object) {
	delay := c.retryDelay()
	print("setting reconnect timeout", delay, e)
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

// Send broadcasts data to a websocket
func (c *Conn) Send(data []byte) {
	c.ws.Call("send", data)
}
