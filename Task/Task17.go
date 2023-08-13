package main

import "fmt"

func main() {
	vowles, cons, speci := func(str string) (int, int, int) {
		vowels := 0
		cons := 0
		spec := 0
		for _, c := range str {
			ch := string(c)
			if ch != " " {
				if c >= 65 && c <= 90 || c >= 97 && c <= 122 {
					if ch == "a" || ch == "e" || ch == "i" || ch == "o" || ch == "u" {
						vowels++
					} else {
						cons++
					}
				} else {
					spec++
				}
			}
		}
		return vowels, cons, spec
	}("Victoria Secert @ and Co.")

	fmt.Println(vowles, cons, speci)
}
