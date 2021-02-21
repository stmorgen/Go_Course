package main

import "fmt"
import "strings"
import "os"
import "bufio"
import "strconv"

func main() {

	data := ReadInput()
	
	BubbleSort(data)
	
	fmt.Print("Your entered list has been sorted to: ", data)
}

func ReadInput() []int {

	fmt.Print("Please enter your up to 10 integer (separated by a single whitespace):")
	scanner := bufio.NewScanner(os.Stdin)	
	scanner.Scan()
	
	inputString := scanner.Text()
	
	return GetIntSliceFromString(inputString)
}

func GetIntSliceFromString(inputString string) []int {
	 
	inputStringParts := strings.Split(inputString, " ")
	
	input := make([]int, len(inputStringParts))
	
	for index, value := range inputStringParts {
		
		convertedValue, _ := strconv.Atoi(value)
		input[index] = convertedValue
	}
	return input
}

func BubbleSort(data []int) {
	
	needsSort := true
	
	for needsSort {
		needsSort = false
		for index, _ := range data {
			
			if (index+1 < len(data) && data[index] > data[index+1]) {
				Swap(data, index)
				needsSort = true
			}
		}
	}
}

func Swap(data []int, index int) {

	temp := data[index]
	data[index] = data[index+1]
	data[index+1] = temp
}