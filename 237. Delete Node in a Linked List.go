package linkedlist

func deleteNode(node *ListNode) {
	*node = *node.Next
}
