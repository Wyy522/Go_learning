package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 多返回值
func TestFunc(t *testing.T) {
	a, b := returnMultiValues()
	t.Log(a, b)
}
func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

// programming function 函数式编程
func TestProgramFunc(t *testing.T) {
	tsSF := timeSpent(SlowFun)
	t.Log(tsSF(10))
}
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}
func SlowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

// 可变长参数
func TestVarParam(t *testing.T) {
	s := Sum(1, 2, 3, 4, 5)
	m := Sum(1, 2, 3, 4, 5)
	t.Log(s, m)
}

func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

// Defer延迟执行函数(清理资源和释放锁类似于try finally)
func TestDefer(t *testing.T) {
	defer Clear()
	fmt.Println("执行结束后")
	panic("err") //无法修复的错误fatal
}

func Clear() {
	fmt.Println("clear resources")
}
