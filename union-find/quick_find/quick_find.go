package main

import "fmt"

// Create an array. The index of the box will identify the number, the value under this index will inform about the index
// of the group to which this number belongs.

func main() {

	p := New(5)

	fmt.Println(p)

}

type QuickFind struct {
	id []int
}

func New(number int) *QuickFind {
	obj := new(QuickFind)

	obj.id = make([]int, number)

	// N array accesses
	for i := 0; i < len(obj.id); i++ {
		obj.id[i] = i
	}

	return obj
}

// 2 arra acceses
func (qf *QuickFind) connected(p, q int) bool {
	return qf.id[p] == qf.id[q]
}

func (qf *QuickFind) union(p, q int) {
	pId := qf.id[p]
	qId := qf.id[q]

	// At most 2N + 2 array accesses
	for i := 0; i < len(qf.id); i++ {
		if qf.id[i] == pId {
			qf.id[i] = qId
		}
	}
}
