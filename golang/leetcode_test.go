package golang

import "testing"

func TestQuickSort(t *testing.T) {
    nums := []int{3,1,13,7,2,87,24}
    QuickSort(nums, 0,  len(nums)-1)
    t.Log(nums)

    nums2 := []int{3,11,13,87,2,87,24}
    QuickSortLoop(nums2, 0, len(nums2)-1)
    t.Log(nums2)
}

func TestTopKFrequent(t *testing.T) {
    nums := []int{1,1,1,2,2,3,4,5}

    result := TopKFrequent(nums, 2)
    t.Logf("result:%+v\n", result)
}

func TestArraySort(t *testing.T) {
    nums := []int{2, 4, 10, 7, 1, 3, 5}
    ArraySort(nums)
}

func TestAddString(t *testing.T) {
    res := addString("3", "21")
    t.Log(res)
}

func TestFindKingNumber(t *testing.T) {
    input := []int{2, 3, 1, 8, 9, 20, 12}
    t.Log(FindKingNumber(input))
}

func TestSubSet(t *testing.T) {
    input := []int{1,2,3}
    t.Log(SubSet(input))
}

func TestQuickSorts(t *testing.T) {
    nums := []int{1,2,5,6,8,7,9,4,3}

    t.Log(FindTopK(nums, 3))
}

func TestAddToArrayForm(t *testing.T) {
    t.Log(AddToArrayForm([]int{9,9,9}, 23))
}


func TestMagicPrint(t *testing.T) {
    MagicPrint([][]int{{1,2,3,4}, {5,6,7,8}, {9,10,11,12}})
}

func TestJumpGame(t *testing.T) {
    t.Log(JumpGame([]int{4,2,1,1,4}))
}

func TestMoveZero(t *testing.T) {
    nums := []int{0, 1, 0, 3, 12}
    MoveZero(nums)
    t.Log(nums)
}

func TestVote(t *testing.T) {
    t.Log(Vote([]int{1,2,3,4,3,3,5,3,4,3}))
}

func TestCoinChange(t *testing.T) {
    t.Log(CoinChange([]int{1,3,6}, 11))
}