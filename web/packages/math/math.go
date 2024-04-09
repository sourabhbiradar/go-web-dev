package math

const PI = 3.142  // exported

var Tax = 0.3 // exported 
// recommanded to use 'const' & constant values

const g = 9.8  // unexported

func Add(x,y int)int{ // exported
	return x + y
}

func sub (x,y int)int{ // unexported
	return x - y
}