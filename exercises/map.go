package exercises

import "fmt"

func Map() {
	family := make(map[string]int)

	// Add a few entries
	family["George"] = 139
	family["Megan"] = 40
	family["Sebastian"] = 5
	family["Pingu"] = 14

	// Update a mistake
	family["George"] = 39

	// Delete non-humans
	delete(family, "Pingu")

	for key := range family {
		fmt.Printf("family[%v] = %v\n", key, family[key])
	}

}
