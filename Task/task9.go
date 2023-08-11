package main

import "fmt"

func main() {
	arr := [6]int{4, 7, 5, 2, 7, 9}
	max, min := getsmallbig(arr)
	fmt.Println("Biggest number in arr", max)
	fmt.Println("smallest number in arr", min)

}
func getsmallbig(arr [6]int) (int, int) {
	var min int = arr[0]
	var max int = arr[0]
	for i := range arr {
		if arr[i] > max {
			max = arr[i]
		}
		if arr[i] < min {
			min = arr[i]
		}
	}
	return max, min
}
