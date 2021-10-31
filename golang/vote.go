package golang

func Vote(nums []int) int {
    num := -1
    count := 0
    for _, n := range nums {
        if count == 0 {
            num = n
            count++
        } else {
            if n == num {
                count++
            } else {
                count--
            }
        }
    }

    if count == 0 {
        return -1
    }

    return num
}
