package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func SmallestArea(l, w, h int) int {
	if (l * w) < (h * w) {
		if (l * w) < (l * h) {
			return l * w
		} else {
			return l * h
		}
	} else {
		if (h * w) < (l * h) {
			return h * w
		} else {
			return l * h
		}
	}

}

func CalculateWrapPaper(dimensions []int) int {
	l, w, h := dimensions[0], dimensions[1], dimensions[2]
	return ((2 * l * w) + (2 * w * h) + (2 * l * h) + SmallestArea(l, w, h))
}

func main() {
	wrappingPaper := 0
	file, err := os.Open("2015.2.input")
	if err != nil {
		log.Fatalf("Unable to open input file: %s\n", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		dimensions := strToIntDimensions(scanner.Text())
		wrappingPaper += CalculateWrapPaper(dimensions)
	}
	fmt.Println(wrappingPaper)
}
