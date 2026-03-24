package L3
/*
an anonymous struct is just like a normal struct, but it is defined without a name
and therefore cannot be referenced elsewhere in the code

to create an anonymous struct, just instantiate the instance immediately 
using a second pair of brackets after declaring the type
*/
var myCar = struct {
	brand string
	model string
}{
	brand: "Toyota",
	model: "Camry",
}


/*
you can even nest anonymous structs as fields within other structs
*/
type car struct {
	brand string
	model string
	doors int
	mileage int
	// wheel is a field containing an anonymous struct
	wheel struct {
		radius int
		material string
	}
}

var myCar2 = car {
	brand: "Rezvani",
	model: "Vengeance",
	doors: 4,
	mileage: 35000,
	// you must repeat the struct definition when initializing it
	wheel: struct {
		radius int
		material string
	}{
		radius: 35,
		material: "alloy",
	},
}

//you do not need to repeat the definition if you initialize it in steps
var myCar3 car
// you cannot execute statements at the package level
func initialize() {
	myCar3.brand = "Toyota"
	myCar3.model = "sadfasd"
	myCar3.doors = 4
	myCar3.mileage = 35000

	myCar3.wheel.radius = 35
	myCar3.wheel.material = "alloy"
}


/*
in general, prefer named structs.
named structs make it easier to read and understand your code, and they have the nice side effect of being reusable.
use anonymous structs if you are sure you won't ever need to use that struct again
and anonymous struct makes it clear that this data structure only exists inside its parent struct
nested anonymous struct is only good for short struct since you need to repeat the definition when initializing it
*/