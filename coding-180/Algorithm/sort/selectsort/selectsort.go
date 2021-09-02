package main

import "fmt"

//选择排序
func main() {
	arr := []int{3, 6, 4, 2, 11, 10, 5}
	fmt.Println(SelectionSort(arr))
}

func SelectionSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		min := i
		for j := i+1; j < length ; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i],arr[min] = arr[min],arr[i]
	}
	return arr
}