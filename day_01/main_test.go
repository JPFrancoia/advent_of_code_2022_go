package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestFindElfMostCalories(t *testing.T) {

	input := "1000\n2000\n3000\n\n4000\n\n5000\n6000\n\n7000\n8000\n9000\n\n10000"

	scanner := bufio.NewScanner(strings.NewReader(input))

	res := findElfMostCalories(scanner)
	target := 24000

	if res != target {
		t.Errorf("Most calories = %d; want %d", res, target)
	}
}


func TestFindTopThreeTotal(t *testing.T) {

	input := "1000\n2000\n3000\n\n4000\n\n5000\n6000\n\n7000\n8000\n9000\n\n10000"

	scanner := bufio.NewScanner(strings.NewReader(input))

	res := findTopThreeTotal(scanner)
	target := 45000

	if res != target {
		t.Errorf("Total top 3 calories = %d; want %d", res, target)
	}

}
