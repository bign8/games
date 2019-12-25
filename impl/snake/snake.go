package snake

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/bign8/games"
	"github.com/bign8/games/player"
)

// Game ...
// https://en.wikipedia.org/wiki/Snake_(video_game_genre)
// https://www.coolmathgames.com/0-snake
// https://www.coolmathgames.com/0-snake/play
var Game = games.Game{
	Name:   "Snake",
	Slug:   "snake",
	Board:  board,
	Counts: []uint8{1},
	Start:  New,
	AI:     player.Random, // TODO: custom planner here
}

const (
	head  = `<svg viewBox="-1 -1 40 28" xmlns="http://www.w3.org/2000/svg">`
	tail  = `</svg>`
	board = head + `<rect x="0" y="0" width="38" height="26" fill="black" stroke="#16324c" stroke-width="1" paint-order="stroke" />` + tail
	body  = `<rect x="%d" y="%d" width="1" height="1" fill="#35de00" />`
	food  = `<rect x="%d" y="%d" width="1" height="1" fill="#ff0000" />`
)

// New constructs a snake game.
func New(actors uint8) games.State {
	if actors != 1 {
		panic("invalid number of players")
	}
	game := &snake{
		body:    []uint{0}, // head --- tail
		width:   38,        // depends on screen size
		height:  26,        // scall to fill maybe?
		heading: 100,       // invalid initial heading
	}
	max := game.width * game.height
	game.body[0] = (max + game.width) / 2
	game.food = uint(rand.Intn(int(max)))
	if game.food == game.body[0] {
		game.food++
	}
	return game
}

type snake struct {
	body    []uint
	food    uint
	heading move
	width   uint
	height  uint
}

func (s *snake) String() string { return "SNAKE!" }
func (s *snake) Player() int    { return 0 }
func (s *snake) Utility() []int { return []int{len(s.body)} }
func (s *snake) Terminal() bool { return s.body[0] < 0 }

func (s *snake) Apply(externalMove games.Action) games.State {
	future := &snake{
		width:  s.width,
		height: s.height,
	}

	// compute next heading
	m := externalMove.(move)
	if s.heading > 10 {
		future.heading = m
	} else {
		future.heading = derive(s.heading, m)
	}

	// find next body position
	x, y := s.pt2xy(s.body[0])
	switch future.heading {
	case up:
		y--
	case down:
		y++
	case left:
		x--
	case right:
		x++
	}
	tip := s.xy2pt(x, y)

	// generate new body
	future.body = append([]uint{tip}, s.body...) // really not performant, but need to copy base array
	if tip == s.food {
		future.spawn()
	} else {
		future.body = future.body[:len(future.body)-1] // didn't hit food, pop tail off
	}
	return future
}

func derive(current, key move) move {
	if key == same {
		return current
	}
	if key == left {
		switch current {
		case up:
			return left
		case left:
			return down
		case down:
			return right
		case right:
			return up
		default:
			panic("wtf")
		}
	}
	// right
	switch current {
	case up:
		return right
	case right:
		return down
	case down:
		return left
	case left:
		return up
	default:
		panic("wtf")
	}
}

func (s *snake) spawn() {
	max := int(s.width * s.height)
	max -= len(s.body) // subtract out body nibs
	if max <= 0 {
		panic("dope player?")
	}
	spot := uint(rand.Intn(max))
place: // n*n worse case to ensure we don't collide with stuff
	for _, part := range s.body {
		if part == spot {
			spot++
			goto place
		}
	}
	s.food = spot
}

func (s *snake) Actions() []games.Action {
	if s.heading > 10 {
		return []games.Action{left, right, up, down}
	}
	return []games.Action{same, left, right}
}

func (s *snake) SVG(bool) string {
	msg := s.pt2svg(s.food, food)
	for _, pt := range s.body {
		msg += s.pt2svg(pt, body)
	}
	return head + msg + tail
}

func (s *snake) pt2svg(pt uint, tpl string) string {
	x, y := s.pt2xy(pt)
	return fmt.Sprintf(tpl, x, y)
}

func (s *snake) pt2xy(pt uint) (x, y uint) { return pt % s.width, pt / s.width }
func (s *snake) xy2pt(x, y uint) uint      { return x + y*s.width }

type move uint

func (m move) String() string { return strings.Title(m.Slug()) }
func (m move) Type() string   { return "" }
func (m move) Slug() string {
	switch m {
	case same:
		return "same"
	case left:
		return "left"
	case right:
		return "right"
	case up:
		return "up"
	case down:
		return "down"
	default:
		panic("invalid state")
	}
}

const (
	same = move(iota)
	left
	right
	up   // only on initial state
	down // only on initial state
)
