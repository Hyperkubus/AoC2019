package main

import (
	"fmt"
	"math"
	"time"
)

type system struct {
	p []int
	v []int
}

func (s system) hash() int {
	hash := 0
	for i := 0; i < len(s.p); i++ {
		h := float64(s.p[i]*100 + s.v[i])
		hash += int(math.Pow(10, float64(i*6)) * h)
	}
	return hash
}

func (s system) step() {
	for i := 0; i < len(s.p); i++ {
		for j := i + 1; j < len(s.p); j++ {
			if s.p[i] < s.p[j] {
				s.v[i]++
				s.v[j]--
			}
			if s.p[i] > s.p[j] {
				s.v[i]--
				s.v[j]++
			}
		}
		s.p[i] += s.v[i]
	}
}

func (s system) print() {
	for i := 0; i < len(s.p); i++ {
		fmt.Printf("%d[%d] ", s.p[i], s.v[i])
	}
	fmt.Printf("\n")
}

func findrepetition(s system) (int, int) {
	hashmap := map[int]int{}
	startHash := s.hash()
	i := 1
	s.step()
	var hash int
	for true {
		hash = s.hash()
		if hashmap[hash] > 0 || hash == startHash {
			break
		}
		i++
		hashmap[hash] = i
		s.step()
	}
	fmt.Println(hashmap[hash])
	return hashmap[hash], i
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func main() {
	sysX := system{
		p: []int{15, -5, 0, 5},
		v: []int{0, 0, 0, 0},
	}
	sysY := system{
		p: []int{-2, -4, -6, 9},
		v: []int{0, 0, 0, 0},
	}
	sysZ := system{
		p: []int{-6, -11, 0, 6},
		v: []int{0, 0, 0, 0},
	}
	start := time.Now()
	_, x := findrepetition(sysX)
	_, y := findrepetition(sysY)
	_, z := findrepetition(sysZ)
	fmt.Println(LCM(x, y, z))
	fmt.Println(time.Since(start))

}
