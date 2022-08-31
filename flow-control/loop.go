package main

import "fmt"

func main() {

	sum := 0

	/**
	 * Unlike other languages like C, Java, or JavaScript there are no parentheses
	 * surrounding the three components of the for statement and the braces { }
	 * are always required.
	 */

	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	/**
	 * The init and post statements are optional.
	 */

	sum_againg := 1
	for sum_againg < 1000 {
		sum_againg += sum_againg
	}
	fmt.Println(sum_againg)

	/**
	 * We can drop the semicolons: C's while is spelled for in Go.
	 */

	sum_with_while := 1
	for sum_with_while < 1000 {
		sum_with_while += sum_with_while
	}
	fmt.Println(sum_with_while)

	/**
	 * If you omit the loop condition it loops forever,
	 * so an infinite loop is compactly expressed.
	 * for {
	 * }
	 */
}
