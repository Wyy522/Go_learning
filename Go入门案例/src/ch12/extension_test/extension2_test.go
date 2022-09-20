package extension

import (
	"fmt"
	"testing"
)

type pet struct {
}

func (p *pet) Speak() {
	fmt.Println("...")
}

func (p *pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(host)
}

// 匿名嵌套类型
type Dog struct {
	pet
}

func TestDog2(t *testing.T) {
	dog := new(Dog)
	dog.SpeakTo("chao")
}
