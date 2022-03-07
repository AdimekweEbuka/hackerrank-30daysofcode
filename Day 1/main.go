//This is a little more complicated as you need a for loop but we can run it (easy).
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//From problem
var i uint64 = 4
var d float64 = 4.0
var s string = "HackerRank "

func main() {
	var newInt uint64
	var newFloat float64
	var newString string
	var inputValues []string

	//This for loop gets multiple input values and adds them to a slice which can be seperated later
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		newInput := scanner.Text()
		inputValues = append(inputValues, newInput)
		if len(inputValues) == 3 {
			break
		}
	}

	newInt, _ = strconv.ParseUint(inputValues[0], 10, 64)
	newFloat, _ = strconv.ParseFloat(inputValues[1], 64)
	newString = inputValues[2]

	fmt.Println(newInt + i)
	fmt.Println(newFloat + d)
	fmt.Println(s + newString)

}
