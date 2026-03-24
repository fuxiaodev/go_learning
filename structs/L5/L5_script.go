package L5

import "fmt"

/*
go is not object oriented language.
go has embedded structs that provide a kind of data only inheritance that can be useful at times.
Go doesn't support classes or inheritance in the complete sense, but embedded structs are a way to elevate and 
share fields between struct definitions
*/
type car struct {
	brand string
	model string
}

type truck struct {
	/*
	car is embedded so the definition of a truck now also additionally contains all
	of the fields of the car struct
	*/
	car
	bedSize int
}

/*
unlike nested structs, an embedded structs' fields are accessed at the top level like normal fields
a nested struct is a struct that appears as a named field inside another struct
an embedded struct is when you include a struct without giving it a field name
use embedded struct = nested struct + field promotion
embedded structs can cause field conflicts, imagining two embedded structs have the same field name
embedded structs also promote methods

nested structs:
enforce clear structure
avoid field conflicts
prevent unintended method promotion
*/
var lanesTruck = truck {
	bedSize: 10,
	car: car {
		brand: "Toyota",
		model: "Tundra",
	},
}

func access_field() {
	fmt.Println(lanesTruck.brand)
	fmt.Println(lanesTruck.model)
}



