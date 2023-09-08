package exercises

import (
	"fmt"
	"slices"
)

func Slices() {
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
