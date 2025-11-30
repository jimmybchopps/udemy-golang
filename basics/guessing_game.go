package basics

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	// Generate random number between 1 and 100
	target := random.Intn(100) +1

	//Welcome message
	fmt.Println("Welcome to the Guessing Game!")
	fmt.Println("I have chosen a number between 1 and 100")
	fmt.Println("Can you guess what it is?")

	var guess int
	for {
		fmt.Println("Enter your guess: ")
		// using just guess would result in value being stored in copy of var and not the var itself
		fmt.Scanln(&guess)

		// Check if guess is correct
		if guess == target {
			fmt.Println("You guessed correctly")
			break
		} else if guess < target {
			fmt.Println("Too low! Try guessing a higher number.")
		} else {
			fmt.Println("Too high! Try guessing a lower number")
		}

	}
}
