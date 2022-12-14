package main

import "fmt"

type Vehicle interface {
	move()
}

func drive(vehicle Vehicle) {
	vehicle.move()
}

type (
	Car      struct{}
	Aircraft struct{}
)

func (c Car) move() {
	fmt.Println("Автомобиль едет")
}

func (a Aircraft) move() {
	fmt.Println("Самолет летит")
}

func main() {
	var tesla Car
	var boing Aircraft

	drive(tesla)
	drive((boing))
}
