package main

import "fmt"
import (
	"os"
	"bufio"
	"strconv"
)

func main() {
	day1 := 0
	day2 := 0
	f, _ := os.Open("day1/day1.txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		val,_ := strconv.Atoi(line)
		day1 += (val / 3) - 2

		day2 += fuelCalc(val)
	}

	fmt.Println(day1)
	fmt.Println(day2)
}

func fuelCalc(mass int) int {
	fuel := (mass / 3) - 2
	if fuel < 0 {
		return 0
	}
	return fuel + fuelCalc(fuel)

}
