package main

import "fmt"

func main() {
	fmt.Println(gdc(8, 6))
	fmt.Println(fib(4))
}

func gdc(x int, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
