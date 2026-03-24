package main

import "fmt"

func main() {
	// initialize variables here
	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool
	var username string

	fmt.Printf("%v %.2f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
	/*
	%v
	value in default format
	works for almost any type: int, bool, float, struct, etc

	%.2f
	floating point formatting
	.2 means 2 decimal places
	f means fixed point notation

	%q  
	quoted string
	prints a string with quotes around it
	*/

	messageStart:= "Happy birthday! you are now"
	age := 21
	messageEnd := "years old!"
	fmt.Println(messageStart, age, messageEnd)

	/*
	:= walrus operator should be used instead of var style declaration basically anywhere possible
	The limitation is that := can't be used outside of a function
	and it is used to declare a new variable and assign it only
	*/

	temperatureFloat := 17.98
	temperatureInt := int(temperatureFloat)
	/*
	casting a float to an integer in this way truncates the floating point proportion
	*/

	/*
	two strings can be concatenated with the + operator. But the compiler will not allow you to concatenate
	a string variable with other types
	*/

	const pi = 1.141592
	/*
	constants are declared with the const keyword. They can't use the := short declaration syntax.
	constants can be primitive types like strings, integers, booleans, and floats, but not the more complex types.
	constants can't be changed after it has been declared.
	constants must be known or computed at compile time.
	*/

	s := fmt.Sprintf("I am %v years old", 10)
	s := fmt.Sprintf("I am %v years old", "way too many")
	fmt.Printf("I'm gay")
	/*
	Sprintf returns the formatted string
	*/

	/*
	Strings are just sequence of bytes, they can hold arbitrary data.
	However, go also has a special type, rune, which is an alias for int32, which is 32 bits. 4 bytes
	Go uses utf-8 encoding, which is a variable length encoding.
	*/
	




}


