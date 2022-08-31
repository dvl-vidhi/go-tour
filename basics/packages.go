// main package
package main

/*
	Groups the imports into a parenthesized,
	"factored" import statement.
*/
import (
	"fmt"
	"math"
)

func exported_names() float32 {
	/*
		In Go, a name is exported if it begins with a capital letter.
		Pi is exported from the math package
	*/
	return math.Pi
}

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

	fmt.Printf("Exported value of Pi %g\n", exported_names())
}
