package main

import (
	"fmt"
	"hyperkubus.dev/Intcomp"
	"math"
	"time"
)

func botBand() []int64 {
	return []int64{3, 8, 1005, 8, 345, 1106, 0, 11, 0, 0, 0, 104, 1, 104, 0, 3, 8, 102, -1, 8, 10, 1001, 10, 1, 10, 4, 10, 108, 1, 8, 10, 4, 10, 102, 1, 8, 28, 1006, 0, 94, 2, 106, 5, 10, 1, 1109, 12, 10, 3, 8, 1002, 8, -1, 10, 1001, 10, 1, 10, 4, 10, 1008, 8, 1, 10, 4, 10, 101, 0, 8, 62, 1, 103, 6, 10, 1, 108, 12, 10, 3, 8, 102, -1, 8, 10, 1001, 10, 1, 10, 4, 10, 1008, 8, 0, 10, 4, 10, 102, 1, 8, 92, 2, 104, 18, 10, 2, 1109, 2, 10, 2, 1007, 5, 10, 1, 7, 4, 10, 3, 8, 102, -1, 8, 10, 1001, 10, 1, 10, 4, 10, 108, 0, 8, 10, 4, 10, 102, 1, 8, 129, 2, 1004, 15, 10, 2, 1103, 15, 10, 2, 1009, 6, 10, 3, 8, 102, -1, 8, 10, 1001, 10, 1, 10, 4, 10, 1008, 8, 1, 10, 4, 10, 101, 0, 8, 164, 2, 1109, 14, 10, 1, 1107, 18, 10, 1, 1109, 13, 10, 1, 1107, 11, 10, 3, 8, 102, -1, 8, 10, 101, 1, 10, 10, 4, 10, 108, 0, 8, 10, 4, 10, 1001, 8, 0, 201, 2, 104, 20, 10, 1, 107, 8, 10, 1, 1007, 5, 10, 3, 8, 102, -1, 8, 10, 101, 1, 10, 10, 4, 10, 1008, 8, 1, 10, 4, 10, 101, 0, 8, 236, 3, 8, 1002, 8, -1, 10, 1001, 10, 1, 10, 4, 10, 108, 0, 8, 10, 4, 10, 1001, 8, 0, 257, 3, 8, 102, -1, 8, 10, 101, 1, 10, 10, 4, 10, 108, 1, 8, 10, 4, 10, 102, 1, 8, 279, 1, 107, 0, 10, 1, 107, 16, 10, 1006, 0, 24, 1, 101, 3, 10, 3, 8, 102, -1, 8, 10, 101, 1, 10, 10, 4, 10, 108, 0, 8, 10, 4, 10, 1002, 8, 1, 316, 2, 1108, 15, 10, 2, 4, 11, 10, 101, 1, 9, 9, 1007, 9, 934, 10, 1005, 10, 15, 99, 109, 667, 104, 0, 104, 1, 21101, 0, 936995730328, 1, 21102, 362, 1, 0, 1105, 1, 466, 21102, 1, 838210728716, 1, 21101, 373, 0, 0, 1105, 1, 466, 3, 10, 104, 0, 104, 1, 3, 10, 104, 0, 104, 0, 3, 10, 104, 0, 104, 1, 3, 10, 104, 0, 104, 1, 3, 10, 104, 0, 104, 0, 3, 10, 104, 0, 104, 1, 21102, 1, 235350789351, 1, 21101, 0, 420, 0, 1105, 1, 466, 21102, 29195603035, 1, 1, 21102, 1, 431, 0, 1105, 1, 466, 3, 10, 104, 0, 104, 0, 3, 10, 104, 0, 104, 0, 21101, 0, 825016079204, 1, 21101, 0, 454, 0, 1105, 1, 466, 21101, 837896786700, 0, 1, 21102, 1, 465, 0, 1106, 0, 466, 99, 109, 2, 21201, -1, 0, 1, 21101, 0, 40, 2, 21102, 1, 497, 3, 21101, 0, 487, 0, 1105, 1, 530, 109, -2, 2106, 0, 0, 0, 1, 0, 0, 1, 109, 2, 3, 10, 204, -1, 1001, 492, 493, 508, 4, 0, 1001, 492, 1, 492, 108, 4, 492, 10, 1006, 10, 524, 1101, 0, 0, 492, 109, -2, 2105, 1, 0, 0, 109, 4, 2102, 1, -1, 529, 1207, -3, 0, 10, 1006, 10, 547, 21102, 1, 0, -3, 21201, -3, 0, 1, 22102, 1, -2, 2, 21101, 1, 0, 3, 21102, 1, 566, 0, 1105, 1, 571, 109, -4, 2106, 0, 0, 109, 5, 1207, -3, 1, 10, 1006, 10, 594, 2207, -4, -2, 10, 1006, 10, 594, 21201, -4, 0, -4, 1106, 0, 662, 21201, -4, 0, 1, 21201, -3, -1, 2, 21202, -2, 2, 3, 21101, 613, 0, 0, 1105, 1, 571, 22101, 0, 1, -4, 21101, 0, 1, -1, 2207, -4, -2, 10, 1006, 10, 632, 21101, 0, 0, -1, 22202, -2, -1, -2, 2107, 0, -3, 10, 1006, 10, 654, 22101, 0, -1, 1, 21102, 654, 1, 0, 105, 1, 529, 21202, -2, -1, -2, 22201, -4, -2, -4, 109, -5, 2105, 1, 0}
}

