package L3

import (
	"fmt"
)

// A type can implement any number of interfaces in Go
// e.g., the empty interface us always implemented by every type

func (e email) cost() int {
	if !e.isSubscribed {
		return len(e.body) * 5
	}
	return len(e.body) * 2
}

func (e email) format() string {
	if e.isSubscribed {
		return fmt.Sprintf("'%s' | Subscribed", e.body)
	}
	return  fmt.Sprintf("'%s' | Not Subscribed", e.body)
}

type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

// TODO: email type implements both the expense and the formatter interfaces
type email struct {
	isSubscribed bool
	body         string
}
