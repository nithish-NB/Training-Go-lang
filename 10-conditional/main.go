package main

import (
	"fmt"
	"reflect"
)

func main() {
	switch day := 4; day {
	case 1:
		fmt.Println("Sunday")
	case 2:
		fmt.Println("Monday")
	case 3:
		fmt.Println("Tuesday")
	default:
		fmt.Println("all day")
	}
	char := 'A'
	fmt.Println(char, reflect.TypeOf(char))
}
