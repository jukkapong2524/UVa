// UVa 108 - Maximum Sum

package main

import (
	"fmt"
	"math"
	"os"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMax(s [][]int) int {
	l := len(s)
	msf := math.MinInt32
	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			meh := 0
			for k := 0; k < l; k++ {
				tmp := s[j][k]
				if i > 0 {
					tmp -= s[i-1][k]
				}
				meh = max(meh+tmp, tmp)
				msf = max(msf, meh)
			}
		}
	}
	return msf
}

func prefixSum(r [][]int) [][]int {
	l := len(r)
	s := make([][]int, l)
	for i := range s {
		s[i] = make([]int, l)
	}
	for i := range s {
		for j := range s[i] {
			s[i][j] = r[i][j]
			if i > 0 {
				s[i][j] += s[i-1][j]
			}
		}
	}
	return s
}

func main() {
	in, _ := os.Open("108.in")
	defer in.Close()
	out, _ := os.Create("108.out")
	defer out.Close()

	var n int
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		r := make([][]int, n)
		for i := range r {
			r[i] = make([]int, n)
		}

		for i := range r {
			for j := range r[i] {
				fmt.Fscanf(in, "%d", &r[i][j])
			}
		}
		fmt.Fprintln(out, findMax(prefixSum(r)))
	}
}
