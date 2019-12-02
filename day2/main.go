package day2

import (
	"strings"
	"strconv"
	"fmt"
)

const data = "1,12,2,3,1,1,2,3,1,3,4,3,1,5,0,3,2,13,1,19,1,6,19,23,2,6," +
		"23,27,1,5,27,31,2,31,9,35,1,35,5,39,1,39,5,43,1,43,10,47,2,6,47,51," +
		"1,51,5,55,2,55,6,59,1,5,59,63,2,63,6,67,1,5,67,71,1,71,6,75,2,75,10," +
		"79,1,79,5,83,2,83,6,87,1,87,5,91,2,9,91,95,1,95,6,99,2,9,99,103,2,9,103," +
		"107,1,5,107,111,1,111,5,115,1,115,13,119,1,13,119,123,2,6,123,127,1,5,127," +
		"131,1,9,131,135,1,135,9,139,2,139,6,143,1,143,5,147,2,147,6,151,1,5,151," +
		"155,2,6,155,159,1,159,2,163,1,9,163,0,99,2,0,14,0"

//const data = "1,9,10,3,2,3,11,0,99,30,40,50"
func Main() {
	fmt.Println("Day2")
	//fmt.Println(stringToList(data))
	ops := stringToList(data)
	//fmt.Println(len(ops))
	for i := 0; i < len(ops); i += 4 {
		//fmt.Println("I am looping! ",i, ops[i])
		op := ops[i]
		if op == 99 {
			fmt.Println("Quiting")
			break
		}

		arg1 := ops[i+1]
		arg2 := ops[i+2]
		arg3 := ops[i+3]
		//fmt.Println("Args ",arg1,arg2,arg3)
		switch op {
		case 1:
			ops[arg3] = ops[arg1] + ops[arg2]
			//fmt.Printf("Putting %d in %d\n", ops[arg3], arg3)
		case 2:
			ops[arg3] = ops[arg1] * ops[arg2]
			//fmt.Printf("Putting %d in %d\n", ops[arg3], arg3)
		default:
			fmt.Println("Something went wrong! ",op)
			break
		}


	}
	fmt.Println("Day2 part 1: ",ops[0])
}

func Part2() {
	computeLoop()
}

func computeLoop() int{
	//fmt.Println(stringToList(data))
	memory := stringToList(data)
	ops := make([]int, len(memory))
	for noun := 0; noun < 100; noun++ {
		for verb:= 0; verb < 100; verb++ {
			copy(ops, memory)
			ops[1] = noun
			ops[2] = verb
			//fmt.Println(noun,verb)
			for i := 0; i < len(ops); i += 4 {
				//fmt.Println("I am looping! ",i, ops[i])
				op := ops[i]
				if op == 99 {
					//fmt.Println("Quiting")
					break
				}

				arg1 := ops[i+1]
				arg2 := ops[i+2]
				arg3 := ops[i+3]
				if arg1 >= len(ops) || arg2 >= len(ops) || arg3 >= len(ops) {
					//fmt.Println("Args out of range")
					break
				}
				//fmt.Println("Args ",arg1,arg2,arg3)
				switch op {
				case 1:
					ops[arg3] = ops[arg1] + ops[arg2]
					//fmt.Printf("Putting %d in %d\n", ops[arg3], arg3)
				case 2:
					//fmt.Println(arg1,arg2,arg3,i)
					ops[arg3] = ops[arg1] * ops[arg2]
					//fmt.Printf("Putting %d in %d\n", ops[arg3], arg3)
				default:
					fmt.Println("Something went wrong! ",op, i, noun, verb)
					fmt.Println(ops)
					break
				}
			}
			if noun == 0 && verb == 0 {
				fmt.Println(ops)
			}
			if ops[0] == 19690720 {
				fmt.Println("Day2 part 2: ",100 * noun + verb)
				return 100 * noun + verb
			}
		}
	}
	fmt.Println(memory)

	return 0
}

func stringToList(str string) (list []int) {
	strList := strings.Split(str,",")
	for _,s := range strList {
		op, _ := strconv.Atoi(s)
		list = append(list, op)
	}

	return list
}