package intcode

type Inputter interface {
	Set(int)
	Input() int
}

type Outputter interface {
	Output(int)
	Close()
}

type nilio struct{}

func (n *nilio) Set(ignored int) {
}

func (n *nilio) Input() int {
	return 0
}

func (n *nilio) Output(ignored int) {
}

func (n *nilio) Close() {
}

type ChannelInput struct {
	input chan int
}

func (c *ChannelInput) Set(value int) {
	c.input <- value
}

func (c *ChannelInput) Input() int {
	return <- c.input
}

type ChannelOutput struct {
	output chan int
}

func (c *ChannelOutput) Output(n int) {
	c.output <- n
}

func (c *ChannelOutput) Close() {
	close(c.output)
}

type ValueInput struct {
	value int
}

func (v *ValueInput) Input() int {
	return v.value
}

func (v *ValueInput) Set(value int) {
	v.value = value
}
