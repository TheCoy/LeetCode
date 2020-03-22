import sys

def sol(n, a):
    if(len(a) < 2):
        return a*n
    res = [a[0]]*a[0]
    for i in range(1,n):
        if len(res) < 2:
            res += [a[i%len(a)]]*a[1]
        else:
            res += [a[i%len(a)]]*res[i]
    return res

if __name__ == "__main__":
    a = sys.stdin.readline().strip().split()
    n , m = int(a[0]),int(a[1])
    line = sys.stdin.readline().strip()
    values = list(map(int, line.split()))
    result = sol(n, values)
    print(result)