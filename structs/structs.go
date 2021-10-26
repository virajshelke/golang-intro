package main

import (
	"fmt"
)

type Dog struct {
	Breed string
	Color string
	Sound string
}

func main() {
	poodle := Dog{Breed: "poodle", Color: "white", Sound: "Woof!"}

	fmt.Printf("%+v\n", poodle)

	poodle.Speak()
	poodle.SpeakThreeTimes()
	poodle.Speak()
}

func (d Dog) Speak() {
	fmt.Println("I bark like", d.Sound)
}

func (d Dog) SpeakThreeTimes() {
	d.Sound = fmt.Sprintf("%v %v %v", d.Sound, d.Sound, d.Sound)
	fmt.Println("I bark like", d.Sound)
}
