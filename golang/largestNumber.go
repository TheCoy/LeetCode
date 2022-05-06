package golang

import (
	"sort"
	"strconv"
)

func LargestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		x, y := nums[i], nums[j]
		sx, sy := 10, 10
		for sx <= x {
			sx *= 10
		}
		for sy <= y {
			sy *= 10
		}
		return sx*y+x < sy*x+y
	})

	ans := ""
	for _, n := range nums {
		ans += strconv.Itoa(n)
	}
	return ans
}
