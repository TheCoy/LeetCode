# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None
        
class Solution(object):
    def addTwoNumbers(self, l1, l2):
        result = cur = ListNode(0)
        carry = 0
        while l1 or l2:
            if l1:
                carry += l1.val
                l1 = l1.next
            if l2:
                carry += l2.val
                l2 = l2.next
            cur.next = ListNode(carry%10)
            cur = cur.next
            carry /= 10
        if carry > 0:
            cur.next = ListNode(carry)
        return result.next
