package main

import "fmt"

func main() {

   fmt.Println("Hello, World!") /* Using the Println function from the fmt package (because Println is exported it is capitalized) */

   fmt.Println("\n")

   /* Syntax in Go */

   	/* Tokens - keywords, identifiers, constants, string literals, symbols that make up a program */

   	/* Don't need ;s in Go! Detects the line separator key as a statement terminator */

   	/* A Go identifier is a name used to identify a variable, function, or any other user-defined item. An identifier starts with a letter A to Z or a to z or an underscore _ followed by zero or more letters, underscores, and digits (0 to 9).

	identifier = letter { letter | unicode_digit }.

	Go does not allow punctuation characters such as @, $, and % within identifiers. Go is a case-sensitive programming language. Thus, Manpower and manpower are two different identifiers in Go. */

   /* Go Data Types */

      /* 4 major categories in Go:
         Boolean types - true/false
         Numeric types - integers, floating point decimals
         String types - strings
         Derived types - Pointer types, Array types, Structure types, Union types, Function types, Slice types, Interface types, Map types, Channel types
      */

      /* Array types and Structure types are referred to as aggregate types */

      /* Int types: uint8, uint16, uint32, uint64, int8, int16, int32, int64 */

      /* Float types: float32, float64, complex64, complex128 */

      /* Other numeric types: byte, rune, uint, int, uintptr */

   /* Go Variables */

      /* Variable names can be composed of letters, numbers, and the underscore character. They must begin with an underscore or a letter. Go is case-sensitive, so upper-case vs. lower-case letters matter */

      /* Valid variable declarations:
         var i, j, k int;
         var u, uh byte;
         var f, salary float32;
         var d = 42;
      */

      /* Static Typing */

   	var x float64 /* x is explicitly given the type of float64 */
   	x = 20.0
   	fmt.Println(x)
   	fmt.Printf("x is of type %T\n", x) /* Holy shit, you can pull the type of a variable so  easily */

   	fmt.Println("\n")

     /* Dynamic Typing */

   	var z float64 = 20.0

   	y := 42 /* y is dynamically typed here */
   	fmt.Println(z)
   	fmt.Println(y)
   	fmt.Printf("z is of type %T\n", z)
   	fmt.Printf("y is of type %T\n", y)

   	fmt.Println("\n") 

   	/* You can use type inference to declare mixed variables at the same time */
   	var a, b, c = 3, 4, "foo" 
	
   	fmt.Println(a)
   	fmt.Println(b)
   	fmt.Println(c)
   	fmt.Printf("a is of type %T\n", a)
   	fmt.Printf("b is of type %T\n", b)
   	fmt.Printf("c is of type %T\n", c)

      /* lvalues and rvalues */
   
   	/* lvalues - expressions that refer to a memory location (ex. variables) */
   	/* rvalues - refers to a data value that is stored at some location in memory (ex. numeric values */


}
