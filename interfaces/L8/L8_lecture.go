package L8

import (
	"io"
	"os"
)

/*
1. keep interfaces small
Interfaces are meant to define the minimal behavior necessary to accurately
represent an idea or concept
*/
type File interface {
    io.Closer
    io.Reader
    io.Seeker
    Readdir(count int) ([]os.FileInfo, error)
    Stat() (os.FileInfo, error)
}

/*
2. interfaces should have no knowledge of satisfying types
an interface should define what is necessary for other types to classify as a
member of that interface.
they should not be aware of any types that happen to satisfy the interface at design time
*/
type car interface {
	Color() string
	Speed() int
	// isFiretruck is an anti-pattern, it is designed for a specific type
	IsFiretruck() bool
}

// developer should rely on the native functionality of type assertion to derive
// the underlying type when given an instance of the car interface
// make car as an embedded interface and add one additional required method to make it a firetruck
type firetruck interface {
	car
	HoseLength() int
}


