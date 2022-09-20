package string_test

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFn(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",") //分割
	t.Log(parts)
	for _, v := range parts {
		t.Log(v)
	}

	t.Log(strings.Join(parts, "-")) //连接

}

func TestStringConv(t *testing.T) {
	//整数转字符串
	s := strconv.Itoa(10)
	t.Log("str" + s)

	//字符串转整数
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + i)
	} else {
		t.Log("something worrng")
	}
}
