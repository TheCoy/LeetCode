package golang

import (
	"strconv"
)

func LetterCombinations(nums string, dict map[int]string) []string {
	result := make([]string, 0)
	var trackBack func(deep int, s string)
	trackBack = func(deep int, s string) {
		if deep == len(nums) {
			result = append(result, s)
			return
		}
		index, _ := strconv.Atoi(string(nums[deep]))
		//fmt.Println(dict[index])
		for _, v := range dict[index] {
			trackBack(deep+1, s+string(v))
		}
	}

	trackBack(0, "")

	return result
}
