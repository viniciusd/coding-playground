package main

import "fmt"

// Create a new type of 'deck',
// which is a slice of strings
type seq []int

func newSeq(start, end, step int) seq {
	s := seq{}
	val := start

	if step > 0 {
		for val < end {
			s = append(s, val)
			val += step
		}
	} else if step < 0 {
		for val > end {
			s = append(s, val)
			val += step
		}
	}
	return s
}

func (s seq) print() {
	for _, val := range s {
		fmt.Println(val)
	}
}

func (s seq) printParity() {
	for _, val := range s {
		if val%2 == 0 {
			fmt.Println(val, "is even")
		} else {
			fmt.Println(val, "is odd")
		}
	}
}
