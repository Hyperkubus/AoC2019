package main

import (
	"fmt"
	"hyperkubus.dev/Intcomp"
)

func hotwire(i *chan int64, o *chan int64, name string, saveLast *int64) {
	for x := range *i {
		fmt.Println(name, x)
		if saveLast != nil {
			*saveLast = x
		}
		*o <- x
	}
}

func initialBand() []int64 {
	return []int64{3, 8, 1001, 8, 10, 8, 105, 1, 0, 0, 21, 34, 51, 64, 81, 102, 183, 264, 345, 426, 99999, 3, 9, 102, 2, 9, 9, 1001, 9, 4, 9, 4, 9, 99, 3, 9, 101, 4, 9, 9, 102, 5, 9, 9, 1001, 9, 2, 9, 4, 9, 99, 3, 9, 101, 3, 9, 9, 1002, 9, 5, 9, 4, 9, 99, 3, 9, 102, 3, 9, 9, 101, 3, 9, 9, 1002, 9, 4, 9, 4, 9, 99, 3, 9, 1002, 9, 3, 9, 1001, 9, 5, 9, 1002, 9, 5, 9, 101, 3, 9, 9, 4, 9, 99, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 99, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 99, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 99}
}

func runFor(phase []int) int64 {
	compA, inA, outA, termA := Intcomp.NewComputer(initialBand())
	compB, inB, outB, termB := Intcomp.NewComputer(initialBand())
	compC, inC, outC, termC := Intcomp.NewComputer(initialBand())
	compD, inD, outD, termD := Intcomp.NewComputer(initialBand())
	compE, inE, outE, termE := Intcomp.NewComputer(initialBand())
	output := new(int64)
	go hotwire(outA, inB, "AB", nil)
	go hotwire(outB, inC, "BC", nil)
	go hotwire(outC, inD, "CD", nil)
	go hotwire(outD, inE, "DE", nil)
	go hotwire(outE, inA, "EA", output)
	go compA.Run()
	go compB.Run()
	go compC.Run()
	go compD.Run()
	go compE.Run()
	*inA <- int64(phase[0])
	*inB <- int64(phase[1])
	*inC <- int64(phase[2])
	*inD <- int64(phase[3])
	*inE <- int64(phase[4])
	*inA <- 0
	<-*termA
	<-*termB
	<-*termC
	<-*termD
	<-*termE

	return *output
}

func main() {
	maxBoost := int64(0)
	settings := []int{}
	for a := 5; a < 10; a++ {
		for b := 5; b < 10; b++ {
			if a == b {
				continue
			}
			for c := 5; c < 10; c++ {
				if a == c || b == c {
					continue
				}
				for d := 5; d < 10; d++ {
					if a == d || b == d || c == d {
						continue
					}
					for e := 5; e < 10; e++ {
						if a == e || b == e || c == e || d == e {
							continue
						}
						boost := runFor([]int{a, b, c, d, e})
						if boost > maxBoost {
							maxBoost = boost
							settings = []int{a, b, c, d, e}
						}
					}
				}
			}
		}
	}
	fmt.Println(settings, maxBoost)
}
