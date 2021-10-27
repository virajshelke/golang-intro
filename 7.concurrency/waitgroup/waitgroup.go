package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// creating a wait group
	var wg sync.WaitGroup

	// Adding the info that we are spwaning two goroutines! Setting the wait counter to 2
	wg.Add(2)

	// creating an Anonymous function & calling the count function inside it and then marking it as done!
	// This makes sure that the function implementation (count) isn't responsible to manage wait groups
	// And makes function count more reusable i.e. we can use it normally or as a goroutine
	go func() {
		count("sheep")
		wg.Done() // decrements the wait counter by 1
	}()

	go func() {
		count("fish")
		wg.Done()
	}()

	// Blocking line that waits for all the goroutines to finish i.e. waits for the count to become zero
	wg.Wait()
}

func count(thing string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
