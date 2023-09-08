package exercises

import "fmt"

func Vars() {
	var name string
	name = "George"
	var age int = 39
	weight := 171.1
	isMale := true

	fmt.Printf("%v is %v years old and weighs %v pounds. Is he male? %v\n", name, age, weight, isMale)

	fmt.Printf("%T is %T years old and weighs %T pounds. Is he male? %T\n", name, age, weight, isMale)
}
