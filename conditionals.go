// >Objective: Use if, else if, and else to determine and print out whether a user-inputted number is positive, negative, or zero.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter a number (press Ctrl+D or Ctrl+Z to end):")

	// Read input line by line
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break // Exit loop if an empty line is entered
		}
		num, err := strconv.ParseFloat(text, 64)
		if err != nil {
			fmt.Println("Your number was mal-formed.")
		} else if num == 0 {
			fmt.Println("Your number was zero")
		} else if num > 0 {
			fmt.Println("Your number was positive")
		} else {
			fmt.Println("Your number was negative")
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}

}
