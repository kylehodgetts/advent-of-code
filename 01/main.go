package main

import (
	"bufio"
	"container/heap"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func ConstructHeap(scanner *bufio.Scanner) *IntHeap {
	h := &IntHeap{}
	heap.Init(h)

	currentTotal := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			current, _ := strconv.Atoi(line)
			currentTotal += current

			continue
		}

		heap.Push(h, currentTotal)
		currentTotal = 0
	}

	return h
}

func GetTopCaloriesForNElves(scanner *bufio.Scanner, numElves int) int {
	calorieHeap := ConstructHeap(scanner)
	sumCalories := 0
	for i := 0; i < numElves; i++ {
		val := heap.Pop(calorieHeap).(int)
		sumCalories += val
	}

	return sumCalories
}

func main() {
	var numElves int
	flag.IntVar(&numElves, "elves", 3, "-elves 3")
	flag.Parse()

	f, err := os.Open("input.txt")
	if err != nil {
		panic(fmt.Errorf("unable to open file: %s", err))
	}

	fmt.Println(
		GetTopCaloriesForNElves(bufio.NewScanner(f), numElves),
	)
}
