package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type resource struct {
	comp  component
	quant int
}

type component struct {
	identifier string
}

type recipie struct {
	inputQuantities []int
	inputs          []component
	outputQuantity  int
	output          component
}

func (r recipie) print(fct int) {
	for i := 0; i < len(r.inputs); i++ {
		if i == 0 {
			fmt.Printf("%d %s", fct*r.inputQuantities[i], r.inputs[i].identifier)
		} else {
			fmt.Printf(", %d %s", fct*r.inputQuantities[i], r.inputs[i].identifier)
		}
	}
	fmt.Printf(" => %d %s\n", fct*r.outputQuantity, r.output.identifier)
}

type nanofactory struct {
	resources      *[]component
	recipies       *[]recipie
	oreConsumed    *int
	spareResources *[]resource
}

func (nf nanofactory) getSpares(c component) int {
	for i := 0; i < len(*nf.spareResources); i++ {
		if (*nf.spareResources)[i].comp.identifier == c.identifier {
			return (*nf.spareResources)[i].quant
		}
	}
	return 0
}

func (nf nanofactory) useSpares(n int, c component) {
	//fmt.Println("using",n,c.identifier)
	for i := 0; i < len(*nf.spareResources); i++ {
		if (*nf.spareResources)[i].comp.identifier == c.identifier {
			(*nf.spareResources)[i].quant -= n
			return
		}
	}
}
func (nf nanofactory) addSpares(n int, c component) {
	//fmt.Println("Adding",n,c.identifier)
	for i := 0; i < len(*nf.spareResources); i++ {
		if (*nf.spareResources)[i].comp.identifier == c.identifier {
			(*nf.spareResources)[i].quant += n
			return
		}
	}
	(*nf.spareResources) = append((*nf.spareResources), resource{c, n})
}

func (nf nanofactory) produce(n int, c component) {
	res := nf.getRecipieFor(c)
	fct := int(math.Ceil(float64(n) / float64(res.outputQuantity)))
	//res.print(fct)
	for i := 0; i < len(res.inputs); i++ {
		c := res.inputs[i]
		if c.identifier == "ORE" {
			*nf.oreConsumed += res.inputQuantities[i] * fct
		} else {
			reqAmt := fct*res.inputQuantities[i] - nf.getSpares(c)
			//fmt.Println(c.identifier, nf.getSpares(c))
			nf.produce(reqAmt, c)
			//fmt.Println(c.identifier, nf.getSpares(c))
			nf.useSpares(fct*res.inputQuantities[i], c)
		}
	}
	nf.addSpares(res.outputQuantity*fct, res.output)
}

func (nf nanofactory) addRecipie(r recipie) {
	(*nf.recipies) = append(*nf.recipies, r)
}

func (nf nanofactory) addComponent(c component) {
	for i := 0; i < len(*nf.resources); i++ {
		if (*nf.resources)[i].identifier == c.identifier {
			return
		}
	}
	*nf.resources = append(*nf.resources, c)
}

func (nf nanofactory) getRecipieFor(c component) recipie {
	for i := 0; i < len(*nf.recipies); i++ {
		if (*nf.recipies)[i].output.identifier == c.identifier {
			return (*nf.recipies)[i]
		}
	}
	return recipie{
		inputQuantities: []int{},
		inputs:          []component{},
		outputQuantity:  0,
		output:          component{"NIL"},
	}
}

func (nf nanofactory) outputUnique() bool {
	for i := 0; i < len(*nf.recipies); i++ {
		for j := i + 1; j < len(*nf.recipies); j++ {
			if (*nf.recipies)[i].output.identifier == (*nf.recipies)[j].output.identifier {
				return false
			}
		}
	}
	return true
}

func main() {
	nf := nanofactory{
		resources:      new([]component),
		recipies:       new([]recipie),
		oreConsumed:    new(int),
		spareResources: new([]resource),
	}

	file, err := os.Open("S140/input_t1.txt")
	check(err)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		InOut := strings.Split(scanner.Text(), " => ")
		r := recipie{
			inputQuantities: []int{},
			inputs:          []component{},
			outputQuantity:  0,
			output:          component{},
		}
		Out := strings.Split(InOut[1], " ")
		r.output = component{identifier: Out[1]}
		r.outputQuantity, _ = strconv.Atoi(Out[0])
		In := strings.Split(InOut[0], ", ")
		for i := 0; i < len(In); i++ {
			part := strings.Split(In[i], " ")
			r.inputs = append(r.inputs, component{identifier: part[1]})
			cnt, _ := strconv.Atoi(part[0])
			r.inputQuantities = append(r.inputQuantities, cnt)
		}
		nf.addRecipie(r)
	}
	fmt.Println("Current Ore:", *nf.oreConsumed)
	i := 0
	for (*nf.oreConsumed) < 1000000000000 {
		nf.produce(1, component{"FUEL"})
		i++
	}
	fmt.Println("Produced", i, "Units of FUEL")
	fmt.Println("Current Ore:", *nf.oreConsumed)

	//fmt.Println(nf.getOreRequirements(1, component{"FUEL"}))
}
