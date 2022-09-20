package map_test

import "testing"

// map与工厂模式
func TestMapWithFunValue(t *testing.T) {
	//key是int ,value是方法 func(op int) int  参数是int 返回值是int
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m[1](2), m[2](2), m[3](2))
}

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	if mySet[1] {
		t.Log("Key is exiscting")
	}
	delete(mySet, 1)
}
