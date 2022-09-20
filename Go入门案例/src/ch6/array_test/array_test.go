package array_test

import (
	"testing"
)

// 数组创建
func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [...]int{1, 2, 3, 4, 5}
	arr3 := [3][1]int{{1}, {2}, {3}}
	t.Log(arr, arr1, arr2, arr3)
}

// 数组遍历
func TestArrayTravel(t *testing.T) {
	arr2 := [...]int{1, 2, 3, 4, 5}

	//最老的写法
	for i := 0; i < len(arr2); i++ {
		t.Log(arr2[i])
	}

	//idx是索引值,e是元素值,_代表不关心这个值
	for idx, e := range arr2 {
		t.Log(idx, e)
	}
	for _, e := range arr2 {
		t.Log(e)
	}
}

// 数组截取(常用)  a[开始索引(包含):结束索引(不包含)]
func TestArraySection(t *testing.T) {
	arr2 := [...]int{1, 2, 3, 4, 5}
	arr2Sec1 := arr2[1:len(arr2)] //2,3,4,5
	arr2Sec2 := arr2[1:]          //2,3,4,5
	arr2Sec3 := arr2[:]           //1,2,3,4,5
	t.Log(arr2Sec1, arr2Sec2, arr2Sec3)
}
