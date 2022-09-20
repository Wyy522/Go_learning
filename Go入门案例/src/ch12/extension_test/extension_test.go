package extension_test

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

type Dog struct {
	p *pet
}

func (d *Dog) Speak() {
	fmt.Println("Wang!")
	//d.p.Speak()
}

func (d *Dog) SpeakTo(host string) {
	d.Speak()
	fmt.Println(host)
}

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.SpeakTo("chao")
}
