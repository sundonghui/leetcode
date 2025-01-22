package copyrandomlist

import "testing"

func TestCopyRandomList(t *testing.T) {
	// Create a linked list with random pointers
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node1.Next = node2
	node2.Next = node3
	node1.Random = node3
	node2.Random = node1
	node3.Random = node2

	// Copy the linked list
	copyNode1 := copyRandomList(node1)

	// Check if the copied linked list is correct
	if copyNode1.Val != 1 || copyNode1.Next.Val != 2 || copyNode1.Random.Val != 3 {
		t.Errorf("CopyRandomList failed")
	}
}
