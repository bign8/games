package assert_test

import (
	"fmt"
	"testing"

	"github.com/bign8/games/util/assert"
)

type chk struct {
	t *testing.T
	m string
}

func (c chk) Errorf(format string, args ...interface{}) {
	if msg := fmt.Sprintf(format, args...); c.m != msg {
		c.t.Errorf("string did not match %q != %q", msg, c.m)
	}
}

func Test(t *testing.T) {
	assert.Equal(chk{t, "string: asdf != jkl;"}, "asdf", "jkl;", "string")
	assert.Equal(chk{t, "bool: true != false"}, true, false, "bool")
	assert.Equal(chk{t, "byte: 97 != 98"}, 'a', 'b', "byte")
}
