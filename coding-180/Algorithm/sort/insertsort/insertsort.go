package main

import "fmt"

func main() {
	arr := []int{3, 6, 4, 2, 11, 10, 5}
	fmt.Println(insertSort(arr))
}

func insertSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j > 0; j-- {
			fmt.Println(j)
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			} else {
				break
			}
		}
	}
	return arr
}


