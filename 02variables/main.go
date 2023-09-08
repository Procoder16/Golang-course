package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to Taking Inputs!")

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the ratings please:")

	input, _ := reader.ReadString('\n')

	fmt.Println("The rating you entered is: ", input)
	fmt.Printf("The type of input is: %T", input)

	newRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(newRating + 1)
	}
}
