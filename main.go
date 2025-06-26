package main

import (
	"fmt"
	"myapp/arraySliceMap"
	"myapp/conditionals"
	"myapp/errors"
	"myapp/pointers"
	"myapp/variables"
)

func main() {
	variables.Variables()
	conditionals.Conditionals()
	arraySliceMap.Array()
	arraySliceMap.Slice()
	arraySliceMap.Map()
	pointers.Pointer()
	fmt.Println(errors.ValidateAge(100))
}
