package L9

import (
	"fmt"
)
/*
the append function changes the underlying array of its parameter 
and returns a new slice
this means that using append on anything other than itself is usually a bad idea
*/

func will_not_break() {
	a := make([]int, 3)
	fmt.Println("len of a:", len(a))
	fmt.Println("cap of a:", cap(a))
	// len of a: 3
	// cap of a: 3

	b := append(a, 4)
	fmt.Println("appending 4 to b from a")
	fmt.Println("b:", b)
	fmt.Println("addr of b:", &b[0])
	// appending 4 to b from a
	// b: [0 0 0 4]
	// addr of b: 0x44a0c0

	c := append(a, 5)
	fmt.Println("appending 5 to c from a")
	fmt.Println("addr of c:", &c[0])
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	// appending 5 to c from a
	// addr of c: 0x44a180
	// a: [0 0 0]
	// b: [0 0 0 4]
	// c: [0 0 0 5]

/*
in this case, the addr of b and c are different because
a's underlying array does not have any capacity left
its capacity is set to 3, same as its length
the append function only create a new array when there isn't any capacity left
*/
}

func will_break() {
	i := make([]int, 3, 8)
	fmt.Println("len of i:", len(i))
	fmt.Println("cap of i:", cap(i))
	// len of i: 3
	// cap of i: 8

	j := append(i, 4)
	fmt.Println("appending 4 to j from i")
	fmt.Println("j:", j)
	fmt.Println("addr of j:", &j[0])
	// appending 4 to j from i
	// j: [0 0 0 4]
	// addr of j: 0x454000

	g := append(i, 5)
	fmt.Println("appending 5 to g from i")
	fmt.Println("addr of g:", &g[0])
	fmt.Println("i:", i)
	fmt.Println("j:", j)
	fmt.Println("g:", g)
	// appending 5 to g from i
	// addr of g: 0x454000
	// i: [0 0 0]
	// j: [0 0 0 5]
	// g: [0 0 0 5]

/*
in this case, j and g ended up having different address
this is because i was allocated to have a capacity of 8
which means it can append 5 more items before the it creates and returns a new underlying array
so j and g are just two pointers pointing to the same underlying array

If len(i) doesn’t change and cap(i) > len(i), then
append(i, value) will keep writing to index len(i) of the same underlying array.
*/
}

