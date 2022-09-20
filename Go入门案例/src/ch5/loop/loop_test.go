package loop

import "testing"

// 普通for
func TestForLoop(t *testing.T) {
	for i := 0; i < 5; i++ {
		t.Log(i)
	}
}

// while
func TestWhileLoop(t *testing.T) {
	n := 0
	for n < 5 {
		t.Log(n)
		n++
	}
	//死循环
	for {
		t.Log(n)
	}
}

// foreach
func TestForEachLoop(t *testing.T) {
	m1 := [1]int{}
	for idx, v := range m1 {
		t.Log(idx, v)
	}
}
