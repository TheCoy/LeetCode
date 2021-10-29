package golang

func makeNext(input []rune) []int {
    length := len(input)
    result := make([]int, length)
    result[0] = 0
    for q, k := 1, 0; q < length; q++ {
        for k > 0 && input[q] != input[k] {
            k = result[k-1]
        }
        if input[q] == input[k] {
            k++
        }
        result[q] = k
    }

    return result
}

func KMP(str, substr string) int {
    next := makeNext([]rune(substr))

    for i, q := 0, 0; i < len(str); i++ {
        for q > 0 && substr[q] != str[i] {
            q = next[q-1]
        }
        if str[i] == substr[q] {
            q++
        }
        if q == len(substr) {
            return i - len(substr) + 1
        }
    }

    return -1
}
