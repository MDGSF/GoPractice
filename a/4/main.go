/*
https://www.geeksforgeeks.org/manachers-algorithm-linear-time-longest-palindromic-substring-part-1/
*/
package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	test("a")
	test("aba")
	test("abaaba")
}

func test(s string) {
	fmt.Printf("longestPalindrome(%v) = %v\n", s, longestPalindrome(s))
}

func longestPalindrome(s string) string {
	N := len(s)
	if N == 0 {
		return ""
	} else if N == 1 {
		return s
	}

	N = 2*N + 1
	L := make([]int, N)
	L[0] = 0
	L[1] = 1
	centerPosition := 1
	centerRight := 2

	var maxLPSLength int
	var maxLPSCenterPosition int

	for i := 2; i < N; i++ {
		needExpand := false
		curRight := i
		curLeft := 2*centerPosition - curRight

		diff := centerRight - curRight
		if diff > 0 {
			if L[curLeft] < diff {
				L[curRight] = L[curLeft]
			} else if L[curLeft] == diff && curRight == N-1 {
				L[curRight] = L[curLeft]
			} else if L[curLeft] == diff && curRight < N-1 {
				L[curRight] = L[curLeft]
				needExpand = true
			} else if L[curLeft] > diff {
				L[curRight] = diff
				needExpand = true
			}
		} else {
			L[curRight] = 0
			needExpand = true
		}

		//fmt.Printf("s=%v, len(s)=%v, N=%v L=%v, center=%v, centerR=%v\n",
		//	s, len(s), N, L, centerPosition, centerRight)
		//fmt.Printf("curL=%v, curR=%v, diff=%v, needExpand=%v\n",
		//	curLeft, curRight, diff, needExpand)

		if needExpand {
			for ((i+L[i]) < N && (i-L[i]) > 0) &&
				(((i+L[i]+1)%2 == 0) ||
					((((i + L[i] + 1) / 2) < len(s)) && s[(i+L[i]+1)/2] == s[(i-L[i]-1)/2])) {
				L[i]++
			}
		}

		if L[curRight] > maxLPSLength {
			maxLPSLength = L[curRight]
			maxLPSCenterPosition = curRight
		}

		if curRight+L[curRight] >= centerRight {
			centerPosition = curRight
			centerRight = curRight + L[curRight]
		}

		//fmt.Println()
	}

	start := (maxLPSCenterPosition - maxLPSLength) / 2
	end := start + maxLPSLength
	fmt.Printf("maxLPSLength = %v, maxLPSCenterPosition = %v, start = %v, end = %v\n",
		maxLPSLength, maxLPSCenterPosition, start, end)
	return s[start:end]
}
