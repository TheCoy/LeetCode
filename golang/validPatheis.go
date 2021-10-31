package golang

import "fmt"

func isValid(s string) bool {
    for c := range s{
        fmt.Printf("%T", c)
    }

    return true
}

func generateParenthesis(n int) []string {
    result := []string{}

    var loop func(left, right int, path string)
    loop = func(left, right int, path string) {
        if len(path) == 2*n {
            result = append(result, path)
        }

        if left > 0 {
            loop(left-1, right, path+"(")
        }

        if left < right {
            loop(left, right-1, path+")")
        }
    }

    loop(n, n, "")
    return result
}