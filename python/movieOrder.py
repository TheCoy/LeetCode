def findResult(n,m):
	result = 0
	if m > 1:
		for i in range(n-m+2):
			result += A(n,i) * findResult(n-i, m-1)
	else:
		result += A(n,n)
	return result
def A(n,m):
	result = 1
	for i in range(m):
		result *= (n-i)
	return result

if __name__ == "__main__":
	import sys
	result = findResult(int(sys.argv[1]),int(sys.argv[2])) * A(int(sys.argv[2]),int(sys.argv[2]))
	print("Result = %d" % result)
