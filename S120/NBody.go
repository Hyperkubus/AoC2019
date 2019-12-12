package main

import (
	"fmt"
	"math"
	"time"
)

type vector3 struct {
	x int
	y int
	z int
}

func (v1 vector3) add(v2 vector3) vector3 {
	return vector3{
		x: v1.x + v2.x,
		y: v1.y + v2.y,
		z: v1.z + v2.z,
	}
}

func (v vector3) getEnergy() int {
	return int(math.Abs(float64(v.x)) + math.Abs(float64(v.y)) + math.Abs(float64(v.z)))
}

type NBody struct {
	pos *vector3
	vel *vector3
}

func (b NBody) print() {
	fmt.Printf("{%d/%d/%d} [%d/%d/%d] (%d)\n", b.pos.x, b.pos.y, b.pos.z, b.vel.x, b.vel.y, b.vel.z, b.getEnergy())
}

func (b NBody) move() {
	*b.pos = (b.pos).add(*b.vel)
}

func (b NBody) getEnergy() int {
	return b.pos.getEnergy() * b.vel.getEnergy()
}

func (b NBody) hash() int {
	hash := b.vel.z
	hash += b.vel.y * 1000
	hash += b.vel.x * 1000 * 1000
	hash += b.pos.z * 1000 * 1000 * 1000
	hash += b.pos.y * 1000 * 1000 * 1000 * 1000
	hash += b.pos.x * 1000 * 1000 * 1000 * 1000 * 1000
	return hash
}

func NewNBody(x int, y int, z int) NBody {
	pos := vector3{x, y, z}
	vel := vector3{0, 0, 0}
	return NBody{&pos, &vel}
}

func applyGravity(a *NBody, b *NBody) {
	if a.pos.x < b.pos.x {
		(*a).vel.x += 1
		(*b).vel.x -= 1
	}
	if a.pos.x > b.pos.x {
		(*a).vel.x -= 1
		(*b).vel.x += 1
	}
	if a.pos.y < b.pos.y {
		(*a).vel.y += 1
		(*b).vel.y -= 1
	}
	if a.pos.y > b.pos.y {
		(*a).vel.y -= 1
		(*b).vel.y += 1
	}
	if a.pos.z < b.pos.z {
		(*a).vel.z += 1
		(*b).vel.z -= 1
	}
	if a.pos.z > b.pos.z {
		(*a).vel.z -= 1
		(*b).vel.z += 1
	}
}

func step(bodies *[]NBody) *[]NBody {
	for i := 0; i < len(*bodies); i++ {
		for j := i + 1; j < len(*bodies); j++ {
			applyGravity(&(*bodies)[i], &(*bodies)[j])
		}
	}
	for i := 0; i < len(*bodies); i++ {
		(*bodies)[i].move()
	}
	return bodies
}

func debugBodies(bodies []NBody) {
	energy := 0
	for i := 0; i < len(bodies); i++ {
		bodies[i].print()
		energy += bodies[i].getEnergy()
	}

}

func getByAxis(bodies []NBody) ([]int, []int, []int) {
	x := []int{}
	y := []int{}
	z := []int{}
	for i := 0; i < len(bodies); i++ {

	}
}

func main() {
	bodies := []NBody{}
	bodies = append(bodies, NewNBody(15, -2, -6))
	bodies = append(bodies, NewNBody(-5, -4, -11))
	bodies = append(bodies, NewNBody(0, -6, 0))
	bodies = append(bodies, NewNBody(5, 9, 6))
	/*bodies = append(bodies,NewNBody(-1,0,2))
	bodies = append(bodies,NewNBody(2,-10,-7))
	bodies = append(bodies,NewNBody(4,-8,8))
	bodies = append(bodies,NewNBody(3,5,-1))*/

	debugBodies(bodies)
	hmap := map[int64]bool{}
	i := 1
	start := time.Now()
	for true {
		if i%1000000 == 0 {
			fmt.Println("Iteration", i, "took", time.Since(start))
			start = time.Now()
		}
		bodies = *step(&bodies)
		i++
	}
	fmt.Println("=================")
	debugBodies(bodies)
}
