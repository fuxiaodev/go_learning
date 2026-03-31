package L5

import (
	"errors"
)

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	switch plan {
	case planPro:
		curr_slice := messages[:]
		return curr_slice, nil
	case planFree:
		curr_slice := messages[:2]
		return curr_slice, nil
	default:
		return nil, errors.New("unsupported plan")		
	}
}

/*
Arrays are fixed in size
Slices are dynamically sized, flexible view of the elements in an array
the zero value of slice is nil
non nil slices always have an underlying array, though it isn't always
specified explicitly
*/