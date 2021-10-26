package main

import (
	"fmt"
	"time"
)

func main() {
	// go count("fish")
	// go count("sheep")
	count("fish")
	count("sheep")

	fmt.Scanln()
}

func count(thing string) {
	for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
