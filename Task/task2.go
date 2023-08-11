package main

import "fmt"

func main() {

	println("after")
	fmt.Println("hello world")      //after exceution goes to new line
	fmt.Print("hello world ")       //fter exceution doesn't goes to new line
	fmt.Printf("%s", "hello world") // we need to use access specifiers
}
