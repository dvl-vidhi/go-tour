package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {

	v := Vertex{1, 2}
	fmt.Println(v)

	/**
	 * Structs fields are accessed using a dot.
	 */
	fmt.Println(v.X)
	fmt.Println(v.Y)

	/**
	 * Struct fields can be accessed through a struct pointer. To access the field X of
	 * a struct when we have the struct pointer p we could write (*p).X.
	 * However, that notation is cumbersome, so the language permits us
	 * instead to write just p.X, without the explicit dereference
	 */

	p := &v
	p.X = 1e9

	fmt.Println(v)
}
