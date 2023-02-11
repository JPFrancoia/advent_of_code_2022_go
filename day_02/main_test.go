package main

import (
	"bufio"
	"strings"
	"testing"
)


func TestComputePlayersScore(t *testing.T) {

	rounds := [3][2]byte{{'A', 'Y'}, {'B', 'X'}, {'C', 'Z'}}
	targets := [3]int{8, 1, 6}

	for i := 0; i < len(rounds); i++ {
		c1, c2 := rounds[i][0], rounds[i][1]
		_, score := computePlayersScore(c1, c2)
		if score != targets[i] {
			t.Errorf("%c vs %c; expected %d got %d", c1, c2, targets[i], score)
		}
	}
}

func TestComputeTotalScore(t *testing.T) {

	input := "A Y\nB X\nC Z"

	scanner := bufio.NewScanner(strings.NewReader(input))
	total := computeTotalScore(scanner)

	target := 15

	if total != target {
		t.Errorf("Total; expected %d got %d", target, total)
	}
}


func TestComputeNonSupisciousTotalScore(t *testing.T) {
	input := "A Y\nB X\nC Z"

	scanner := bufio.NewScanner(strings.NewReader(input))
	total := computeNonSupisciousTotalScore(scanner)

	target := 12

	if total != target {
		t.Errorf("Total; expected %d got %d", target, total)
	}

}
