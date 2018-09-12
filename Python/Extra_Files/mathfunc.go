package main

import (
"fmt"
"math"
)
func main(){
//Like float32, float64 is another float data type in Go.Most of the functions in Math library expects float input to be of float64 type.
var num1 float64=234.01
var num3 float64=-27.01
fmt.Println("The smallest integer greater than or equal to num1:",math.Ceil(num1))
fmt.Println("The largest integer smaller than or equal to num1:",math.Floor(num1))
fmt.Println("The absolute value of num3:",math.Abs(num3))

}