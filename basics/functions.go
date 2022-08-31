package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

/*
 *	When two or more consecutive named function
 *	parameters share a type, you can omit the type from all but the last
 */
func add_v2(x, y int) int {
	return x + y
}

/*
 * A function can return any number of strings.
 * The swap function returns two strings.
 */
func swap(x, y string) (string, string) {
	return y, x
}

/*
 * Go's return values may be named.
 * If so, they are treated as variables defined at the top of the function.
 */
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(add_v2(42, 13))

	//Return two strings
	a, b := swap("Hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))
}
