package main

import (
	"fmt"
	"hyperkubus.dev/Intcomp"
	"log"
	"math"
	"time"
)

func inputBand() []int64 {
	return []int64{109, 424, 203, 1, 21101, 11, 0, 0, 1105, 1, 282, 21101, 0, 18, 0, 1106, 0, 259, 1202, 1, 1, 221, 203, 1, 21101, 0, 31, 0, 1105, 1, 282, 21102, 1, 38, 0, 1106, 0, 259, 20101, 0, 23, 2, 22102, 1, 1, 3, 21101, 1, 0, 1, 21101, 0, 57, 0, 1106, 0, 303, 1202, 1, 1, 222, 21002, 221, 1, 3, 21001, 221, 0, 2, 21102, 1, 259, 1, 21101, 80, 0, 0, 1105, 1, 225, 21102, 1, 117, 2, 21102, 1, 91, 0, 1105, 1, 303, 1202, 1, 1, 223, 20102, 1, 222, 4, 21101, 0, 259, 3, 21101, 0, 225, 2, 21101, 225, 0, 1, 21101, 118, 0, 0, 1105, 1, 225, 21001, 222, 0, 3, 21101, 20, 0, 2, 21102, 1, 133, 0, 1105, 1, 303, 21202, 1, -1, 1, 22001, 223, 1, 1, 21101, 0, 148, 0, 1106, 0, 259, 2101, 0, 1, 223, 20102, 1, 221, 4, 21001, 222, 0, 3, 21101, 0, 16, 2, 1001, 132, -2, 224, 1002, 224, 2, 224, 1001, 224, 3, 224, 1002, 132, -1, 132, 1, 224, 132, 224, 21001, 224, 1, 1, 21102, 195, 1, 0, 105, 1, 108, 20207, 1, 223, 2, 21002, 23, 1, 1, 21102, -1, 1, 3, 21101, 0, 214, 0, 1105, 1, 303, 22101, 1, 1, 1, 204, 1, 99, 0, 0, 0, 0, 109, 5, 1201, -4, 0, 249, 22102, 1, -3, 1, 22101, 0, -2, 2, 21202, -1, 1, 3, 21102, 1, 250, 0, 1106, 0, 225, 22102, 1, 1, -4, 109, -5, 2105, 1, 0, 109, 3, 22107, 0, -2, -1, 21202, -1, 2, -1, 21201, -1, -1, -1, 22202, -1, -2, -2, 109, -3, 2106, 0, 0, 109, 3, 21207, -2, 0, -1, 1206, -1, 294, 104, 0, 99, 21202, -2, 1, -2, 109, -3, 2105, 1, 0, 109, 5, 22207, -3, -4, -1, 1206, -1, 346, 22201, -4, -3, -4, 21202, -3, -1, -1, 22201, -4, -1, 2, 21202, 2, -1, -1, 22201, -4, -1, 1, 21201, -2, 0, 3, 21101, 343, 0, 0, 1105, 1, 303, 1105, 1, 415, 22207, -2, -3, -1, 1206, -1, 387, 22201, -3, -2, -3, 21202, -2, -1, -1, 22201, -3, -1, 3, 21202, 3, -1, -1, 22201, -3, -1, 2, 21201, -4, 0, 1, 21101, 0, 384, 0, 1105, 1, 303, 1105, 1, 415, 21202, -4, -1, -4, 22201, -4, -3, -4, 22202, -3, -2, -2, 22202, -2, -4, -4, 22202, -3, -2, -3, 21202, -4, -1, -2, 22201, -3, -2, 1, 22101, 0, 1, -4, 109, -5, 2105, 1, 0}
}

func output(i *chan int64) {
	cnt := 0
	for y := int64(0); y < 50; y++ {
		for x := int64(0); x < 50; x++ {
			n := <-*i
			fmt.Print(n)
			if n > 0 {
				cnt++
			}
		}
		fmt.Print("\n")
	}
	fmt.Println(cnt)
}

func input(i *chan int64) {
	for y := int64(0); y < 50; y++ {
		for x := int64(0); x < 50; x++ {
			*i <- x
			*i <- y
		}
	}
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func main() {
	start := time.Now()
	cnt := 0
	lastXmap := map[int64]int64{}
	var startX int64 = 250
	for y := int64(750); true; y = y + 1 {
		var firstX int64 = math.MaxInt64
		var lastX int64 = math.MaxInt64
		for x := startX; lastX == math.MaxInt64; x++ {
			comp, in, out, term := Intcomp.NewComputer(inputBand())

			ret := make(chan bool)
			isTermed := false
			go comp.Run(&ret)
			*in <- x
			*in <- y
			n := <-*out
			//fmt.Print(n)
			if n > 0 {
				cnt++
				firstX = min(firstX, x)
			} else {
				if firstX != math.MaxInt64 && lastX == math.MaxInt64 {
					lastX = x - 1
				}
			}

			for !isTermed {
				<-*term
				isTermed = true
			}

			<-ret

			time.Sleep(10 /* 1000*/ * 1000)
		}
		//3810986 to high
		//3790981 WRONG
		startX = firstX
		//if firstX > 282 {log.Fatal("to far")}
		lastXmap[y] = lastX
		if y > 850 {
			L := firstX
			R := lastXmap[y-99]
			if L+99 <= R {
				U := y - 99
				D := y
				log.Fatal("HIT ", []int64{L, U}, []int64{R, D}, L*10000+(U), "took:", time.Since(start))
			}
		}
		//if(lastX-firstX > 100){
		//fmt.Printf(" (%d)[%d-%d|%d\n",y,firstX,lastX,lastX-firstX)
		//}
	}
	fmt.Println(cnt)
}