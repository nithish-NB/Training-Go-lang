package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	//ch4 := make(chan int)
	ch4 := variadicchan(ch1, ch2, ch3)
	go func() {
		ch1 <- 10
		//close(ch1)
		ch2 <- 20
		//close(ch2)
		ch3 <- 30
		//close(ch3)

	}()

	for sum1 := range ch4 {
		fmt.Println(sum1)
	}

}

func variadicchan(chs ...chan int) chan int {
	ch := make(chan int)
	go func() {
		sum := 0
		for _, v := range chs {
			a := <-v
			sum = sum + a
		}
		ch <- sum
		close(ch)
	}()
	return ch
}
