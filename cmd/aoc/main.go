package main

import (
	"fmt"

	"github.com/todorovski-vasil/aoc2023/internal/trebuchet"
)

func main() {
	path := "./internal/trebuchet/calibration.txt"
	// path := "./internal/trebuchet/testCalibration.txt"
	calibrationValue := trebuchet.GetDocumentAndCalibrate(path)
	fmt.Printf("The calibration value is, %d.\n", calibrationValue)
}
