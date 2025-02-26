//exercise 1

/*package add

// Add takes two integers and returns their sum.
func Add(a, b int) int {
	return a + b
}
*/

//to commit changes:

//git add .
//git commit -m "Added Add function"
//git tag v1.0.0
//git push origin main
//git push origin v1.0.0

//exercise 2

/*package add

// Add takes two integers and returns their sum.
// For more details, visit: https://www.mathsisfun.com/numbers/addition.html.
func Add(a, b int) int {
	return a + b
}
*/

// Package add provides basic mathematical operations with support for generics.
package add

import "golang.org/x/exp/constraints"

// Number is a generic type constraint that allows both integers and floats.
type Number interface {
	constraints.Integer | constraints.Float
}

// Add takes two numbers (integers or floats) and returns their sum.
func Add[T Number](a, b T) T {
	return a + b
}
