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
	N, _ := strconv.ParseInt(input, 10, 64)

	//In the challenge we were asked to check if n is even in the subsequent if statements but in Go
	//that's not neccessary this is because if the N variable fufills the first if statement the code
	//breaks meaning any variable that passes is even.
	if N%2 != 0 {
		fmt.Println("Weird")
	} else if N >= 2 && N <= 5 {
		fmt.Println("Not Weird")
	} else if N >= 6 && N <= 20 {
		fmt.Println("Weird")
	} else if N > 20 {
		fmt.Println("Not Weird")
	}
}
