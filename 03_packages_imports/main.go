package main

// imports brings external packages into the file that we are working on it and
// where actually we need it.
import (
  "fmt"
	"math"
)

func main(){

	// we use packageName.functionName -> call a function from a package.
	// we always use the packageName.functionName like for printing any thing we use
	// fmt.Println("Hello")
	// In this "fmt" is packageName and Println is functionName.

	fmt.Println("Square Root of (25): ", math.Sqrt(25));
	// Here in this case math is packageName and Sqrt is functionName. 

	// OUTPUT : Square Root of (25):  5
  
}