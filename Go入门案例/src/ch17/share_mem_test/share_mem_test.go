package share_mem_test

import (
	"sync"
	"testing"
	"time"
)

// 有并发安全问题
func TestCounter(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter =%d", counter)
}

// 通过mutex锁来保证并发安全
func TestCounterThreadSafe(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	time.Sleep(1 * time.Second) //之所以要sleep是因为上面协程计算需要时间，也可以用waitGroup实现
	t.Logf("counter =%d", counter)
}

// waitGroup 类似于 join
func TestCounterWaitGroup(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup

	counter := 0
	now := time.Now()
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	t.Log("time is", time.Since(now).String())
	t.Logf("counter =%d", counter)
}
