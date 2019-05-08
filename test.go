package main

import "fmt"



func main() {
	//var  i, j, k int
	var f int
	f = 5
	fmt.Println(f)
	fmt.Println("f is of type %T\n", f)
	
	var x float64 = 20.0
	
	y := 42
	fmt.Println(x)
	fmt.Println(y)
	
	
	
	var a, b, c = 3, 4, "foo"  
	
   fmt.Println(a)
   fmt.Println(b)
   fmt.Println(c)
   fmt.Printf("a is of type %T\n", a)
   fmt.Printf("b is of type %T\n", b)
   fmt.Printf("c is of type %T\n", c)
   
   
   //CONSTANTS
   
   const LENGTH int = 10
   const WIDTH int = 5   
   var area int

   area = LENGTH * WIDTH
   fmt.Printf("value of area : %d\n", area)
   area = max(a,b)
   fmt.Printf("value of area : %d", area)
}
/* function returning the max between two numbers */
func max(num1, num2 int) int {
   /* local variable declaration */
   var result int

   if (num1 > num2) {
      result = num1
   } else {
      result = num2
   }
   return result
}


/*
RETURNING MULTIPLE VALUES FROM FUNCTION
package main

import "fmt"

func swap(x, y string) (string, string) {
   return y, x
}
func main() {
   a, b := swap("Mahesh", "Kumar")
   fmt.Println(a, b)
}
*/

/*
CALL BY REFERENCE
package main

import "fmt"

func main() {
   /* local variable definition 
   var a int = 100
   var b int = 200

   fmt.Printf("Before swap, value of a : %d\n", a )
   fmt.Printf("Before swap, value of b : %d\n", b )

   /* calling a function to swap the values.
   * &a indicates pointer to a ie. address of variable a and 
   * &b indicates pointer to b ie. address of variable b.
   
   swap(&a, &b)

   fmt.Printf("After swap, value of a : %d\n", a )
   fmt.Printf("After swap, value of b : %d\n", b )
}
func swap(x *int, y *int) {
   var temp int
   temp = *x    /* save the value at address x 
   *x = *y    /* put y into x
   *y = temp    /* put temp into y 
}
*/

/*
STRING LENGTH
var greeting =  "Hello world!"
   
   fmt.Printf("String Length is: ")
   fmt.Println(len(greeting))  
*/

/* CONCATENATING STRINGS

import ("fmt" "math" )"fmt" "strings")

func main() {
   greetings :=  []string{"Hello","world!"}   
   fmt.Println(strings.Join(greetings, " "))
}
*/

/* ARRAY
var balance [10] float32
var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
var balance = []float32{1000.0, 2.0, 3.4, 7.0, 50.0}
balance[4] = 50.0
float32 salary = balance[9]

var n [10]int /* n is an array of 10 integers */
   //var i,j int

   /* initialize elements of array n to 0 */         
   //for i = 0; i < 10; i++ {
    //  n[i] = i + 100 /* set element at location i to i + 100 */
   }
   
   /* output each array element's value */
   //for j = 0; j < 10; j++ {
    //  fmt.Printf("Element[%d] = %d\n", j, n[j] )
   }
//*/
/*
MULTIDIMENS ARRAY
var threedim [5][10][4]int
var arrayName [ x ][ y ] variable_type
a = [3][4]int{  
   {0, 1, 2, 3} ,   /*  initializers for row indexed by 0 */
   //{4, 5, 6, 7} ,   /*  initializers for row indexed by 1 */
   //{8, 9, 10, 11}   /*  initializers for row indexed by 2 */
}
//*/
/* int val = a[2][3]
/* an array with 5 rows and 2 columns
   var a = [5][2]int{ {0,0}, {1,2}, {2,4}, {3,6},{4,8}}
   var i, j int

   /* output each array element's value 
   for  i = 0; i < 5; i++ {
      for j = 0; j < 2; j++ {
         fmt.Printf("a[%d][%d] = %d\n", i,j, a[i][j] )
      }
   }
*/

