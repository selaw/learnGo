package main

import (
	"bufio"
	"fmt"
	"os"
	s "strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

type Actions interface {
	Speak()
	Eat()
	Move()
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	cow := Animal{"grass", "walk", "moo"}
	snake := Animal{"mice", "slither", "hsss"}
	bird := Animal{"worms", "fly", "peep"}
	for {
		fmt.Printf(">")
		text, _ := reader.ReadString('\n')
		input := s.Split(s.ToLower(s.TrimSpace(text)), " ")
		var t Actions
		switch input[0] {
		case "cow":
			t = cow
		case "snake":
			t = snake
		case "bird":
			t = bird
		case "exit":
			return
		default:
			fmt.Println("Incorrect Input. Please try again.")
			continue
		}
		switch input[1] {
		case "eat":
			t.Eat()
		case "move":
			t.Move()
		case "speak":
			t.Speak()
		default:
			fmt.Println("Incorrect Input. Please try again.")
			continue
		}
	}
}
