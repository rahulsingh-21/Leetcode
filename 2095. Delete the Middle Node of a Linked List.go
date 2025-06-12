package main

func deleteMiddle(head *ListNode) *ListNode {

	slow := head
	fast := head

	if head == nil || head.Next == nil {
		return nil
	}
	var temp *ListNode
	for fast != nil && fast.Next != nil {
		temp = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	temp.Next = slow.Next
	return head

}
