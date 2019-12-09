package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func terminate(band *[]int, pos *int) int {
	debugprint("TERM")
	return (*band)[0]
}

func add(band *[]int, pos *int, modes int) {
	arg1 := getArgument(band, pos, modes, 1)
	arg2 := getArgument(band, pos, modes, 2)
	arg3 := putArgument(band, pos, modes, 3)
	debugprint("ADDI", (*band)[*pos:*pos+4], arg1, arg2, arg3)
	writeAddr(band, &arg3, arg1+arg2)
	//(*band)[arg3] = arg1 + arg2
	*pos += 4
	return
}

func mul(band *[]int, pos *int, modes int) {
	arg1 := getArgument(band, pos, modes, 1)
	arg2 := getArgument(band, pos, modes, 2)
	arg3 := putArgument(band, pos, modes, 3)
	debugprint("MULT", (*band)[*pos:*pos+4], arg1, arg2, arg3)
	writeAddr(band, &arg3, arg1*arg2)
	//(*band)[arg3] = arg1 * arg2
	*pos += 4
	return
}

func jumpIfTrue(band *[]int, pos *int, modes int) {
	arg1 := getArgument(band, pos, modes, 1)
	arg2 := getArgument(band, pos, modes, 2)
	debugprint("JUMP", (*band)[*pos:*pos+3], arg1, arg2)
	if arg1 != 0 {
		*pos = arg2
	} else {
		*pos += 3
	}
	return
}

func jumpIfFalse(band *[]int, pos *int, modes int) {
	arg1 := getArgument(band, pos, modes, 1)
	arg2 := getArgument(band, pos, modes, 2)
	debugprint("JUMF", (*band)[*pos:*pos+3], arg1, arg2)
	if arg1 == 0 {
		*pos = arg2
	} else {
		*pos += 3
	}
	return
}

func getArgument(band *[]int, pos *int, modes int, i int) int {
	if i < 1 {
		log.Fatal("argument position must be at least 1")
	}
	newPos := (*pos) + i
	value := readAddr(band, &newPos)
	modes /= int(math.Pow(10, float64(i-1)))
	modes %= 10
	var out int
	if modes == 0 {
		//positional mode
		out = readAddr(band, &value)
		//fmt.Println("Getting Argument",i,"as Positional",value,out)
	} else if modes == 1 {
		out = value
		//fmt.Println("Getting Argument",i,"as Immediate:",out)
		//immediate mode
	} else {
		value += *relativeBase
		out = readAddr(band, &value)
		//fmt.Println("Getting Argument",i,"as Relative:",out)
		//relative mode
	}
	return out
}

func putArgument(band *[]int, pos *int, modes int, i int) int {
	if i < 1 {
		log.Fatal("argument position must be at least 1")
	}
	newPos := (*pos) + i
	value := readAddr(band, &newPos)
	modes /= int(math.Pow(10, float64(i-1)))
	modes %= int(math.Pow(10, float64(i)))
	var out int
	if modes == 0 {
		//positional mode
		out = value
		//fmt.Println("Getting Argument",i,"as Positional",value,out)
	} else if modes == 1 {
		//immediate mode
		log.Fatal("Trying to write in Immediate Mode")
		//fmt.Println("Getting Argument",i,"as Immediate:",out)
	} else {
		//relative mode
		out = *relativeBase + value
		//fmt.Println("Getting Argument",i,"as Relative:",out)
	}
	return out
}

func lessThan(band *[]int, pos *int, modes int) {
	arg1 := getArgument(band, pos, modes, 1)
	arg2 := getArgument(band, pos, modes, 2)
	arg3 := putArgument(band, pos, modes, 3)

	debugprint("LESS", (*band)[*pos:*pos+4], arg1, arg2, arg3)
	if arg1 < arg2 {
		writeAddr(band, &arg3, 1)
		//(*band)[arg3] = 1
	} else {
		writeAddr(band, &arg3, 0)
		//(*band)[arg3] = 0
	}
	*pos += 4
	return
}

