package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func strToIntDimensions(dimensions string) (intDimensions []int) {
	strDimensions := strings.Split(dimensions, "x")

	for _, str := range strDimensions {
		i, err := strconv.Atoi(str)
		if err != nil {
			fmt.Printf("Unable to parse dimensions: %s\n", err)
			return []int{}
		}
		intDimensions = append(intDimensions, i)
	}
	return
}

func CalculateRibbon(dimensions []int) int {
	sort.Ints(dimensions)
	return ((2 * dimensions[0]) + (2 * +dimensions[1]) + (dimensions[0] * dimensions[1] * dimensions[2]))
}

func main() {
	ribbon := 0
	file, err := os.Open("2015.2.input")
	if err != nil {
		log.Fatalf("Unable to open input file: %s\n", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		dimensions := strToIntDimensions(scanner.Text())
		ribbon += CalculateRibbon(dimensions)
	}
	fmt.Println(ribbon)
}
