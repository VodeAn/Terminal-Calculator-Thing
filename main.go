package main

import (
	"fmt"
	"time"
)

var i, ii int
var o string

func main() {
Loop:
	for {
		fmt.Println("Enter Two Integer Numbers:")
		fmt.Scan(&i, &ii)
		fmt.Println("Enter one of the operators(+,-,*,/,%)")
		fmt.Scan(&o)
		switch o {
		case "+":
			fmt.Print(i + ii)
			break Loop
		case "-":
			fmt.Print(i - ii)
			break Loop

		case "*":
			fmt.Print(i * ii)
			break Loop
		case "/":
			fmt.Print(i / ii)
			break Loop
		case "%":
			fmt.Print(i % ii)
			break Loop
		default:
			fmt.Println("Not a valid operator")
			time.Sleep(1000)
		}

	}

}
