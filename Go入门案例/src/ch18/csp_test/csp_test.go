package csp_test

import (
	"fmt"
	"testing"
	"time"
)

// 没有使用csp 串行执行
func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done.")
}

// 使用了csp，异步返回结果
func TestAsynService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh) //从channel中取数据
	time.Sleep(time.Second * 1)
}

func AsyncService() chan string {
	retCh := make(chan string, 1) //声明一个buffer-channel 容量为1 协程不会阻塞
	//retCh := make(chan string)//channel 会阻塞
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret //往channel中存数据
		fmt.Println("service exited.")
	}()
	return retCh
}

//
