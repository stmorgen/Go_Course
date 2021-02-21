package main

import "fmt"

func main() {

	go fmt.Printf("first\n")
	go fmt.Printf("second\n")
	
	fmt.Printf("main\n")
	
	// with the go keyword additional go routines are started (additional to the main goroutine)
	// if the code would be executed sequentially it would print first - second - main
	// but because of interleaving of the routines the order can differ with every run
	// it also might be the case that the main goroutine is finished first and therefore the others aren't executed at all
}