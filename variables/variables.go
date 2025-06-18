package variables

import "fmt"

func Variables(){
	// string
	var name string = "ram"
	lastName := "krishnan"
	fmt.Println(name,lastName)

	// int
	var a,b int = 10 , 20
	fmt.Println(a,b)

	// default
	var stringOne string // ""
	var numberOne int // 0
	var boolOne bool // false
	fmt.Println(stringOne, numberOne, boolOne)

	// const
	// const needs to be initialized , no default value and canot be reassigned
	const constValue int = 90
	fmt.Println(constValue)

}
