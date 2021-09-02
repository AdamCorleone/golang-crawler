package main

import (
	"fmt"
)

func main() {
	fmt.Println("不重复子串长度为:",lengthOfLongestSubstring("pwwkew"))
}

//滑动窗口
func lengthOfLongestSubstring(s string) int {
	strLen := len(s)
	if strLen == 0 {
		return 0
	}
	left, right, ans := 0, 0, 0
	m := make(map[byte]int)
	for right < strLen {
		if _, ok := m[s[right]]; !ok {
			fmt.Println(string(s[right]),right)
			m[s[right]] = right//用字符（s[right]）表示map的key，用字符所在的下标（right）表示 map的value （此处有妙用）
		} else {
			//1、我们应该找到重复元素第一次（最近一次更新）添加到map中的value值，然后把该value值加1，既为最新的 “左窗棱” 的位置
			//2、更新后的“左窗棱”不能比之前的 “左窗棱” 小
			//3、更新map中重复元素最新的下标
			fmt.Println("未更新left坐标:",left)
			if m[s[right]]+1 >= left {
				left = m[s[right]] + 1
				fmt.Println("right:",right)
				fmt.Println("更新后left坐标:",left)
			}
			m[s[right]] = right //更新map中value值  其中 key=s[right] = w value = 1,2,5  如果没有则会从[w,1]取值，导致left不能更新为3，结果错误
			fmt.Println("map中重复元素下标:",m[s[right]])
			fmt.Println("更新后map:",string(s[right]),right)
		}
		ans = max(right-left+1, ans)
		right++
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
