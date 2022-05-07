package golang

import (
	"sort"
	"strconv"
)

func CombinationSum2(candidates []int, number int) [][]int {
	ans := make([][]int, 0)
	visit := make([]int, len(candidates))
	mm := map[string]bool{}
	sort.Ints(candidates)
	var fn func([]int, int)
	fn = func(digits []int, target int) {

		if target == 0 {
			arr := []int{}
			key := ""
			for k, n := range digits {
				if n > 0 {
					arr = append(arr, candidates[k])
					key += strconv.Itoa(candidates[k])
				}
			}
			if !mm[key] {
				//fmt.Println(digits)
				ans = append(ans, arr)
				mm[key] = true
			}
			return
		}
		for k, n := range candidates {
			if digits[k] == 0 && target >= n {
				digits[k]++
				fn(digits, target-n)
				digits[k]--
			}
		}
	}
	fn(visit, number)
	return ans
}
