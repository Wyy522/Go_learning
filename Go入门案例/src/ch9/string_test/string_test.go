package string_test

import "testing"

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	c := []rune(s) //取s的unicode
	t.Log("s的unicode=", c)
	for _, x := range s {
		t.Logf("%[1]x %[1]d %[1]c", x)
	}

}
