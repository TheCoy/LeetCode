class Solution(object):
    def twoSum(self, nums, target):
        res = {}
        for i in range(len(nums)):
            if (target-nums[i]) in res:
                return [res[target-nums[i]],i]
            else :
                res[nums[i]] = i