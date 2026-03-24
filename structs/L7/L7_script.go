package L7
import ("reflect")

/*
in Go, structs sit in memory in a contiguous block, with fields placed one after another as defined in the struct
*/
type stats_good struct {
	Reach uint16
	NumPosts uint8
	NumLikes uint8
}

/*
Go aligns the fields, meaning that it has added some padding (wasted space) to make up for the size difference
between the uint16 and uint8 types.
It is done for execution speed, but it can lead to increased memory usage
*/
type stats_bad struct {
	NumPosts uint8
	Reach uint16
	NumLikes uint8
}

/*
Normally you should not stress about memory layout.
However, if you have a specific reason to be concerned about memory usage, aligning the fields by size
(largest -> smallest) can help.
you can also use the reflect package to debug the memory layout of the struct
*/
func mem_layout(){
	typ := reflect.TypeOf(stats_good{})
	fmt.Printf("Struct is %d bytes\n", typ.Size())
}


/*
empty structs are used in Go as unary value
empty structs are the smallest possible type in Go
they can take up zero bytes of memory
*/
//anonymous empty struct
var empty_ano = struct{}{}

//named empty struct
type emptyStruct struct{}
var empty_named = emptyStruct{}



