package main

import "fmt"

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (animal Animal) Eat() {
	fmt.Println("Animal ate: " + animal.food)
}

func (animal Animal) Move() {
	fmt.Println("Animal moved: " + animal.locomotion)
}

func (animal Animal) Speak() {
	fmt.Println("Animal made noise: " + animal.noise)
}

func main() {
	fmt.Println("Please enter the name of the animal, the action you want to perform after prompt \">\" in one line as two separate words: ")
	fmt.Println("Use sigkill to end program execution - ctrl/cmd + C")
	bird := Animal{"worms", "fly", "peep"}
	cow := Animal{"grass", "walk", "moo"}
	snake := Animal{"mice", "slither", "hsss"}

	for {
		var animalTypeFromUser, animalSoundFromUser string
		var animal Animal

		fmt.Print(">")
		if _, err := fmt.Scanln(&animalTypeFromUser, &animalSoundFromUser); err != nil {
			fmt.Println("Error:", err)
			break
		}

		switch animalTypeFromUser {
		case "bird":
			animal = bird
		case "cow":
			animal = cow
		case "snake":
			animal = snake
		default:
			fmt.Println("unrecognized animal")
			continue
		}

		switch animalSoundFromUser {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Println("unrecognized sound")
			continue
		}
	}

}
