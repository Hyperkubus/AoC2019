package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Node struct {
	name   string
	parent *Node
}

type NodeList struct {
	all []*Node
}

func newNode(n *NodeList, name string) *Node {
	node := Node{name, nil}
	(*n).all = append((*n).all, &node)
	return &node
}

func findOrCreateNode(n *NodeList, name string) *Node {
	for i := 0; i < len((*n).all); i++ {
		if ((*n).all[i]).name == name {
			return (*n).all[i]
		}
	}
	return newNode(n, name)
}

func (n NodeList) print() {
	for i := 0; i < len(n.all); i++ {
		fmt.Println(n.all[i].name)
	}
}

func getEdgeCount(n *NodeList) int {
	count := 0
	for i := 0; i < len(n.all); i++ {
		node := (*n).all[i]
		for node.name != "COM" {
			fmt.Println("At:", node.name, "Parent:", node.parent.name)
			count++
			node = node.parent
		}
	}
	return count
}

func getAllParents(node Node) []string {
	out := []string{}
	for node.name != "COM" {
		out = append(out, node.name)
		node = *node.parent
	}
	return out
}

func newOrbit(nodeList *NodeList, name string, parentName string) {
	parent := findOrCreateNode(nodeList, parentName)
	node := findOrCreateNode(nodeList, name)
	(*node).parent = parent
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func main() {
	var nodeList *NodeList = &NodeList{}

	file, err := os.Open("S060/input.txt")
	check(err)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		splitString := strings.Split(scanner.Text(), ")")
		newOrbit(nodeList, splitString[1], splitString[0])
	}

	san := findOrCreateNode(nodeList, "SAN")
	you := findOrCreateNode(nodeList, "YOU")
	sanList := getAllParents(*san)
	youList := getAllParents(*you)
	sanListB := []string{}
	youListB := []string{}
	cnt := 0
	i := 0
	for !stringInSlice(sanList[i], youList) {
		sanListB = append(sanListB, sanList[i])
		cnt++
		i++
	}
	i = 0
	for !stringInSlice(youList[i], sanList) {
		youListB = append(youListB, youList[i])
		cnt++
		i++
	}
	fmt.Println(sanListB)
	fmt.Println(youListB)
	cnt--
	cnt--
	fmt.Println(cnt)

	//fmt.Println(getEdgeCount(nodeList))
	//nodeList.print()
}
