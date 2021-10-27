package golang

import (
	"fmt"
	"strings"
)
// WordPattern("abba", "北京 杭州 杭州 北京")
func WordPattern(pattern string, str string) bool {
	arrInput := strings.Split(str, " ")
	mm := map[rune]string{}
	valMap := map[string]int{}
	for i, c := range pattern {
		fmt.Println("i:", c, "str:", arrInput[i])
		if val, ok := mm[c]; ok {
			return val == arrInput[i]
		}else{
			if _, ok := valMap[arrInput[i]]; ok {
				return false
			}else{
				mm[c] = arrInput[i]
				valMap[arrInput[i]] = 1
			}
		}
	}

	return true
}
