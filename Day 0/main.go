//This first challenge is pretty straight forward, get input from stdin
//which also means standard input and print

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	//ReadString method only takes in '' single quote strings
	message, _ := reader.ReadString('\n')
	fmt.Println("Hello World")
	fmt.Println(message)

	//Alternatively
	newMessage:= readFromInput()
	fmt.Println("Hello World")
	fmt.Println(newMessage)
}

//Alternatively you can use the scanner instead of the reader
func readFromInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	return input
}
