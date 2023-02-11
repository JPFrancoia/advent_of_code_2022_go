package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Rock A, X
// Paper B, Y
// Scissors C, Z
var SHAPE_TO_POINTS = map[byte]int{
	'X': 1,
	'Y': 2,
	'Z': 3,
	'A': 1,
	'B': 2,
	'C': 3,
}

var OUTCOME_TO_POINTS = map[string]int{
	"win":  6,
	"draw": 3,
	"lose": 0,
}

func main() {
	dat, err := os.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(strings.NewReader(string(dat)))

	// totalScore := computeTotalScore(scanner)
	// fmt.Println(totalScore)

	totalScore := computeNonSupisciousTotalScore(scanner)
	fmt.Println(totalScore)
}

func computePlayersScore(c1 byte, c2 byte) (int, int) {

	outcomeScore1, outcomeScore2 := 0, 0
	shapeScore1, shapeScore2 := 0, 0

	combi := [2]byte{c1, c2}

	switch combi {
	case [2]byte{'A', 'X'}:
		outcomeScore1, outcomeScore2 = OUTCOME_TO_POINTS["draw"], OUTCOME_TO_POINTS["draw"]
	case [2]byte{'A', 'Y'}:
		outcomeScore1, outcomeScore2 = OUTCOME_TO_POINTS["lose"], OUTCOME_TO_POINTS["win"]
	case [2]byte{'A', 'Z'}:
		outcomeScore1, outcomeScore2 = OUTCOME_TO_POINTS["win"], OUTCOME_TO_POINTS["lose"]

	case [2]byte{'B', 'X'}:
		outcomeScore1, outcomeScore2 = OUTCOME_TO_POINTS["win"], OUTCOME_TO_POINTS["lose"]
	case [2]byte{'B', 'Y'}:
		outcomeScore1, outcomeScore2 = OUTCOME_TO_POINTS["draw"], OUTCOME_TO_POINTS["draw"]
	case [2]byte{'B', 'Z'}:
		outcomeScore1, outcomeScore2 = OUTCOME_TO_POINTS["lose"], OUTCOME_TO_POINTS["win"]

	case [2]byte{'C', 'X'}:
		outcomeScore1, outcomeScore2 = OUTCOME_TO_POINTS["lose"], OUTCOME_TO_POINTS["win"]
	case [2]byte{'C', 'Y'}:
		outcomeScore1, outcomeScore2 = OUTCOME_TO_POINTS["win"], OUTCOME_TO_POINTS["lose"]
	case [2]byte{'C', 'Z'}:
		outcomeScore1, outcomeScore2 = OUTCOME_TO_POINTS["draw"], OUTCOME_TO_POINTS["draw"]
	}

	shapeScore1, shapeScore2 = SHAPE_TO_POINTS[c1], SHAPE_TO_POINTS[c2]

	return outcomeScore1 + shapeScore1, outcomeScore2 + shapeScore2
}

// For part 1
func computeTotalScore(scanner *bufio.Scanner) int {
	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, " ")

		_, score := computePlayersScore(fields[0][0], fields[1][0])

		total += score
	}

	return total
}

// For part 2
func computeNonSupisciousTotalScore(scanner *bufio.Scanner) int {

	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, " ")

		outcome := fields[1][0]
		shapeAdversary := fields[0][0]

		var shape byte

		switch outcome {
		case 'X':
			shape = findShapeForLoss(shapeAdversary)
		case 'Y':
			shape = findShapeForDraw(shapeAdversary)
		case 'Z':
			shape = findShapeForWin(shapeAdversary)
		}

		_, score := computePlayersScore(shapeAdversary, shape)

		total += score

	}

	return total
}

func findShapeForWin(c byte) byte {

	var out byte

	switch c {
	case 'A':
		out = 'Y'
	case 'B':
		out = 'Z'
	case 'C':
		out = 'X'
	}

	return out
}

func findShapeForDraw(c byte) byte {

	var out byte

	switch c {
	case 'A':
		out = 'X'
	case 'B':
		out = 'Y'
	case 'C':
		out = 'Z'
	}

	return out
}

func findShapeForLoss(c byte) byte {

	var out byte

	switch c {
	case 'A':
		out = 'Z'
	case 'B':
		out = 'X'
	case 'C':
		out = 'Y'
	}

	return out
}
