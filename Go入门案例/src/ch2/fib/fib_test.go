package fib

import "testing"

// 斐波那契数列
func TestFibList(t *testing.T) {
	//var a int = 1
	//var b int = 1

	//var (
	//	a int =1
	//	b =1
	//)
	//

	a := 1 //自动类型推断
	b := 1

	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log(b)
		tmp := a
		a = b
		b = tmp + a
	}
}

// 交换两个数
func TestExchange(t *testing.T) {
	a := 1
	b := 2
	a, b = b, a
	t.Log(a, b)
}
