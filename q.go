package main

import (
	"fmt"
)

func producer(c chan rune, max int) {
	defer close(c)
	for i := 0; i < max; i++ {
		select {
		case c <- ' ':
		case c <- '%':
		}
	}
}

func consumer(c chan rune) {
	var v rune
	ok := true
	for ok {
		if v, ok = <-c; ok {
			fmt.Print(string(v))
		}
	}
}

func main() {
	c := make(chan rune)
	go producer(c, 1000)
	consumer(c)
}
