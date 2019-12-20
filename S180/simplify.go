package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func removeDeadEnds(myMap [82][82]rune) ([82][82]rune, int) {
	repls := 0
	for y := 0; y < 82; y++ {
		for x := 0; x < 82; x++ {
			if myMap[x][y] == '.' || unicode.IsUpper(myMap[x][y]) {
				i := 0
				if myMap[x+1][y] == '#' {
					i++
				}
				if myMap[x][y+1] == '#' {
					i++
				}
				if myMap[x-1][y] == '#' {
					i++
				}
				if myMap[x][y-1] == '#' {
					i++
				}
				if i > 2 {
					myMap[x][y] = '#'
					repls++
				}
			}
		}
	}
	return myMap, repls
}
func print(myMap [82][82]rune) {
	for y := 0; y < 82; y++ {
		for x := 0; x < 82; x++ {
			if myMap[x][y] == '#' {
				fmt.Print(" ")
			} else if myMap[x][y] == '.' {
				fmt.Print("#")
			} else {
				fmt.Print(string(myMap[x][y]))
			}
		}
		fmt.Print("\n")
	}
}

func main() {
	fileIn, _ := os.Open("S180/input.txt")
	scanner := bufio.NewScanner(fileIn)
	myMap := [82][82]rune{}
	y := 0
	for scanner.Scan() {
		runes := []rune(scanner.Text())
		for x := 0; x < len(runes); x++ {
			myMap[x][y] = runes[x]
		}
		y++
	}
	i := 1
	for i > 0 {
		myMap, i = removeDeadEnds(myMap)
	}
	fileOut, _ := os.Create("S180/input.simplified.txt")
	writer := bufio.NewWriter(fileOut)
	for y := 0; y < 81; y++ {
		for x := 0; x < 81; x++ {
			_, err := writer.WriteRune(myMap[x][y])
			check(err)
		}
		_, err := writer.WriteRune('\n')
		check(err)
	}
	err := writer.Flush()
	check(err)

	print(myMap)
}
