package algos_test

import (
	. "algorithms"
	"fmt"
	"math/rand"
	"testing"
)

type Int struct {
	DataType
	Val int64
}

func (i Int) Equals(other DataType) bool {
	return i.Val == other.(Int).Val
}

func (i Int) Less(other DataType) bool {
	return i.Val < other.(Int).Val
}

func (i Int) More(other DataType) bool {
	return i.Val > other.(Int).Val
}

func (i Int) Divide(other DataType) bool {
	return i.Val/other.(Int).Val == 0
}

func Work(d DataType) {
	fmt.Printf(" %v", d.(Int).Val)
}

func TestPlayWithBinarySearchTree(t *testing.T) {
	fmt.Println("\n######")
	var bst = &Bst[Int]{Data: Int{Val: 1}, Left: nil, Parent: nil, Right: nil}
	bst.Insert(bst, Int{Val: 5})
	bst.Insert(bst, Int{Val: 14})
	bst.Insert(bst, Int{Val: 56})
	bst.Insert(bst, Int{Val: 9})
	bst.Insert(bst, Int{Val: 10})
	bst.Insert(bst, Int{Val: 2})
	bst.Insert(bst, Int{Val: 4})
	bst.Insert(bst, Int{Val: -2})
	bst.Traversal(bst, Work)
	printMaxMin(bst)
}

func printMaxMin(bst *Bst[Int]) {
	fmt.Println("\n######")
	b, _ := bst.Min(bst)
	if nil != b {
		fmt.Printf("Min element of tree is %v\n", b.Data.Val)
	}
	b, _ = bst.Max(bst)
	if nil != b {
		fmt.Printf("Max element of tree is %v\n", b.Data.Val)
	}
}

func TestLargeNumberOfElementsInTree(t *testing.T) {
	var elementsCount int64 = 1000000
	var count int64 = 0

	var bst = &Bst[Int]{Data: Int{Val: 1}, Left: nil, Parent: nil, Right: nil}

	for count < elementsCount {
		count++
		bst.Insert(bst, Int{Val: int64(rand.Intn(1000000))})
	}
	printMaxMin(bst)
	searchForRandomValues(count, bst)

}

func printElementCountInTree(bst *Bst[Int]) int64 {
	var eleCount int64 = 0
	bst.Traversal(bst, func(d DataType) {
		eleCount++
	})
	fmt.Printf("%v Elements are there in tree\n", eleCount)
	return eleCount
}

func searchForRandomValues(count int64, bst *Bst[Int]) {
	var latestSize int64 = 0
	var lastKnownSize int64 = 0
	var sc int64 = 100
	count = 0
	for count < sc {
		count++
		e := int64(rand.Intn(1000000))
		s := bst.Search(bst, Int{Val: e})
		if nil != s {
			lastKnownSize = latestSize
			fmt.Printf("Found element %v in tree\n", s.Data.Val)
			bst.Delete(bst, Int{Val: e})
			latestSize = printElementCountInTree(bst)
			if lastKnownSize-latestSize > 1 {
				fmt.Printf("#### Here is the bug\n")
			}
		} else {
			fmt.Printf("Element %v not found in tree\n", e)
		}
	}
}
