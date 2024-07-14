package linklist

type Node struct {
	Value string
	Next  *Node
}

func reverse(head *Node) *Node {
	var prev, current, next *Node = nil, head, nil
	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}
