package L2

/*
interfaces are implemented implicitly. If an interface exists and a type
has the proper methods defined, then the type automatically fulfils that interface

A quick way of checking whether a struct implements an interface is to declare a function
that takes an interface as an argument. if the function can take the struct as an argument
then the struct implements the interface

Implicit interfaces decouple the definition of an interface from its implementation. 
You may add methods to a type and in the process be unknowingly implementing various interfaces, 
and that's okay.
*/
type employee interface {
	getName() string
	getSalary() int
}

type contractor struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}

func (c contractor) getName() string {
	return c.name
}
func (c contractor) getSalary() int {
	return c.hourlyPay * c.hoursPerYear
}

type fullTime struct {
	name   string
	salary int
}

func (ft fullTime) getSalary() int {
	return ft.salary
}

func (ft fullTime) getName() string {
	return ft.name
}
