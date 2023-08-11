/*
Create and assign values to 2d array 3X3 array

store rows to columns and columns to rows in the array

# Write a func to perform those operations

Input:

1 2 3

4 5 6

7 8 9

Output:

	1 4 7

	2 5 8

	3 6 9
*/
package main

import "fmt"

func main() {
	arr := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	transverse(arr)

}

func transverse(arr [3][3]int) {

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0])/2; j++ {
			temp := arr[i][j]
			arr[i][j] = arr[j][i]
			arr[j][i] = temp

		}
	}
	fmt.Println(arr)

}
