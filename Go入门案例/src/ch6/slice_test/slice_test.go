package slice_test

import "testing"

// 切片创建  len表示切片当前长度(默认是0) cap表示切片的总容量
func TestSliceInit(t *testing.T) {
	var s0 []int //区别与数组的是没有声明大小 所以切片是可以伸缩的
	t.Log(len(s0), cap(s0))
	//添加元素
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
}

// 切片是可变长的 cap的增长是前一次cap的两倍
func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i) //每次扩容其实是进行拷贝一个新的切片而不是连续存储空间增长
		t.Log(len(s), cap(s))
	}
}

// 共享内存，同一块连续存储空间
func TestSliceShareMemory(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	m1 := s[3:6]
	t.Log(m1, len(m1), cap(m1)) //[4,5,6],3,9
	m2 := s[5:8]
	t.Log(m2, len(m2), cap(m2)) //[6,7,8],3,7
	m2[0] = 1000
	t.Log(m1) //[4 5 1000] 共享内存

}

// slice是不可以比较的
func TestSliceComparing(t *testing.T) {
	//s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	//m := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	//
	//if a == b {
	//	t.Log()
	//}
}
