package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 100
	fmt.Println(len(ch))
	ch <- 200
	fmt.Println(len(ch))

	close(ch)
	for c := range ch {
		fmt.Println(c)
	}

	// // chennelの取り出し
	// x := <-ch
	// fmt.Println(x)
	// fmt.Println(len(ch))

	// ch <- 300
	// fmt.Println(len(ch))
}
