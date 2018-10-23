package main

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

type Line struct {
	p1 Point
	p2 Point
}

func main() {
	allPoints := make([]Point, 0)
	var n int
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		var p Point
		fmt.Scanf("%d %d", &p.X, &p.Y)
		allPoints = append(allPoints, p)
	}

	fmt.Printf("%v", calcSquare(allPoints)/4)
}

func calcSquare(allPoints []Point) int {
	num := 0
	m := make(map[Line]struct{})
	for _, p1 := range allPoints {
		for _, p2 := range allPoints {
			l := Line{p1: p1, p2: p2}
			l1, l2 := calcLine(l)

			if _, ok := m[l1]; ok {
				num++
			}

			if _, ok := m[l2]; ok {
				num++
			}

			m[l] = struct{}{}
		}
	}
	return num
}

func calcLine(l Line) (l1, l2 Line) {
	x1 := l.p1.X
	y1 := l.p1.Y
	x2 := l.p2.X
	y2 := l.p2.Y

	l1.p1.X = x1 + (y2 - y1)
	l1.p1.Y = y1 - (x2 - x1)
	l1.p2.X = x2 + (y2 - y1)
	l1.p2.Y = y2 - (x2 - x1)

	l2.p1.X = x1 - (y2 - y1)
	l2.p1.Y = y1 + (x2 - x1)
	l2.p2.X = x2 - (y2 - y1)
	l2.p2.Y = y2 + (x2 - x1)
	return
}

/*
4
0 0
0 1
1 0
1 1

1
*/
