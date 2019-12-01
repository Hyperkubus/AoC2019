package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func calculate(mass int64) int64 {
	return int64(math.Floor(float64(mass)/3) - 2)
}

func calculateTotal(mass int64) int64 {
	totalFuel := int64(0)
	newFuel := calculate(mass)
	for newFuel > 0 {
		totalFuel += newFuel
		newFuel = calculate(newFuel)
	}
	return totalFuel
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	file, err := os.Open("input.txt")
	check(err)

	scanner := bufio.NewScanner(file)
	fuel := int64(0)
	totalFuel := int64(0)
	for scanner.Scan() {
		mass, err := strconv.ParseInt(scanner.Text(), 10, 64)
		check(err)
		fuel += calculate(mass)
		totalFuel += calculateTotal(mass)
	}
	fmt.Println("fuel: ", fuel)
	fmt.Println("totalFuel: ", totalFuel)
}
