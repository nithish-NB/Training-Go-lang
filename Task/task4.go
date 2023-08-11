package main

import "fmt"

var a int = 10 // global variable

func main() {

	var b int = 20 // local variable

	fmt.Println("Local variable:", b)

	fmt.Println("Global Variable:", a)
	glb()
}

func glb() {
	c := 10 //local variable
	fmt.Println("Local varibale 2:", c)
}
