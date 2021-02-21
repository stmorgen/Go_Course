package main

import "fmt"
import "strings"
import "os"
import "bufio"
import "strconv"

func main() {

	data := ReadInput()
	
	c := make(chan []int, 4)
	
	subSize := len(data) / 4
	
	go BubbleSortRoutine(data[0:subSize], c)
	go BubbleSortRoutine(data[subSize:2*subSize], c)
	go BubbleSortRoutine(data[2*subSize:3*subSize], c)
	go BubbleSortRoutine(data[3*subSize:len(data)], c)
	
	combinedSortedParts := CombineSortedParts(len(data), c)
	
	BubbleSort(combinedSortedParts)
	
	fmt.Println("Your entered list has been sorted to: ", combinedSortedParts)
}

func BubbleSortRoutine(data []int, c chan []int) {

	fmt.Println("Goroutine will sort: ", data)
	
	BubbleSort(data)
	
	c <- data
}

func CombineSortedParts(length int, c chan []int) []int {

	result := make([]int, length)
	subSize := length / 4
	
	for partIndex := 0; partIndex < 4; partIndex++ {
	
		receivedSortedPart := <- c		
		for index, value := range receivedSortedPart {
		
			newIndex := partIndex * subSize + index
			result[newIndex] = value
		}
	}
	return result
}

func ReadInput() []int {

	fmt.Print("Please enter your integer values (separated by a single whitespace):")
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