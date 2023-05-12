package main

import (
	"fmt"
	"strings"
)

type Animal interface {
	Speak()
	Eat()
	Move()
}

func main() {

	fmt.Println("Please enter the newanimal or query command \">\" in one (ex. newanimal <name> <type> or query <name> <activity>: ")
	fmt.Println("ex. newanimal tom bird, query tom speak. syntax is: newanimal <name> [bird|snake|cow] and query <name> [eat|speak|move]")
	fmt.Println("Use sigkill to end program execution - ctrl/cmd + C")
	bird := Bird{"worms", "fly", "peep"}
	cow := Cow{"grass", "walk", "moo"}
	snake := Snake{"mice", "slither", "hsss"}

	var cmd, name, parameter string
	var animalsMap = make(map[string]Animal)
	for {
		fmt.Print(">")
		if _, err := fmt.Scanln(&cmd, &name, &parameter); err != nil {
			fmt.Println("Error:", err)
			break
		}
		switch cmd {
		case "newanimal":
			if animalsMap[name] == nil {
				var newAnimal Animal
				switch parameter {
				case "bird":
					newAnimal = bird
				case "cow":
					newAnimal = cow
				case "snake":
					newAnimal = snake
				default:
					fmt.Println("unrecognized animal")
					continue
				}
				animalsMap[name] = newAnimal
				fmt.Println("Created it！")
			}
		case "query":
			animal := animalsMap[name]
			if animal == nil {
				fmt.Println("Animal not found!")
				continue
			}
			switch strings.ToUpper(parameter) {
			case "EAT":
				animal.Eat()
			case "MOVE":
				animal.Move()
			case "SPEAK":
				animal.Speak()
			}
		default:
			fmt.Println("Unrecognized command")
			continue
		}
	}

}

type Cow struct {
	food, locomotion, noise string
}

func (a Cow) Eat() {
	fmt.Println(a.food)
}

func (a Cow) Move() {
	fmt.Println(a.locomotion)
}

func (a Cow) Speak() {
	fmt.Println(a.noise)
}

type Bird struct {
	food, locomotion, noise string
}

func (a Bird) Eat() {
	fmt.Println(a.food)
}

func (a Bird) Move() {
	fmt.Println(a.locomotion)
}

func (a Bird) Speak() {
	fmt.Println(a.noise)
}

type Snake struct {
	food, locomotion, noise string
}

func (a Snake) Eat() {
	fmt.Println(a.food)
}

func (a Snake) Move() {
	fmt.Println(a.locomotion)
}

func (a Snake) Speak() {
	fmt.Println(a.noise)
}
