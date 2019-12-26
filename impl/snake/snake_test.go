package snake

import (
	"testing"
	"time"

	"github.com/bign8/games/util/assert"
)

func TestNewPanic(t *testing.T) {
	die := make(chan interface{}, 1)
	go func() {
		defer func() {
			die <- recover()
		}()
		New(0)
	}()
	select {
	case v := <-die:
		assert.Equal(t, v.(string), "invalid number of players", "invalid panic")
	case <-time.After(50 * time.Millisecond):
		t.Fatal("timeout dying")
	}
}

func TestStory(t *testing.T) {
	s := New(1).(*snake)
	assert.Equal(t, s.String(), "SNAKE!", "dummy string")
	assert.Equal(t, s.Player(), 0, "always p0")
	assert.Equal(t, s.Utility()[0], 1, "utility is snake length")
	assert.Equal(t, s.Terminal(), false, "start is not dead")
}

func TestDerive(t *testing.T) {
	assert.Equal(t, derive(up, same), up, "same 1")
	assert.Equal(t, derive(left, same), left, "same 2")
	assert.Equal(t, derive(right, same), right, "same 3")
	assert.Equal(t, derive(down, same), down, "same 4")

	assert.Equal(t, derive(up, left), left, "left 1")
	assert.Equal(t, derive(left, left), down, "left 2")
	assert.Equal(t, derive(down, left), right, "left 3")
	assert.Equal(t, derive(right, left), up, "left 4")

	assert.Equal(t, derive(up, right), right, "right 1")
	assert.Equal(t, derive(right, right), down, "right 2")
	assert.Equal(t, derive(down, right), left, "right 3")
	assert.Equal(t, derive(left, right), up, "right 4")
}

func TestSpawn(t *testing.T) {
	s := &snake{
		height: 2,
		width:  2,
		body:   []uint{0, 1, 2},
	}
	s.spawn()
	assert.Equal(t, s.food, uint(3), "only spot left")
}

func TestMove(t *testing.T) {
	assert.Equal(t, same.Type(), "", "no type")
	assert.Equal(t, same.String(), "Same", "same")
	assert.Equal(t, left.String(), "Left", "left")
	assert.Equal(t, right.String(), "Right", "right")
	assert.Equal(t, up.String(), "Up", "up")
	assert.Equal(t, down.String(), "Down", "down")
}

func TestPointMagic(t *testing.T) {
	s := &snake{width: 10}
	x, y := s.pt2xy(23)
	assert.Equal(t, x, uint(3), "x")
	assert.Equal(t, y, uint(2), "y")
	assert.Equal(t, s.xy2pt(x, y), uint(23), "pt")
	assert.Equal(t, s.pt2svg(23, "%d-%d"), "3-2", "svg")
}

func TestActions(t *testing.T) {
	s := &snake{heading: 123}
	assert.Equal(t, len(s.Actions()), 4, "first move")
	s.heading = down
	assert.Equal(t, len(s.Actions()), 3, "subsequent")
}

func TestApply(t *testing.T) {
	s := &snake{
		width:   10,
		height:  5,
		food:    23,
		heading: right,
		body:    []uint{21, 20},
	}
	s = s.Apply(same).(*snake)
	assert.Equal(t, s.width, uint(10), "width 1")
	assert.Equal(t, s.height, uint(5), "width 1")
	assert.Equal(t, s.food, uint(23), "food 1")
	assert.Equal(t, s.heading, right, "heading 1")
	assert.Equal(t, len(s.body), 2, "body len 1")
	assert.Equal(t, s.body[0], uint(22), "body 0 1")
	assert.Equal(t, s.body[1], uint(21), "body 1 1")
	s = s.Apply(same).(*snake)
	s = s.Apply(same).(*snake)
	assert.Equal(t, len(s.body), 3, "body len grow")
}
