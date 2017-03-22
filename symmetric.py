class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution(object):
    def isSymmetric(self, root):
        if root == None:
        	return True
        else:
        	return self.isMirror(root.left,root.right)
    def isMirror(self,p,q):
    	if p == None and q ==None:
<<<<<<< HEAD
    		return True
    	if p == None or q == None:
    		return False
=======
    			return True
    	if p == None or q == None:
    			return False
>>>>>>> 8247000c45f3b0cbbb338b16a35329a6dfbc15d7
    	return p.val == q.val and self.isMirror(p.left,q.right) and self.isMirror(p.right,q.left)