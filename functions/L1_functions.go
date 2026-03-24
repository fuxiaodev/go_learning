package main

import "fmt"
import "errors"
import "os"

func concat(s1 string, s2 string) string {
	return s1 + s2
}

func main() {
	test("Lane,", " happy birthday!")
	test("Zuck,", " hope that Metaverse thing works out")
	test("Go", " is fantastic")

	increment_value()
	
	/*
	A function can return a value that the caller doesn't care about. 
	We can explicitly ignore variables by using an underscore, or more precisely,
	the blank identifier
	*/
	x, _ := getPoint()
	fmt.Println(x)

	mul, div, err := calculator(1, 2)
	fmt.Println("calculator function:", mul, div, err)
	mul, div, err = calculator(1, 0)
	fmt.Println("calculator function:", mul, div, err)

	fmt.Println(aggregate(2, 3, 4, add))
	fmt.Println(aggregate(2, 3, 4, multiply))

	newX, newY, newZ := conversions(func(a int) int {
		return a + a
	}, 1, 2, 3)
	fmt.Printf("conversions results: %v, %v, %v \n", newX, newY, newZ)

	//harryPotterAggregator now holds the returned function, can call it multiple times
	// and it will remember its private doc across calls
	harryPotterAggregator := concatter()
	harryPotterAggregator("Mr.")
	harryPotterAggregator("and")
	harryPotterAggregator("Mrs.")
	harryPotterAggregator("Dursley")
	harryPotterAggregator("of")
	harryPotterAggregator("number")
	harryPotterAggregator("four,")
	harryPotterAggregator("Privet")

	fmt.Println(harryPotterAggregator("Drive"))
	// Mr. and Mrs. Dursley of number four, Privet Drive

	squareFunc := selfMath(multiply2)
	doubleFunc := selfMath(addition)

	fmt.Printf("square of 5 is %v \n", squareFunc(5))
	fmt.Printf("double of 5 is %v \n", doubleFunc(5))
}

func test(s1 string, s2 string) {
	fmt.Println(concat(s1, s2))
}

/*
When multiple arguments are of the same type, and are next to each other in the function signature, 
the type only needs to be declared after the last argument.
*/
func addToDatabase(hp, damage int, name string) {
  // ?
}

/*
Variables in Go are passed by value (except for a few data types we haven't covered yet). 
"Pass by value" means that when a variable is passed into a function, 
that function receives a copy of the variable. 
The function is unable to mutate the caller's original data.
*/

func increment_value(){
    x := 5
    increment(x)

    fmt.Println(x)
    // still prints 5,
    // because the increment function
    // received a copy of x
}

func increment(x int) {
    x++
}

func getPoint() (x int, y int) {
    return 3, 4
}

/*
Named return values are what enable naked returns. 
Use naked returns only in short functions where the purpose of the returned values is obvious.
*/
func getCoords_named() (x, y int) {
	// x and y are initialized with zero values

	return // automatically returns x and y
}

func getCoords_override(flag bool) (x, y int) {
	if flag{
		return 
		// this returns 0, 0 because go initializes x and y as 0
	}
    return 5, 6 // this is explicit, x and y are NOT returned
}

func getCoords() (int, int) {
	var x int
	var y int
	return x, y
}

/*
Named return parameters are great for documenting a function. 
We know what the function is returning directly from its signature, no need for a comment.
Named return parameters are particularly important in longer functions with many return values.
*/
func calculator(a, b int) (mul, div int, err error) {
    if b == 0 {
      return 0, 0, errors.New("can't divide by zero")
    }
    mul = a * b
    div = a / b
	//nil is the zero value of an error.
    return mul, div, nil
}

/*
functions as values
go supports first class and higher order functions
*/
func add(x, y int) int {
	return x + y
}

func multiply(x, y int) int {
	return x * y
}

func aggregate(a, b, c int, arithmetic func(int, int) int) int {
  firstResult := arithmetic(a, b)
  secondResult := arithmetic(firstResult, c)
  return secondResult
}

/*
anonymous functions are useful when defining a function that will only be used once or create a quick closure
*/
func conversions(converter func(int) int, x, y, z int) (int, int, int) {
	convertedX := converter(x)
	convertedY := converter(y)
	convertedZ := converter(z)
	return convertedX, convertedY, convertedZ
}

/*
defer keyword allows a function to be executed automatically just before its enclosing function returns
the deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding
function returns
if you have multiple defer statements, they are executed in last in first out order.
defer is especially useful when a function could end in multiple ways.
In that case defer ensure cleanup always happens
*/
func CreateTempFile() {
	f, _ := os.Create("temp-42.txt")
	defer os.Remove(f.Name()) // executed second
	defer f.Close()           // executed first

	// FPrintln writes to any io.Writer you provide
	fmt.Fprintln(f, "How many roads must a man walk down?")
}

/*
unlike python, go is not function scoped, its block-scoped.
Variables declared inside a block are only accessible within that block, and its nested blocks.
It's a bit unusual, but occasionally you'll see a plain old explicit block. 
It exists for no other reason than to create a new scope.
*/
func blocks(){
	// create a new scope
	{
		age := 15
		fmt.Println("my age is ", age)
	}

	fmt.Println("age variable is out of scope now")
}

/*
a closure is a function that reference variables from outside its own function body.
the function may access and reassign to the referenced variables
why use closure:
1. encapsulation: doc is fully private, nothing except the harryPotterAggregator can access or modify it
2. consistency: doc remains consistent throughout the code
3. multiple independent accumulators: you can create other closures.
*/
func concatter() func(string) string {
	doc := ""
	return func(word string) string {
		doc += word + " "
		return doc
	}
}

/*
currying allows a function with multiple arguments to be transformed into a sequence of functions, 
each taking a single argument
*/
func multiply2(x, y int) int {
	return x * y
}

func addition(x, y int) int {
	return x + y
}

func selfMath(mathFunc func(int, int) int) func(int) int {
	return func(x int) int {
		return mathFunc(x, x)
	}
}