package main

import (
	"bufio"
	"container/heap"
	"log"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

func main() {
	dat, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	mostCal := findElfMostCalories(scanner)

	scanner = bufio.NewScanner(strings.NewReader(string(dat)))
	total := findTopThreeTotal(scanner)

	log.Default().Println(mostCal)
	log.Default().Println(total)
}

func findElfMostCalories(scanner *bufio.Scanner) int {

	maxCal := 0
	curCal := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			maxCal = Max(curCal, maxCal)
			curCal = 0
			continue
		}

		cal, _ := strconv.Atoi(line)
		curCal += cal
	}

	return maxCal
}

func findTopThreeTotal(scanner *bufio.Scanner) int {

	h := &MyHeap[int]{}
	heap.Init(h)
	curCal := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			heap.Push(h, curCal)
			if len(*h) > 3 {
				heap.Pop(h)
			}

			curCal = 0
			continue
		}

		cal, _ := strconv.Atoi(line)
		curCal += cal
	}

	// Last block that wasn't processed
	heap.Push(h, curCal)
	if len(*h) > 3 {
		heap.Pop(h)
	}

	// log.Println(h)

	total := 0

	for i := 0; i < 3; i++ {
		total += heap.Pop(h).(int)
	}

	return total
}

func Max[T constraints.Ordered](a, b T) T {
	if a < b {
		return b
	}
	return a
}

type MyHeap[T constraints.Ordered] []T

func (h MyHeap[T]) Len() int           { return len(h) }
func (h MyHeap[T]) Less(i, j int) bool { return h[i] < h[j] }
func (h MyHeap[T]) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MyHeap[T]) Push(x any) {
	*h = append(*h, x.(T))
}

func (h *MyHeap[T]) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
