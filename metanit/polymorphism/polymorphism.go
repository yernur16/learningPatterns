package main

import "fmt"

type Vehicle interface {
	move()
}

type Car struct {
	model string
}

type Aircraft struct {
	model string
}

func (c Car) move() {
	fmt.Println(c.model, "driving")
}

func (a Aircraft) move() {
	fmt.Println(a.model, "flying")
}

func main() {
	tesla := Car{"Tesla"}
	volvo := Car{"Volvo"}
	boeing := Aircraft{"Boeing"}

	vehicles := [...]Vehicle{tesla, volvo, boeing}

	for _, values := range vehicles {
		values.move()
	}
}
