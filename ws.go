package main

import (
	"errors"
	"fmt"
)

type series struct {
	year           int
	opponentScores [7]int
	playerRuns     int
}

var years = [2]series{
	{1925, [7]int{1, 2, 3, 4, 5, 6, 7}, 21},
	{1926, [7]int{1, 2, 3, 4, 5, 6, 7}, 21},
}

type GameResult int

const (
	WIN GameResult = iota
	LOSS
	UNPLAYED
)

// ask player for number of runs, determine winner of game, return runs expended by player
func playGame(game, opScore, maxRuns int) (GameResult, int) {
	var number int
	var err1 = errors.New("No value entered")

	fmt.Printf("How many runs will you score in this game? (max %d)\n", maxRuns)

	// scan until we get a number
	for err1 != nil {
		_, err := fmt.Scanf("%d", &number)
		err1 = err
	}

	var result GameResult
	switch {
	case (number > opScore):
		fmt.Printf("You won Game %d: %d to %d\n", game+1, number, opScore)
		result = WIN
	case (number < opScore):
		fmt.Printf("You lost Game %d: %d to %d\n", game+1, number, opScore)
		result = LOSS
	case (number == opScore):
		// in the case of tie, will give the player the win if they have more runs available
		if number < maxRuns {
			number++
			fmt.Printf("You won Game %d in extra innings: %d to %d\n", game+1, number, opScore)
			result = WIN
		} else {
			fmt.Printf("You lost Game %d in extra innings: %d to %d\n", game+1, number, opScore+1)
			result = LOSS
		}
	}
	return result, number
}

func countWins(results [7]GameResult) (int, int) {
	var wins, losses int
	i := 0
	for ; i < 7; i++ {
		if results[i] == WIN {
			wins++
		} else if results[i] == LOSS {
			losses++
		}
	}
	return wins, losses
}

func playSeries(s series) bool {
	results := [7]GameResult{UNPLAYED, UNPLAYED, UNPLAYED, UNPLAYED, UNPLAYED, UNPLAYED, UNPLAYED}
	for i := 0; i < 7; i++ {
		fmt.Println("Play Game", i+1)
		results[i], _ = playGame(i, s.opponentScores[i], 5)

		var wins, losses int
		wins, losses = countWins(results)

		if wins >= 4 {
			fmt.Printf("You won the World Series %d games to %d!\n", wins, losses)
			return true
		} else if losses >= 4 {
			fmt.Printf("You lost the World Series %d games to %d!\n", wins, losses)
			return true
		}
	}
	return false
}

func main() {
	playSeries(years[0])
}
