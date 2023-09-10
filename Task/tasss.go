package main

import (
	"fmt"
	"os"
	"runtime"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main2() {
	f, err := os.ReadFile("hi.txt")
	check(err)
	ch := make(chan string)
	func() {
		ch <- string(f)
		ch1 := <-ch
		fmt.Print(ch1)
	}()

	fmt.Print(ch)
	fmt.Println()
	runtime.Goexit()

}
