package main

import "fmt"
import "bufio"
import "os"
import "strings"

type Animal interface {
	Eat ()
	Move ()
	Speak ()
}

type Cow struct {
}
func (a *Cow) Eat () {
	fmt.Printf("grass\n")
}
func (a *Cow) Move () {
	fmt.Printf("walk\n")
}
func (a *Cow) Speak () {
	fmt.Printf("moo\n")
}

type Bird struct {
}
func (a *Bird) Eat () {
	fmt.Printf("worms\n")
}
func (a *Bird) Move () {
	fmt.Printf("fly\n")
}
func (a *Bird) Speak () {
	fmt.Printf("peep\n")
}

type Snake struct {
}
func (a *Snake) Eat () {
	fmt.Printf("mice\n")
}
func (a *Snake) Move () {
	fmt.Printf("slither\n")
}
func (a *Snake) Speak () {
	fmt.Printf("hsss\n")
}

func main() {

	fmt.Printf("Two commands are allowed:\n")
	fmt.Printf("- 'newanimal <name> <type>'\n")
	fmt.Printf("- 'query <name> <methode>'\n")	
	fmt.Printf("\n\n")	
	
	animals := make(map[string]Animal)

	for {		
		fmt.Printf(">")
		scanner := bufio.NewScanner(os.Stdin)	
		scanner.Scan()
		s := scanner.Text()
		var line = strings.ToLower(s)
		lineParts := strings.Split(line, " ")
		
		command := lineParts[0]
		name := lineParts[1]
		
		if(command == "newanimal") {
		
			animalType := lineParts[2]		
			if(animalType == "cow") {
				animals[name] = new(Cow)
				fmt.Printf("Created it!\n")	
			} else if(animalType == "bird") {
				animals[name] = new(Bird)
				fmt.Printf("Created it!\n")	
			} else if(animalType == "snake") {
				animals[name] = new(Snake)
				fmt.Printf("Created it!\n")	
			} else {
				fmt.Printf("The type '%s' is not supported\n", animalType)
			}
			 		
		} else if(command == "query") {
			
			action := lineParts[2]
			tmp := animals[name]
			if(tmp != nil) {				
				if(action == "speak") {
					tmp.Speak()
				} else if(action == "eat") {
					tmp.Eat()
				} else if(action == "move") {
					tmp.Move()
				} else {
					fmt.Printf("The action '%s' is not supported\n", action)
				}
				
			} else {
				fmt.Printf("There is not animal with name %s\n", name)
			}
		} else {
			fmt.Printf("The command '%s' is not supported\n", command)
		}		
	}
}