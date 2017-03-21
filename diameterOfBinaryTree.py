class Solution(object):
    maxLen = 0
    def diameterOfBinaryTree(self, root):
        self.maxDepth(root)
        return self.maxLen
    def maxDepth(self,root):
        if root == None:
            return 0
        left = self.maxDepth(root.left)
        right = self.maxDepth(root.right)
        
        self.maxLen = max(self.maxLen,left+right)
        
        return max(left,right)+1