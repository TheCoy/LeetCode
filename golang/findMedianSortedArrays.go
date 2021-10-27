package golang

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n, m:= len(nums1), len(nums2)
	index, hasTwo := (n + m) / 2, (n + m) % 2 == 0

	var result float64
	if n <= 0 {
		if hasTwo {
			result = float64(nums2[index] + nums2[index-1]) / 2
		}else{
			result = float64(nums2[index])
		}
	}else if m <= 0 {
		if hasTwo {
			result = float64(nums1[index] + nums1[index-1]) / 2
		}else{
			result = float64(nums1[index])
		}
	}else {
		var i, j, k int
		var last, cur int
		if nums1[0] < nums2[0] {
			last, cur = nums1[0], nums1[0]
			i++
		}else {
			last, cur = nums2[0], nums2[0]
			j++
		}
		for ; k <= index ;k++ {
			if i >= n - 1 {
				j++
				last, cur = cur, nums2[j]
			}else if j >= m - 1 {
				i++
				last, cur = cur, nums1[i]
			}else {
				if nums1[i] <= nums2[j] {
					last, cur = cur, nums1[i]
					i++
				}else {
					last, cur = cur, nums2[j]
					j++
				}
			}
		}
		if hasTwo {
			result = float64(cur + last) / 2
		}else {
			result = float64(cur)
		}
	}

	return result
}
