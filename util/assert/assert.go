package assert

// TestingT matches the *testing.T interface for what is consumed here.
type TestingT interface {
	Errorf(string, ...interface{})
}

// Equal checks if two types are the same.
// Note: this only supports base types, so we do not allow reflect.DeepEqual.
func Equal(t TestingT, a, b interface{}, msg string) {
	if a != b {
		t.Errorf("%s: %v != %v", msg, a, b)
	}
}
