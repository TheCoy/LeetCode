package golang

func AddToArrayForm(nums []int, k int) []int {
    lenN := len(nums)
    index := 0
    result := make([]int, 0)
    for k > 0 {
        x := k % 10
        var temp int
        if index <= lenN-1 {
            temp = x + nums[lenN-1-index]

        } else {
            temp = x
        }
        result = append(result, temp)
        k /= 10
        index++
    }

    for index < len(nums) {
        result = append(result, nums[lenN-1-index])
        index++
    }

    for i := 0; i < len(result); i++ {
        if result[i] > 9 {
            if i < len(result)-1 {
                result[i+1] += result[i] / 10
            } else {
                result = append(result, result[i]/10)
            }
            result[i] = result[i] % 10
        }
    }

    for i := 0; i < len(result)/2; i++ {
        result[i], result[len(result)-1-i] = result[len(result)-1-i], result[i]
    }

    return result
}
