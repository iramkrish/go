package arraySliceMap

import "fmt"

func Array(){
	// fixed size cannot extend
	arr := [5]int{10, 20, 30, 40, 50}
	fmt.Println(arr,len(arr),cap(arr))
}
