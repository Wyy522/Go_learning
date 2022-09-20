package error_test

import (
	"errors"
	"fmt"
	"testing"
)

func GetFibonacci(n int) ([]int, error) {
	if n < 0 || n > 100 {
		return nil, errors.New("输入不正确")
	}
	fibList := []int{1, 1}

	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
	t.Log(GetFibonacci(10))
	if v, err := GetFibonacci(-10); err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}
}

// painc和Os.Exit的区别
// recover有可能是僵尸进程(没有可用的资源) 可以通过let it crash来恢复不确定性
func TestPanicVxExit(t *testing.T) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from ", err)
		}
	}()
	fmt.Println("Start")
	panic(errors.New("Something wrong!"))
	//os.Exit(-1)
	//fmt.Println("End")
}
