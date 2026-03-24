package L1

// L1
type car struct {
	brand string
	model string
	doors int
	mileage int
	frontWheel wheel
	backWheel wheel
}

/*
nested structs in go
*/
type wheel struct {
	radius int
	material string
}

// := can only be used inside a function
var myCar = car {
	brand: "toyota",
	model: "blsd",
	doors: 2,
	mileage: 50000,
	frontWheel: wheel {
		radius: 5,
		material: "rubber",
	},
	backWheel: wheel {
		radius: 6,
		material: "rubber",
	},
}
