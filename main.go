package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	greet()
	comChoice := rand.IntN(100)
	var chances int
	choice := choice()
	switch choice {
	case 1:
		chances = 10
		fmt.Println("Great! You have selected the Easy difficulty level.")
		fmt.Println("Let's start the game!")
	case 2:
		chances = 5
		fmt.Println("Great! You have selected the Medium difficulty level.")
		fmt.Println("Let's start the game!")
	case 3:
		chances = 3
		fmt.Println("Great! You have selected the Hard difficulty level.")
		fmt.Println("Let's start the game!")
	}

	game(comChoice, chances)
}

func greet() {
	fmt.Print("Welcome to the Number Guessing Game!\n")
	fmt.Print("I'm thinking of a number between 1 and 100.\n")
	fmt.Print("You have 5 chances to guess the correct number.\n\n")

	fmt.Print("Please select the difficulty level:\n")
	fmt.Print("1. Easy (10 chances)\n")
	fmt.Print("2. Medium (5 chances)\n")
	fmt.Print("3. Hard (3 chances) \n\n")

}

func choice() int {
	fmt.Print("Enter Your Choice:")
	var choice int
	fmt.Scan(&choice)
	return choice
}
func game(comp int, chances int) {
	var guess int
	var attempts int
	for {
		attempts += 1
		fmt.Print("Enter your guess:")
		fmt.Scan(&guess)
		if guess > comp {
			fmt.Printf("Incorrect! The number is less than %v \n", guess)
		} else if guess < comp {
			fmt.Printf("Incorrect! The number is greater than %v\n", guess)
		} else if guess == comp {
			fmt.Printf("Congratulations! You guessed the correct number in %v attempts.\n", attempts)
			break
		}
	}
}
