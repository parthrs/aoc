package main

import (
	"fmt"
	"log"
	"os"
)

type DeliveryStatus struct {
	Location    [2]int
	DeliveryMap map[[2]int]int
}

func NewDeliveryStatus() *DeliveryStatus {
	return &DeliveryStatus{
		Location: [2]int{0, 0},
		DeliveryMap: map[[2]int]int{
			{0, 0}: 1,
		},
	}
}

var SantaDelivery *DeliveryStatus

func init() {
	SantaDelivery = NewDeliveryStatus()
}

func DeliverGift(direction string) {
	for _, _direction := range direction {
		direction := string(_direction)

		switch direction {
		case ">":
			SantaDelivery.Location[1] += 1
		case "<":
			SantaDelivery.Location[1] -= 1
		case "v":
			SantaDelivery.Location[0] -= 1
		case "^":
			SantaDelivery.Location[0] += 1
		}

		SantaDelivery.DeliveryMap[SantaDelivery.Location] += 1
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
