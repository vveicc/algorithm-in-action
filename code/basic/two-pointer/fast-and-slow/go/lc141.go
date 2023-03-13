package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head != nil {
		slow, fast := head, head
		for fast.Next != nil && fast.Next.Next != nil {
			slow, fast = slow.Next, fast.Next.Next
			if slow == fast {
				return true
			}
		}
	}
	return false
}
