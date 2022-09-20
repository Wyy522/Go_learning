package series

import "fmt"

// GetFibonacci 第一个字母大写声明是可以被包外访问的，反之不行
func GetFibonacci(n int) []int {
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList
}

// 小写则不行
func getFibonacci(n int) []int {
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList
}

func init() {
	fmt.Println("first init")
}
func init() {
	fmt.Println("second init")
}
