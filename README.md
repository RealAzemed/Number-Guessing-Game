# Number Guessing Game Go-Implementation

A simple command-line number guessing game built in Go. The user has to guess a randomly selected number within a limited number of attempts based on the difficulty level chosen.

## Project Information

This project is categorized under **Backend Development** and is suitable for **beginners**. You can find more details about this project on the Roadmap website:

[Number Guessing Game Project](https://roadmap.sh/projects/number-guessing-game)

## Features

- Random number selection between 1 and 100
- User-defined difficulty levels (Easy, Medium, Hard)
- Feedback on guesses indicating whether the number is higher or lower
- Congratulatory messages for correct guesses
- Ends when the user guesses correctly or runs out of attempts

## Project Structure

```
number-guessing-game/
└── main.go
```

## Getting Started

### Prerequisites

- Go (version 1.26 or later)

### Installation

1. Clone the repository:

```bash
git clone https://github.com/yourusername/number-guessing-game.git
cd number-guessing-game
```

2. Run the application:

```bash
go run main.go
```

### How to Play

1. When prompted, select a difficulty level:
- Easy (10 chances)
- Medium (5 chances)
- Hard (3 chances)

2. Enter your guess when prompted.

3. The game will provide feedback on whether your guess is too high or too low.

4. Try to guess the correct number within the given number of attempts.

## Example Output

```bash
Welcome to the Number Guessing Game!
I'm thinking of a number between 1 and 100.
Please select the difficulty level:
1. Easy (10 chances)
2. Medium (5 chances)
3. Hard (3 chances)

Enter your choice: 2

Great! You have selected the Medium difficulty level.
Let's start the game!

Enter your guess: 50
Incorrect! The number is less than 50.

Enter your guess: 25
Incorrect! The number is greater than 25.

Enter your guess: 35
Congratulations! You guessed the correct number in 3 attempts.
```
