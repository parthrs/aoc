package main

import (
	"fmt"
	"log"
	"os"
)

type DeliveryStatus struct {
	Location    [2][2]int // Santa coordinages at 0, Robo-Santa coordinates at 1
	DeliveryMap map[[2]int]int
}

func NewDeliveryStatus() *DeliveryStatus {
	return &DeliveryStatus{
		Location: [2][2]int{{0, 0}, {0, 0}},
		DeliveryMap: map[[2]int]int{
			{0, 0}: 2,
		},
	}
}

var SantaDelivery *DeliveryStatus

func init() {
	SantaDelivery = NewDeliveryStatus()
}

func DeliverGift(direction string) {
	for i, _direction := range direction {
		direction := string(_direction)
		evenOrOdd := i % 2

		switch direction {
		case ">":
			SantaDelivery.Location[evenOrOdd][1] += 1
		case "<":
			SantaDelivery.Location[evenOrOdd][1] -= 1
		case "v":
			SantaDelivery.Location[evenOrOdd][0] -= 1
		case "^":
			SantaDelivery.Location[evenOrOdd][0] += 1
		}

		SantaDelivery.DeliveryMap[SantaDelivery.Location[evenOrOdd]] += 1
	}
}

func main() {
	_directions, err := os.ReadFile("2015.3.input")
	directions := string(_directions)
	if err != nil {
		log.Fatalf("Unable to read input file: %s. Exiting..", err)
	}
	DeliverGift(directions)

	fmt.Println(len(SantaDelivery.DeliveryMap))
}
