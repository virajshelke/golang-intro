package main

import (
	"fmt"
)

// There is no concept of classes in go lang & no inheritence
// We have structs!

// Note the captial letter in the start of these member indicates they are public and can be accessed directly i.e. they are exported!
type Dog struct {
	Breed string
	Color string
	Sound string
}

func main() {
	var poodle Dog = Dog{Breed: "poodle", Color: "white", Sound: "Woof!"}
	// we can declare it in the following manner as well
	// poodle := Dog{Breed: "poodle", Color: "white", Sound: "Woof!"}

	// print the complete instance of the type
	fmt.Printf("%+v\n", poodle)

	// calling the methods available on the poodle object (instance of struct)
	poodle.Speak()           // prints - woof!
	poodle.SpeakThreeTimes() // prints - woof! woof! woof!
	poodle.Speak()           // prints - woof!
}

// This method is available on the Dog type because we are using reference here (d Dog)
// the d here is the copy of the object that will be passed to this function when we call this method on the object
func (d Dog) Speak() {
	fmt.Println("I bark like", d.Sound)
}

// In this example we try to modify the object member value but still in the main when we look at the output
// we can see that the original poodle object is not modified
// This brings us to the concept that the objects are passed a value
// We can also use pointers and pass by reference if we need the original instance & not copy of the object
func (d Dog) SpeakThreeTimes() {
	d.Sound = fmt.Sprintf("%v %v %v", d.Sound, d.Sound, d.Sound)
	fmt.Println("I bark like", d.Sound)
}
