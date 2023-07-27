package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	theirRock     = "A"
	theirPaper    = "B"
	theirScissors = "C"

	myRock     = "X"
	myPaper    = "Y"
	myScissors = "Z"

	// scores
	Loss = 0
	Draw = 3
	Win  = 6

	// part 2
	NeedToLose = "X"
	NeedToDraw = "Y"
	NeedToWin  = "Z"
)

var (
	beats = map[string]string{
		myRock:     theirScissors,
		myPaper:    theirRock,
		myScissors: theirPaper,

		theirRock:     myScissors,
		theirPaper:    myRock,
		theirScissors: myPaper,
	}

	// part 2
	beatenBy = map[string]string{
		theirRock:     myPaper,
		theirPaper:    myScissors,
		theirScissors: myRock,
	}

	scores = map[string]int{
		myRock:     1,
		myPaper:    2,
		myScissors: 3,

		theirRock:     1,
		theirPaper:    2,
		theirScissors: 3,
	}
)

func GetTotalScore(rounds []string) int {
	totalScore := 0

	for _, round := range rounds {
		moves := strings.Split(round, " ")
		theirMove := moves[0]
		myMove := moves[1]

		totalScore += getScoreForRound(myMove, theirMove)
	}

	return totalScore
}

func GetTotalScorePart2(rounds []string) int {
	totalScore := 0

	for _, round := range rounds {
		moves := strings.Split(round, " ")
		theirMove := moves[0]
		myMove := moves[1]

		totalScore += getScoreForRoundPart2(myMove, theirMove)
	}
	return totalScore
}

func getScoreForRound(myMove, theirMove string) int {
	return scores[myMove] + getRoundsResult(myMove, theirMove)
}

func getRoundsResult(myMove, theirMove string) int {

	if beatenMove := beats[myMove]; beatenMove == theirMove {
		return Win
	}

	if beatenMove := beats[theirMove]; beatenMove == myMove {
		return Loss
	}

	return Draw
}

func getScoreForRoundPart2(myMove, theirMove string) int {
	if myMove == NeedToWin {
		beaten := beatenBy[theirMove]
		return Win + scores[beaten]
	}

	if myMove == NeedToLose {
		beaten := beats[theirMove]
		return Loss + scores[beaten]
	}

	return Draw + scores[theirMove]
}

func main() {
	bInput, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Fprint(os.Stderr, fmt.Errorf("unable to read file contents: %w", err))
		os.Exit(1)
	}

	rounds := strings.Split(string(bInput), "\n")
	totalScore := GetTotalScore(rounds)
	fmt.Println("part 1:", totalScore)

	totalScore = GetTotalScorePart2(rounds)
	fmt.Println("part 2:", totalScore)
}
