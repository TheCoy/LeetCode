package golang

func lengthOfLongestSubstring(s string) int {
	hashMap := make(map[byte]int)
	right, max, length := -1, 0, len(s)

	for i := 0; i < length; i++ {
		if i != 0 {
			delete(hashMap, s[i])
		}

		for right + 1 < length && hashMap[s[right]] == 0 {
			right++
			hashMap[s[right]]++
		}

		if right - i + 1 > max {
			max = right - i + 1
		}
	}

	return max
}


