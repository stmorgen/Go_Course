package main

import "fmt"

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 { 
   return func(t float64) float64 {
		return 0.5 * a * t * t + v0*t + s0
   }
}

func main() {

	var acc float64	
	fmt.Printf("Enter acceleration: ")
	fmt.Scan(&acc)
	
	var vel float64	
	fmt.Printf("Enter initial velocity: ")
	fmt.Scan(&vel)
	
	var dis float64	
	fmt.Printf("Enter initial displacement: ")
	fmt.Scan(&dis)

	fn := GenDisplaceFn(acc, vel, dis)
	
	var time float64	
	fmt.Printf("Enter time: ")
	fmt.Scan(&time)
	
	res := fn(time)
	
	fmt.Printf("Result is: %f", res)
}