func equals(band *[]int, pos *int, modes int) {
	arg1 := getArgument(band, pos, modes, 1)
	arg2 := getArgument(band, pos, modes, 2)
	arg3 := putArgument(band, pos, modes, 3)

	debugprint("EQUA", (*band)[*pos:*pos+4], arg1, arg2, arg3)
	if arg1 == arg2 {
		writeAddr(band, &arg3, 1)
	} else {
		writeAddr(band, &arg3, 0)
	}
	*pos += 4
	return
}

func setBase(band *[]int, pos *int, modes int) {
	arg1 := getArgument(band, pos, modes, 1)
	debugprint("SETB", (*band)[*pos:*pos+2], arg1)
	*relativeBase += arg1
	*pos += 2
	return
}

func bandInput(band *[]int, pos *int, modes int) {
	arg1 := putArgument(band, pos, modes, 1)
	debugprint("READ", (*band)[*pos:*pos+2], arg1)
	value := pop(inputBand)
	writeAddr(band, &arg1, value)
	*pos += 2
	return
}

func bandOutput(band *[]int, pos *int, modes int) {
	arg1 := getArgument(band, pos, modes, 1)
	debugprint("WRIT", (*band)[*pos:*pos+2], arg1)
	*outputBand = append(*outputBand, arg1)
	*pos += 2
	return
}

