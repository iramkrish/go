package arraySliceMap

import "fmt"

func Slice(){
	// dynamic size can be extended using append
	arr := []int{10, 20, 30, 40, 50}
	fmt.Println(arr,len(arr),cap(arr))
	arr = append(arr,60)
	fmt.Println(arr,len(arr),cap(arr))
}
