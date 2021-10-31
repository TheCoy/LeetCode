package golang

import (
    "strconv"
)

func addString(a, b string) string{
    byteA  := []byte(a)
    byteB  := []byte(b)
    lenA := len(byteA)
    lebB := len(byteB)

    result := make([]byte ,0)
    res := ""
    var flag int
    for i, j := lenA-1, lebB-1; i >= 0 || j >= 0;{
        var cur int
        if i >= 0 {
            cur += int(byteA[i] - '0')
        }
        if j >= 0 {
            cur += int(byteB[j] - '0')
        }
        cur +=  flag
        result  = append(result, strconv.Itoa(cur%10)[0])
        res =  strconv.Itoa(cur%10) + res
        flag = cur / 10
        i--
        j--
    }

    byteResult := make([]byte, 0)
    for i := len(result) - 1; i >=0; i--{
        byteResult =  append(byteResult, result[i])
    }
    return string(byteResult)
}
