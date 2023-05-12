package main

import "fmt"

type Mammal struct {
	food       string
	locomotion string
	noise      string
}

func (animal Mammal) Eat() {
	fmt.Println("Mammal ate: " + animal.food)
}

func (animal Mammal) Move() {
	fmt.Println("Mammal moved: " + animal.locomotion)
}

func (animal Mammal) Speak() {
	fmt.Println("Mammal made noise: " + animal.noise)
}

func main() {
	fmt.Println("Please enter the name of the animal, the action you want to perform after prompt \">\" in one line as two separate words: ")
	fmt.Println("Use sigkill to end program execution - ctrl/cmd + C")
	bird := Mammal{"worms", "fly", "peep"}
	cow := Mammal{"grass", "walk", "moo"}
	snake := Mammal{"mice", "slither", "hsss"}

	for {
		var animalTypeFromUser, animalSoundFromUser string
		var animal Mammal

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
