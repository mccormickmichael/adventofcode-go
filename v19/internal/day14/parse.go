package day14

import (
	"strconv"
	"strings"
)

type componentDescriptor struct {
	name string
	count int
}

type reactionDescriptor struct {
	output componentDescriptor
	inputs []componentDescriptor
}

func parse(line string) reactionDescriptor {
	inout := strings.Split(line, " => ")

	return reactionDescriptor{
		inputs: parseInputs(inout[0]),
		output:parseComponent(inout[1]),
	}
}

func parseInputs(ins string) []componentDescriptor {
	tokens := strings.Split(strings.TrimSpace(ins), ", ")
	inputs := make([]componentDescriptor, len(tokens))
	for i, t := range tokens {
		inputs[i] = parseComponent(t)
	}
	return inputs
}

func parseComponent(c string) componentDescriptor {
	tokens := strings.Split(strings.TrimSpace(c), " ")
	count, _ := strconv.Atoi(tokens[0])
	return componentDescriptor{tokens[1], count}
}