type orientation int8

const (
	NORTH orientation = 0
	WEST  orientation = 1
	SOUTH orientation = 2
	EAST  orientation = 3
)

type Position struct {
	x int
	y int
}

type Bot struct {
	pos         *Position
	orientation *orientation
}

func NewBot(pos Position, ori orientation) Bot {
	return Bot{
		pos:         &pos,
		orientation: &ori,
	}
}

func serveInput(bot *Bot, hull *map[Position]bool, i *chan int64) {
	defer close(*i)
	for true {
		out := int64(0)
		if (*hull)[*bot.pos] == true {
			out = 1
			continue
		}
		fmt.Println("in: ", out)
		*i <- out
	}
}

func processOutput(bot *Bot, hull *map[Position]bool, o *chan int64) {
	defer close(*o)
	for {
		p := <-*o
		t := <-*o
		fmt.Println("out: ", p, t)
		if p == 1 {
			(*hull)[*bot.pos] = true
		}
		if t == 0 {
			(*bot.orientation) = ((*bot.orientation) + 1) % 4
		}
		if t == 1 {
			(*bot.orientation) = ((*bot.orientation) - 1) % 4
		}

	}
}

func checkTerm(t *bool, i *chan bool) {
	for {
		<-*i
		*t = true
	}
}

func countHull(hull *map[Position]bool) int {
	i := 0
	for _, v := range *hull {
		if v {
			i++
		}
	}
	return i
}

func left(o orientation) orientation {
	switch o {
	case NORTH:
		return WEST
	case WEST:
		return SOUTH
	case SOUTH:
		return EAST
	case EAST:
		return NORTH
	}
	return NORTH
}

func right(o orientation) orientation {
	switch o {
	case NORTH:
		return EAST
	case WEST:
		return NORTH
	case SOUTH:
		return WEST
	case EAST:
		return SOUTH
	}
	return NORTH
}

func printHull(hull map[Position]bool, minX, maxX, minY, maxY int64) {
	for y := maxY; y >= minY; y-- {
		for x := minX; x <= maxX; x++ {
			if (hull[Position{int(x), int(y)}]) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func botExe(hpBot *Bot, hull *map[Position]bool, in *chan int64, out *chan int64, isTermed *bool, output *[]int64) {
	var maxX int64 = math.MinInt64
	var maxY int64 = math.MinInt64
	var minX int64 = math.MaxInt64
	var minY int64 = math.MaxInt64
	for {
		c := int64(0)
		if (*hull)[*hpBot.pos] {
			c = 1
		}
		*in <- c
		p := <-*out
		t := <-*out
		if p == 1 {
			(*hull)[*hpBot.pos] = true
		} else {
			(*hull)[*hpBot.pos] = false
		}
		if t == 0 {
			(*hpBot.orientation) = left(*hpBot.orientation)
		}
		if t == 1 {
			(*hpBot.orientation) = right(*hpBot.orientation)
		}
		switch *hpBot.orientation {
		case NORTH:
			(*hpBot.pos).y++
			break
		case WEST:
			(*hpBot.pos).x--
			break
		case SOUTH:
			(*hpBot.pos).y--
			break
		case EAST:
			(*hpBot.pos).x++
			break
		}
		maxX = max(maxX, int64((*hpBot.pos).x))
		minX = min(minX, int64((*hpBot.pos).x))
		maxY = max(maxY, int64((*hpBot.pos).y))
		minY = min(minY, int64((*hpBot.pos).y))
		*output = []int64{minX, maxX, minY, maxY}
		if *isTermed {
			printHull(*hull, minX, maxX, minY, maxY)
			return
		}
	}
}

func main() {
	comp, in, out, term := Intcomp.NewComputer(botBand())
	hpBot := NewBot(Position{0, 0}, NORTH)
	hull := make(map[Position]bool)
	hull[Position{int(0), int(0)}] = true
	ret := make(chan bool)
	isTermed := false
	go comp.Run(&ret)
	output := []int64{}
	go botExe(&hpBot, &hull, in, out, &isTermed, &output)
	for !isTermed {
		<-*term
		isTermed = true
	}
	<-ret
	time.Sleep(10 * 1000 * 1000)
	printHull(hull, output[0], output[1], output[2], output[3])
}
