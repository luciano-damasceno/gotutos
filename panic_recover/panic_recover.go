package main

import (
	"fmt"
	"log"
)

func main() {
	divideByZero(4)
	fmt.Println("We survived dividing by zero!")
}

func panicRecover() {
	if err := recover(); err != nil {
		log.Println("Panic occurred:", err)
	}
}

func divideByZero(a int) {
	for i := 0; i < 5; i = i + 1 {
		/*
		 * Case 1 - Panic interrupts divideByZero() funtion.
		 *
		 * Recover from a panic automatically exits the current function
		 * execution, even inside a for loop. In that case, divideByZero().
		 */
		// defer panicRecover()
		// log.Println(a / i)

		/*
		 * Case 2 - For's body as a function.
		 *
		 * It's possible to encapsulate the for's body into another function to
		 * avoid stopping the execution of divideByZero() while keeping the for
		 * loop.
		 */
		divide(a, i)
		log.Println("Continues executing?")

		/*
		 * Case 3 - Anonimous function of Case 2
		 *
		 * It's possible to reach the same outcome with anonimous function.
		 */
		// func() {
		// 	defer panicRecover()
		// 	log.Println(a, i)
		// }()
		log.Println("Continues executing?")
	}
}

func divide(a, b int) {
	defer panicRecover()
	log.Println(a / b)
}
