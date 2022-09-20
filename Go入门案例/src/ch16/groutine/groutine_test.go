package groutine_test

import (
	"fmt"
	"testing"
)

// go的协程(一个轻量级的线程 和内核空间实体成多对多的关系，而java是一比一的关系)
func TestGoroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
}
