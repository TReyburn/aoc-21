package day1

type Node struct {
	value int
	prev  *Node
	next  *Node
}

func (n *Node) Value() int {
	return n.value
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (n *Node) Next() *Node {
	return n.next
}

func DoublyLinkedList(items []int) (start *Node, end *Node) {
	var curr, prev *Node

	start = &Node{
		value: items[0],
		prev:  nil,
		next:  nil,
	}

	prev = start

	for _, item := range items[1:] {
		curr = &Node{
			value: item,
			prev:  prev,
			next:  nil,
		}

		prev.next = curr

		prev = curr
	}
	end = curr

	return start, end
}
