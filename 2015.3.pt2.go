package main

import (
	"fmt"
	"log"
	"os"
)

/*
	We now track Robo-Santa's coordinates as well. DeilveryStatus.Location is a list of lists [ [], [] ]
	representing Santa and Robo-Santa's location coordinates respectively.
	A directional char at even position is taken up by Santa while at odd position is taken up by Robo.
	This logic is added to the for loop where in the beginning we calculate the even or odd function
	for the input, and within the case, use it to access coordinates either at the 0th of 1st position.

	Both the Santa make deliveries in the same home grid and hence the Delivery map is shared.
*/

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
