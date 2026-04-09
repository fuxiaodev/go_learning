package L4

import(
	"fmt"
)

func l4 () {
	// most of the time we don't need to think about the underlying array of a slice.
	// we can create a new slice using the make function
	// func make([]T, len, cap)
	// mySlice_w_capacity := make([]int, 5, 10)
	// slices created with make will be filled with the zero value of the type
	// mySlice_wn_capacity := make([]int, 5)
	mySlice_literal := []string{"I", "love", "go"}

	// the capacity of a slice is the number of elements in the underlying array
	// counting from the first element in the slice
	// it can be accessed using the built-in cap() function
	fmt.Println(cap(mySlice_literal))
}
