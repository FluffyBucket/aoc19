package day1

import "fmt"
import (
	"os"
	"bufio"
	"strconv"
)

func Main() {
	fmt.Println("Day1")
	day1 := 0
	day1p2 := 0
	f, _ := os.Open("day1/day1.txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		val,_ := strconv.Atoi(line)
		day1 += (val / 3) - 2

		day1p2 += fuelCalc(val)
	}

	fmt.Println(day1)
	fmt.Println(day1p2)
}

func fuelCalc(mass int) int {
	fuel := (mass / 3) - 2
	if fuel < 0 {
		return 0
	}
	return fuel + fuelCalc(fuel)

}
