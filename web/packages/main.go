package main

import(
	"fmt"  // package fmt
	"packages/math"  // our math package
)

func main(){
	
	fmt.Println("PI =",math.PI)
	// fmt.Println(math.g)  // can NOT import

	income:=100000.00
	fmt.Println("Tax on income :",income*math.Tax) // math.Tax

    // math.Add()
	a:=math.Add(3,21)
	fmt.Println("Ans =",a)

	math.Sub(5,4)
}