package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var stringArray []string
	var a string
	fmt.Scan(&a)
	noOfStrings, _ := strconv.Atoi(a)
	for {
		var b string
		fmt.Scan(&b)
		stringArray = append(stringArray, b)

		if noOfStrings == len(stringArray) {
			break
		}
	}

	for _, stringValue := range stringArray {
		newString := strings.Split(stringValue, "")

		var evenString []string
		var oddString []string
		for i := 0; i < len(newString); {
			if i%2 == 0 {
				evenString = append(evenString, newString[i])
			} else {
				oddString = append(oddString, newString[i])
			}
			i++
		}

		printEven := strings.Join(evenString, "")
		printOdd := strings.Join(oddString, "")
		fmt.Printf("%s %s\n", printEven, printOdd)
	}

}
