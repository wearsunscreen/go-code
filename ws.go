package main

import (
	"errors"
	"fmt"
)

type series struct {
	year           int
	opponentScores [7]int
	loser          string
	loserRuns      int
	winner         string
	winnerRuns     int
}

var years = [8]series{
	{1945, [7]int{0, 4, 0, 4, 8, 7, 9}, "Chicago Cubs", 29, "Detroit Tigers", 32},
	{1957, [7]int{3, 2, 12, 5, 0, 3, 0}, "Milwaukee Braves", 23, "New York Yankees", 25},
	{1959, [7]int{11, 3, 1, 4, 1, 3, 0}, "Los Angeles Dodgers", 21, "Chicago White Sox", 23},
	{1925, [7]int{1, 3, 3, 0, 6, 3, 9}, "Washington Senators", 26, "Pittsburgh Pirates", 25},
	{1948, [7]int{1, 1, 0, 1, 11, 3, 0}, "Cleveland Indians", 18, "Boston Braves", 17},
	{1931, [7]int{2, 2, 5, 0, 5, 1, 4}, "Philadelphia Athletics", 22, "St. Louis Cardinals", 19},
	{1940, [7]int{2, 5, 4, 5, 0, 4, 2}, "Detroit Tigers", 28, "Cincinnati Reds", 22},
	{1960, [7]int{6, 3, 0, 3, 5, 0, 10}, "New York Yankees", 55, "Pittsburgh Pirates", 27},
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

	if maxRuns < 1 {
		number = 0
	} else {
		var err1 = errors.New("No value entered")

		fmt.Printf("You have %d runs remaining. How many runs will you score in this game?\n", maxRuns)
		// scan until we get a number
		for err1 != nil {
			_, err := fmt.Scanf("%d", &number)
			err1 = err
			if err1 == nil {
				if number < 0 {
					number = 0
				}
				if number > maxRuns {
					err1 = errors.New("Invalid number of runs")
					fmt.Printf("Enter a number equal to or less than %d\n", maxRuns)
				}
			}
		}
	}

	var result GameResult
	switch {
	case (number > opScore):
		fmt.Printf("You won Game %d: %d to %d.\n", game+1, number, opScore)
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
	// print intro
	fmt.Println("\n\nIn", s.year, s.winner, "scored", s.winnerRuns, "and defeated the", s.loser, "who scored", s.loserRuns, "runs.")
	fmt.Println("Can you repeat this feat?\n")
	results := [7]GameResult{UNPLAYED, UNPLAYED, UNPLAYED, UNPLAYED, UNPLAYED, UNPLAYED, UNPLAYED}
	runsRemaining := s.winnerRuns
	var runsUsed int
	for i := 0; i < 7; i++ {
		fmt.Println("\nGame", i+1)
		results[i], runsUsed = playGame(i, s.opponentScores[i], runsRemaining)
		runsRemaining -= runsUsed

		var wins, losses int
		wins, losses = countWins(results)

		if wins >= 4 {
			fmt.Printf("You won the World Series %d games to %d!\n", wins, losses)
			return true
		} else if losses >= 4 {
			fmt.Printf("You lost the World Series %d games to %d!\n", wins, losses)
			return true
		} else {
			if wins > losses {
				fmt.Printf("You lead the series %d games to %d.\n", wins, losses)
			} else if losses > wins {
				fmt.Printf("You trail in the series %d games to %d.\n", wins, losses)
			} else {
				fmt.Printf("The series tied %d games each.\n", wins)
			}
		}
	}
	return false
}

func main() {
	for i := 0; i < len(years); i++ {
		playSeries(years[i])
	}
}
