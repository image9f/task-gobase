package main

import (
	"fmt"
)

func findShortestStringLength(strs []string) int {
	if len(strs) == 0 {
		return 0
	}
	shortest := strs[0]
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < len(shortest) {
			shortest = strs[i]
		}
	}
	return len(shortest)
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	var shortest_len int = findShortestStringLength(strs)

	for i := 0; i < shortest_len; i++ {
		for j := 1; j < len(strs); j++ {
			if strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}

	return strs[0][:shortest_len]
}

func main() {
	fmt.Println("测试字符数组最长公共前缀")
	s1 := []string{"test1", "testt2", "testtt3"}
	fmt.Println("s1 最长公共前缀:", longestCommonPrefix(s1))
	s2 := []string{"dog", "racecar", "car"}
	fmt.Println("s2 最长公共前缀:", longestCommonPrefix(s2))

	s3 := []string{}
	fmt.Println("s3 最长公共前缀:", longestCommonPrefix(s3))

}
