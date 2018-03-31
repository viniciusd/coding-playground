package main

import "testing"

func TestNewSeq_whenAskedForAIncreasingSequence(t *testing.T) {
	start := 0
	stop := 10
	step := 1
	s := newSeq(start, stop, step)

	if len(s) != 10 {
		t.Errorf("Expected sequence length of 10, but got %d", len(s))
	}

	for i, _ := range s {
		if i == 0 {
			if s[i] != start {
				t.Errorf("Expected %d-th value of %d, but got %d", i, start, s[0])
			}
		} else {
			if s[i]-s[i-1] != step {
				t.Errorf("Expected a difference of %d between consecutive values, but got %d", step, s[i]-s[i-1])
			}
		}
	}
	if s[len(s)-1] == stop {
		t.Errorf("Expected stop value of %d not to be included in the sequence", stop)
	}
}

func TestNewSeq_whenAskedForADecreasingSequence(t *testing.T) {
	start := 10
	stop := 0
	step := -1
	s := newSeq(start, stop, step)

	if len(s) != 10 {
		t.Errorf("Expected sequence length of 10, but got %d", len(s))
	}

	for i, _ := range s {
		if i == 0 {
			if s[i] != start {
				t.Errorf("Expected %d-th value of %d, but got %d", i, start, s[0])
			}
		} else {
			if s[i]-s[i-1] != step {
				t.Errorf("Expected a difference of %d between consecutive values, but got %d", step, s[i]-s[i-1])
			}
		}
	}
	if s[len(s)-1] == stop {
		t.Errorf("Expected stop value of %d not to be included in the sequence", stop)
	}
}

func TestNewSeq_whenGivenAStepOfZero(t *testing.T) {
	start := 0
	stop := 10
	step := 0
	s := newSeq(start, stop, step)

	if len(s) != 0 {
		t.Errorf("Expected sequence length of 0, but got %d", len(s))
	}
}
