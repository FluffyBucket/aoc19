package day5

import (
	"fmt"
	"github.com/fluffybucket/aoc19/helpers"
	"strconv"
	"strings"
)

const data = "3,3,1105,-1,9,1101,0,0,12,4,12,99,1"

func Part1() {
	fmt.Println("Day5")
	//fmt.Println(stringToList(data))
	strData := helpers.LoadFile("day5/input.txt")
	//_ = strData
	ops := stringToList(strData)
	//ops := stringToList(data)
	argCount := 4
	// Input = 1 for part 1
	input := 5
	// Input = 5 for part 2
	//input := 5
	output := 0

	for i := 0; i < len(ops); i += argCount {
		//fmt.Println("I am looping! ",i, ops[i])
		op := getOp(ops[i])
		if op[0] == 99 {
			fmt.Println("Quiting",i)
			fmt.Println("Prev inst", ops[i-5:i+1])
			break
		}

		//fmt.Println("Args ",args)
		switch op[0] {
		case 1:
			args := getArgs(i,3, op[1:], ops,true)
			ops[args[2]] = args[0] + args[1]
			argCount = 4
			//fmt.Printf("Putting %d in %d\n", ops[arg3], arg3)
		case 2:
			args := getArgs(i,3, op[1:], ops, true)
			ops[args[2]] = args[0] * args[1]
			argCount = 4
			//fmt.Printf("Putting %d in %d\n", ops[arg3], arg3)
		case 3:
			args := getArgs(i,1,op[1:],ops,true)
			ops[args[0]] = input
			argCount = 2
		case 4:
			args := getArgs(i,1,op[1:],ops,true)
			output = ops[args[0]]
			argCount = 2
		case 5:
			argCount = 3
			args := getArgs(i,2,op[1:],ops,false)
			if args[0] != 0 {
				i = args[1]
				argCount = 0
			}
		case 6:
			argCount = 3
			args := getArgs(i,2,op[1:],ops,false)
			if args[0] == 0 {
				i = args[1]
				argCount = 0
			}
		case 7:
			args := getArgs(i,3,op[1:],ops,true)
			if args[0] < args[1] {
				ops[args[2]] = 1
			} else {
				ops[args[2]] = 0
			}
			argCount = 4
		case 8:
			args := getArgs(i,3,op[1:],ops,true)
			if args[0] == args[1] {
				ops[args[2]] = 1
			} else {
				ops[args[2]] = 0
			}
			argCount = 4
		default:
			fmt.Println("Something went wrong! ", op)
			break
		}

	}
	fmt.Println("Day5 part 1 or 2: ", output)
}

func stringToList(str string) (list []int) {
	strList := strings.Split(str, ",")
	for _, s := range strList {
		op, _ := strconv.Atoi(s)
		list = append(list, op)
	}

	return list
}

// Will always return a op code of length 4
func getOp(op int) []int {
	digits := make([]int, 5)
	num := op
	// Will get the digits in reverse
	// i.e. EDCBA
	for i := 0; i < 5; i++ {
		digits[i] = num % 10
		num = num / 10
	}

	return []int{digits[1]*10 + digits[0], digits[2], digits[3], digits[4]}
}

// Assumes that it is a digit
func toDigit(r rune) int {
	return int(r) - 48
}

// Will return all values that is used
// If output is set to true it will return the address as the last arg
func getArgs(index,num_args int, modes []int, ops []int, output bool) []int {
	args := make([]int, num_args)
	for i := 0; i<num_args; i++ {
		m := modes[i]
		if m == 1 {
			args[i] = ops[index+i+1]
		} else if output && i == num_args - 1 {
			args[i] = ops[index+i+1]
		} else {
			pos := ops[index+i+1]
			args[i] = ops[pos]
		}
	}

	return args
}
