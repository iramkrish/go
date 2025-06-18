package conditionals

import "fmt"

// starts with lowercase character so it's not exported by default
// functions starting with capital character are exported by default
func MultiplyByTwo(arg int) int {
	return arg * 2
}

func Conditionals() {
	for i := 0; i < 5; i++ {
		fmt.Println(MultiplyByTwo(i))
	}

	// switch cases
	num := 1

	switch num {
	case 1:
		fmt.Println("One")
		fallthrough
	case 2:
		fmt.Println("Two (because of fallthrough)")
	}

}
