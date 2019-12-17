package main

import (
	"fmt"
	"log"
)

type accessMode int

const (
	POSITIONAL accessMode = 0
	IMMEDIATE  accessMode = 1
	RELATIVE   accessMode = 2
)

type operation int

const (
	ADDI operation = 1
	MULT operation = 2
	READ operation = 3
	WRIT operation = 4
	JUMP operation = 5
	JUMF operation = 6
	LESS operation = 7
	EQUA operation = 8
	SETR operation = 9
	TERM operation = 99
)

type Computer struct {
	head         *int64
	relativeBase int64
	memory       *[]int64
	input        *<-chan int64
	output       *chan<- int64
	debug        *bool
	terminated   *chan bool
}

func (c Computer) term() {
	if *c.debug {
		fmt.Println("TERM")
	}
	*(c.terminated) <- true
}

func (c Computer) addi(modes int64) {
	args := c.getArguments(3, modes, 100)
	if *c.debug {
		fmt.Println("ADDI", (*c.memory)[(*c.head):(*c.head)+4], args)
	}
	c.writeToAddress(args[2], args[1]+args[0])
	*(c.head) += 4
}

func (c Computer) mult(modes int64) {
	args := c.getArguments(3, modes, 100)
	fmt.Println("MULTI", (*c.memory)[(*c.head):(*c.head)+4], args)
	c.writeToAddress(args[2], args[1]*args[0])
	*(c.head) += 4
}

func (c Computer) read(modes int64) {
	args := c.getArguments(1, modes, 1)
	fmt.Println("READ", (*c.memory)[(*c.head):(*c.head)+1], args)
	data := <-(*c.input)
	c.writeToAddress(args[0], data)
	*(c.head) += 2
}

func (c Computer) writ(modes int64) {
	args := c.getArguments(1, modes, 0)
	fmt.Println("WRIT", (*c.memory)[(*c.head):(*c.head)+1], args)
	(*c.output) <- args[0]
	*(c.head) += 2
}

func (c Computer) step() {
	if *(c.debug) {
		c.dumpMemory()
	}
	instruction := c.readAtAddress(*(c.head))
	opCode := instruction % 100
	modes := instruction / 100
	switch operation(opCode) {
	case ADDI:
		c.addi(modes)
		break
	case MULT:
		c.mult(modes)
		break
	case READ:
		log.Fatal("INVALID OPCODE")
	case WRIT:
		log.Fatal("INVALID OPCODE")
	case JUMP:
		log.Fatal("INVALID OPCODE")
	case JUMF:
		log.Fatal("INVALID OPCODE")
	case LESS:
		log.Fatal("INVALID OPCODE")
	case EQUA:
		log.Fatal("INVALID OPCODE")
	case SETR:
		log.Fatal("INVALID OPCODE")
	case TERM:
		c.term()
	default:
		log.Fatal("INVALID OPCODE")
	}

}

func (c Computer) getArguments(cnt int64, modes int64, rw int64) []int64 {
	var out []int64
	for i := int64(0); i < cnt; i++ {
		var arg int64
		var addr int64 = 0
		switch accessMode(modes % 10) {
		case POSITIONAL:
			addr = c.readAtAddress(*(c.head) + 1 + i)
			if rw%10 == 0 {
				arg = c.readAtAddress(addr)
			} else {
				arg = addr
			}
			break
		case IMMEDIATE:
			if rw%10 == 0 {
				arg = c.readAtAddress(*(c.head) + 1 + i)
			} else {
				log.Fatal("Cannot write in Immediate Mode")
			}
			break
		case RELATIVE:
			addr = c.readAtAddress(*(c.head) + 1 + i)
			addr += c.relativeBase
			if rw%10 == 0 {
				arg = c.readAtAddress(addr)
			} else {
				arg = addr
			}
			break
		default:
			log.Fatal("INVALID ACCESS MODE")
		}
		modes /= 10
		rw /= 10
		out = append(out, arg)
	}
	return out
}

func (c Computer) readAtAddress(addr int64) int64 {
	return (*c.memory)[addr]
}
func (c Computer) writeToAddress(addr int64, value int64) {
	(*c.memory)[addr] = value
}

func (c Computer) setDebug(d bool) {
	*(c.debug) = d
}
func (c Computer) dumpMemory() {
	fmt.Println(*c.memory)
}

func (c Computer) setInput(in *<-chan int64) {
	c.input = in
}

func (c Computer) setOutput(out *chan<- int64) {
	c.output = out
}

func NewComputer(data []int64) Computer {
	var head int64 = 0
	computer := Computer{
		head:         &head,
		relativeBase: 0,
		memory:       &data,
		input:        nil,
		output:       nil,
		debug:        new(bool),
		terminated:   nil,
	}
	return computer
}

func run(comp *Computer) {
	for true {
		(*comp).step()
	}
}

func main() {
	comp := NewComputer(day2input(12, 2))
	comp.setDebug(true)
	input := make(<-chan int64)
	output := make(chan<- int64)
	termination := make(chan bool)
	comp.setInput(&input)
	comp.setOutput(&output)

	fmt.Println("GO:")

	comp.Run()

	<-termination

	fmt.Println("terminated", (*(comp).memory)[0])
}

func day2input(x int64, y int64) []int64 {
	return []int64{1, x, y, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 1, 10, 19, 1, 6, 19, 23, 1, 10, 23, 27, 2, 27, 13, 31, 1, 31, 6, 35, 2, 6, 35, 39, 1, 39, 5, 43, 1, 6, 43, 47, 2, 6, 47, 51, 1, 51, 5, 55, 2, 55, 9, 59, 1, 6, 59, 63, 1, 9, 63, 67, 1, 67, 10, 71, 2, 9, 71, 75, 1, 6, 75, 79, 1, 5, 79, 83, 2, 83, 10, 87, 1, 87, 5, 91, 1, 91, 9, 95, 1, 6, 95, 99, 2, 99, 10, 103, 1, 103, 5, 107, 2, 107, 6, 111, 1, 111, 5, 115, 1, 9, 115, 119, 2, 119, 10, 123, 1, 6, 123, 127, 2, 13, 127, 131, 1, 131, 6, 135, 1, 135, 10, 139, 1, 13, 139, 143, 1, 143, 13, 147, 1, 5, 147, 151, 1, 151, 2, 155, 1, 155, 5, 0, 99, 2, 0, 14, 0}
}
