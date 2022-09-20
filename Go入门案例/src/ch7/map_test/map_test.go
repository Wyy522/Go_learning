package map_test

import "testing"

// Map创建
func TestMapInit(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log("m1[2]=", m1[2])       //4
	t.Logf("len m1=%d", len(m1)) //3

	m2 := map[int]int{}
	m2[2] = 16
	t.Logf("len m2=%d", len(m2)) //1

	m3 := make(map[int]int, 10)  //这个10是cap
	t.Logf("len m3=%d", len(m3)) //0  如果是Cap(m3)则会报错
}

// 是否存在Value
func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1]) //无值返回0
	m1[2] = 0
	t.Log(m1[2]) //有值就是0

	//如何区分Value是否有值?
	if v, ok := m1[3]; ok {
		t.Log("key is existing =", v)
	} else {
		t.Log("key is not existing")
	}
}

func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	for k, v := range m1 {
		t.Log(k, v)
	}
}
