package pointers

import "fmt"

func Pointer() {
	var p *int
	var x int = 11
	p = &x
	fmt.Println(p, x)
	*p = 12
	fmt.Println(p, x, &x)
}
