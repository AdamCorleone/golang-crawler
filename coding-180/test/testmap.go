package main

import "fmt"

func main() {
	type deviceData struct {
		Temperature float64 `json:"temperature"`
		Stirrer     float64 `json:"stirrer"`
		Timestamp   int64   `json:"timestamp"`
	}
	m := make(map[int]deviceData)
	dataList := deviceData{
		Temperature: 5,
		Stirrer: 10,
		Timestamp: 23123123,
	}
	m[1] = dataList
	for k, v := range m {
		fmt.Println(k,v)
	}
}
