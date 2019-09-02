package main

import "testing"

func Test1(t *testing.T) {
	result := st([]int{1})
	if result != 1 {
		t.Fatal(result)
	}
}

func Test2(t *testing.T) {
	result := st([]int{1, 2})
	if result != 2 {
		t.Fatal(result)
	}
}

func Test3(t *testing.T) {
	result := st([]int{1, 2, 3})
	if result != 4 {
		t.Fatal(result)
	}
}

func Test4(t *testing.T) {
	result := st([]int{1, 2, 3, 4})
	if result != 6 {
		t.Fatal(result)
	}
}

func Test5(t *testing.T) {
	result := st([]int{1, 2, 3, 4, 5, 6})
	if result != 12 {
		t.Fatal(result)
	}
}

func Test6(t *testing.T) {
	result := st2([]int{1})
	if result != 1 {
		t.Fatal(result)
	}
}

func Test7(t *testing.T) {
	result := st2([]int{1, 2})
	if result != 2 {
		t.Fatal(result)
	}
}

func Test8(t *testing.T) {
	result := st2([]int{1, 2, 3})
	if result != 4 {
		t.Fatal(result)
	}
}

func Test9(t *testing.T) {
	result := st2([]int{1, 2, 3, 4})
	if result != 6 {
		t.Fatal(result)
	}
}

func Test10(t *testing.T) {
	result := st2([]int{1, 2, 3, 4, 5, 6})
	if result != 12 {
		t.Fatal(result)
	}
}
