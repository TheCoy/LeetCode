package golang

//input : "[a-z]*"
//output: wordCount
func WordCount(input string) int {
    cnt :=  0
    skip := true
    for _, v := range []byte(input) {
        c := string(v)
        if c == " " {
            skip = true
            continue
        }
        if skip && c != " " {
            cnt++
            skip = false
        }
    }

    return cnt
}
