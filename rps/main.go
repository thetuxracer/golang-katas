package main

import (
	"fmt"
	"strings"
)


// P R 

// R S

// S P

// P,R,S
var scores = map[string]string{
	"paper" 	: "rock",
	"rock" 		: "scissors",
	"scissor" 	: "paper",
}

func rps(p1 string) string {
	
	fmt.Println("--")
	parts := strings.Split(p1, ", ")
	player_1 := strings.ToLower(parts[0])
	player_2 := strings.ToLower(parts[1])

	fmt.Println(fmt.Sprintf("Between %s and %s", player_1, player_2))

	if player_1 == player_2 {
		return "It's a Tie"
		
	}
	
	if scores[player_1] == player_2 {
		return "Player 1 wins."
	}

	if scores[player_2] == player_1 {
		return "Player 2 wins."
	}
	
	return "Invalid Move."
}

func main()  {
	round_1 := "Rock, Paper"
	round_2 := "Scissor, Paper"
	round_3 := "Paper, Paper"

	fmt.Println(rps(round_1))
	fmt.Println(rps(round_2))
	fmt.Println(rps(round_3))
}