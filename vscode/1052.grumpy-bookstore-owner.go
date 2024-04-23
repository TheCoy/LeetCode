/*
 * @lc app=leetcode.cn id=1052 lang=golang
 * @lcpr version=30122
 *
 * [1052] 爱生气的书店老板
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	right := 0
	maxSum := 0
	for i := 0; i < len(customers); i++ {
		if grumpy[i] == 0 {
			maxSum += customers[i]
		}
	}
	extraSum := 0
	for right < minutes {
		if grumpy[right] == 1 {
			extraSum += customers[right]
		}
		right++
	}

	sumOfWindow := extraSum
	for ; right < len(customers); right++ {
		left := right - minutes
		if grumpy[left] == 1 {
			sumOfWindow -= customers[left]
		}
		if grumpy[right] == 1 {
			sumOfWindow += customers[right]
		}
		extraSum = max(sumOfWindow, extraSum)
	}

	//fmt.Println("start =", start)
	return maxSum + extraSum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

/*
// @lcpr case=start
// [1,0,1,2,1,1,7,5]\n[0,1,0,1,0,1,0,1]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1]\n[0]\n1\n
// @lcpr case=end

*/

