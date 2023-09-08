// Objective: Familiarize yourself with Go's concurrency primitives, specifically goroutines and channels.
// Create a program that uses goroutines to calculate the sum of two slices concurrently and then combines the results.

// 1. Create two slices of integers.
// 2. Use one goroutine to calculate the sum of the first slice.
// 3. Use another goroutine to calculate the sum of the second slice.
// 4. Both goroutines should send their results to a single channel.
// 5. The main function should read from the channel twice to get both sums and then combine them to get the total sum.
// 6. Print out the individual sums of both slices as well as the combined total.

// This exercise will help you get hands-on experience with launching goroutines and synchronizing them using channels.
package exercises

import "fmt"

func add(s []int, c chan int) {
	total := 0
	for i := range s {
		total += s[i]
	}
	c <- total
}

func Concurrency() {
	m := make(chan int)
	total := 0

	s1 := []int{1, 3, 5, 7}
	s2 := []int{2, 4, 6, 8}

	go add(s1, m)
	go add(s2, m)

	total += <-m
	total += <-m
	close(m)
	fmt.Println(total)
}
