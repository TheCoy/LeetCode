package golang

//输入：[1,2]
//输出：[[], [1], [2], [1,2]]

func SubSet(arr []int) [][]int {
    result := make([][]int, 0)

    empty := make([]int, 0)
    trackBack(empty, arr, 0, &result)

    return result
}

func trackBack(input, arr []int, deep int, result *[][]int) {
    if deep == len(arr) {
        *result = append(*result, input)
        return
    }

    newInput := make([]int, len(input))
    copy(newInput, input)
    trackBack(newInput, arr, deep+1, result)
    newInput = append(newInput, arr[deep])
    trackBack(newInput, arr, deep+1, result)

}
