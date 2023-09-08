# Using ChatGPT as a Go Tutor

## Exercise 1: Hello World!
> Objective: Set up your Go environment and print "Hello, World!" to the console.

## Exercise 2: Variables and Types
> Objective: Declare different types of variables (string, integer, float, boolean) and print their values and types.

## Exercise 3: Conditional Statements
> Objective: Use if, else if, and else to determine and print out whether a user-inputted number is positive, negative, or zero.

## Exercise 4: Loops
> Objective: Use a for loop to print the first 10 numbers in the Fibonacci series.

## Exercise 5: Functions
> Objective: Write a function that calculates the factorial of a number and returns it. Test your function with various inputs.

## Exercise 6: Slices
> Objective: Create a slice, append 5 values to it, and print the slice in reverse order.

## Exercise 7: Maps
> Objective: Create a map that stores the age of people by name. Add a few entries, update an entry, delete an entry, and print the map.

## Exercise 8: Goroutines and Channels
Objective: Familiarize yourself with Go's concurrency primitives, specifically goroutines and channels. Create a program that uses goroutines to calculate the sum of two slices concurrently and then combines the results.

1. Create two slices of integers.
2. Use one goroutine to calculate the sum of the first slice.
3. Use another goroutine to calculate the sum of the second slice.
4. Both goroutines should send their results to a single channel.
5. The main function should read from the channel twice to get both sums and then combine them to get the total sum.
6. Print out the individual sums of both slices as well as the combined total.

This exercise will help you get hands-on experience with launching goroutines and synchronizing them using channels.