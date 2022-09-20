package operator_test

import "testing"

const (
	Readable = 1 << iota
	Writable
	Executable
)

// 比较数组需要相同维度和相同元素，每个元素相同才相等
func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 3, 5}
	c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	//t.Log(a == b) flase
	//t.Log(a == c) 长度不同的数组比较会编译错误
	//t.Log(a == d) true
	t.Log(a, b, c, d)
}

// &^按位清零运算符
func TestBitClear(t *testing.T) {
	a := 7
	a = a &^ Readable

	/*按位清零运算符  下面为1就置上面为0
	     1100
	  &^ 0110
	  =  1000

	*/
	t.Log(a&Readable == Readable, a&Writable == Writable)
}
