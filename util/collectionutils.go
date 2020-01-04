package util

import (
	"github.com/sunmeng90/leetcode-go/model"
	"strconv"
)

func IntSliceToString(data []int) string {
	result := ""
	if len(data) == 0 {
		return result
	}
	for i := 0; i < len(data); i++ {
		result += strconv.Itoa(data[i]) + ","
	}
	return result[:len(result)-1]
}

func SliceToSinglyLinkedList(data []int) *model.ListNode {
	preHead := &model.ListNode{}
	cur := preHead
	for _, v := range data {
		cur.Next = &model.ListNode{Val: v}
		cur = cur.Next
	}
	return preHead.Next
}
