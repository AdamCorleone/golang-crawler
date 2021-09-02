package main

import "fmt"

func main() {
	arr := []int{3, 6, 4, 2, 11, 10, 5}
	fmt.Println(MergeSort(arr))
}

// MergeSort 归并排序
func MergeSort(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}
	mid := length / 2
	left := MergeSort(arr[0:mid])
	right := MergeSort(arr[mid:])
	res := merge(left, right)
	return res
}

//合并数组
func merge(left, right []int) []int {
	result := make([]int, 0)
	m, n := 0, 0
	l, r := len(left), len(right)
	//比较两个数组，谁小把元素值添加到结果集内
	for m < l && n < r {
		if left[m] > right[n] {
			result = append(result, right[n])
			n++
		} else {
			result = append(result, left[m])
			m++
		}
	}
	//如果有一个数组比完了，另一个数组还有元素的情况，则将剩余元素添加到结果集内
	result = append(result, right[n:]...)
	result = append(result, left[m:]...)
	return result
}