func input(band *[]int, pos *int, modes int) {
	arg1 := putArgument(band, pos, modes, 1)
	debugprint("READ", (*band)[*pos:*pos+2], arg1)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("-> ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	value, err := strconv.Atoi(text)
	check(err)
	writeAddr(band, &arg1, value)
	//(*band)[arg1] = value
	*pos += 2
	return
}

func output(band *[]int, pos *int, modes int) {
	arg1 := getArgument(band, pos, modes, 1)
	debugprint("WRIT", (*band)[*pos:*pos+2], arg1)
	fmt.Println(arg1)
	*pos += 2
	return
}

func parseOpCode(band *[]int, pos *int) (int, int) {
	fullcode := (*band)[*pos]
	//fmt.Println(fullcode)
	opcode := fullcode % 100
	modes := fullcode / 100
	return opcode, modes
}

func writeAddr(band *[]int, pos *int, value int) {
	if len(*band) <= *pos {
		for i := len(*band); i <= *pos; i++ {
			*band = append(*band, 0)
		}
	}
	(*band)[*pos] = value
	return
}

func readAddr(band *[]int, pos *int) int {
	//fmt.Println("value at",*pos,(*band)[*pos])
	if len(*band) < *pos {
		for i := len(*band); i <= *pos; i++ {
			*band = append(*band, 0)
		}
	}
	return (*band)[*pos]
}

func debugprint(a ...interface{}) {
	if globalDebug {
		fmt.Println(a)
	}
}

func run(band []int) int {
	*relativeBase = 0
	iP := 0
	debugprint(band)
	for true {
		opcode, modes := parseOpCode(&band, &iP)
		if modes > 99 {
			debugBreak = true

		}
		if debugBreak {
			debugprint()
		}
		switch opcode {
		case 99:
			return terminate(&band, &iP)
		case 1:
			add(&band, &iP, modes)
			break
		case 2:
			mul(&band, &iP, modes)
			break
		case 3:
			if bandMode {
				bandInput(&band, &iP, modes)
			} else {
				input(&band, &iP, modes)
			}
			break
		case 4:
			if bandMode {
				bandOutput(&band, &iP, modes)
			} else {
				output(&band, &iP, modes)
			}
			break
		case 5:
			jumpIfTrue(&band, &iP, modes)
			break
		case 6:
			jumpIfFalse(&band, &iP, modes)
			break
		case 7:
			lessThan(&band, &iP, modes)
			break
		case 8:
			equals(&band, &iP, modes)
			break
		case 9:
			setBase(&band, &iP, modes)
			break
		default:
			log.Fatal("Unsupported OpCode")
			return -1
		}
	}
	return -2 //this should NEVER happen
}

func pop(a *[]int) int {
	var x int
	x, *a = (*a)[0], (*a)[1:]
	return x
}

var relativeBase *int = new(int)
var inputBand *[]int = new([]int)
var outputBand *[]int = new([]int)

/** settings **/
var globalDebug = false
var debugBreak = false
var bandMode = true

/** executive stuff below **/

func part1() {
	band := orgBand()
	band[1] = 12
	band[2] = 2
	fmt.Println("Solution: ", run(band))
}

func part2() {
	band := orgBand()
	i := 0
	j := 0
	for run(band) != 19690720 {
		if j > 99 {
			log.Fatal("Out of Bounds")
			return
		}
		if i == 99 {
			i = 0
			j++
		} else {
			i++
		}
		band = orgBand()
		band[1] = i
		band[2] = j
	}
	fmt.Println("Solution: ", 100*i+j)

}

func day5() {
	band := newBand()
	run(band)
}

func day7Band() []int {
	return []int{3, 8, 1001, 8, 10, 8, 105, 1, 0, 0, 21, 34, 51, 64, 81, 102, 183, 264, 345, 426, 99999, 3, 9, 102, 2, 9, 9, 1001, 9, 4, 9, 4, 9, 99, 3, 9, 101, 4, 9, 9, 102, 5, 9, 9, 1001, 9, 2, 9, 4, 9, 99, 3, 9, 101, 3, 9, 9, 1002, 9, 5, 9, 4, 9, 99, 3, 9, 102, 3, 9, 9, 101, 3, 9, 9, 1002, 9, 4, 9, 4, 9, 99, 3, 9, 1002, 9, 3, 9, 1001, 9, 5, 9, 1002, 9, 5, 9, 101, 3, 9, 9, 4, 9, 99, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 99, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 99, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 99}
}

func day7() {
	var ra, rb, rc, rd, re int
	out := 0
	for a := 0; a < 5; a++ {
		*inputBand = []int{a, 0}
		run(day7Band())
		ra = pop(outputBand)
		for b := 0; b < 5; b++ {
			if a == b {
				continue
			}
			*inputBand = []int{b, ra}
			run(day7Band())
			rb = pop(outputBand)
			for c := 0; c < 5; c++ {
				if a == c || b == c {
					continue
				}
				*inputBand = []int{c, rb}
				run(day7Band())
				rc = pop(outputBand)
				for d := 0; d < 5; d++ {
					if a == d || b == d || c == d {
						continue
					}
					*inputBand = []int{d, rc}
					run(day7Band())
					rd = pop(outputBand)
					for e := 0; e < 5; e++ {
						if a == e || b == e || c == e || d == e {
							continue
						}
						*inputBand = []int{e, rd}
						run(day7Band())
						re = pop(outputBand)
						if re > out {
							out = re
						}
					}
				}
			}
		}
	}
	fmt.Println("BoostSignal:", out)
}

func day9() {
	band := boostBand()
	//fmt.Println(len(band))
	//band := []int{109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99}
	run(band)
}

func orgBand() []int {
	return []int{1, 12, 2, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 1, 10, 19, 1, 6, 19, 23, 1, 10, 23, 27, 2, 27, 13, 31, 1, 31, 6, 35, 2, 6, 35, 39, 1, 39, 5, 43, 1, 6, 43, 47, 2, 6, 47, 51, 1, 51, 5, 55, 2, 55, 9, 59, 1, 6, 59, 63, 1, 9, 63, 67, 1, 67, 10, 71, 2, 9, 71, 75, 1, 6, 75, 79, 1, 5, 79, 83, 2, 83, 10, 87, 1, 87, 5, 91, 1, 91, 9, 95, 1, 6, 95, 99, 2, 99, 10, 103, 1, 103, 5, 107, 2, 107, 6, 111, 1, 111, 5, 115, 1, 9, 115, 119, 2, 119, 10, 123, 1, 6, 123, 127, 2, 13, 127, 131, 1, 131, 6, 135, 1, 135, 10, 139, 1, 13, 139, 143, 1, 143, 13, 147, 1, 5, 147, 151, 1, 151, 2, 155, 1, 155, 5, 0, 99, 2, 0, 14, 0}
}

func newBand() []int {
	return []int{3, 225, 1, 225, 6, 6, 1100, 1, 238, 225, 104, 0, 1101, 81, 30, 225, 1102, 9, 63, 225, 1001, 92, 45, 224, 101, -83, 224, 224, 4, 224, 102, 8, 223, 223, 101, 2, 224, 224, 1, 224, 223, 223, 1102, 41, 38, 225, 1002, 165, 73, 224, 101, -2920, 224, 224, 4, 224, 102, 8, 223, 223, 101, 4, 224, 224, 1, 223, 224, 223, 1101, 18, 14, 224, 1001, 224, -32, 224, 4, 224, 1002, 223, 8, 223, 101, 3, 224, 224, 1, 224, 223, 223, 1101, 67, 38, 225, 1102, 54, 62, 224, 1001, 224, -3348, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 1, 224, 1, 224, 223, 223, 1, 161, 169, 224, 101, -62, 224, 224, 4, 224, 1002, 223, 8, 223, 101, 1, 224, 224, 1, 223, 224, 223, 2, 14, 18, 224, 1001, 224, -1890, 224, 4, 224, 1002, 223, 8, 223, 101, 3, 224, 224, 1, 223, 224, 223, 1101, 20, 25, 225, 1102, 40, 11, 225, 1102, 42, 58, 225, 101, 76, 217, 224, 101, -153, 224, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 5, 224, 1, 224, 223, 223, 102, 11, 43, 224, 1001, 224, -451, 224, 4, 224, 1002, 223, 8, 223, 101, 6, 224, 224, 1, 223, 224, 223, 1102, 77, 23, 225, 4, 223, 99, 0, 0, 0, 677, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1105, 0, 99999, 1105, 227, 247, 1105, 1, 99999, 1005, 227, 99999, 1005, 0, 256, 1105, 1, 99999, 1106, 227, 99999, 1106, 0, 265, 1105, 1, 99999, 1006, 0, 99999, 1006, 227, 274, 1105, 1, 99999, 1105, 1, 280, 1105, 1, 99999, 1, 225, 225, 225, 1101, 294, 0, 0, 105, 1, 0, 1105, 1, 99999, 1106, 0, 300, 1105, 1, 99999, 1, 225, 225, 225, 1101, 314, 0, 0, 106, 0, 0, 1105, 1, 99999, 8, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 329, 1001, 223, 1, 223, 7, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 344, 101, 1, 223, 223, 108, 677, 677, 224, 1002, 223, 2, 223, 1006, 224, 359, 101, 1, 223, 223, 1107, 226, 677, 224, 1002, 223, 2, 223, 1005, 224, 374, 101, 1, 223, 223, 1008, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 389, 101, 1, 223, 223, 1007, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 404, 1001, 223, 1, 223, 1107, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 419, 1001, 223, 1, 223, 108, 677, 226, 224, 102, 2, 223, 223, 1006, 224, 434, 1001, 223, 1, 223, 7, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 449, 1001, 223, 1, 223, 107, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 464, 101, 1, 223, 223, 107, 677, 226, 224, 102, 2, 223, 223, 1006, 224, 479, 101, 1, 223, 223, 1007, 677, 677, 224, 1002, 223, 2, 223, 1006, 224, 494, 1001, 223, 1, 223, 1008, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 509, 101, 1, 223, 223, 7, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 524, 1001, 223, 1, 223, 1007, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 539, 101, 1, 223, 223, 8, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 554, 101, 1, 223, 223, 1008, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 569, 101, 1, 223, 223, 1108, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 584, 101, 1, 223, 223, 107, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 599, 1001, 223, 1, 223, 1108, 677, 677, 224, 1002, 223, 2, 223, 1006, 224, 614, 1001, 223, 1, 223, 1107, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 629, 1001, 223, 1, 223, 108, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 644, 101, 1, 223, 223, 8, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 659, 101, 1, 223, 223, 1108, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 674, 101, 1, 223, 223, 4, 223, 99, 226}
}

func boostBand() []int {
	return []int{1102, 34463338, 34463338, 63, 1007, 63, 34463338, 63,
		1005, 63, 53,
		1101, 3, 0, 1000,
		109, 988,
		209, 12,
		9, 1000,
		209, 6,
		209, 3,
		203, 0,

		1008, 1000, 1, 63, //
		1005, 63, 65, //
		1008, 1000, 2, 63, //
		1005, 63, 904, //
		1008, 1000, 0, 63, //
		1005, 63, 58, //
		4, 25, // Output 25
		104, 0, // Output 0
		99, // Terminate

		4, 0, 104, 0, 99, 4, 17, 104, 0, 99, 0, 0, 1101, 25, 0, 1016, 1102, 760, 1, 1023, 1102, 1, 20, 1003, 1102, 1, 22, 1015, 1102, 1, 34, 1000, 1101, 0, 32, 1006, 1101, 21, 0, 1017, 1102, 39, 1, 1010, 1101, 30, 0, 1005, 1101, 0, 1, 1021, 1101, 0, 0, 1020, 1102, 1, 35, 1007, 1102, 1, 23, 1014, 1102, 1, 29, 1019, 1101, 767, 0, 1022, 1102, 216, 1, 1025, 1102, 38, 1, 1011, 1101, 778, 0, 1029, 1102, 1, 31, 1009, 1101, 0, 28, 1004, 1101, 33, 0, 1008, 1102, 1, 444, 1027, 1102, 221, 1, 1024, 1102, 1, 451, 1026, 1101, 787, 0, 1028, 1101, 27, 0, 1018, 1101, 0, 24, 1013, 1102, 26, 1, 1012, 1101, 0, 36, 1002, 1102, 37, 1, 1001, 109, 28, 21101, 40, 0, -9, 1008, 1019, 41, 63, 1005, 63, 205, 1001, 64, 1, 64, 1105, 1, 207, 4, 187, 1002, 64, 2, 64, 109, -9, 2105, 1, 5, 4, 213, 1106, 0, 225, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, -9, 1206, 10, 243, 4, 231, 1001, 64, 1, 64, 1105, 1, 243, 1002, 64, 2, 64, 109, -3, 1208, 2, 31, 63, 1005, 63, 261, 4, 249, 1106, 0, 265, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 5, 21108, 41, 41, 0, 1005, 1012, 287, 4, 271, 1001, 64, 1, 64, 1105, 1, 287, 1002, 64, 2, 64, 109, 6, 21102, 42, 1, -5, 1008, 1013, 45, 63, 1005, 63, 307, 1105, 1, 313, 4, 293, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, -9, 1201, 0, 0, 63, 1008, 63, 29, 63, 1005, 63, 333, 1106, 0, 339, 4, 319, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, -13, 2102, 1, 4, 63, 1008, 63, 34, 63, 1005, 63, 361, 4, 345, 1105, 1, 365, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 5, 1201, 7, 0, 63, 1008, 63, 33, 63, 1005, 63, 387, 4, 371, 1105, 1, 391, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 7, 1202, 1, 1, 63, 1008, 63, 32, 63, 1005, 63, 411, 1105, 1, 417, 4, 397, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 20, 1205, -7, 431, 4, 423, 1106, 0, 435, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 2, 2106, 0, -3, 1001, 64, 1, 64, 1105, 1, 453, 4, 441, 1002, 64, 2, 64, 109, -7, 21101, 43, 0, -9, 1008, 1014, 43, 63, 1005, 63, 479, 4, 459, 1001, 64, 1, 64, 1105, 1, 479, 1002, 64, 2, 64, 109, -5, 21108, 44, 43, 0, 1005, 1018, 495, 1105, 1, 501, 4, 485, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, -7, 1205, 9, 517, 1001, 64, 1, 64, 1105, 1, 519, 4, 507, 1002, 64, 2, 64, 109, 11, 1206, -1, 531, 1106, 0, 537, 4, 525, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, -15, 1208, 0, 36, 63, 1005, 63, 557, 1001, 64, 1, 64, 1106, 0, 559, 4, 543, 1002, 64, 2, 64, 109, 7, 2101, 0, -7, 63, 1008, 63, 35, 63, 1005, 63, 581, 4, 565, 1106, 0, 585, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, -3, 21107, 45, 46, 4, 1005, 1015, 607, 4, 591, 1001, 64, 1, 64, 1105, 1, 607, 1002, 64, 2, 64, 109, -16, 2102, 1, 10, 63, 1008, 63, 31, 63, 1005, 63, 631, 1001, 64, 1, 64, 1106, 0, 633, 4, 613, 1002, 64, 2, 64, 109, 1, 2107, 33, 10, 63, 1005, 63, 649, 1106, 0, 655, 4, 639, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 17, 2101, 0, -9, 63, 1008, 63, 31, 63, 1005, 63, 679, 1001, 64, 1, 64, 1106, 0, 681, 4, 661, 1002, 64, 2, 64, 109, -6, 2107, 34, 0, 63, 1005, 63, 703, 4, 687, 1001, 64, 1, 64, 1106, 0, 703, 1002, 64, 2, 64, 109, 5, 1207, -5, 34, 63, 1005, 63, 719, 1105, 1, 725, 4, 709, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, -15, 1202, 6, 1, 63, 1008, 63, 20, 63, 1005, 63, 751, 4, 731, 1001, 64, 1, 64, 1105, 1, 751, 1002, 64, 2, 64, 109, 21, 2105, 1, 5, 1001, 64, 1, 64, 1106, 0, 769, 4, 757, 1002, 64, 2, 64, 109, 5, 2106, 0, 5, 4, 775, 1001, 64, 1, 64, 1106, 0, 787, 1002, 64, 2, 64, 109, -27, 1207, 4, 35, 63, 1005, 63, 809, 4, 793, 1001, 64, 1, 64, 1106, 0, 809, 1002, 64, 2, 64, 109, 13, 2108, 33, -1, 63, 1005, 63, 831, 4, 815, 1001, 64, 1, 64, 1106, 0, 831, 1002, 64, 2, 64, 109, 4, 21107, 46, 45, 1, 1005, 1014, 851, 1001, 64, 1, 64, 1105, 1, 853, 4, 837, 1002, 64, 2, 64, 109, 3, 21102, 47, 1, -3, 1008, 1013, 47, 63, 1005, 63, 875, 4, 859, 1106, 0, 879, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, -9, 2108, 28, 2, 63, 1005, 63, 895, 1106, 0, 901, 4, 885, 1001, 64, 1, 64, 4, 64, 99, 21101, 27, 0, 1, 21102, 1, 915, 0, 1106, 0, 922, 21201, 1, 59074, 1, 204, 1, 99, 109, 3, 1207, -2, 3, 63, 1005, 63, 964, 21201, -2, -1, 1, 21102, 942, 1, 0, 1105, 1, 922, 21201, 1, 0, -1, 21201, -2, -3, 1, 21102, 1, 957, 0, 1105, 1, 922, 22201, 1, -1, -2, 1106, 0, 968, 22102, 1, -2, -2, 109, -3, 2105, 1, 0}
}

func main() {
	day7()
}
