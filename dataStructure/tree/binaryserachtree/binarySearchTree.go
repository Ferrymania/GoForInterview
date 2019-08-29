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
//递归法增加元素
func(bst *BST) Add(v int){
	if bst.root ==nil  {
		bst.root = CreateNode(v)
		bst.size++
	}else{
		bst.add(bst.root,v)
	}
}
//向以n为根结点的二分搜索树中插入元素v
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
		return
	}
	if v < n.Value {
		bst.add(n.left,v)
	}else {
		bst.add(n.right,v)
	}
}
//判断二叉搜索树是否存在元素v
func(bst *BST) contains(v int)bool {
	return containRecursively(bst.root,v)
}

func containRecursively(node *Node,v int)bool{
	if node == nil {
		return false
	}
	if node.Value == v {
		return true
	}else if v < node.Value{
		return containRecursively(node.left,v)
	}else {
		return containRecursively(node.right,v)
	}
}

//寻找二分搜索树的最小元素
func(bst *BST)Min() int{
	if bst.size == 0{
		panic("binary serach tree is empty!")
	}
	return minmum(bst.root).Value
}

func minmum(node *Node)*Node{
	if node.left == nil {
		return node
	}
	return minmum(node.left)
}

//寻找二分搜索树的最大元素
func(bst *BST)Max() int {
	if bst.size == 0 {
		panic("binary is empty!")
	}
	return maximum(bst.root).Value
}

func maximum(node *Node)*Node{
	if node.right == nil{
		return node
	}
	return maximum(node.right)
}

//删除二分搜索树删除最小值所在的结点，返回最小值
func(bst *BST)RemoveMin() int{
	min := bst.Min()
	bst.root = bst.removeMin(bst.root)
	return min
}
func(bst *BST)removeMin(node *Node) *Node{
	if node.left == nil {
		rightNode := node.right
		node.right = nil
		bst.size--
		return rightNode
	}
	node.left = bst.removeMin(node.left)
	return node
}

//删除二分搜索树最大值的结点
func(bst *BST)RemoveMax() int{
	max :=bst.Max()
	bst.root = bst.removeMax(bst.root)
	return max
}
func(bst *BST)removeMax(node *Node)*Node{
	if node.right == nil {
		leftNode := node.left
		node.left = nil
		bst.size--
		return leftNode
	}
	node.right = bst.removeMax(node.right)
	return node
}