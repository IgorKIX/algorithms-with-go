package main

import "fmt"

func main() {
	p := New(5)

	fmt.Println(p)
}

// At beggining every index is a root of a tree. Connect method moves chenge the roots elements of given box -> change the root of this element
// So when we merge to elements, we are trully merging the roots of this elements. union(1, 2) union(2, 3) will give a tree
// with a root of 3 and the highes element of this tree will have value 2 (3-1-2)
type QuickUnion struct {
	id []int
}

func New(number int) *QuickUnion {
	obj := new(QuickUnion)

	obj.id = make([]int, number)

	for i := 0; i < len(obj.id); i++ {
		obj.id[i] = i
	}

	return obj
}

// Find the root element of the given value
func (qu *QuickUnion) findRoot(i int) int {
	for i != qu.id[i] {
		i = qu.id[i]
	}
	return i
}

func (qu *QuickUnion) connected(p, q int) bool {
	return qu.findRoot(p) == qu.findRoot(q)
}

func (qu *QuickUnion) union(p, q int) {
	pRoot := qu.findRoot(p)
	qRoot := qu.findRoot(q)

	qu.id[pRoot] = qRoot
}
