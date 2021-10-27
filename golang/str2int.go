package golang

import (
	"math"
)

func MyAtoi(s string) int {
	result:= 0
	flag := 1
	started := false
	for _, c := range s {
		if !started {
			if c == ' ' {
				continue
			}
			if c == '+' {
				flag = 1
				started = true
				continue
			}
			if c == '-' {
				flag = -1
				started = true
				continue
			}

		}
		started = true
		if started && (c < '0' || c > '9') {
			break
		}

		tmp := c - '0'
		result = result * 10 + int(tmp)
		if result*flag > math.MaxInt32 {
			return math.MaxInt32
		}else if result * flag < math.MinInt32{
			return math.MinInt32
		}

	}

	return result*flag
}
