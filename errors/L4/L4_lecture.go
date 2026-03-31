package main
import (
	"fmt"
)

func enrichUser(userID string) (User, error) {
    user, err := getUser(userID)
    if err != nil {
        // fmt.Errorf is GOATed: it wraps an error with additional context
		// it creates a new error
		// %w is used with fmt.Errorf so that the new error remembers the original error
		// you can later unwrap it using errors.Unwrap or errors.Is/errors.As
        // can only wrap one error per fmt.Errorf
		return User{}, fmt.Errorf("failed to get user: %w", err)
    }
    return user, nil
}

/*
there is another way to deal with errors: the panic function
when a function calls panic, the program crashes and prints a stack trace

the panic function yeets control out of the current function and up the call stack
until it reaches a function that defers a recover
if no function calls recover, the goroutine crashes
as a general rule, do not use panic!
*/
func enrichUser_panic(userID string) User {
	user, err := getUser(userID)
	if err != nil {
		panic(err)
	}
	return user
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover from panic", r)
		}
	}()
	// this panics, but the defer/recover block catches it
	//a truly astonishingly bad way to handle errors
	enrichUser("123")
}

//use error values for all "normal" error handling, and if I have a truly unrecoverable
// error, use log.Fatal to print a message and exit the program