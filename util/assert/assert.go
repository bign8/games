package assert

// TestingT matches the *testing.T interface for what is consumed here.
type TestingT interface {
	Error(args ...interface{})
}

// String verifies a string is the same.
func String(t TestingT, a, b, msg string) {
	if a != b {
		t.Error("string: " + msg + ": '" + a + "' != '" + b + "'")
	}
}

func bos(bit bool) string {
	if bit {
		return "true"
	}
	return "false"
}

// Bool verifies a boolean is the same.
func Bool(t TestingT, a, b bool, msg string) {
	if a != b {
		t.Error("bool: " + msg + ": '" + bos(a) + "' != '" + bos(b) + "'")
	}
}

func bys(char byte) string { return string(char) }

// Byte verifies a byte is the same.
func Byte(t TestingT, a, b byte, msg string) {
	if a != b {
		t.Error("byte: " + msg + ": '" + bys(a) + "' != '" + bys(b) + "'")
	}
}
