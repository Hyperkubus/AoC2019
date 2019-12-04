package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getDigit(number int, pos int) int {
	r := number % int(math.Pow(10, float64(pos)))
	return r / int(math.Pow(10, float64(pos-1)))
}

func doubleAdjacency(number int) bool {
	numberS := strconv.Itoa(number)
	for i := 0; i < 5; i++ {
		if numberS[i] == numberS[i+1] {
			return true
		}
	}
	return false
}

func trippleAdjacency(number int) bool {
	numberS := strconv.Itoa(number)
	if numberS[1] != numberS[2] && numberS[0] == numberS[1] {
		return true
	}
	for i := 1; i < 4; i++ {
		if numberS[i] != numberS[i-1] && numberS[i] == numberS[i+1] && numberS[i+1] != numberS[i+2] {
			return true
		}
	}
	if numberS[3] != numberS[4] && numberS[4] == numberS[5] {
		return true
	}
	return false
}

func outOfBounds(number int) bool {
	return number < 1000000 && number > 99999
}

func ascending(number int) bool {
	lastDigit := -1
	for i := 0; i < 6; i++ {
		if getDigit(number, 6-i) < lastDigit {
			return false
		}
		lastDigit = getDigit(number, 6-i)
	}
	return true

}

func meetsConditions(number int) bool {
	return outOfBounds(number) && ascending(number) && doubleAdjacency(number) && trippleAdjacency(number)
}

func main() {
	fmt.Println("112233", meetsConditions(112233))
	fmt.Println("123444", meetsConditions(123444))
	fmt.Println("111122", meetsConditions(111122))

	count := 0
	for i := 273025; i <= 767253; i++ {
		if meetsConditions(i) {
			fmt.Println(i)
			count++
		}
	}
	fmt.Println("Found:", count)
}
