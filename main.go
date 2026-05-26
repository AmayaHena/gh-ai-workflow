// Tiny FizzBuzz CLI. Prints FizzBuzz from 1 to N (default 20).
package main

import (
	"fmt"
	"os"
	"strconv"
)

func fizzbuzz(n int) string {
	if n%15 == 0 {
		return "FizzBuzz"
	}
	if n%3 == 0 {
		return "Fizz"
	}
	if n%5 == 0 {
		return "Buzz"
	}
	return strconv.Itoa(n)
}

func run(args []string) int {
	n := 20
	if len(args) >= 1 {
		v, err := strconv.Atoi(args[0])
		if err != nil || v < 1 {
			fmt.Fprintln(os.Stderr, "usage: fizzbuzz [N]  (N must be a positive integer)")
			return 2
		}
		n = v
	}
	for i := 1; i <= n; i++ {
		fmt.Println(fizzbuzz(i))
	}
	return 0
}

func main() {
	os.Exit(run(os.Args[1:]))
}
