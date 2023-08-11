package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num1 int = 7
	var num2 int8 = 101
	var num3 int16 = 19020
	var num4 int32 = 1426024011
	var num5 int64 = 4238713284952114251
	var num6 float32 = 2.3
	var num7 float64 = 4.66
	var num8 uint = 2
	var num9 uint8 = 221
	var num10 uint16 = 7522
	var num11 uint32 = 987653
	var num12 uint64 = 9876543211234567898
	//var num8 bool = true
	//var num9 string = "Nithish"
	var num13 byte = 222       //alias for int8
	var num14 rune = 687332454 //alias for int32

	var inttype any = int(num1) + int(num10) + int(num11) + int(num4) + int(num5) + int(num6) + int(num8)

	var float32type any = float64(num12) - float64(num13) + float64(num14) - float64(num2) + float64(num3) + float64(num4) - float64(num9) - float64(num7)

	fmt.Println("interface of inttype", inttype, reflect.TypeOf(inttype))
	fmt.Println("interface of float32 type", float32type, reflect.TypeOf(float32type))
}
