class Solution(object):
    def reverse(self, x):
        result = 0
        flag = False
        if x < 0:
            x *= -1
            flag = True
        while x:
            val = x % 10
            result = (result*10 + val)
            x /= 10
        if result > 0x7FFFFFFF:
            return 0
        if flag:
            result *= -1
            
        return result