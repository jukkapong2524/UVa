// UVa 324 - Factorial Frequencies

package main

import (
	"fmt"
	"math/big"
	"os"
)

func factorial(n int) string {
	var res big.Int
	res.SetInt64(1)
	for i := 1; i <= n; i++ {
		res.Mul(&res, big.NewInt(int64(i)))
	}
	return res.String()
}

func output(out *os.File, str string) {
	var res [10]int
	for i := range str {
		digit := str[i] - '0'
		res[digit]++
	}
	for i, v := range res {
		fmt.Fprintf(out, "   (%d)%5d", i, v)
		if i != 4 && i != 9 {
			fmt.Fprint(out, " ")
		} else {
			fmt.Fprintln(out)
		}
	}
}

func main() {
	in, _ := os.Open("324.in")
	defer in.Close()
	out, _ := os.Create("324.out")
	defer out.Close()

	var n int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		fmt.Fprintf(out, "%d! --\n", n)
		output(out, factorial(n))
	}
}
