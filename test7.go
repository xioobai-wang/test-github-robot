package main

import (
	"fmt"
)

func funUsePointer(paramSlice *[]int) {
	paramSliceV := *paramSlice
	paramSliceV[2] = 9
	for index := 0; index < 6; index++ {
		*paramSlice = append(*paramSlice, index)
	}
	fmt.Printf("funUsePointer paramSlice =%v\n", paramSlice)
}

func main() {
	slice := make([]int, 5)
	slice[0] = 2
	slice[1] = 3
	fmt.Println(slice)
	funUsePointer(&slice)
	fmt.Println(slice)
}
