package main

import (
	"fmt"
	"practice/calculator"
	"practice/filehandler"
)

func calculation(calc *calculator.Calculator) {
	var a, b float64
	var op string
	fmt.Print("Enter a var: ")
	if _, err := fmt.Scan(&a); err != nil {
		fmt.Println("❌ Invalid input for 'a'. Please enter a number.")
		return
	}

	// Input b
	fmt.Print("Enter b var: ")
	if _, err := fmt.Scan(&b); err != nil {
		fmt.Println("❌ Invalid input for 'b'. Please enter a number.")
		return
	}

	// Input operator
	fmt.Print("Enter Operator (+ - * /): ")
	if _, err := fmt.Scan(&op); err != nil {
		fmt.Println("❌ Invalid input for operator.")
		return
	}
	var result float64
	switch op {
	case "+":
		result = calc.Add(a, b)
		fmt.Println(result)
	case "-":
		result = calc.Subtract(a, b)
		fmt.Println(result)
	case "*":
		result = calc.Multiply(a, b)
		fmt.Println(result)
	case "/":
		result = calc.Division(a, b)
		fmt.Println(result)
	default:
		fmt.Println("Invalid Operator")
		return
	}
	operation := calculator.Operation{
		Operands: [2]float64{a, b},
		Operator: op,
		Result:   result,
	}
	// Important: append() adds items to the end of a slice and returns the new slice.
	calc.History = append(calc.History, operation)
}

func calculatorApp() {
	calc := calculator.NewCalculator()
outer:
	for {
		var command string
		fmt.Println("c - calculate\th-print history\te-exit\td-delete history")
		fmt.Scan(&command)
		switch command {
		case "c":
			calculation(calc)
		case "h":
			calc.PrintHistory()
		case "d":
			calc.ClearHistory()
		case "e":
			break outer
		}
	}
}

func filehandle(filePath string) {
	fhanlder := filehandler.NewFileReader(filePath)
	if k := fhanlder.Open(); k != nil {
		fmt.Println("Error loading File", k.Error())
		return
	}
	fmt.Println("File opened successfully")
	// Your current code has a bug. What if Open() succeeds but something crashes before Close()? The file never closes! USE "defer"
	// DEFER ensures Close is called when function exits. No matter what happens (return, panic, normal exit)
	// DEFER must come before the panic
	defer func() {
		// Recover from panic

		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
		if close := fhanlder.Close(); close != nil {
			fmt.Println("Error closing file")
			return
		}
		fmt.Println("File Close Successfully")
	}()
	var inp int32
	fmt.Scan(&inp)
	k := 10 / inp // err if 0
	fmt.Println(k)
}

func main() {
	// calculatorApp()
	filehandle("./sample.txt")
}
