package main

import "fmt"

func main() {
	var arr [2]string

	arr[0] = "Hello"
	arr[1] = "World!"

	fmt.Println(arr[0], arr[1])

	fmt.Println(arr)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	/**
	 * The type []T is a slice with elements of type T. A slice is formed by specifying
	 * two indices, a low and high bound, separated by a colon: a[low : high]
	 * This selects a half-open range which includes the first element,
	 * but excludes the last one.
	 */

	var s []int = primes[1:4]
	fmt.Println(s)
}
