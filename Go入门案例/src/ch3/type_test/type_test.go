package type_test

import "testing"

type myInt int64

// go语言不支持隐式类型转换
func TestImpleicit(t *testing.T) {

	var b int32 = 1
	var c int64
	//b = a
	//b = c
	b = int32(c) //显式类型转换可以
	//var d myInt
	//c = d  别名也不可以转换
	t.Log(b, c)
}

// Go支持指针的使用
func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a //取内存地址
	//aPtr = aPtr+1,Go不支持指针的运算
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr) //取类型
}

// string是值类型，默认是空字符串，而不是nil值
func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))
}
