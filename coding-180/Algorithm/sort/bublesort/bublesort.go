package main

import "fmt"

func main() {
	arr := []int{3, 6, 4, 2, 11, 10, 5}
	fmt.Println(BubbleSort(arr))
}

// BubbleSort 每次和相邻的元素比较，内层每循环一次会把最大的循环到最后
func BubbleSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		//j < length -i -1 原因：每循环一次，最后一位数已排好，不用再比
		for j := 0; j < length-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
