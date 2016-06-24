// 符合6-1/6-2/6-3/6-4/6-5需求
// 支持32/64位word字长

// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 165.

// Package intset provides a set of integers based on a bit vector.
package intset

import (
	"bytes"
	"fmt"
)

//!+intset
// pc is helper table for popCount function
// pc表格用于记录每个8bit宽度的数字含二进制的1bit的bit个数
//在处理64bit宽度的数字时就没有必要循环64次，只需要8次查表就可以了
var pc [256]byte

// 字的宽度，在64位平台是64位，在32位平台是32位。为了满足6-5的需求定义，
// 下面=号右边是平台的自动判断的一个智能表达式：32 << (^uint(0) >> 63)
const WORDSIZE = 32 << (^uint(0) >> 63)

func init() {
	// pc[i] is the population count of i.
	pc = func() (pc [256]byte) {
		for i := range pc {
			pc[i] = pc[i/2] + byte(i&1)
		}
		return
	}()
	_ = pc[0]

}


// popCount returns the population count (number of set bits) of x.
func popCount(x uint) int {
	if WORDSIZE == 64 {
		return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
	} else {
		return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))])
	}

}

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint
}


// requirment of 6-1 begin
// return the number of elements
func (s *IntSet) Len() int {
	total := 0
	for _, tword := range s.words {
		total += popCount(tword)
	}
	return total
}

// remove x from the set
func (s *IntSet) Remove(x int) {
	word, bit := x/WORDSIZE, uint(x%WORDSIZE)
	if word < len(s.words) {

		s.words[word] &= ^(1 << bit)
	}
}

// remove all elements from the set
func (s *IntSet) Clear() {
	for i, _ := range s.words {
		s.words[i] = 0
	}
}

// return a copy of the set
func (s *IntSet) Copy() *IntSet {
	if s != nil {
		t := &IntSet{words: make([]uint, len(s.words))}
		copy(t.words, s.words)
		return t
	} else {
		return nil
	}
}
// requirment of 6-1 end

// requirment of 6-2 begin
func (s *IntSet) AddAll(in ...int) {
	if len(in) == 0 {
		return
	} else {
		for _, v := range in {
			s.Add(v)
		}
	}
}
// requirment of 6-2 end

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/WORDSIZE, uint(x%WORDSIZE)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/WORDSIZE, uint(x%WORDSIZE)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}
// requirment of 6-3 begin
// IntersectWith sets s to the union of s and t.
func (s *IntSet) IntersectWith(t *IntSet) {
	c := make(map[int] int)
	for _, v := range s.Elems() {
		c[v]++
	}
	for _, v := range t.Elems() {
		c[v]++
	}
	if len(c) > 0 {
		s.Clear()
		for k, v := range c {
			if v > 1 {
				s.Add(k)
			}
		}
	}
}

// DifferenceWith sets s to the union of s and t.
func (s *IntSet) DifferenceWith(t *IntSet) {
	c := make(map[int] int)
	for _, v := range s.Elems() {
		c[v]++
	}
	for _, v := range t.Elems() {
		if _, ok := c[v]; ok {
			delete(c, v)
		}
	}

	s.Clear()
	for k, _ := range c {
		s.Add(k)
	}
}

// SymmetricDifference sets s to the union of s and t.
func (s *IntSet) SymmetricDifference(t *IntSet) {
	d := s.Copy()
	d.DifferenceWith(t)
	t.DifferenceWith(s)
	s.Clear()
	s.UnionWith(d)
	s.UnionWith(t)
}
// requirment of 6-3

// requirment of 6-4 end
func (s *IntSet) Elems() []int {
	r := make([]int, 0)
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < WORDSIZE; j++ {
			if word&(1<<uint(j)) != 0 {
				r = append(r, WORDSIZE*i+j)
			}
		}
	}
	return r
}
// requirment of 6-4 end
//!-intset

//!+string

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < WORDSIZE; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", WORDSIZE*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

//!-string

