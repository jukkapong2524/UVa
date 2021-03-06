// UVa 10235 - Simply Emirp

package main

import (
	"fmt"
	"os"
)

func isPrime(n int) bool {
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func reverse(n int) int {
	r := 0
	for n != 0 {
		r = r*10 + n%10
		n /= 10
	}
	return r
}

func main() {
	in, _ := os.Open("10235.in")
	defer in.Close()
	out, _ := os.Create("10235.out")
	defer out.Close()

	var n int
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		if isPrime(n) {
			r := reverse(n)
			if n != r && isPrime(r) {
				fmt.Fprintf(out, "%d is emirp.\n", n)
			} else {
				fmt.Fprintf(out, "%d is prime.\n", n)
			}
		} else {
			fmt.Fprintf(out, "%d is not prime.\n", n)
		}
	}
}
