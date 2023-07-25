package main

func main() {

}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	nodes := make(map[*Node]*Node)
	curr := head

	for curr != nil {
		clone := &Node{Val: curr.Val}
		nodes[curr] = clone

		curr = curr.Next
	}

	curr = head
	for curr != nil {
		clone := nodes[curr]
		clone.Next = nodes[curr.Next]
		clone.Random = nodes[curr.Random]

		curr = curr.Next
	}

	return nodes[head]
}