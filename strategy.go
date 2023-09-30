package main

import "fmt"

func main() {
	context := NewContext(&AddStrategy{})

	result := context.ExecuteOperation(10, 5)
	fmt.Printf("res of addition: %d\n", result)

	context.SetStrategy(&SubtractStrategy{})

	result = context.ExecuteOperation(10, 5)
	fmt.Printf("res of subtraction: %d\n", result)
}

type Strategy interface {
	ExecuteStrategy(int, int) int
}
type AddStrategy struct{}

func (s *AddStrategy) ExecuteStrategy(a, b int) int {
	return a + b
}

type SubtractStrategy struct{}

func (s *SubtractStrategy) ExecuteStrategy(a, b int) int {
	return a - b
}

type Context struct {
	strategy Strategy
}

func NewContext(strategy Strategy) *Context {
	return &Context{strategy}
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) ExecuteOperation(a, b int) int {
	return c.strategy.ExecuteStrategy(a, b)
}
