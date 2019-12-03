package day3

import (
	"github.com/fluffybucket/aoc19/helpers"
	"strings"
	"fmt"
	"strconv"
	"math"
)

type pos struct {
	y int
	x int
}

type line struct {
	s pos
	e pos
}

func Part1() {
	fmt.Println("Day3")

	input := strings.Split(helpers.LoadFile("day3/input.txt"), "\n")
	line1Inst := strings.Split(input[0], ",")
	line2Inst := strings.Split(input[1], ",")

	var lines1 []line
	linePos1 := pos{0, 0}
	for _, inst := range line1Inst {
		//fmt.Println(inst)
		steps, err := strconv.Atoi(strings.Trim(inst[1:], " \r"))
		if err != nil {
			panic(err)
		}
		//fmt.Println(steps)
		var x, y int
		switch inst[0] {
		case 'U':
			y = steps
		case 'R':
			x = steps
		case 'D':
			y = -steps
		case 'L':
			x = -steps
		}
		tmp := linePos1
		linePos1.x += x
		linePos1.y += y
		lines1 = append(lines1, line{tmp, linePos1})
	}

	var intersections []pos
	var lines2 []line
	linePos2 := pos{0, 0}
	for _, inst := range line2Inst {
		//fmt.Println(inst)
		steps, err := strconv.Atoi(strings.Trim(inst[1:], " \r"))
		if err != nil {
			panic(err)
		}
		//fmt.Println(steps)
		var x, y int
		switch inst[0] {
		case 'U':
			y = steps
		case 'R':
			x = steps
		case 'D':
			y = -steps
		case 'L':
			x = -steps
		}
		tmp := linePos2
		linePos2.x += x
		linePos2.y += y
		intersections = append(intersections, intersects(line{tmp, linePos2}, lines1)...)
		lines2 = append(lines2, line{tmp, linePos2})
	}
	//yuck
	if len(intersections) > 0 {
		min := math.Abs(float64(intersections[1].x)) + math.Abs(float64(intersections[1].y))
		for _, i := range intersections[1:] {
			d := math.Abs(float64(i.x)) + math.Abs(float64(i.y))
			if d < min {
				min = d
			}
		}
		fmt.Println("Part1")
		fmt.Println(min)

		minD := distanceTo(lines1,intersections[1]) + distanceTo(lines2,intersections[1])
		minI := intersections[1]
		for _,i := range intersections[1:]{
			d := distanceTo(lines1,i) + distanceTo(lines2,i)
			if d < minD {
				minD = d
				minI = i
			}
			//fmt.Println(distanceTo(lines1,i))
			//fmt.Println(distanceTo(lines2,i))
		}

		fmt.Println("Part2")
		fmt.Println(minI,minD)
	}

}

// Returns a list of pos where all the intersections for a line are
func intersects(l line, lines []line) []pos {
	var list []pos

	for _, l2 := range lines {
		intersection, b := intersect(l, l2)
		if b {
			//fmt.Println("These intersect",l,l2,intersection)
			list = append(list, intersection)
		}
	}

	return list
}

// Can not handle the same line intersecting at multiple positions :/
func intersect(l1, l2 line) (pos, bool) {
	if l1.s.x == l1.e.x {
		// If they are both vertical lines
		if l2.s.x == l2.e.x && l2.s.x != l1.s.x {
			return pos{0, 0}, false
		}
		// Its to the left
		if l2.s.x < l1.s.x && l2.e.x < l1.s.x {
			return pos{0, 0}, false
		}
		// Its to the right
		if l2.s.x > l1.s.x && l2.e.x > l1.e.x {
			return pos{0, 0}, false
		}
		// Check if l2 y cord is within l1's range
		if l2.s.y <= l1.s.y && l2.s.y >= l1.e.y {
			return pos{l2.s.y, l1.s.x}, true
		} else if l2.s.y <= l1.e.y && l2.s.y >= l1.s.y {
			return pos{l2.s.y, l1.s.x}, true
		}
	} else if l1.s.y == l1.e.y {
		// If they both are horizontal lines
		if l2.s.y == l2.e.y && l2.s.y != l1.s.y {
			return pos{0, 0}, false
		}
		// Its above
		if l2.s.y < l1.s.y && l2.e.y < l1.s.y {
			return pos{0, 0}, false
		}
		// Its below
		if l2.s.y > l1.s.y && l2.e.y > l1.e.y {
			return pos{0, 0}, false
		}
		// Check if l2 x cord is within l1's range
		if l2.s.x <= l1.s.x && l2.s.x >= l1.e.x {
			return pos{l1.s.y, l2.s.x}, true
		} else if l2.s.x <= l1.e.x && l2.s.x >= l1.s.x {
			return pos{l1.s.y, l2.s.x}, true
		}
	}

	return pos{0, 0}, false
}

func onLine(l line, p pos) bool {
	if l.s.x == l.e.x {
		if p.x == l.s.x {
			if p.y <= l.s.y && p.y >= l.e.y {
				return true
			} else if p.y <= l.e.y && p.y >= l.s.y {
				return true
			}
		}
	} else if l.s.y == l.e.y {
		if p.y == l.s.y {
			if p.x <= l.s.x && p.x >= l.e.x {
				return true
			} else if p.x <= l.e.x && p.x >= l.s.x {
				return true
			}
		}
	}
	return false
}

func distance(s,e pos) int {
	return int(math.Abs(float64(s.x - e.x))+ math.Abs(float64(s.y - e.y)))
}

func distanceTo(ls []line,p pos) int {
	steps := 0
	for _,l := range ls {
		if onLine(l,p) {
			//fmt.Println("Is on line",l,p,distance(l.s,p),steps)
			return steps + distance(l.s,p)
		} else {
			//fmt.Println("Not on line",l,distance(l.s,l.e),steps)
			steps += distance(l.s,l.e)
		}
	}

	return steps
}