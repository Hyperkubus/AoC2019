package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func Str2IntA(str string) []int {
	runes := []rune(str)
	intA := []int{}
	for i := 0; i < len(runes); i++ {
		intI, _ := strconv.Atoi(string(runes[i]))
		intA = append(intA, intI)
	}
	return intA
}

func getFactor(r, c int) int {
	base := []int{0, 1, 0, -1}
	flt := float64(c+1) / float64(r+1)
	flt = math.Floor(flt)
	idx := int64(flt)
	idx %= 4
	return base[idx]
}

func phase(input []int, offset int) []int {
	list := []int{}
	for i := 0; i < len(input); i++ {
		e := 0
		for j := i; j < len(input); j++ {
			fct := getFactor(i, j+offset)
			e += fct * input[j]
		}
		e = int(math.Abs(float64(e))) % 10
		list = append(list, e)
	}
	return list
}

func pseudoPhase(input []int) []int {
	var out = make([]int, len(input))
	tmp := 0
	for i := len(input) - 1; i >= 0; i-- {
		tmp += input[i]
		out[i] = tmp % 10
	}
	return out
}

func printFirstEight(list []int) {
	for i := 0; i < 16; i++ {
		fmt.Printf("%d ", list[i])
	}
	fmt.Printf("\n")
}

func getOffset(input []int) int {
	offset := 0
	for i := 0; i < 7; i++ {
		offset += int(math.Pow(10, float64(6-i))) * input[i]
	}
	return offset
}

func main() {
	//input := Str2IntA("59796737047664322543488505082147966997246465580805791578417462788780740484409625674676660947541571448910007002821454068945653911486140823168233915285229075374000888029977800341663586046622003620770361738270014246730936046471831804308263177331723460787712423587453725840042234550299991238029307205348958992794024402253747340630378944672300874691478631846617861255015770298699407254311889484508545861264449878984624330324228278057377313029802505376260196904213746281830214352337622013473019245081834854781277565706545720492282616488950731291974328672252657631353765496979142830459889682475397686651923318015627694176893643969864689257620026916615305397")
	//input := Str2IntA("03036732577212944063491565474664")
	//input := Str2IntA("02935109699940807407585447034323")
	input := Str2IntA("03081770884921959731165446850517")
	inputA := append([]int{}, input...)
	offset := getOffset(input)
	fmt.Println("offset:", offset)
	start := time.Now()
	for i := 0; i < (9999 - 1); i++ {
		if offset > len(input) {
			offset -= len(inputA)
		} else {
			inputA = append(inputA, input...)
		}
	}
	inputA = inputA[offset:]
	fmt.Println(len(inputA))
	fmt.Println(time.Since(start))

	for i := 0; i < 100; i++ {
		//inputA = phase(inputA,phaseOffset)
		inputA = pseudoPhase(inputA)
	}

	inputA = append(inputA, inputA...)
	fmt.Println(time.Since(start))
	printFirstEight(inputA)
}
