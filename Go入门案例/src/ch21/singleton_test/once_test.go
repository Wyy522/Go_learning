package singleton_test

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type Singleton struct {
}

var singleInstance *Singleton
var Once sync.Once
var wg sync.WaitGroup

// 单例模式 懒汉
func GetSingletonObj() *Singleton {
	Once.Do(func() {
		fmt.Println("creat obj")
		singleInstance = new(Singleton)
	})
	return singleInstance
}

func TestGetSingletionObj(t *testing.T) {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		i := i
		go func() {
			obj := GetSingletonObj()
			//fmt.Printf("%d &      :%d\n", i, &obj)
			//fmt.Printf("%d *      :%d\n", i, &obj)
			fmt.Printf("%d unsafe :%d\n", i, unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}
