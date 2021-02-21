package main

import "fmt"
import "os"
import "bufio"
import "strings"

type Name struct {
	fname string
	lname string
}

func main() {

	var filename string	
	fmt.Printf("Enter the name of your file: ")
	fmt.Scan(&filename)
	
	file, err := os.Open(filename)
	
	if (err != nil) {
		fmt.Printf("Could not find file: %s", filename)
	} else {
		
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		names := make([]Name, 0, 100)
	 
		for scanner.Scan() {
			line := scanner.Text()
			lineParts := strings.Split(line, " ")
			v := Name{fname: lineParts[0], lname: lineParts[1]}
			names = append(names, v)
		}	
		file.Close()
		
		fmt.Printf("The names from the file are: \n")
		
		for _, n := range names {
			fmt.Printf("%s/%s\n", n.fname, n.lname)
		}
	}
}