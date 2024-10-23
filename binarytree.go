package algos

import "errors"

type DataType interface {
	Equals(other DataType) bool
	Less(other DataType) bool
	More(other DataType) bool
	Divide(other DataType) bool
}

type Bst[t DataType] struct {
	Data   t
	Left   *Bst[t]
	Right  *Bst[t]
	Parent *Bst[t]
}

func (ty Bst[t]) Search(tree *Bst[t], key t) (subtree *Bst[t]) {
	if nil == tree {
		return nil
	}
	if key.Equals(tree.Data) {
		return tree
	} else if key.Less(tree.Data) {
		return ty.Search(tree.Left, key)
	} else if key.More(tree.Data) {
		return ty.Search(tree.Right, key)
	}
	return nil
}

func (ty Bst[t]) Traversal(tree *Bst[t], apply func(d DataType)) {
	if nil != tree {
		ty.Traversal(tree.Left, apply)
		apply(tree.Data)
		ty.Traversal(tree.Right, apply)
	}
}

func (ty Bst[T]) Min(tree *Bst[T]) (*Bst[T], error) {
	if nil == tree {
		return nil, errors.New(`Tree is nil`)
	}
	m := tree
	for m.Left != nil {
		m = m.Left
	}
	return m, nil
}

func (ty Bst[T]) Max(tree *Bst[T]) (*Bst[T], error) {
	if nil == tree {
		return nil, errors.New(`Tree is nil`)
	}
	m := tree
	for m.Right != nil {
		m = m.Right
	}
	return m, nil
}

func (ty Bst[T]) Insert(tree *Bst[T], data T) (tr *Bst[T], err error) {
	if nil == tree {
		return &Bst[T]{Data: data, Left: nil, Right: nil, Parent: nil}, nil
	}

	currNode := tree
	for currNode != nil {
		if currNode.Data.Less(data) {
			if currNode.Right != nil {
				currNode = currNode.Right
			} else {
				n := &Bst[T]{Data: data, Left: nil, Right: nil, Parent: currNode}
				currNode.Right = n
				return tree, nil
			}
		} else if currNode.Data.More(data) {
			if currNode.Left != nil {
				currNode = currNode.Left
			} else {
				n := &Bst[T]{Data: data, Left: nil, Right: nil, Parent: currNode}
				currNode.Left = n
				return tree, nil
			}
		} else {
			return nil, errors.New(`Duplicate not allowed`)
		}
	}
	return tree, nil
}

func (ty Bst[T]) Delete(tree *Bst[T], data T) (*Bst[T], error) {
	if nil == tree {
		return nil, errors.New(`tree is nil`)
	}

	node := ty.Search(tree, data)
	if nil != node {

		nodeIsLeftChildOfItsParent := false
		if node.Parent != nil && node.Parent.Left == node {
			nodeIsLeftChildOfItsParent = true
		}

		// Deleting leaf node
		if node.Left == nil && node.Right == nil {
			if nodeIsLeftChildOfItsParent {
				node.Parent.Left = nil
			} else {
				node.Parent.Right = nil
			}
			node.Parent = nil
			return node, nil
		}
		// Deleting a node which has only left child tree
		if node.Left != nil && node.Right == nil {
			// This node is left child of its parent
			if nodeIsLeftChildOfItsParent {
				node.Parent.Left = node.Left
				node.Left.Parent = node.Parent
			} else {
				node.Parent.Right = node.Left
				node.Left.Parent = node.Parent
			}
			node.Left = nil
			node.Parent = nil
			return node, nil
		}

		//Deleting a node which has only right child tree
		if node.Left == nil && node.Right != nil {
			if nodeIsLeftChildOfItsParent {
				node.Parent.Left = node.Right
				node.Right.Parent = node.Parent
			} else {
				node.Parent.Right = node.Right
				node.Right.Parent = node.Parent
			}
			node.Right = nil
			node.Parent = nil
			return node, nil
		}

		//Deleting a node which has both children
		if node.Left != nil && node.Right != nil {
			// Let's make left child as root of new subtree
			var n = node.Left
			if nodeIsLeftChildOfItsParent {
				node.Parent.Left = n
			} else {
				node.Parent.Right = n
			}
			n.Parent = node.Parent
			if n.Right != nil {
				// take right children of lN and insert it at min position in right subtree of node
				leftMostNodeInRightChildrenOfNode, _ := ty.Min(node.Right)
				leftMostNodeInRightChildrenOfNode.Left = n.Right
				n.Right.Parent = leftMostNodeInRightChildrenOfNode
				n.Right = nil
			}
			node.Parent = nil
			node.Left = nil
			node.Right = nil
			return node, nil
		}
	}
	return nil, errors.New(`No node found to delete`)
}