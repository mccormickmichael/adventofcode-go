package day14

import (
	"math"
)

type component struct {
	name     string
	produced int
	consumed int
	silo     int
	producer *reaction
}

type input struct {
	c *component
	n int
}

type reaction struct {
	output int
	inputs []input
}


var components = make(map[string]component)
var reactions = make(map[string]reaction)


func doit() {
	components["ORE"] = component{name:"ORE", silo:math.MaxInt64}

	
	f := components["FUEL"]

	f.take(1)

	o := components["ORE"]

	println(o.consumed)
}

func (c *component) take(n int) int {

	for c.silo < n {
		made := c.producer.make()
		c.produced += made
		c.silo += made
	}
	c.consumed += n
	c.silo -= n
	return n
}

func (r *reaction) make() int {
	for _, in := range r.inputs {
		in.c.take(in.n)
	}
	return r.output
}

