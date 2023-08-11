package main

import "fmt"

func main() {
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
