package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	inputInt, _ := strconv.ParseInt(input, 10, 64)

	for i := 1; i <= 10; i++ {
		newValue := inputInt * int64(i)
		fmt.Printf("%d x %d = %d \n", inputInt, i, newValue)
	}

}
