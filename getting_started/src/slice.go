package main

import "fmt"
import "strings"
import "strconv"
import "sort"

func main() {

	v := make([]int, 3)

	for {
		fmt.Printf("Enter next number:")
		
		var s string
		fmt.Scan(&s)
		
		if(strings.ToLower(s) == "x") {
			break
		}
		
		y, _ := strconv.Atoi(s)
			
		v = append(v, y)
		
		sort.Ints(v)
		
		fmt.Printf("Currently the slice looks like %d\n", v)
	}
}