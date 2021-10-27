package main

import (
	"fmt"
	"time"
)

func main() {

	// Block 1
	// The following code prints fish only! Since it's an infinte loop we never finish the first function call
	count("fish")
	count("sheep")

	// Block 2
	// If you comment the block 1 & uncomment this block 2 then we can see the output as
	// fish & sheep both because writing go creates a goroutine (a function that runs concurrently)
	// Note that since we sleep for sometime in between we can execute both these calls concurrently!
	/*
		// go count("fish")
		// count("sheep")
	*/

	// Block 3
	// If you execute the following block where both are goroutines then nothing is displayed on screen
	// and program terminates becuase we say to go runtime that execute these concurrently and then there is
	// nothing in main to execute so we finish!
	/*
		// go count("fish")
		// go count("sheep")
	*/

	// This is a hack where we introduce a blocking function waiting for user input to keep main alive
	// while we execute the goroutines
	// Not the proper way to do things! Refer wait group or channels or worker pools!
	// fmt.Scanln()
}

func count(thing string) {
	for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
