package main

// Node
type Node struct {

	// the connected document id
	doc string

	// indexes where the term exists in the document
	indexes []int

	// next points to the next node
	next *Node
}

// PostingList
type PostingList struct {

	// n is the total number of nodes in the list
	n int32

	// idf refers to inverse document frequency (tf-idf)
	idf float64

	// hNode is the head/first node in the list
	hNode *Node

	// tNode is the tail/last node in the list
	tNode *Node
}

// Size returns the number of nodes in the list
func (l PostingList) Size() int32 {
	return l.n
}

// Insert adds a new node to the linked list
func (l *PostingList) Insert(n *Node) {

	l.n++

	if l.tNode != nil {
		l.tNode.next = n
		l.tNode = n
	}

	if l.hNode == nil {
		l.hNode = n
		l.tNode = n
	}
}
