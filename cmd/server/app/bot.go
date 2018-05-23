package app

import (
	"io"
	"time"

	"github.com/bign8/games/util/markov"
)

// https://talks.golang.org/2012/chat.slide
// https://talks.golang.org/2012/chat/both/chat.go

var chain = markov.NewChain(2) // 2-word prefixes

func cp(w io.Writer, r io.Reader, errc chan<- error) {
	_, err := io.Copy(io.MultiWriter(w, chain), r) // copy chats to markov chain
	errc <- err
}

// Bot returns an io.ReadWriteCloser that responds to
// each incoming write with a generated sentence.
func Bot() io.ReadWriteCloser {
	r, out := io.Pipe() // for outgoing data
	return bot{r, out}
}

type bot struct {
	io.ReadCloser
	out io.Writer
}

func (b bot) Write(buf []byte) (int, error) {
	if len(buf) > 0 && buf[0] == 'u' {
		go b.speak()
	}
	return len(buf), nil
}

func (b bot) speak() {
	time.Sleep(time.Second)
	msg := chain.Generate(10) // at most 10 words
	b.out.Write([]byte("u" + msg + "\n"))
}
