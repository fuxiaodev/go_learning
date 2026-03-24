package L1

/*
go programs express errors with error values. 
An error is any type that implements the simple built-in error interface
when something can go wrong in a function, that function should return an error
as its last return value
any code that calls a function that can return an error should handle errors by testing whether
the error is nil
*/
type error interface {
    Error() string
}