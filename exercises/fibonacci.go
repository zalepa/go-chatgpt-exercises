package exercises

import "fmt"

func Fibonacci() {

	nums := make([]int, 10)

	for i := 0; i < 10; i++ {
		if i == 0 {
			nums[0] = i
		} else if i == 1 {
			nums[1] = i
		} else {
			nums[i] = nums[i-1] + nums[i-2]
		}
	}

	for i := range nums {
		fmt.Printf("%v ", nums[i])
	}

	fmt.Print("\n")
}
