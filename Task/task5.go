package main

import "fmt"

func main() {
	var num1 int = 7
	var num2 int8 = 101
	var num3 int16 = 19020
	var num4 int32 = 1426024011
	var num5 int64 = 4238713284952114251
	var num6 float32 = 2.3
	var num7 float64 = 4.66
	var num12 uint = 2
	var num13 uint8 = 221
	var num14 uint16 = 7522
	var num15 uint32 = 987653
	var num16 uint64 = 9876543211234567898
	//var num8 bool = true
	//var num9 string = "Nithish"
	var num10 byte = 222       //alias for int8
	var num11 rune = 687332454 //alias for int32

	add := float64(num1) + float64(num10) + float64(num11) + float64(num4) + float64(num5) + float64(num6) + float64(num15)

	sub := float64(num12) - float64(num13) - float64(num14) - float64(num2) - float64(num3) - float64(num4) - float64(num16) - float64(num7)

	fmt.Println("Addition of vairables:", add)
	fmt.Println("subtraction of varibales", sub)
}
