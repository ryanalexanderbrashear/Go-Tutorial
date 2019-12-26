package main

import "fmt"

func main() {

	fmt.Println("Hello World pt. 2")

	/* Go Constants */

	/* Constants/literals are fixed values that cannot be modified during program execution. Otherwise, constants are treated just like variables. */

	/* Integer Literals */

	/* Integer literals can be decimal, hexadecimal, or octal constants.

	   Prefixes are used to specify the base (radix) of the constant: 0x or 0X for hexadecimal, 0 for octal, and nothing for decimal.

	   An integer literal can also have a suffix that is a combination of U and L, for unsigned and long, respectively. The suffix can be uppercase or lowercase and can be in any order.

	   212         Legal
	   215u        Legal
	   0xFeeL      Legal
	   078         Illegal: 8 is not an octal digit
	   032UU       Illegal: cannot repeat a suffix

	   85         Decimal
	   0213       Octal
	   0x4b       Hexadecimal
	   30         Int
	   30u        Unsigned int
	   30l        Long
	   30ul       Unsigned long
	*/

	/* Floating-point Literals */

	/* A floating-point literal has an integer part, a decimal point, a fractional part, and an exponent part. You can represent floating point literals either in decimal form or exponential form.

	   While representing using decimal form, you must include the decimal point, the exponent, or both and while representing using exponential form, you must include the integer part, the fractional part, or both. The signed exponent is introduced by e or E.

	   3.14159       Legal
	   314159E-5L    Legal
	   510E          Illegal: incomplete exponent
	   210f          Illegal: no decimal or exponent
	   .e55          Illegal: missing integer or fraction
	*/

	/* Escape Character Sequences */

	/* When certain characters are preceded by a backslash, they will have a special meaning in Go. These are known as Escape Sequence codes which are used to represent newline (\n), tab (\t), backspace, etc.

	 \\ - \ character
	 \' - ' character
	 \" - " character
	 \? - ? character
	s \a - alert or bell
	 \b - backspace
	 \f - form feed
	 \n - newline
	 \r - carriage return
	 \t - horizontal tab
	 \v - vertical tab
	 \ooo - octal number of 1-3 digits
	 \xhh... - hexadecimal number of 1 or more digits

	*/

	fmt.Println("Hello\tWorld!") /* Example of using \t */

	/* String Literals */

	/* String literals or constants are enclosed in double quotes "". A string contains characters that are similar to character literals: plain characters, escape sequences, and universal characters.

	You can break a long line into multiple lines using string literals and separating them using whitespaces.

	            "hello, dear"

	            "hello, \

	            dear"

	            "hello, " "d" "ear"

	*/

	/* Const Keyword */

	/* Like var for variables, const is the keyword used when declaring constants in Go. */

	const LENGTH int = 10 /* It is considered best practice to declare constants using all capital letters */
	const WIDTH int = 5
	var area int

	area = LENGTH * WIDTH
	fmt.Printf("value of area : %d", area)

}
