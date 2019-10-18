/*
https://www.geeksforgeeks.org/manachers-algorithm-linear-time-longest-palindromic-substring-part-1/

example 1:
s = "abababa"
s.length = 7

news = "|a|b|a|b|a|b|a|"
news.length = 2 * s.length + 1 = 15

string:         |   a   |   b   |   a   |   b   |   a   |   b   |   a   |
LSP length L:   0   1   0   3   0   5   0   7   0   5   0   3   0   1   0
Position:       0   1   2   3   4   5   6   7   8   9  10  11  12   13  14

example 2:
s = "abaaba"
s.length = 6

news = "|a|b|a|a|b|a|"
news.length = 2 * s.length + 1 = 13

string:        |   a   |   b   |   a   |   a   |   b   |   a   |
LSP length L:  0   1   0   3   0   1   6   1   0   3   0   1   0
Position:      0   1   2   3   4   5   6   7   8   9  10  11  12


index(N): 0 -> N-1
position(2*N+1): 0 -> 2*N

L[i] = d means:
1. substring from position i-d to i+d is a palindrome of length d. (in terms of position)
2. substring from index (i-d)/2 to [(i+d)/2-1] is a palindrome of length d. (in terms of index)

in string "abaaba" L[3] = 3 means:
1. substring from position 3-3=0 to 3+3=6 is a palindrome of length 3.
2. substring from index (3-3)/2=0 to [(3+3)/2-1]=2 is a palindrome of length 3.

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
	centerPosition := 1 // centerPosition 是已经计算完成的
	centerRight := 2

	var maxLPSLength int
	var maxLPSCenterPosition int

	for i := 2; i < N; i++ {
		needExpand := false
		curRight := i // curRight 是正要计算的
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
