package L6

import "fmt"

/*
While go is not object oriented, it does support methods that can be defined on structs.
Methods are just functions that have a receiver. 
A receiver is a special parameter that syntactically goes before the name of the function
*/
type rect struct {
	width int
	height int
}

/*
area has a receiver of (r rect)
rect is the struct
r is the placeholder
by convention, go code will often use the first letter of the struct's name
*/
func (r rect) area() int {
	return r.width * r.height
}

var r = rect {
	width: 5,
	height: 10,
}

func print_area(){
	fmt.Println(r.area())
}

type authenticationInfo struct {
	username string
	password string
}

func (a authenticationInfo) auth_string() string {
	return "Authorization: Basic " + a.username + ":" + a.password
}


