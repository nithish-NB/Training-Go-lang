package main

import "fmt"

func main() {
	//	a := 20
	//b := 10

	var fmap map[string]func(int, int) int
	fmap = make(map[string]func(int, int) int)
	fadd := func(a int, b int) int {
		return a + b
	}
	fsub := func(a, b int) int {
		an := a - b
		return an
	}
	mul := func(a, b int) int {
		return a * b
	}
	fmap["add"] = fadd
	fmap["sub"] = fsub
	fmap["mul"] = mul
	//fmap["div"] = div()
	for key, f := range fmap {
		fmt.Println(key, f(12, 14))
	}

}
