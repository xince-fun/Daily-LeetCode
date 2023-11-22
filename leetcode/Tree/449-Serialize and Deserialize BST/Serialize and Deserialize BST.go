package main

import (
	"strconv"
	"strings"

	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	ans := ""
	if root == nil {
		return "[]"
	}
	// 层序
	ans += "["
	q := []*TreeNode{root}
	for len(q) > 0 {
		length := len(q)
		for i := 0; i < length; i++ {
			node := q[0]
			q = q[1:]
			if node != nil {
				ans += strconv.Itoa(node.Val) + ","
				q = append(q, node.Left)
				q = append(q, node.Right)
			} else {
				ans += "null,"
			}
		}
	}
	ans = strings.TrimSuffix(ans, ",")
	ans += "]"
	return ans
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "[]" {
		return nil
	}
	data = strings.Trim(data, "[]")
	arr := strings.Split(data, ",")
	val, _ := strconv.Atoi(arr[0])
	root := &TreeNode{Val: val}
	i := 1
	q := []*TreeNode{root}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if arr[i] != "null" {
			val, _ := strconv.Atoi(arr[i])
			node.Left = &TreeNode{Val: val}
			q = append(q, node.Left)
		}
		i++
		if arr[i] != "null" {
			val, _ := strconv.Atoi(arr[i])
			node.Right = &TreeNode{Val: val}
			q = append(q, node.Right)
		}
		i++
	}
	return root
}
