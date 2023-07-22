package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestGetMostCalories(t *testing.T) {
	input := `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

	sr := strings.NewReader(input)
	scanner := bufio.NewScanner(sr)
	mostCalories := GetTopCaloriesForNElves(scanner, 1)

	if mostCalories != 24000 {
		t.Errorf("Expected 24000 got %d", mostCalories)
	}
}
