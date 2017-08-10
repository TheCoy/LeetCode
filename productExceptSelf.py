def productExceptSelf(nums):
    size = len(nums)
    output = [1] * size
    print output
    left = right = 1
    for i in range(size-1):
        print 'output[', i, ']' 
        left *= nums[i]
        output[i+1] *= left
    for j in range(size-1,0,-1):
        right *= nums[j]
        output[j-1] *= right
    return output
nums = [2, 4, 3, 1]
print productExceptSelf(nums)
