package condition

import (
	"testing"
)

func TestIfMultiSec(t *testing.T) {
	//普通用法
	if a := 1 == 1; a {
		t.Log("true")
	}

	//正常使用
	//if v, err := TestSomeOn(flag); err == nil {
	//	t.Log("err==nil")
	//} else {
	//	t.Log("err=" + err)
	//}
}

// Switch Case 与C语言的相反，默认是有break的
func TestSwitchMultiSec(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("Even")
		case 1, 3:
			t.Log("Odd")
		default:
			t.Log("it is not 0-3")
		}
	}

}

func TestSwitchMultICondition(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("Even")
		case i%2 == 1:
			t.Log("Odd")
		default:
			t.Log("it is not 0-3")
		}
	}

}
