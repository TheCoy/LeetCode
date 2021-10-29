package golang

import "testing"

func TestKMP(t *testing.T) {
    str := "ACAFADFAS"
    subStr := "DAD"
    t.Log(KMP(str, subStr))
}
