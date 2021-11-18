/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 *
 * https://leetcode-cn.com/problems/median-of-two-sorted-arrays/description/
 *
 * algorithms
 * Hard (40.94%)
 * Likes:    4610
 * Dislikes: 0
 * Total Accepted:    537.9K
 * Total Submissions: 1.3M
 * Testcase Example:  '[1,3]\n[2]'
 *
 * 给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums1 = [1,3], nums2 = [2]
 * 输出：2.00000
 * 解释：合并数组 = [1,2,3] ，中位数 2
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums1 = [1,2], nums2 = [3,4]
 * 输出：2.50000
 * 解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums1 = [0,0], nums2 = [0,0]
 * 输出：0.00000
 *
 *
 * 示例 4：
 *
 *
 * 输入：nums1 = [], nums2 = [1]
 * 输出：1.00000
 *
 *
 * 示例 5：
 *
 *
 * 输入：nums1 = [2], nums2 = []
 * 输出：2.00000
 *
 *
 *
 *
 * 提示：
 *
 *
 * nums1.length == m
 * nums2.length == n
 * 0
 * 0
 * 1
 * -10^6
 *
 *
 *
 *
 * 进阶：你能设计一个时间复杂度为 O(log (m+n)) 的算法解决此问题吗？
 *
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}
	len1, len2 := len(nums1), len(nums2)
	left, right := 0, len1
	median1, median2 := 0, 0

	for left <= right {
		i := (left + right) / 2
		j := (len1+len2+1)/2 - i
		leftAMax := math.MinInt32
		if i != 0 {
			leftAMax = nums1[i-1]
		}
		rightAMin := math.MaxInt32
		if i != len1 {
			rightAMin = nums1[i]
		}
		rightBMin := math.MaxInt32
		if j != len2 {
			rightBMin = nums2[j]
		}
		leftBMax := math.MinInt32
		if j != 0 {
			leftBMax = nums2[j-1]
		}
		if leftAMax <= rightBMin {
			median1 = max(leftAMax, leftBMax)
			median2 = min(rightAMin, rightBMin)
			left = i + 1
		} else {
			right = i - 1
		}
	}
	if (len1+len2)%2 == 0 {
		return float64(median1+median2) / 2
	} else {
		return float64(median1)
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

