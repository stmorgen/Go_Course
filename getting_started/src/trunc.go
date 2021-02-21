package main

import "fmt"

func main() {

	var x float64	

	fmt.Printf("Input your floating point number:")

	fmt.Scan(&x)
	
	var y int = int(x)	
	
	fmt.Printf("Your truncated number is: %v", y)
}