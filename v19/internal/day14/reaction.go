package day14

import (
	"fmt"
)

type component struct {
	name     string
	produced int
	consumed int
	silo     int
	producer *reaction
}

type reactionInput struct {
	c *component
	n int
}

type reaction struct {
	description string
	output int
	inputs []reactionInput
}

type refinery map[string]*component

func (r refinery) find(name string) *component {
	if c, ok := r[name]; ok {
		return c
	}
	c := &component{name:name}
	r[name] = c
	return c
}

func (r refinery) init(ore int) {
	r.find("ORE").silo = ore
}

func (r *refinery) makeReaction(line string) {
	rd := parse(line)

	inputs := make([]reactionInput, len(rd.inputs))
	for i, cd := range rd.inputs {
		component := r.find(cd.name)
		inputs[i] = reactionInput{c: component, n:cd.count}
	}
	reaction := reaction{output:rd.output.count, inputs:inputs, description:line}
	outputComponent := r.find(rd.output.name)
	outputComponent.producer = &reaction
}

func (r *refinery) refine(amount int) error {
	f := r.find("FUEL")
	_, err := f.take(amount)
	return err
}

type RefineError string

func (e RefineError) Error() string {
	return string(e)
}

func (c *component) take(n int) (int, error) {

	if c.silo < n && c.producer == nil {
		return 0, RefineError(fmt.Sprintf("%s needed %d but only %d available", c.name, n, c.silo))
	}

	for c.silo < n {
		made, err := c.producer.make()
		if err != nil {
			return 0, err
		}
		c.produced += made
		c.silo += made
	}
	c.consumed += n
	c.silo -= n
	return n, nil
}

func (r *reaction) make() (int, error) {
	for _, in := range r.inputs {
		if _, err := in.c.take(in.n); err != nil {
			return 0, err
		}
	}
	return r.output, nil
}

