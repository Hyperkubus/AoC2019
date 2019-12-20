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

type direction int

const (
	NORTH direction = 0
	WEST  direction = 1
	SOUTH direction = 2
	EAST  direction = 3
)

type node struct {
	key   *int
	pos   *[2]int
	north *edge
	south *edge
	east  *edge
	west  *edge
}

type edge struct {
	dist  *int
	doors *[]int
	node  *node
}

func (e *edge) print() {
	//	fmt.Println("   ",*(*e).doors)
}

func (n *node) print() {
	if n.key != nil {
		fmt.Printf("[%d/%d] key: %d\n", (*n.pos)[0], (*n.pos)[1], *n.key)
	} else {
		fmt.Printf("[%d/%d] no key\n", (*n.pos)[0], (*n.pos)[1])
	}
	n.north.print()
	n.west.print()
	n.south.print()
	n.east.print()
}

func NewNode() *node {
	return &node{}
}

type position [2]int

func (p position) adjacent() []position {
	up := position{p[0], p[1] + 1}
	right := position{p[0] + 1, p[1]}
	down := position{p[0], p[1] - 1}
	left := position{p[0] - 1, p[1]}
	return []position{up, right, down, left}
}

func (n *node) toNextNode(board [128][128]rune, dir direction) {
	curPos := *n.pos
	nextPos := [2]int{}
	doors := []int{}
	stps := 0
	fmt.Println("going ", dir)
	for {
		fmt.Println("step")
		switch dir {
		case NORTH:
			nextPos[0] = curPos[0]
			nextPos[1] = curPos[1] + 1
			break
		case WEST:
			nextPos[0] = curPos[0] - 1
			nextPos[1] = curPos[1]
			break
		case SOUTH:
			nextPos[0] = curPos[0]
			nextPos[1] = curPos[1] - 1
			break
		case EAST:
			nextPos[0] = curPos[0] + 1
			nextPos[1] = curPos[1]
			break
		default:
			log.Fatal("We only know 4 directions")
		}
		r := board[nextPos[0]][nextPos[1]]
		stps++
		if isNode(board, nextPos, (dir+4)%4) {
			fmt.Println("got node")
			newNode := NewNode()
			if isKey(r) {
				*newNode.key = int(r)
			}
			switch (dir + 4) % 4 {
			case NORTH:
				(*newNode).north = &edge{
					dist:  &stps,
					doors: &doors,
					node:  n,
				}
				(*n).south = &edge{
					dist:  &stps,
					doors: &doors,
					node:  newNode,
				}
			}
			return
		}
		if isDoor(r) {
			doors = append(doors, int(r))
		}
		if r == '#' {
			return
		}
		curPos = nextPos
	}
}

func generateGraph(board [128][128]rune, spwn [2]int) (*node, []int) {
	k := 0
	spawn := node{key: &k, pos: &spwn}
	keys := []int{}

	spawn.toNextNode(board, SOUTH)

	return &spawn, keys
}

func isKey(r rune) bool {
	return unicode.IsLower(r)
}
func isDoor(r rune) bool {
	return unicode.IsUpper(r)
}
func isPath(r rune) bool {
	return unicode.IsUpper(r) || r == '.'
}
func isNode(board [128][128]rune, p [2]int, from direction) bool {
	if isKey(board[p[0]][p[1]]) {
		return true
	}
	switch from {
	case NORTH:
		if isPath(board[p[0]][p[1]-1]) && !isPath(board[p[0]-1][p[1]]) && !isPath(board[p[0]+1][p[1]]) {
			return false
		}
		break
	case WEST:
		if isPath(board[p[0]-1][p[1]]) && !isPath(board[p[0]][p[1]-1]) && !isPath(board[p[0]][p[1]+1]) {
			return false
		}
		break
	case SOUTH:
		if isPath(board[p[0]][p[1]+1]) && !isPath(board[p[0]-1][p[1]]) && !isPath(board[p[0]+1][p[1]]) {
			return false
		}
		break
	case EAST:
		if isPath(board[p[0]+1][p[1]]) && !isPath(board[p[0]][p[1]-1]) && !isPath(board[p[0]][p[1]+1]) {
			return false
		}
		break
	default:
		log.Fatal("We only know four directions")
	}
	return true
}

func parseMap(filePath string) *node {
	var keys []int
	spawn := NewNode()
	file, err := os.Open("S180/" + filePath)
	check(err)
	board := [128][128]rune{}
	scanner := bufio.NewScanner(file)
	spwnps := [2]int{0, 0}
	y := 0
	for scanner.Scan() {
		runes := []rune(scanner.Text())
		for x := 0; x < len(runes); x++ {
			board[x][y] = runes[x]
			if runes[x] == '@' {
				spwnps[0] = x
				spwnps[1] = y
			}
		}
		y++
	}

	spawn, keys = generateGraph(board, spwnps)
	fmt.Println(keys)
	return spawn
}

func main() {
	board := parseMap("input.simplified.txt")
	board.print()
	board.south.node.print()

	/*
		myMap := make(map[position]glyph)
		file, err := os.Open("S180/input1.txt")
		check(err)

		scanner := bufio.NewScanner(file)
		dim := [2]int{0,0}
		y:=0
		keys := []rune{}
		doors := []rune{}
		var spawn position;
		for scanner.Scan() {
			runes := []rune(scanner.Text())
			for x:=0;x<len(runes);x++{
				elposition := position{x,y}

				dim[0] = x+1
			}
			y++
		}
		dim[1] = y

		fmt.Println(dim, doors, keys, spawn)
		myMap[spawn] = EMPTY
		steps := search(myMap,dim,spawn,keys,[]rune{})
		fmt.Println(steps)
	*/
}
