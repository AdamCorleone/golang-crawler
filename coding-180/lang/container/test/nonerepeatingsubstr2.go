package main

import (
	"fmt"
)

func main() {
	fmt.Println("不重复子串长度为:", lengthOfLongestSubstring("pwwkew"))
}

//滑动窗口
func lengthOfLongestSubstring(s string) int {
	strLen := len(s)
	left, right, ans := 0, 0, 0
	m := make(map[byte]int)
	for right < strLen {
		if _, ok := m[s[right]]; !ok {
			m[s[right]] = right
		} else {
			if m[s[right]]+1 >= left {
				fmt.Println("right:",right)
				left = m[s[right]] +1
			}
			m[s[right]] = right  //更新值
		}
		ans = max(right-left+1,ans)
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
