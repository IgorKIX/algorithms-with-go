package main

import "fmt"

func main() {
	p := New(5)

	fmt.Println(p)
}

// the idea of trees stay the same, but now We will collect the data about the high (weight) of tree.
// Thanks to this info we will be ble to connect smaller tree to the root of the bigger one, and avoid the to tall trees
// Cost of algo -> deph of tree at most log(N)
// + Path compression - During the search the code will teach the numbers in the branch what is the main root of this branch.
// Thanks to this we will be able to create almost flat trees.
type WQuickUnion struct {
	id   []int
	size []int
}

func New(number int) *WQuickUnion {
	obj := new(WQuickUnion)

	obj.id = make([]int, number)
	obj.size = make([]int, number)

	for i := 0; i < len(obj.id); i++ {
		obj.id[i] = i
		obj.size[i] = 1
	}

	return obj
}

// Find the root element of the given value
func (qu *WQuickUnion) findRoot(i int) int {
	for i != qu.id[i] {
		qu.id[i] = qu.id[qu.id[i]] // path commpression implementation
		i = qu.id[i]
	}
	return i
}

func (qu *WQuickUnion) connected(p, q int) bool {
	return qu.findRoot(p) == qu.findRoot(q)
}

func (qu *WQuickUnion) union(p, q int) {
	pRoot := qu.findRoot(p)
	qRoot := qu.findRoot(q)

	if pRoot == qRoot {
		return
	}

	// Find which tree is smaller and connect the smaller one to the root of bigger one
	if qu.size[pRoot] < qu.size[qRoot] {
		qu.id[pRoot] = qRoot
		qu.size[qRoot] += qu.size[pRoot]
	} else {
		qu.id[qRoot] = pRoot
		qu.size[pRoot] += qu.size[qRoot]
	}
}
