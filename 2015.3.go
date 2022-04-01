package main

import (
	"fmt"
	"log"
	"os"
)

/*
	The delivery grid is thought of as a 2-d coordinate system, with x, y axis.
	We then track Santa's coordinates (DeliveryStatus.Location) and the coordinates
	of the home where a gift was dropped.

	A map is used - with keys as the coordinates and values as the number of gifts dropped
	at that coordinate - to keep track of the homes that recieved gifts.

	Since we only care about the count of homes that recieved more than one gift, we simply
	take the count of homes (len(Map)) in the map i.e. the keys in the map.
*/

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

// File level var
var SantaDelivery *DeliveryStatus

// Idea is to not have this pass around as a parameter (i.e. if declared in main)
func init() {
	SantaDelivery = NewDeliveryStatus()
}

// DeliverGift takes a series of directions in the form of a single string.
// It then iterates over each 'direction' step, incrementing Santa's positional
// coordinates in the process. At the end of the loop it adds/increments the
// gift count to the resulting (house) coordinate.
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

// The main function parses the input from a file.
// Input file should contain a string of characters denoting the
// direction for Santa to deliver gifts. This function prints the
// number of homes that recieved atleast 1 gift at the
// end of parsing all the directions.
func main() {
	_directions, err := os.ReadFile("2015.3.input")
	directions := string(_directions)
	if err != nil {
		log.Fatalf("Unable to read input file: %s. Exiting..", err)
	}
	DeliverGift(directions)

	fmt.Println(len(SantaDelivery.DeliveryMap))
}
