package main

import "fmt"
import "encoding/json"
import "os"
import "bufio"

func main() {

	var name string	
	fmt.Printf("Enter a name: ")
	nameScanner := bufio.NewScanner(os.Stdin)	
	nameScanner.Scan()
	name = nameScanner.Text()
	
	var address string	
	fmt.Printf("Enter a address: ")
	addressScanner := bufio.NewScanner(os.Stdin)	
	addressScanner.Scan()
	address = addressScanner.Text()

	valueMap := map[string]string{"name": name, "address": address}
	
	json, _ := json.Marshal(valueMap)
	
	fmt.Printf("Your entered values result in the following json:\n")
	
	os.Stdout.Write(json)
}