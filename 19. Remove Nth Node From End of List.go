package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	fast := head
	slow := head

	if head == nil || head.Next == nil {
		return nil
	}

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	if fast == nil {
		return head.Next
	}

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next
	return head
}
