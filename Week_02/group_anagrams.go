package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 异位词分组

// 1.排序+map：先对每个字符排序，然后构建一个map，来分组，时间复杂度为：n(klogk)  空间复杂度为o(nk)
// 2 计数分类：将字符串转为字符数count，由26个字母组成，由字符数count来构建分组key，时间复杂度：o(nk),空间复杂度为o(nk)

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return nil
	}

	base := 'a'
	var builder strings.Builder
	m := make(map[string][]string)
	for _, v := range strs {
		tmp := [26]int{}
		builder.Reset()
		for _, t := range v {
			tmp[int(t-base)]++
		}

		for _, v1 := range tmp {
			builder.WriteString("#")
			if v1 != 0 {
				builder.WriteString(strconv.Itoa(v1))
			}
		}

		s := builder.String()
		if _, ok := m[s]; ok {
			m[s] = append(m[s], v)
		} else {
			m[s] = []string{v}
		}
	}

	var result [][]string
	for _, v := range m {
		result = append(result, v)
	}

	return result
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
