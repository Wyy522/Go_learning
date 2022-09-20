package main

import (
	"fmt"
	"reflect"
	"unicode/utf8"
	"unsafe"
)

func main() {
	//指针
	//uintptr----unsafe.Pointer JAVA大端
	var a int16 = 0x6162
	var ptr *int16 = &a
	var b = uint64(uintptr(unsafe.Pointer(ptr)))
	b++
	var p *uint8 = (*uint8)(unsafe.Pointer(uintptr(b)))
	fmt.Print(*p)

	//type
	type size int // 相当于新建一个类型
	var xx size = 1
	var yy int = 1
	type zsize = int //相当于赋值
	var zz zsize
	fmt.Println(reflect.TypeOf(xx))        // main.size
	fmt.Println(reflect.TypeOf(yy))        // int
	fmt.Println(reflect.TypeOf(zz))        // int
	fmt.Println(reflect.TypeOf(xx).Kind()) // int
	fmt.Println(reflect.TypeOf(yy).Kind()) // int

	//byte == ascii 一字节
	//rune == utf-8 四字节int
	//go所有字母都是byte  中文都是rune

	//string 是由StringHeader构成 有两个属性len和data 16字节
	var str string = "abc"
	var str2 string = "中"
	//len()统计字节数
	fmt.Println(len(str))  //3
	fmt.Println(len(str2)) //3
	//tf8.RuneCountInString() //统计字符数
	fmt.Println(utf8.RuneCountInString(str))  //3
	fmt.Println(utf8.RuneCountInString(str2)) //1

	fmt.Println("-----------------------------------")
	var s string = "abc"
	r := (*reflect.StringHeader)(unsafe.Pointer(&s))
	fmt.Println(r.Len)
	fmt.Println(r.Data)
	i := *(*uint8)(unsafe.Pointer(r.Data))
	fmt.Printf("%c", i)

	//函数
	plus := func(a, b int64) int64 {
		return a + b
	}
	ops := func(f func(int64, int64) int64, a, b int64) int64 {
		return f(a, b)
	}
	fmt.Println(ops(plus, 4, 2))

	//闭包
	incr := func() func() int64 {
		var a int64 = 0
		return func() int64 {
			a++
			return a
		}
	}
	x := incr()
	x1 := incr()
	fmt.Println(x())
	fmt.Println(x1())

	//defer函数 lazy调用
	//defer 相当于压栈 压引用  fmt.Println(0)  fmt.Println(1)  1,0
	for i := 0; i < 2; i++ {
		defer fmt.Println(i)
	}
	//fmt.Println(i)  fmt.Println(i)  2,2
	for i := 0; i < 2; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}

	//struct
	type Person struct {
		Id   int    `aaa=bbb` //8
		Name string // len +data =8+8 16
	} //sizeof= 24 =8+8+8
	var arr [3]int64 = [3]int64{}
	arr[0] = 99
	arr[2] = 3
	strr := "abc"
	ptrr := (*reflect.StringHeader)(unsafe.Pointer(&strr))
	arr[1] = int64(ptrr.Data)
	ppp := (*Person)(unsafe.Pointer(&arr))
	fmt.Println(ppp.Id, ppp.Name)

	//继承
	type stu struct {
		Person
	}

	t := Test{}
	t.SetTestOk(1, "xx")
	fmt.Println(t)

}

type Test struct {
	Id   int    `aaa=bbb` //8
	Name string // len +data =8+8 16
}

// 读写
func (t *Test) SetTestOk(i int, n string) {
	t.Id = i
	t.Name = n
}

// 只读
func (t Test) SetTestFalse(i int, n string) {
	t.Id = i
	t.Name = n
}
