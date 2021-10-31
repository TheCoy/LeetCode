package golang

func CoinChange(coins []int, amount int) int {
    result := make([]int, amount+1)
    for i := 1; i <= amount; i++ {
        result[i] = amount+1
        for _, coin := range coins {
            if i >= coin {
                result[i] = min(result[i], result[i-coin]+1)

            }
        }
    }

    return result[amount]
}
