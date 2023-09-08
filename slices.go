// Objective: Create a slice, append 5 values to it, and print the slice in reverse order.
package main

import (
	"fmt"
	"slices"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	for i := len(s); i > 0; i-- {
		fmt.Printf("%v ", s[i-1])
	}
	fmt.Print("\n")

	// Alterative
	// Warning: modifies `s`
	slices.Reverse(s)
	for i := range s {
		fmt.Printf("%v ", s[i])
	}
}
