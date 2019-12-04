package day4

import (
	"fmt"
	"strconv"
)

func Part1() {
	fmt.Println("Day4")
	count1 := 0
	count2 := 0

	for i := 123257; i <= 647015; i++ {
		if isPassword(strconv.Itoa(i)) {
			count1 += 1

		}
		if isPassword2(strconv.Itoa(i)) {
			count2 += 1

		}
	}

	fmt.Println("Part1")
	fmt.Println(count1)

	fmt.Println("Part2")
	fmt.Println(count2)
}

// Fort part 1
func isPassword(pw string) bool {
	runes := []rune(pw)
	hasDouble := false
	ascending := true
	for i := 1; i < len(runes); i++ {
		d := toDigit(runes[i])
		prev := toDigit(runes[i-1])

		if d == prev {
			hasDouble = true
		}
		if d < prev {
			ascending = false
		}
	}

	return hasDouble && ascending
}

// For part 2
func isPassword2(pw string) bool {
	runes := []rune(pw)
	hasDouble := false
	ascending := true
	groups := make(map[int]int,len(runes))
	for i := 0; i < len(runes)-1; i++ {
		d := toDigit(runes[i])
		next := toDigit(runes[i+1])

		if d == next {
			groups[d] += 1
			hasDouble = true
		}
		if d > next {
			ascending = false
		}
	}
	hasGroups := false

	for _,v := range groups {
		// We have to have at least 1 group that is just a pair
		if v == 1 {
			hasGroups = true
		}
	}

	return hasDouble && ascending && hasGroups
}

// Assumes that it is a digit
func toDigit(r rune) int {
	return int(r)-48
}



