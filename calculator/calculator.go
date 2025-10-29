package calculator // ✅ lowercase package name

import (
	"fmt"
)

// Calculator struct
type Calculator struct {
	Memory  float64
	History []Operation
}

// Operation struct

type Operation struct {
	Operands [2]float64
	Operator string
	Result   float64
}

// Constructor

func NewCalculator() *Calculator {
	return &Calculator{
		Memory:  0,
		History: make([]Operation, 0), // creating empty slice
	}
}

// Add
func (c *Calculator) Add(a float64, b float64) float64 {
	return a + b
}

// Subtract
func (c *Calculator) Subtract(a float64, b float64) float64 {
	return a - b
}

// Multiply
func (c *Calculator) Multiply(a float64, b float64) float64 {
	return a * b
}

// Division
func (c *Calculator) Division(a float64, b float64) float64 {
	return a / b
}

// Print History
func (c *Calculator) PrintHistory() {
	fmt.Println("Printing History")
	fmt.Println("History Size", c.History)
	for i := len(c.History) - 1; i > -1; i-- {
		fmt.Printf("%v %s %v = %v\n", c.History[i].Operands[0], c.History[i].Operator, c.History[i].Operands[1], c.History[i].Result)
	}
}

func (c *Calculator) ClearHistory() {
	c.History = make([]Operation, 0)
}
