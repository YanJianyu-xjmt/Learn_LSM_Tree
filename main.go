package main 

import (
	"fmt"
)

type u = int

func util() u{
	var p int
	p = 1
	return p
}
func main(){
	fmt.Println(util())
}