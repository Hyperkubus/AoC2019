package main

import (
	"fmt"
	"log"
	"time"
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
	Head         *int64
	RelativeBase *int64
	Memory       *map[int64]int64
	Input        *chan int64
	Output       *chan int64
	Debug        *bool
	Terminated   *chan bool
	Sigterm      *bool
}

func (c Computer) term() {
	if *c.Debug {
		fmt.Println("TERM")
		c.panic()
	}
	*c.Sigterm = true
	time.Sleep(1000000)
	*c.Terminated <- true
}

func (c Computer) addi(modes int64) {
	args := c.getArguments(3, modes, 100)
	if *c.Debug {
		fmt.Println("ADDI", (*c.Memory)[(*c.Head)], (*c.Memory)[(*c.Head)+1], (*c.Memory)[(*c.Head)+2], (*c.Memory)[(*c.Head)+3], args)
	}
	c.writeToAddress(args[2], args[1]+args[0])
	*(c.Head) += 4
}

func (c Computer) mult(modes int64) {
	args := c.getArguments(3, modes, 100)
	if *c.Debug {
		fmt.Println("MULTI", (*c.Memory)[(*c.Head)], (*c.Memory)[(*c.Head)+1], (*c.Memory)[(*c.Head)+2], (*c.Memory)[(*c.Head)+3], args)
	}
	c.writeToAddress(args[2], args[1]*args[0])
	*(c.Head) += 4
}

func (c Computer) read(modes int64) {
	args := c.getArguments(1, modes, 1)
	if *c.Debug {
		fmt.Println("READ", (*c.Memory)[(*c.Head)], (*c.Memory)[(*c.Head)+1], args)
	}
	data := <-*c.Input
	c.writeToAddress(args[0], data)
	*(c.Head) += 2
}

func (c Computer) writ(modes int64) {
	args := c.getArguments(1, modes, 0)
	if *c.Debug {
		fmt.Println("WRIT", (*c.Memory)[(*c.Head)], (*c.Memory)[(*c.Head)+1], args)
	}
	*c.Output <- args[0]
	*(c.Head) += 2
}

func (c Computer) jump(modes int64) {
	args := c.getArguments(2, modes, 0)
	if *c.Debug {
		fmt.Println("JUMP", (*c.Memory)[(*c.Head)], (*c.Memory)[(*c.Head)+1], (*c.Memory)[(*c.Head)+2], args)
	}
	if args[0] != 0 {
		*(c.Head) = args[1]
	} else {
		*(c.Head) += 3
	}
}

func (c Computer) jumf(modes int64) {
	args := c.getArguments(2, modes, 0)
	if *c.Debug {
		fmt.Println("JUMF", (*c.Memory)[(*c.Head)], (*c.Memory)[(*c.Head)+1], (*c.Memory)[(*c.Head)+2], args)
	}
	if args[0] == 0 {
		*(c.Head) = args[1]
	} else {
		*(c.Head) += 3
	}
}

func (c Computer) less(modes int64) {
	args := c.getArguments(3, modes, 100)
	if *c.Debug {
		fmt.Println("LESS", (*c.Memory)[(*c.Head)], (*c.Memory)[(*c.Head)+1], (*c.Memory)[(*c.Head)+2], (*c.Memory)[(*c.Head)+3], args)
	}
	if args[0] < args[1] {
		c.writeToAddress(args[2], 1)
	} else {
		c.writeToAddress(args[2], 0)
	}
	*(c.Head) += 4
}

func (c Computer) equa(modes int64) {
	args := c.getArguments(3, modes, 100)
	if *c.Debug {
		fmt.Println("EQUA", (*c.Memory)[(*c.Head)], (*c.Memory)[(*c.Head)+1], (*c.Memory)[(*c.Head)+2], (*c.Memory)[(*c.Head)+3], args)
	}
	if args[0] == args[1] {
		c.writeToAddress(args[2], 1)
	} else {
		c.writeToAddress(args[2], 0)
	}
	*(c.Head) += 4
}

func (c Computer) setr(modes int64) {
	args := c.getArguments(1, modes, 0)
	if *c.Debug {
		fmt.Println("SETR", (*c.Memory)[(*c.Head)], (*c.Memory)[(*c.Head)+1], args)
	}
	*(c.RelativeBase) += args[0]
	*(c.Head) += 2
}

func (c Computer) Step() {
	instruction := c.readAtAddress(*(c.Head))
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
		c.read(modes)
		break
	case WRIT:
		c.writ(modes)
		break
	case JUMP:
		c.jump(modes)
		break
	case JUMF:
		c.jumf(modes)
		break
	case LESS:
		c.less(modes)
		break
	case EQUA:
		c.equa(modes)
		break
	case SETR:
		c.setr(modes)
		break
	case TERM:
		c.term()
	default:
		c.panic()
		log.Fatal("INVALID OPCODE ", modes, opCode)
	}

}

func (c Computer) getArguments(cnt int64, modes int64, rw int64) []int64 {
	var out []int64
	for i := int64(0); i < cnt; i++ {
		var arg int64
		var addr int64 = 0
		switch accessMode(modes % 10) {
		case POSITIONAL:
			addr = c.readAtAddress(*(c.Head) + 1 + i)
			if rw%10 == 0 {
				arg = c.readAtAddress(addr)
			} else {
				arg = addr
			}
			break
		case IMMEDIATE:
			if rw%10 == 0 {
				arg = c.readAtAddress(*(c.Head) + 1 + i)
			} else {
				log.Fatal("Cannot write in Immediate Mode")
			}
			break
		case RELATIVE:
			addr = c.readAtAddress(*(c.Head) + 1 + i)
			addr += *(c.RelativeBase)
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

func (c *Computer) readAtAddress(addr int64) int64 {
	return (*c.Memory)[addr]
}
func (c *Computer) writeToAddress(addr int64, value int64) {
	(*c.Memory)[addr] = value
}

func (c *Computer) SetDebug(d bool) {
	*(c.Debug) = d
}
func (c *Computer) DumpMemory() {
	fmt.Println(*c.Memory)
}

func (c *Computer) SetInput(in *chan int64) {
	*c.Input = *in
}

func (c *Computer) SetOutput(out *chan int64) {
	c.Output = out
}

func (c *Computer) SetTerminated(term *chan bool) {
	*c.Terminated = *term
}

func (c *Computer) panic() {
	fmt.Printf("PANIC:\nhead: %d\nbase: %d\n", *c.Head, c.RelativeBase)
}

func (c *Computer) Run(term *chan bool) {
	for !*c.Sigterm {
		c.Step()
	}
	*term <- true
}

func (c *Computer) Reset(data []int64) {
	myMap := make(map[int64]int64)
	for i := 0; i < len(data); i++ {
		myMap[int64(i)] = data[i]
	}
	var head int64 = 0
	var relb int64 = 0
	(*c).Memory = &myMap
	(*c).Head = &head
	(*c).RelativeBase = &relb
	(*c).Sigterm = new(bool)
}

func NewComputer(data []int64) (Computer, *chan int64, *chan int64, *chan bool) {
	myMap := make(map[int64]int64)
	for i := 0; i < len(data); i++ {
		myMap[int64(i)] = data[i]
	}
	var head int64 = 0
	var relb int64 = 0
	input := make(chan int64)
	output := make(chan int64)
	terminated := make(chan bool)
	computer := Computer{
		Head:         &head,
		RelativeBase: &relb,
		Memory:       &myMap,
		Input:        &input,
		Output:       &output,
		Debug:        new(bool),
		Terminated:   &terminated,
		Sigterm:      new(bool),
	}
	return computer, &input, &output, &terminated
}
