package main

import (
	"fmt"
	"math/rand"
)

func main()  {

	var guessedNumber int
	var numberToGuess int = rand.Intn(100) + 1
	var attempts int
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I have selected a number between 1 and 100. Can you guess it?")

	for {
		fmt.Print("Enter your guess: ")
		_, err := fmt.Scan(&guessedNumber)
		attempts++
		if err != nil {
			fmt.Println("Please enter a valid number.")
			continue
		}

		if guessedNumber < numberToGuess {
			fmt.Println("Too low! Try again.")
		} else if guessedNumber > numberToGuess {
			fmt.Println("Too high! Try again.")
		} else {
			fmt.Printf("Congratulations! You've guessed the number %d in %d attempts.\n", numberToGuess, attempts)
			break
		}
	}

}