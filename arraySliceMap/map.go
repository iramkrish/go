package arraySliceMap

import "fmt"

func Map() {
	// Create a map
	myMap := make(map[string]int)

	myMap["apple"] = 5
	myMap["banana"] = 10
	myMap["orange"] = 15

	fmt.Println("Map:", myMap)

	// Access a value by key
	value,bool := myMap["banana"]
	fmt.Println("Value for 'banana':", value , "Exists:", bool)

	for key, val := range myMap {
		fmt.Println("Key:", key, "Value:", val)
	}
}
