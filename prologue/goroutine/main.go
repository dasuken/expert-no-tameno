package main

import "fmt"

func main() {
	ch := make(chan int)
	go send(ch, 100)
	fmt.Println(receive(ch))
}

func send(ch chan <- int, n int) {
	ch <- n
}

func receive(ch <- chan int) int {
	return <- ch
}

