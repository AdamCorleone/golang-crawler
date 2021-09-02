package main

import (
	"fmt"
)

func main() {
	a := []int{1, 5, 6, 10, 3, 0, 8,4}
	quickSort(a, 0, len(a)-1)
	fmt.Println(a)
}

func quickSort(arr []int, start int, end int) {
	if start < end {
		left, right := start, end
		pivot := arr[(start+end)/2]
		for left <= right {
			for arr[left] < pivot {
				left++
			}
			for arr[right] > pivot {
				right--
			}

			if left <= right {
				arr[left], arr[right] = arr[right], arr[left]
				left++
				right--
			}
		}

		if start < left {
			quickSort(arr, start, right)
		}

		if end > left {
			quickSort(arr, left, end)
		}
	}
}
