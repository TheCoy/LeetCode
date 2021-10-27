package golang

import (
	_ "math"
	"sort"
)

//幂集
func SubSet(nums []int) [][]int {
	result := make([][]int, 0)
	list := make([]int, 0)

	subbacktrack(nums, 0, list, &result)

	return result
}

func subbacktrack(nums []int, pos int, list []int, result *[][]int) {
	res := make([]int, len(list))
	copy(res, list)
	*result = append(*result, res)
	for i := pos; i < len(nums); i++ {
		list = append(list, nums[i])
		subbacktrack(nums, i+1, list, result)
		list = list[0:len(list)-1]
	}
}

//全排列
func permute(nums []int) [][]int {
	result := make([][]int, 0)
	visited := make([]bool, len(nums))
	list := make([]int, 0)
	sort.Ints(nums)

	perBackTrack(nums, 0, list, visited, &result)

	return result
}

func perBackTrack(nums []int, pos int, list []int, visited []bool, result *[][]int) {
	if len(list) == len(nums) {
		res := make([]int, len(list))
		copy(res, list)
		*result = append(*result, res)
		return
	}
	for i := pos; i < len(nums); i++ {
		if i != pos && nums[i] == nums[i-1] {
			continue
		}
		if !visited[i] {
			list = append(list, nums[i])
			visited[i] = true
			perBackTrack(nums, i+1, list, visited, result)
			list = list[0:len(list)-1]
			visited[i] = false
		}
	}
}
