package main

import (
	"fmt"
	"strings"
)

var scores = map[string]string{
	"paper" 	: "rock",
	"rock" 		: "scissors",
	"scissor" 	: "paper",
}

func rps(round string) string {
	
	fmt.Println("New round begins.")

	parts := strings.Split(round, ", ")
	player_1 := strings.ToLower(parts[0])
	player_2 := strings.ToLower(parts[1])

	fmt.Println(fmt.Sprintf("Between %s and %s", player_1, player_2))

	if player_1 == player_2 {
		return "It's a Tie"
	}
	
	// Opportunity to make more readable: Move this into an object and replace
	// comparision with a .beats() function call that returns boolean or string.
	
	if scores[player_1] == player_2 {
		return "Player 1 wins."
	}

	if scores[player_2] == player_1 {
		return "Player 2 wins."
	}
	
	return "Invalid Move."
}

func main()  {
	fmt.Println(rps("Rock, Paper"))
	fmt.Println(rps("Scissor, Paper"))
	fmt.Println(rps("Paper, Paper"))
}