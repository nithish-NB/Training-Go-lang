/*
- Find the second biggest number in an array/slice

- Find the second smallest number in an array/slice

- Write everything using functions.

- SecondBiggest(arr []int)int
- SecondSmallest(arr []int)int
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{13, 12, 7, 3, 5, 2, 1}
	sb := SecondBiggest1(arr)
	sb2 := SecondBiggest2(arr)
	ss := SecondSmallest(arr)
	fmt.Println(sb, sb2)
	fmt.Println(ss)

}
func SecondBiggest1(arr []int) int { //time complexity O(2n)
	max1 := math.MinInt
	max2 := math.MinInt
	for i := 0; i < len(arr); i++ {
		if arr[i] > max1 {
			max1 = arr[i]
		}

	}
	for i := 0; i < len(arr); i++ {
		if arr[i] != max1 {
			if arr[i] > max2 {
				max2 = arr[i]
			}
		}
	}
	return max2
}
func SecondBiggest2(arr []int) int { //time complexity O(n)
	max1 := math.MinInt
	max2 := math.MinInt
	for i := 0; i < len(arr); i++ {

		if max2 < max1 && arr[i] > max2 && arr[i] > max1 {
			max1 = arr[i]
			max2 = max1
		}

	}
	return max2
}
func SecondSmallest(arr []int) int {
	min1 := arr[0]
	min2 := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] < min1 {
			min1 = arr[i]
		}

	}
	for i := 0; i < len(arr); i++ {
		if arr[i] != min1 {
			if arr[i] < min2 {
				min2 = arr[i]
			}
		}
	}

	return min2
}
