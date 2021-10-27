package golang

// 测试用例：
// [
// [2],
// [3,4],
// [6,5,7],
// [4,1,8,3]
// ]
func MinSumOfTriangle(triangle [][]int) int {
	if len(triangle) <= 0 {
		return 0
	}
	output := make([][]int, len(triangle))
	for i := 0; i < len(triangle); i++ {
		output[i] = make([]int, len(triangle[i]))
	}
	output[0][0] = triangle[0][0]

	for i := 1; i < len(triangle); i++ {
		output[i][0] = output[i-1][0] + triangle[i][0]
		output[i][len(triangle[i])-1] = output[i-1][len(triangle[i-1])-1] + triangle[i][len(triangle[i])-1]
		for j := 1; j < len(triangle[i])-1 ; j++ {
			output[i][j] = min(output[i-1][j], output[i-1][j-1]) + triangle[i][j]
		}
	}

	ans := output[len(triangle)-1][0]
	for i := 1; i < len(output[len(triangle)-1]); i++ {
		ans = min(ans, output[len(triangle)-1][i])
	}
	return ans
}

