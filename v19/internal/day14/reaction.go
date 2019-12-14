package day14

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

func (r refinery) init() {
	r.find("ORE").silo = 1000000000000
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

func (r *refinery) refine() {
	f := r.find("FUEL")
	f.take(1)
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

