package linklist

import (
	"strings"
)

func linklistFromSlice(values []string) *Node {
	sentinel := &Node{Next: nil}
	cursor := sentinel
	for _, value := range values {
		cursor.Next = &Node{Value: value, Next: nil}
		cursor = cursor.Next
	}
	return sentinel.Next
}

func compareLinklist(a, b *Node) bool {
	for a != nil && b != nil {
		if a.Value != b.Value {
			return false
		}
		a = a.Next
		b = b.Next
	}
	return a == nil && b == nil
}

func printLinklist(head *Node) string {
	sb := strings.Builder{}
	for head != nil {
		sb.WriteString(head.Value + " -> ")
		head = head.Next
	}
	sb.WriteString("nil")
	return sb.String()
}
