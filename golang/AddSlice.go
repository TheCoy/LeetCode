package golang

func AddSlice(a,b []int) []int {
	lenA, lenB := len(a), len(b)
	temp := make([]int, len(a) + len(b))
	for i := 0;i <len(a);i++{
		temp[i] += a[lenA-1-i]
	}
	for j := 0;j<lenB;j++{
		temp[j] += b[lenB-1-j]
	}
	for i := 0;i <len(temp);i++{
		if temp[i] >9 {
			temp[i+1] += temp[i]/10
			temp[i] = temp[i]%10
		}
	}

	result := make([]int, 0)
	flag, lenT := 0, len(temp)
	for i := 0;i<lenT;i++{
		if temp[lenT-1-i] == 0 && flag == 0 {
			continue
		}
		flag = 1
		result = append(result, temp[lenT-1-i])
	}

	return result
}
