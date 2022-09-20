package interface_test

import (
	"testing"
)

type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {
}

// DuckType 鸭子类型
func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"helloworld\")"
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())

	var prog Programmer = &GoProgrammer{}
	t.Log(prog.WriteHelloWorld())

	//&GoProgrammer{} 等同于 new(GoProgrammer) 接口对应的是指针类型
}
