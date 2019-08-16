package binaryserachtree

type Node struct {
	Value int
	left *Node
	right *Node
}

func CreateNode(v int) *Node{
	return &Node{Value:v}
}

type BST struct {
	root *Node
	size int
}

func CreateBST() *BST {
	return &BST{}
}

func(bst *BST) GetSize() int {
	return bst.size
}

func(bst *BST) IsEmpty() bool {
	return bst.size == 0
}

func(bst *BST) Add(v int){
	if bst.root ==nil  {
		bst.root = CreateNode(v)
		bst.size++
	}else{
		bst.add(bst.root,v)
	}
}

func(bst *BST)add(n *Node,v int){
	if v == n.Value {
		return
	}else if v < n.Value && n.left == nil {
		n.left = CreateNode(v)
		bst.size++
		return
	}else if v >n.Value && n.right == nil {
		n.right  = CreateNode(v)
		bst.size++
	}

	if v < n.Value {
		bst.add(n.left,v)
	}else {
		bst.add(n.right,v)
	}
}