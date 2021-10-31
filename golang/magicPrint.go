package golang

import "fmt"

//input:
//1 2 3 4
// 5 6 7 8 
//9 10 11 12 

// output:
// 4 (0,3)
// 3(0,2) 8(1,3)
//  2(0,1) 7(1,2) 12(2,3)
//  1(0,0) 6(1,1) 11(2,2) 
// 5(1,0) 10(2,1) 
// 9 (2,0)
func MagicPrint(nums [][]int) {
    col := len(nums[0]) //4
    row := len(nums)    //3

    for k := col - 1; k > 0; k-- {
        for i, j := 0, k; i < row && j < col; {
            fmt.Printf("%d\t", nums[i][j])
            i++
            j++
        }
        fmt.Println()
    }

    for k := 0; k < row; k++ {
        for i, j := k, 0; i < row && j < col; {
            fmt.Printf("%d\t", nums[i][j])
            i++
            j++
        }
        fmt.Println()
    }
}
