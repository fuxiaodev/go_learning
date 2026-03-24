package L2

import "fmt"
/*
because errors are just interfaces, you can build your own custom types that
implement the error interface.
*/
type userError struct {
	name string
}

func (e userError) Error() string {
	return fmt.Sprintf("%v has a problem with their account", e.name)
}

func sendSMS(msg, userName string) error {
	if !canSendToUser(userName) {
		return userError{name: userName}
	}
	return nil
}

func canSendToUser(username string) bool {
	return false
}