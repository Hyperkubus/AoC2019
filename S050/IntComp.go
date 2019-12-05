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
	return (*band)[0]
}

func add(band *[]int, pos *int, modes int) {
	arg1 := getArgument(band, pos, modes, 1)
	arg2 := getArgument(band, pos, modes, 2)
	arg3 := getArgument(band, pos, 100, 3)
	(*band)[arg3] = arg1 + arg2
	*pos += 4
	return
}

func mul(band *[]int, pos *int, modes int) {
	arg1 := getArgument(band, pos, modes, 1)
	arg2 := getArgument(band, pos, modes, 2)
	arg3 := getArgument(band, pos, 100, 3)
	(*band)[arg3] = arg1 * arg2
	*pos += 4
	return
}

func jumpIfTrue(band *[]int, pos *int, modes int) {
	if getArgument(band, pos, modes, 1) != 0 {
		*pos = getArgument(band, pos, modes, 2)
	} else {
		*pos += 3
	}
	return
}

func jumpIfFalse(band *[]int, pos *int, modes int) {
	if getArgument(band, pos, modes, 1) == 0 {
		*pos = getArgument(band, pos, modes, 2)
	} else {
		*pos += 3
	}
	return
}

func getArgument(band *[]int, pos *int, modes int, i int) int {
	if i < 1 {
		log.Fatal("argument position must be at least 1")
	}
	value := readAddr(*band, (*pos)+i)
	tmp := int(math.Pow(10, float64(i-1)))
	tmp2 := int(math.Pow(10, float64(i)))
	modes /= tmp
	var out int
	if modes%(tmp2) == 0 {
		//positional mode
		out = readAddr(*band, value)
		//fmt.Println("Getting Argument",i,"as Positional",value,out)
	} else {
		out = value
		//fmt.Println("Getting Argument",i,"as Immediate:",out)
		//immediate mode
	}
	return out
}

func lessThan(band *[]int, pos *int, modes int) {
	arg1 := getArgument(band, pos, modes, 1)
	arg2 := getArgument(band, pos, modes, 2)
	arg3 := getArgument(band, pos, 100, 3)

	if arg1 < arg2 {
		(*band)[arg3] = 1
	} else {
		(*band)[arg3] = 0
	}
	*pos += 4
	return
}

func equals(band *[]int, pos *int, modes int) {
	arg1 := getArgument(band, pos, modes, 1)
	arg2 := getArgument(band, pos, modes, 2)
	arg3 := getArgument(band, pos, 100, 3)

	if arg1 == arg2 {
		(*band)[arg3] = 1
	} else {
		(*band)[arg3] = 0
	}
	*pos += 4
	return
}

