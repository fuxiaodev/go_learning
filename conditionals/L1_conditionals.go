package main

import "fmt"

func main(){
	height := 5
	if height > 4 {
    	fmt.Println("You are tall enough!")
	}

	if height > 6 {
    	fmt.Println("You are super tall!")
	} else if height > 4 {
    	fmt.Println("You are tall enough!")
	} else {
    	fmt.Println("You are not tall enough!")
	}

	/*
	if statements in Go do not use parentheses around the condition
	You must put the opening brace on the same line as the condition and not on a new line
	initial statement of an if block is used to keep the code concise and the scope limited
	*/
	email := "jj@gmail.com"

	if length := getLength(email); length < 10 {
		fmt.Printf("Email must be at least 10 characters, is %d\n", length)
	}
	
	/*
	in Go, break statement is not required at th end of a case to stop it from falling through to the next case.
	The break statement is implicit in Go
	if you do want a case to fall through, you can use the fallthrough keyword
	*/
	getCreator := func(os string) string {
		switch os {
		case "linux":
			return "Linus Torvalds"
		case "windows":
			return "Bill Gates"
		// all three of these cases will set creator to "A Steve"
		case "macOS":
			fallthrough
		case "Mac OS X":
			fallthrough
		case "mac":
			return "A Steve"
		default:
			return "Unknown"
		}
	}

	fmt.Println(getCreator("macOS"))

}

func getLength(email string) int {
	return len(email)
}

