package main

import (
	"fmt"
	"sort"
)

// 待排序的元素
type CustomItem struct {
	name     string   // user name
	age   	 int      // user age
	level 	 int      // user level
}
type lessFunc func(p1, p2 *CustomItem) bool // less的封装，方便对自定义结构所有字段进行比较

// 针对多字段排序的类型封装，又此类型来实现sort.Interface的三个接口方法
type MutiSorter struct {
	items	    []CustomItem
	less        []lessFunc
}

// Sort sorts the argument slice according to the less functions passed to OrderedBy.
func (ms *MutiSorter) Sort(its []CustomItem) {
	ms.items = its
	sort.Sort(ms)
}

// OrderedBy returns a Sorter that sorts using the less functions, in order.
// Call its Sort method to sort the data.
func OrderedBy(less ...lessFunc) *MutiSorter {
	return &MutiSorter{
		less: less,
	}
}

// 实现sort.Interface需要的接口
// A type, typically a collection, that satisfies sort.Interface can be
// sorted by the routines in this package.  The methods require that the
// elements of the collection be enumerated by an integer index.
/*
type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}
*/
func (ms *MutiSorter) Len() int {
	return len(ms.items)
}

func (ms *MutiSorter) Swap(i, j int) {
	ms.items[i], ms.items[j] = ms.items[j], ms.items[i]
}

func (ms *MutiSorter) Less(i, j int) bool {
	p, q := &ms.items[i], &ms.items[j]
	var k int
	for k = 0; k < len(ms.less) - 1; k++ {
		f := ms.less[k]
		switch {
		case f(p, q): // return value is true, p < q
			return true
		case f(q, p): // p > q
			return false
		default:
			// means p equal to q in this field, compare that by next field
		}
	}
	// All comparisons to here said "equal", so just return whatever
	// the final comparison reports.
	return ms.less[k](p, q)
}

var testItems = []CustomItem{
	{"gri", 23, 100},
	{"ken", 23, 150},
	{"glenda", 23, 200},
	{"rsc", 28, 200},
	{"r", 26, 100},
	{"ken", 28, 200},
	{"dmr", 28, 100},
	{"r", 26, 150},
	{"gri", 26, 80},
}

func main() {
	// define less compare function value
	name := func(p1, p2 *CustomItem) bool {
		return p1.name < p2.name
	}
	age := func(p1, p2 *CustomItem) bool {
		return p1.age < p2.age
	}

	level := func(p1, p2 *CustomItem) bool {
		return p1.level < p2.level
	}

	OrderedBy(name, age).Sort(testItems)
	fmt.Println("Sort by name & age:")
	fmt.Println(testItems)
	OrderedBy(name, level).Sort(testItems)
	fmt.Println("Sort by name & level:")
	fmt.Println(testItems)
}