func input(band *[]int, pos *int, modes int) {
	arg1 := getArgument(band, pos, 1, 1)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("-> ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	value, err := strconv.Atoi(text)
	check(err)
	(*band)[arg1] = value
	*pos += 2
	return
}

func output(band *[]int, pos *int, modes int) {
	arg1 := getArgument(band, pos, modes, 1)
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

func readAddr(band []int, pos int) int {
	//fmt.Println("value at",*pos,(*band)[*pos])
	return (band)[pos]
}

func debugprint(debug bool, line string) {
	if debug {
		fmt.Println(line)
	}
}

func run(band []int) int {
	debug := false
	iP := 0
	debugprint(debug, fmt.Sprint(band))
	for true {
		//fmt.Println(band,iP)
		opcode, modes := parseOpCode(&band, &iP)
		switch opcode {
		case 99:
			debugprint(debug, "Terminate")
			return terminate(&band, &iP)
		case 1:
			debugprint(debug, fmt.Sprint("Add", modes))
			add(&band, &iP, modes)
			break
		case 2:
			debugprint(debug, fmt.Sprint("Mul", modes))
			mul(&band, &iP, modes)
			break
		case 3:
			debugprint(debug, fmt.Sprint("Read", modes))
			input(&band, &iP, modes)
			break
		case 4:
			debugprint(debug, fmt.Sprint("Write", modes))
			output(&band, &iP, modes)
			break
		case 5:
			debugprint(debug, "JumpIfTrue")
			jumpIfTrue(&band, &iP, modes)
			break
		case 6:
			debugprint(debug, "JumpIfFalse")
			jumpIfFalse(&band, &iP, modes)
			break
		case 7:
			debugprint(debug, "lessThan")
			lessThan(&band, &iP, modes)
			break
		case 8:
			debugprint(debug, "Equal")
			equals(&band, &iP, modes)
			break
		default:
			debugprint(debug, "Unsupported OpCode")
			return -1
		}
	}
	return -2 //this should NEVER happen
}

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

func orgBand() []int {
	return []int{1, 12, 2, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 1, 10, 19, 1, 6, 19, 23, 1, 10, 23, 27, 2, 27, 13, 31, 1, 31, 6, 35, 2, 6, 35, 39, 1, 39, 5, 43, 1, 6, 43, 47, 2, 6, 47, 51, 1, 51, 5, 55, 2, 55, 9, 59, 1, 6, 59, 63, 1, 9, 63, 67, 1, 67, 10, 71, 2, 9, 71, 75, 1, 6, 75, 79, 1, 5, 79, 83, 2, 83, 10, 87, 1, 87, 5, 91, 1, 91, 9, 95, 1, 6, 95, 99, 2, 99, 10, 103, 1, 103, 5, 107, 2, 107, 6, 111, 1, 111, 5, 115, 1, 9, 115, 119, 2, 119, 10, 123, 1, 6, 123, 127, 2, 13, 127, 131, 1, 131, 6, 135, 1, 135, 10, 139, 1, 13, 139, 143, 1, 143, 13, 147, 1, 5, 147, 151, 1, 151, 2, 155, 1, 155, 5, 0, 99, 2, 0, 14, 0}
}

func newBand() []int {
	return []int{3, 225, 1, 225, 6, 6, 1100, 1, 238, 225, 104, 0, 1101, 81, 30, 225, 1102, 9, 63, 225, 1001, 92, 45, 224, 101, -83, 224, 224, 4, 224, 102, 8, 223, 223, 101, 2, 224, 224, 1, 224, 223, 223, 1102, 41, 38, 225, 1002, 165, 73, 224, 101, -2920, 224, 224, 4, 224, 102, 8, 223, 223, 101, 4, 224, 224, 1, 223, 224, 223, 1101, 18, 14, 224, 1001, 224, -32, 224, 4, 224, 1002, 223, 8, 223, 101, 3, 224, 224, 1, 224, 223, 223, 1101, 67, 38, 225, 1102, 54, 62, 224, 1001, 224, -3348, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 1, 224, 1, 224, 223, 223, 1, 161, 169, 224, 101, -62, 224, 224, 4, 224, 1002, 223, 8, 223, 101, 1, 224, 224, 1, 223, 224, 223, 2, 14, 18, 224, 1001, 224, -1890, 224, 4, 224, 1002, 223, 8, 223, 101, 3, 224, 224, 1, 223, 224, 223, 1101, 20, 25, 225, 1102, 40, 11, 225, 1102, 42, 58, 225, 101, 76, 217, 224, 101, -153, 224, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 5, 224, 1, 224, 223, 223, 102, 11, 43, 224, 1001, 224, -451, 224, 4, 224, 1002, 223, 8, 223, 101, 6, 224, 224, 1, 223, 224, 223, 1102, 77, 23, 225, 4, 223, 99, 0, 0, 0, 677, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1105, 0, 99999, 1105, 227, 247, 1105, 1, 99999, 1005, 227, 99999, 1005, 0, 256, 1105, 1, 99999, 1106, 227, 99999, 1106, 0, 265, 1105, 1, 99999, 1006, 0, 99999, 1006, 227, 274, 1105, 1, 99999, 1105, 1, 280, 1105, 1, 99999, 1, 225, 225, 225, 1101, 294, 0, 0, 105, 1, 0, 1105, 1, 99999, 1106, 0, 300, 1105, 1, 99999, 1, 225, 225, 225, 1101, 314, 0, 0, 106, 0, 0, 1105, 1, 99999, 8, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 329, 1001, 223, 1, 223, 7, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 344, 101, 1, 223, 223, 108, 677, 677, 224, 1002, 223, 2, 223, 1006, 224, 359, 101, 1, 223, 223, 1107, 226, 677, 224, 1002, 223, 2, 223, 1005, 224, 374, 101, 1, 223, 223, 1008, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 389, 101, 1, 223, 223, 1007, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 404, 1001, 223, 1, 223, 1107, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 419, 1001, 223, 1, 223, 108, 677, 226, 224, 102, 2, 223, 223, 1006, 224, 434, 1001, 223, 1, 223, 7, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 449, 1001, 223, 1, 223, 107, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 464, 101, 1, 223, 223, 107, 677, 226, 224, 102, 2, 223, 223, 1006, 224, 479, 101, 1, 223, 223, 1007, 677, 677, 224, 1002, 223, 2, 223, 1006, 224, 494, 1001, 223, 1, 223, 1008, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 509, 101, 1, 223, 223, 7, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 524, 1001, 223, 1, 223, 1007, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 539, 101, 1, 223, 223, 8, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 554, 101, 1, 223, 223, 1008, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 569, 101, 1, 223, 223, 1108, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 584, 101, 1, 223, 223, 107, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 599, 1001, 223, 1, 223, 1108, 677, 677, 224, 1002, 223, 2, 223, 1006, 224, 614, 1001, 223, 1, 223, 1107, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 629, 1001, 223, 1, 223, 108, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 644, 101, 1, 223, 223, 8, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 659, 101, 1, 223, 223, 1108, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 674, 101, 1, 223, 223, 4, 223, 99, 226}
}

func main() {
	day5()
}
