package main

import "fmt"
import "bufio"
import "os"
import "strings"

type Animal struct {
	eats  string
	moves string
	says  string
}

func (a *Animal) InitAnimal(eats, moves, says string) {
	a.eats = eats
	a.moves = moves
	a.says = says
}

func (a Animal) Eat() {
	fmt.Println(a.eats)
}

func (a Animal) Move() {
	fmt.Println(a.moves)
}

func (a Animal) Speak() {
	fmt.Println(a.says)
}

func main() {

	var buf string

	reader := bufio.NewReader(os.Stdin)

	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}

	for {
		fmt.Print("> ")
		buf, _ = reader.ReadString('\n')
		bufs := strings.Split(buf, " ")
		xa := bufs[0]
		xo := bufs[1]
		xo = strings.TrimSpace(xo)

		//fmt.Printf("[%s]\n",xa)
		//fmt.Printf("[%s]\n",xo)

		switch xa {
		case "cow":
			switch xo {
			case "eat":
				cow.Eat()
			case "move":
				cow.Move()
			case "speak":
				cow.Speak()
			}
		case "bird":
			switch xo {
			case "eat":
				bird.Eat()
			case "move":
				bird.Move()
			case "speak":
				bird.Speak()
			}
		case "snake":
			switch xo {
			case "eat":
				snake.Eat()
			case "move":
				snake.Move()
			case "speak":
				snake.Speak()
			}
		}

	}

}
