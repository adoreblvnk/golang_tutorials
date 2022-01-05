package main

import "fmt"

func main() {
	// floating points
	fmt.Printf("Number: %e \n", 2.23235235) // scientific notation
	fmt.Printf("Number: %f \n", 2.23235235) // decimal no exponent
	fmt.Printf("Number: %g \n", 2.23235235) // large exponent

	// strings
	fmt.Printf("String: %s \n", "blvnk") // default
	fmt.Printf("String: %q \n", "blvnk") // double quoted string

	// width & precision
	fmt.Printf("String: %9s \n", "blvnk")  // width 9
	fmt.Printf("String: %-9s \n", "blvnk") // width -9
	fmt.Printf("Number: %.2f \n", 3.1415)  // precision to 2

	// Sprintf
	var random string = fmt.Sprintf("Number: %07d is cool", 45)
	fmt.Println(random)
}
