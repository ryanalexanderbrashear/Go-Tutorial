package main

import "fmt"

func main() {

	/* Go Operators */

	/* An operator is a symbol that tells the compiler to perform specific mathematical or logical manipulations. */

	/* Go Operator types:

	* Arithmetic Operators
	* Relational Operators
	* Logical Operators
	* Bitwise Operators
	* Assignment Operators
	* Misc. Operators

	 */

	/* Arithmetic Operators */

	/*

		+ adds two operands
		- subtracts two operands
		* multiplies two operands
		/ divides two operands
		% modulus operator
		++ incrememnt operator
		-- decrement operator

	*/

	/* Relational Operators */

	/*

		== equivalence
		!= not equivalent
		> greater than
		< less than
		>= greater than or equal
		<= less than or equal

	*/

	/* Logical Operators */

	/*

		&& logical and
		|| logical or
		! logical not

	*/

	/* Bitwise Operators */

	/*

		& bitwise and
		| bitwise or
		^ bitwise exclusive or
		<< bitwise left shift
		>> bitwise right shift

	*/

	/* Assignment Operators */

	/*

		= assignment
		+= add + assign
		-= subtract + assign
		*= multiply + assign
		/= divide + assign
		%= modulus + assign
		<<= left shift + assign
		>>= right shift + assign
		&= bitwise and + assign
		^= bitwise exclusive or + assign
		|= bitwise inclusive or + assign

	*/

	/* Misc. Operators */

	/*

		& - returns the address of a variable - &a; provides actual address of the variable
		* - pointer to a variable - *a; provides pointer to a variable

	*/

	/* Operator Precedence in Go */

	/*

		Operator precedence determines the grouping of terms in an expression. This affects how an expression is evaluated. Certain operators have higher precedence than others; for example, the multiplication operator has higher precedence than the addition operator.

		Operator types in order of precedence:
		Postfix
		Unary
		Multiplicative
		Additive
		Shift
		Relational
		Equality
		Bitwise AND
		Bitwise XOR
		Bitwise OR
		Logical AND
		Logical OR
		Conditional
		Assignment
		Comma

	*/

	/* Go Decision Making */

	/*

		Decision making structures require that the programmer specify one or more conditions to be evaluated or tested by the program, along with a statement or statements to be executed if the condition is determined to be true, and optionally, other statements to be executed if the condition is determined to be false.

		Types of decision statements in Go:

		If Statement
		If...Else Statement
		Nested If Statements
		Switch Statement
		Select Statement

	*/

	/* If Statement */

	/* local variable definition */
	var a int = 10

	/* check the boolean condition using if statement */
	if a < 20 {
		/* if condition is true then print the following */
		fmt.Printf("a is less than 20\n")
	}
	fmt.Printf("value of a is : %d\n", a)

	/* If...Else Statement */

	/* local variable definition */
	var b int = 100

	/* check the boolean condition */
	if b < 20 {
		/* if condition is true then print the following */
		fmt.Printf("b is less than 20\n")
	} else {
		/* if condition is false then print the following */
		fmt.Printf("b is not less than 20\n")
	}
	fmt.Printf("value of b is : %d\n", b)

	/* Nested If Statement */

	/* local variable definition */
	var c int = 100
	var d int = 200

	/* check the boolean condition */
	if c == 100 {
		/* if condition is true then check the following */
		if d == 200 {
			/* if condition is true then print the following */
			fmt.Printf("Value of c is 100 and d is 200\n")
		}
	}
	fmt.Printf("Exact value of c is : %d\n", c)
	fmt.Printf("Exact value of d is : %d\n", d)

	/* Switch Statement */

	/* local variable definition */
	var grade string = "B"
	var marks int = 90

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}
	switch {
	case grade == "A":
		fmt.Printf("Excellent!\n")
	case grade == "B", grade == "C":
		fmt.Printf("Well done\n")
	case grade == "D":
		fmt.Printf("You passed\n")
	case grade == "F":
		fmt.Printf("Better try again\n")
	default:
		fmt.Printf("Invalid grade\n")
	}
	fmt.Printf("Your grade is  %s\n", grade)

	/* another switch example */
	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Printf("type of x :%T", i)
	case int:
		fmt.Printf("x is int")
	case float64:
		fmt.Printf("x is float64")
	case func(int) float64:
		fmt.Printf("x is func(int)")
	case bool, string:
		fmt.Printf("x is bool or string")
	default:
		fmt.Printf("don't know the type")
	}

}
