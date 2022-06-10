package main

import "fmt"

func Fib(n int) int {
	//นานจัดๆ
	/*
		if n <= 1 {
			return n
		}
		return Fib(n-1) + Fib(n-2)
	*/
	f := make([]int, n+1, n+2)
	if n < 2 {
		f = f[0:2]
	}
	f[0] = 0
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
		//fmt.Println(f[i])
	}
	return f[n]
}
func main() {
	var n int
	var f int
	fmt.Print("Enter :")
	fmt.Scanln(&n)
	f = Fib(n)
	fmt.Println(f)
}
