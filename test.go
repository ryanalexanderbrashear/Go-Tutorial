package main

import "fmt"

func main() {

   fmt.Println("Hello, World!")

   fmt.Println("\n")

   /* Tokens in Go */
   /* Tokens - keywords, identifiers, constants, string literals, symbols that make up a program */

   /* Don't need ;s in Go! Detects the line separator key as a statement terminator */ 

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
