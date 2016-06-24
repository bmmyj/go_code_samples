// 符合6-1/6-2/6-3/6-4/6-5需求

package main

import (
	"fmt"
	. "./intSet" // can ingore package prefix
)

func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"
	fmt.Println("Len of x: ", x.Len())

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"
	fmt.Println("Len of y: ", y.Len())

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println("Len of x|y : ", x.Len())
	fmt.Println(x.Has(9), x.Has(123)) // "true false"

	z := x.Copy()
	if z != nil {
		fmt.Println("z(copy of x) : ", z.String()) // z(copy of x) :  {1 9 42 144}
	} else {
		fmt.Println("error, x.Copy return nil !!!")
	}
	//fmt.Printf("addr of x is %p;  addr of z is %p \n", &x, &z)

	x.Remove(9)
	fmt.Println("after remove value 9, Len of x is: ", x.Len()) // after remove value 9, Len of x is:  3
	fmt.Println("Values of x is: ", x.String()) // Values of x is:  {1 42 144}

	x.Clear()
	fmt.Println("Len of x(after Clear): ", x.Len()) // Len of x(after Clear):  0

	x.AddAll(1,2,3)
	fmt.Println("after call x.AddAll(1,2,3) : ", x.String()) // after call x.AddAll(1,2,3) :  {1 2 3}

	fmt.Println("Print all Elements of set x:")
	for i, v := range x.Elems() {
		fmt.Printf("index = %d; value = %d\n", i, v)
	}
	// Print all Elements of set x:
	// index = 0; value = 1
	// index = 1; value = 2
	// index = 2; value = 3


	y.Clear()
	y.AddAll(3,194,2,5)

	fmt.Println("set x is: ", x.String()) // set x is:  {1 2 3}
	fmt.Println("set y is: ", y.String()) // set y is:  {2 3 5 194}

	x.IntersectWith(&y)

	fmt.Println("Intersect for x and y is: ", x.String()) // Intersect for x and y is:  {2 3}

	x.Clear()
	x.AddAll(1,2,3,4)
	fmt.Println("set x is: ", x.String()) // set x is:  {1 2 3 4}
	fmt.Println("set y is: ", y.String()) // set y is:  {2 3 5 194}
	x.DifferenceWith(&y)
	fmt.Println("x.DifferenceWith(&y) is: ", x.String()) // x.DifferenceWith(&y) is:  {1 4}

	x.Clear()
	x.AddAll(1,2,3,4)
	fmt.Println("set x is: ", x.String()) // set x is:  {1 2 3 4}
	fmt.Println("set y is: ", y.String()) // set y is:  {2 3 5 194}
	x.SymmetricDifference(&y)
	fmt.Println("x.SymmetricDifference(&y) is: ", x.String()) // x.SymmetricDifference(&y) is:  {1 4 5 194}

	fmt.Println(WORDSIZE)
}
