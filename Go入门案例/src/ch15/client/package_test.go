package client

import (
	"HelloWorld/src/ch15/series"
	"testing"
)

// * 获取远程包
// * 获取远程包
// * go get github.com/easierway/concurrent_map
// * 更新远程包（第一次获取也可以这样用）
// * go get -u github.com/easierway/concurrent_map
// 1. 通过go get来获取远程依赖
// * go get -u 强制从网络更新远程依赖

func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacci(10))
}
