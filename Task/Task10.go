/*Take 2 dimentional array  3 X 3 and assign values .

Get the sume of Each row and each column data

Write a func to perform those operations

Input:

    1 2 3

    4 5 6

    7 8 9

Output:

   Row 1 Sum: 6

   Row 2 Sum: 15

   Row 3 Sum: 24

   Column 1 Sum: 12

   Column 2 Sum: 15

   Column 3 Sum: 18 */

package main

import "fmt"

func main10() {
	arr := [3][3]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}
	vertical, horzi := addrowcol(arr)
	fmt.Println(vertical, horzi)
}

func addrowcol(arr [3][3]int) (int, int) {
	vertical := 0
	horzi := 0
	for i, val := range arr {
		for j, _ := range val {
			horzi = horzi + arr[i][j]
		}
	}
	for i, val := range arr {
		for j, _ := range val {
			vertical = vertical + arr[j][i]
		}
	}
	return vertical, horzi
}
