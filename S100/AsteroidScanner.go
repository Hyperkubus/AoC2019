package main

import (
	"bufio"
	"fmt"
	"log"
	"math"

	//"math"
	"os"
)

type position struct {
	x    int
	y    int
	hits int
}

type vector struct {
	x int
	y int
}

func (v vector) norm() vector {
	return vector{
		x: int((float64(v.x) / v.len()) * 1000),
		y: int((float64(v.y) / v.len()) * 1000),
	}
}

func (v vector) len() float64 {
	return math.Sqrt(math.Pow(float64(v.x), 2) + math.Pow(float64(v.y), 2))
}

func (v vector) theta() float64 {
	return math.Tanh(float64(v.y) / float64(v.x))
}

func inLine(a vector, b vector, c vector) bool {
	dxc := a.x - b.x
	dyc := a.y - b.y
	dxl := c.x - b.x
	dyl := c.y - b.y
	cross := dxc*dyl - dyc*dxl
	return cross == 0
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getVector(from position, to position) vector {
	return vector{
		x: to.x - from.x,
		y: to.y - from.y,
	}
}

func unique(intSlice []vector) []vector {
	keys := make(map[vector]bool)
	list := []vector{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func toSpiral(list []position) {
	myMap := map[float64][]vector{}
	for i := 0; i < len(list); i++ {
		v := vector{list[i].x, list[i].y}
		mag := v.len()
		rad := 2*math.Pi - v.theta()

	}

}

func calcHits(list *[]position) {
	for i := 0; i < len(*list); i++ {
		//vlist := []vector{}
		vmap := new(map[vector]bool)
		*vmap = map[vector]bool{}
		here := vector{(*list)[i].x, (*list)[i].y}
		cnt := 0
		for j := 0; j < len(*list); j++ {
			if i == j {
				continue
			}
			there := vector{(*list)[j].x, (*list)[j].y}
			diff := vector{there.x - here.x, there.y - here.y}
			diff = diff.norm()
			if (*vmap)[diff] {
				continue
			}
			cnt++
			(*vmap)[diff] = true
			//vlist = append(vlist, diff.norm())
		}

		(*list)[i].hits = cnt
	}
}

func main() {
	file, err := os.Open("S100/input.txt")
	check(err)

	asteroids := []position{}

	y := 0
	maxX := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := []rune(scanner.Text())
		for x := 0; x < len(line); x++ {
			if line[x] == '#' {
				asteroids = append(asteroids, position{x, y, 0})
			}
			if x > maxX {
				maxX = x
			}
		}
		y++
	}
	fmt.Printf("Found %d Asteroids in %dx%d grid\n", len(asteroids), maxX+1, y)
	calcHits(&asteroids)
	for i := 0; i < len(asteroids); i++ {
		fmt.Println(asteroids[i].x, asteroids[i].y, asteroids[i].hits)
	}
	maxHits := -1
	j := -1
	for i := 0; i < len(asteroids); i++ {
		if asteroids[i].hits > maxHits {
			maxHits = asteroids[i].hits
			j = i
		}
	}
	fmt.Printf("Best location for Observatory %d/%d with %d Asteroids in sight\n", asteroids[j].x, asteroids[j].y, asteroids[j].hits)

}
