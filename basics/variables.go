package main

import "fmt"

/**
 * The var statement declares a list of variables; as in function argument lists,
 * the type is last. A var statement can be at package or function level.
 * We see both in this example.
 */

var c, python, java bool

/**
 * A var declaration can include initializers, one per variable. If an
 * initializer is present, the type can be omitted; the variable will
 * take the type of the initializer.
 */

var k, j int = 1, 2

func main() {
	var i int
	fmt.Println(i, c, python, java)

	var c, python, java = true, false, "no!"
	fmt.Println(k, j, c, python, java)

	/**
	 * Inside a function, the := short assignment statement
	 * can be used in place of a var declaration with implicit type.
	 */

	l := 3

	fmt.Println(l)

	/**
	 * Variables declared without an explicit initial value are
	 * given their zero value.
	 */

	var ii int
	var ff float64
	var bb bool
	var ss string

	fmt.Printf("%v %v %v %q\n", ii, ff, bb, ss)

}
