package main

import "fmt"
import "bufio"
import "os"
import "strings"

type Animal struct {
	food string
	locomotion string
	noise string
}

func (a *Animal) InitMe(f, l, n string) {
	a.food = f
	a.locomotion = l
	a.noise = n
}

func (a *Animal) Eat () {
	fmt.Printf("%s\n", a.food)
}

func (a *Animal) Move () {
	fmt.Printf("%s\n", a.locomotion)
}

func (a *Animal) Speak () {
	fmt.Printf("%s\n", a.noise)
}

func main() {

	var cow Animal
	cow.InitMe("grass", "walk", "moo")
	var bird Animal
	bird.InitMe("worms", "fly", "peep")
	var snake Animal
	snake.InitMe("mice", "slither", "hsss")
	
	for {
		
		fmt.Printf(">")
		scanner := bufio.NewScanner(os.Stdin)	
		scanner.Scan()
		s := scanner.Text()
		var lowerInput = strings.ToLower(s)
		
		var x Animal
		
		if(strings.HasPrefix(lowerInput, "cow")) {
			x = cow
		} else if(strings.HasPrefix(lowerInput, "bird")) {
			x = bird
		} else if(strings.HasPrefix(lowerInput, "snake")) {
			x = snake
		}
		
		if(strings.HasSuffix(lowerInput, "eat")) {
			x.Eat()
		} else if(strings.HasSuffix(lowerInput, "move")) {
			x.Move()
		} else if(strings.HasSuffix(lowerInput, "speak")) {
			x.Speak()
		}
	}
}