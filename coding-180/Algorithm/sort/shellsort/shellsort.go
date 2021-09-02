package main

import "fmt"

func main() {
	arr := []int{3, 6, 4, 2, 11, 10, 5}
	shellSort(arr)
	fmt.Println("arr:", arr)
	s := []int{9, 0, 6, 5, 8, 2, 1, 7, 4, 3}
	shellSort(s)
	fmt.Println("s:", s)
}

func shellSort(arr []int) {
	length := len(arr)
	h := 1
	for h < length/3 {
		h = 3*h + 1
	}
	for gap := h; gap > 0; gap = (gap - 1) / 3 {
		for i := gap; i < length; i++ {
			for j := i; j >= gap && arr[j] < arr[j-gap]; j -= gap {
				arr[j-gap], arr[j] = arr[j], arr[j-gap]
			}
		}
	}
}
