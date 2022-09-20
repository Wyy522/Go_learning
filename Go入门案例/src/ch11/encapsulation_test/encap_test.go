package encapsulation_test

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

// 创建实例(封装)
func TestFirstTry(t *testing.T) {

	e1 := Employee{"0", "Bob", 18}
	fmt.Printf("Address is :%x", unsafe.Pointer(&e1.Name))
	s := e1.String()
	t.Log(s)

	e2 := Employee{Name: "Mike", Age: 20}
	t.Log(e2)

	e3 := new(Employee)
	e3.Id = "2"
	e3.Name = "Rose"
	e3.Age = 19
	t.Log(*e3)
}

//两种传递的区别
//func (e Employee) String() string {
//	fmt.Printf("Address is :%x", unsafe.Pointer(&e.Name))//复制了新的实例，地址不同
//	return fmt.Sprintf("ID:%s,Name:%s,Age:%d", e.Id, e.Name, e.Age)
//	//ID:0,Name:Bob,Age:18
//}

func (e *Employee) String() string {
	fmt.Printf("Address is :%x", unsafe.Pointer(&e.Name)) //没有复制新的实例，同一个实例，地址相同
	return fmt.Sprintf("ID:%s,Name:%s,Age:%d", e.Id, e.Name, e.Age)
	//ID:0,Name:Bob,Age:18
}
