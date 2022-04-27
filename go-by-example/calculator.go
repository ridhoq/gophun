package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
	getOperation(reader)
}

func getOperation(reader *bufio.Reader) {
	fmt.Println("What operation do you want to do? Valid options are: add, subtract, multiply, divide")
    operation, _ := reader.ReadString('\n')
	fmt.Printf("Operation selected: %s", operation)
}