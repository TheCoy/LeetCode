package golang

import (
	"fmt"
	"math"
	"testing"
)

func TestQuickSort(t *testing.T) {
	nums := []int{3, 1, 13, 7, 2, 87, 24}
	QuickSort(nums, 0, len(nums)-1)
	t.Log(nums)

	nums2 := []int{3, 11, 13, 87, 2, 87, 24}
	QuickSortLoop(nums2, 0, len(nums2)-1)
	t.Log(nums2)
}

func TestTopKFrequent(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3, 4, 5}

	result := TopKFrequent(nums, 2)
	t.Logf("result:%+v\n", result)
}

func TestTopK(t *testing.T) {
	nums := []int{1, 1, 1, 2, 6, 2, 2, 3, 4, 5}

	result := TopK(nums, 2)
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
	input := []int{1, 2, 3}
	t.Log(SubSet(input))
}

func TestQuickSorts(t *testing.T) {
	nums := []int{1, 2, 5, 6, 8, 7, 9, 4, 3}
	QuickSort(nums, 0, len(nums)-1)
	t.Log(nums)
}

func TestAddToArrayForm(t *testing.T) {
	t.Log(AddToArrayForm([]int{9, 9, 9}, 23))
}

func TestJumpGame(t *testing.T) {
	t.Log(JumpGame([]int{4, 2, 1, 1, 4}))
}

func TestMoveZero(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	MoveZero(nums)
	t.Log(nums)
}

func TestVote(t *testing.T) {
	t.Log(Vote([]int{1, 2, 3, 4, 3, 3, 5, 3, 4, 3}))
}

func TestCoinChange(t *testing.T) {
	t.Log(CoinChange([]int{1, 3, 6}, 11))
}

func TestIsNumberic(t *testing.T) {
	s := "1 "
	t.Log(isNumber(s))
}

func TestReorderOddEven(t *testing.T) {
	nums := []int{2, 12, 3, 41, 6, 124, 12, 1, 4, 7, 3, 6}
	ReorderOddEven(nums)
	t.Log(nums)
	t.Log(math.Pow(2, 128))
}

func TestBinarySearch(t *testing.T) {
	nums := []int{2, 3, 4, 12, 15, 17, 22}
	for i := 0; i < 23; i++ {
		t.Logf("target=%d, search result=%d\n", i, BinarySearchIntverse(0, len(nums)-1, i, nums))
	}
}

func TestBinarySearchLastOne(t *testing.T) {
	nums := []int{2, 3, 5, 5, 5, 6, 22}
	t.Log(BinarySearchLastOne(5, nums))
}

func TestLetterCombinations(t *testing.T) {
	nums := "123"
	dict := map[int]string{
		1: "abc",
		2: "de",
		3: "f",
	}
	t.Log(LetterCombinations(nums, dict))
}

func TestSearchInsert(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	t.Log(SearchInsert(nums, 2))
}

func TestRotate(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	//Reverse(nums, 0, len(nums)-1)
	//t.Log(nums)
	//Reverse(nums, 0, 2)
	//t.Log(nums)
	//Reverse(nums, 3, len(nums)-1)
	Rotate(nums, 3)
	t.Log(nums)
}

func TestAddSlice(t *testing.T) {
	t.Log(AddSlice([]int{7, 7, 9}, []int{9}))
}

func TestWordCount(t *testing.T) {
	t.Log(WordCount(" 12 3sa 12421 sada "))
}

func TestMagicPrint(t *testing.T) {
	MagicPrint([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}})
}

func TestMyAtoi(t *testing.T) {
	t.Log(MyAtoi("9223372036854775808"))
	t.Log("minInt", math.MinInt32)
}

func TestWordPattern(t *testing.T) {
	t.Log(WordPattern("abba", "北京 杭州 杭州 北京"))
	t.Log(WordPattern("aabb", "北京 杭州 杭州 北京"))
	t.Log(WordPattern("baab", "北京 杭州 杭州 北京"))
}

func TestFloodFill(t *testing.T) {
	image := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	FloodFill(image, 1, 1, 2)
	t.Log(image)
}

func TestInsertIntoTree(t *testing.T) {
	root := &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 3},
	}

	fmt.Println("before insert:", preOrderTree(root))

	newRoot := InsertIntoTree(root, 2)
	InsertIntoTree(root, 2)
	fmt.Println("after insert:", preOrderTree(newRoot))

}

func TestMinSumOfTriangle(t *testing.T) {
	input := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	t.Log(MinSumOfTriangle(input))
}

func TestWordBreak(t *testing.T) {
	t.Log(WordBreak("dogs", []string{"dog", "s", "gs"}))
}