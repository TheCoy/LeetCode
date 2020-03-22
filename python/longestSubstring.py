class Solution(object):
    def lengthOfLongestSubstring(self, s):
        if len(s) == 0:
            return 0
        maxlen = 0
        res = {}
        j = 0
        for i in range(len(s)):
            if s[i] in res:
            	'''
            	三目运算符
            	d = b if a else c 
            	d = a and b or c 
            	'''
                j = j > res[s[i]]+1 and j or res[s[i]]+1
            res[s[i]] = i
            maxlen = max(maxlen, (i-j+1))
        return maxlen