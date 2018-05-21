package assert

import "testing"

type checker struct {
	msg [][]interface{}
}

func (c *checker) Error(args ...interface{}) {
	c.msg = append(c.msg, args)
}

func (c *checker) check(t *testing.T, msgs ...string) {
	if len(msgs) != len(c.msg) {
		t.Errorf("check missmatched error lengths: %d != %d", len(msgs), len(c.msg))
	}
	for i, msg := range msgs {
		if len(c.msg[i]) != 1 {
			t.Errorf("check did not receive a singular error: %q", c.msg[i])
		} else if core, ok := c.msg[i][0].(string); !ok {
			t.Errorf("first arg was not a string")
		} else if core != msg {
			t.Errorf("string did not match %q != %q", core, msg)
		}
	}
}

func TestString(t *testing.T) {
	c := &checker{}
	String(c, "asdf", "jkl;", "whuhh???")
	c.check(t, "string: whuhh???: 'asdf' != 'jkl;'")
}

func TestBool(t *testing.T) {
	c := &checker{}
	Bool(c, true, false, "whuhh???")
	c.check(t, "bool: whuhh???: 'true' != 'false'")
}

func TestByte(t *testing.T) {
	c := &checker{}
	Byte(c, 'a', 'b', "whuhh???")
	c.check(t, "byte: whuhh???: 'a' != 'b'")
}
