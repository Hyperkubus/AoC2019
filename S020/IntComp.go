package main

import (
	"fmt"
	"log"
)

func terminate(band *[]int, pos *int) int {
	return (*band)[0]
}

func add(band *[]int, pos *int) {
	pos1 := (*band)[*pos+1]
	pos2 := (*band)[*pos+2]
	pos3 := (*band)[*pos+3]
	//log.Println(pos1,"(",(*band)[pos1],")",pos2,"(",(*band)[pos2],")",pos3,"(",(*band)[pos3],")"," Add:",(*band)[pos1],"+",(*band)[pos2],"=",(*band)[pos1] * (*band)[pos2],"->",(*band)[pos3])
	(*band)[pos3] = (*band)[pos1] + (*band)[pos2]
	*pos += 4
	return
}

func mul(band *[]int, pos *int) {
	pos1 := (*band)[*pos+1]
	pos2 := (*band)[*pos+2]
	pos3 := (*band)[*pos+3]
	//log.Println(pos1,"(",(*band)[pos1],")",pos2,"(",(*band)[pos2],")",pos3,"(",(*band)[pos3],")"," Mul:",(*band)[pos1],"*",(*band)[pos2],"=",(*band)[pos1] * (*band)[pos2],"->",(*band)[pos3])
	(*band)[pos3] = (*band)[pos1] * (*band)[pos2]
	*pos += 4
	return
}

func run(band []int) int {
	iP := 0
	for true {
		switch band[iP] {
		case 99:
			return terminate(&band, &iP)
		case 1:
			add(&band, &iP)
			break
		case 2:
			mul(&band, &iP)
			break
		default:
			log.Fatal("Unsupported OpCode")
			return -1
		}
	}
	return -2 //this should NEVER happen
}

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

func orgBand() []int {
	return []int{1, 12, 2, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 1, 10, 19, 1, 6, 19, 23, 1, 10, 23, 27, 2, 27, 13, 31, 1, 31, 6, 35, 2, 6, 35, 39, 1, 39, 5, 43, 1, 6, 43, 47, 2, 6, 47, 51, 1, 51, 5, 55, 2, 55, 9, 59, 1, 6, 59, 63, 1, 9, 63, 67, 1, 67, 10, 71, 2, 9, 71, 75, 1, 6, 75, 79, 1, 5, 79, 83, 2, 83, 10, 87, 1, 87, 5, 91, 1, 91, 9, 95, 1, 6, 95, 99, 2, 99, 10, 103, 1, 103, 5, 107, 2, 107, 6, 111, 1, 111, 5, 115, 1, 9, 115, 119, 2, 119, 10, 123, 1, 6, 123, 127, 2, 13, 127, 131, 1, 131, 6, 135, 1, 135, 10, 139, 1, 13, 139, 143, 1, 143, 13, 147, 1, 5, 147, 151, 1, 151, 2, 155, 1, 155, 5, 0, 99, 2, 0, 14, 0}
}

func main() {
	part1()
	part2()
}
