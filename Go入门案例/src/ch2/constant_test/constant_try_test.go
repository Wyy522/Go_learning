package constant_test

import "testing"

const (
	Monday = iota + 1
	Tuesday
	wednesday
)

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday)
	a := 1
	t.Log(a&Readable == Readable, a&Writable == Writable)
}
