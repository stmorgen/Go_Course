package main

import "fmt"
import "strings"
import "bufio"
import "os"

func main() {

	var s string	

	fmt.Printf("Enter your string:")
											
	scanner := bufio.NewScanner(os.Stdin)
	
	scanner.Scan()
	s = scanner.Text()

	var lowerInput = strings.ToLower(s)
	
	var containsA = strings.Contains(lowerInput, "a")
	var startsWithI = strings.HasPrefix(lowerInput, "i")
	var endsWithN = strings.HasSuffix(lowerInput, "n")
	
	if (containsA && startsWithI && endsWithN) {	
		fmt.Printf("Found!") 
	} else {
		fmt.Printf("Not Found!")
	}
}