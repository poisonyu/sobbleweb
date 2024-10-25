// 6.5 示例：Bit数组
// 用bit数组来表示集合
// todo 可以给这个包写一个单元测试，来验证方法实现的准确性
package main

import (
	"bytes"
	"fmt"
)

// An IntSet is a set of small non-negative integers
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) AddAll(g ...int) {
	for _, x := range g {
		s.Add(x)
	}
}
func (s *IntSet) Len() int {
	var count int
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				count++
			}
		}
	}
	return count
}
func (s *IntSet) Remove(x int) {
	if s.Has(x) {
		word, bit := x/64, x%64
		s.words[word] &^= 1 << bit
	}

}

func (s *IntSet) Clear() {
	// s = new(IntSet)
	s.words = make([]uint64, 0)
}

func (s *IntSet) Copy() *IntSet {
	// c := new(IntSet)
	var c IntSet
	c.words = append(c.words, s.words...)
	return &c
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

func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		}
	}
	if len(s.words) > len(t.words) {
		s.words = s.words[:len(t.words)]
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		}
	}
}

func (s *IntSet) SymmetricDifference(t *IntSet) {
	s.DifferenceWith(t)
	t.DifferenceWith(s)
	s.UnionWith(t)
}
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Elems() (setSlice []int) {
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				setSlice = append(setSlice, 64*i+j)
			}
		}
	}
	return
}

func main() {
	bitSet := new(IntSet)
	fmt.Println("bitSet len:", len(bitSet.words))
	bitSet.Add(1)
	bitSet.Add(9)
	bitSet.Add(141)
	fmt.Println(bitSet)
	bitSet.Remove(9)
	fmt.Println(bitSet)
	// bitSet.Clear()
	// fmt.Println(bitSet)
	bitSet2 := bitSet.Copy()
	fmt.Println(bitSet2)
	bitSet2.Add(111)
	bitSet2.AddAll(3, 22, 34)
	fmt.Println()
	bitSet.AddAll(3, 77, 777)
	fmt.Println(bitSet, "\n", bitSet2)
	bitSet.IntersectWith(bitSet2)
	fmt.Println(bitSet)
}
