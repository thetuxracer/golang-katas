package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func calculateTotalPoints(season_scores []string) int {
	total_points := 0

	if len(season_scores) > 10 {
		log.Fatal("Season scores cannot be longer than 10 games.")
	}

	for i := range season_scores {
		game_score := strings.Split(season_scores[i], ":")
		
		x, err_x := strconv.Atoi(game_score[0])
		y, err_y := strconv.Atoi(game_score[1])

		if err_x != nil || err_y != nil {
			log.Fatal("Error converting string to integer", err_x, err_y)
		}

		if x >= 4 || y >= 4 {
			log.Fatal("Scores cannot be higher than 4.")
		}
		if x > y {
			total_points += 3
		}
		if x < y {
			total_points += 0
		}
		if x == y {
			total_points += 1
		}
	}

	return total_points
}

func main() {
	fmt.Println("Total Points:", calculateTotalPoints([]string{"3:1", "2:2", "0:1", "2:0", "1:3", "2:0", "3:3", "0:0", "1:2", "2:1"}))
}