class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution(object):
	def levelOrder(self,root):
		if not root:
			return []
		res,level = [],[root]
		while level:
			res.append([node.val for node in level])
			temp = []
			for node in level:
				temp.extend([node.left,node.right])
			level = [leaf for leaf in temp if leaf]
		return res