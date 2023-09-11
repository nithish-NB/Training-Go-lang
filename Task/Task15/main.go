package main

import (
	interfaces "demo/Task/Task15/interface"
	"demo/Task/Task15/shape"
	"fmt"
	"reflect"
)

func main() {
	var num shape.Circle
	fmt.Println(reflect.TypeOf(num))
	num = shape.Circle{R: 6}
	Area(&num)
	Perimeter(&num)

}

func Area(iarea interfaces.IArea) { // interface as a parameter.It can be a form of dependency injecttion
	a := iarea.Area()
	fmt.Println("Area :", a)
}
func Perimeter(ipere interfaces.Iperemeter) {
	p := ipere.Perimeter()
	fmt.Println("perimeter:", p)
}
