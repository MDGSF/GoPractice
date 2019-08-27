package main

import (
	"testing"
)

func SliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for k := range a {
		if a[k] != b[k] {
			return false
		}
	}
	return true
}

func SliceElemEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	ma := make(map[int]int)
	for k := range a {
		if _, ok := ma[k]; ok {
			ma[k]++
		} else {
			ma[k] = 1
		}
	}

	for k := range b {
		if _, ok := ma[k]; ok {
			if ma[k] > 1 {
				ma[k]--
			} else {
				delete(ma, k)
			}
		} else {
			return false
		}
	}

	return len(ma) == 0
}

func TestPartition1(t *testing.T) {
	a := []int{2, 1, 0, 4, 5, 6, 3}
	i := partition(a)
	if i != 3 {
		t.Fatal(i)
	}
	expectResult := []int{2, 1, 0, 3, 5, 6, 4}
	if !SliceEqual(a, expectResult) {
		t.Fatal(a)
	}
}

func TestPartition2(t *testing.T) {
	a := []int{2, 1}
	i := partition(a)
	if i != 0 {
		t.Fatal(i)
	}
	expectResult := []int{1, 2}
	if !SliceEqual(a, expectResult) {
		t.Fatal(a)
	}
}

func TestPartition3(t *testing.T) {
	a := []int{1, 2}
	i := partition(a)
	if i != 1 {
		t.Fatal(i)
	}
	expectResult := []int{1, 2}
	if !SliceEqual(a, expectResult) {
		t.Fatal(a)
	}
}

func TestPartition4(t *testing.T) {
	a := []int{1}
	i := partition(a)
	if i != 0 {
		t.Fatal(i)
	}
}

func TestPartition5(t *testing.T) {
	a := []int{}
	i := partition(a)
	if i != 0 {
		t.Fatal(i)
	}
}

func TestTopK1(t *testing.T) {
	a := []int{2, 4, 0, 1, 5, 6, 3}
	out := TopK(a, 1)
	expectResult := []int{0}
	if !SliceElemEqual(out, expectResult) {
		t.Fatal(a)
	}
}

func TestTopK2(t *testing.T) {
	a := []int{2, 4, 0, 1, 5, 6, 3}
	out := TopK(a, 2)
	expectResult := []int{0, 1}
	if !SliceElemEqual(out, expectResult) {
		t.Fatal(a)
	}
}

func TestTopK3(t *testing.T) {
	a := []int{2, 4, 0, 1, 5, 6, 3}
	out := TopK(a, 3)
	expectResult := []int{0, 1, 2}
	if !SliceElemEqual(out, expectResult) {
		t.Fatal(out)
	}
}

func TestTopK4(t *testing.T) {
	a := []int{2, 4, 0, 1, 5, 6, 3}
	out := TopK(a, 4)
	expectResult := []int{0, 1, 2, 3}
	if !SliceElemEqual(out, expectResult) {
		t.Fatal(out)
	}
}

func TestTopK5(t *testing.T) {
	a := []int{2, 4, 0, 1, 5, 6, 3}
	out := TopK(a, 5)
	expectResult := []int{0, 1, 2, 3, 4}
	if !SliceElemEqual(out, expectResult) {
		t.Fatal(out)
	}
}

func TestTopK6(t *testing.T) {
	a := []int{2, 4, 0, 1, 5, 6, 3}
	out := TopK(a, 6)
	expectResult := []int{0, 1, 2, 3, 4, 5}
	if !SliceElemEqual(out, expectResult) {
		t.Fatal(out)
	}
}

func TestTopK7(t *testing.T) {
	a := []int{2, 4, 0, 1, 5, 6, 3}
	out := TopK(a, 7)
	expectResult := []int{0, 1, 2, 3, 4, 5, 6}
	if !SliceElemEqual(out, expectResult) {
		t.Fatal(out)
	}
}

func TestTopK8(t *testing.T) {
	a := []int{2, 4, 0, 1, 5, 6, 3}
	out := TopK(a, 8)
	expectResult := []int{0, 1, 2, 3, 4, 5, 6}
	if !SliceElemEqual(out, expectResult) {
		t.Fatal(out)
	}
}

func TestTopK9(t *testing.T) {
	a := []int{2, 4, 0, 1, 5, 6, 3}
	out := TopK(a, -1)
	if out != nil {
		t.Fatal(out)
	}
}

func TestSort1(t *testing.T) {
	a := []int{2, 4, 0, 1, 5, 6, 3}
	sort(a)
	expectResult := []int{0, 1, 2, 3, 4, 5, 6}
	if !SliceEqual(a, expectResult) {
		t.Fatal(a)
	}
}

func TestSort2(t *testing.T) {
	a := []int{1}
	sort(a)
	expectResult := []int{1}
	if !SliceEqual(a, expectResult) {
		t.Fatal(a)
	}
}

func TestSort3(t *testing.T) {
	a := []int{1, 3, 1, 2}
	sort(a)
	expectResult := []int{1, 1, 2, 3}
	if !SliceEqual(a, expectResult) {
		t.Fatal(a)
	}
}

func TestSort4(t *testing.T) {
	a := []int{1, 3, 1, 2, 4}
	sort(a)
	expectResult := []int{1, 1, 2, 3, 4}
	if !SliceEqual(a, expectResult) {
		t.Fatal(a)
	}
}
