package main

import (
	"strconv"
	"strings"
)

type CodecX struct {
}

func ConstructorXIV() CodecX {
	return CodecX{}
}

// Serializes a tree to a single string.
func (this *CodecX) serialize(root *TreeNode) string {

	var tree []string
	if root == nil {
		return ""
	}

	q := []*TreeNode{root}

	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node != nil {
			tree = append(tree, strconv.Itoa(node.Val))
			q = append(q, node.Left)
			q = append(q, node.Right)
		} else {
			tree = append(tree, "null")
		}
	}
	return strings.Join(tree, ",")
}

// Deserializes your encoded data to tree.
func (this *CodecX) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}

	values := strings.Split(data, ",")
	rootVal, _ := strconv.Atoi(values[0])
	root := &TreeNode{Val: rootVal}
	q := []*TreeNode{root}

	i := 1

	for i < len(values) {

		node := q[0]
		q = q[1:]

		if values[i] != "null" {
			val, _ := strconv.Atoi(values[i])
			node.Left = &TreeNode{Val: val}
			q = append(q, node.Left)
		}
		i++

		if i < len(values) && values[i] != "null" {
			val, _ := strconv.Atoi(values[i])
			node.Right = &TreeNode{Val: val}
			q = append(q, node.Right)
		}
		i++

	}
	return root
}
