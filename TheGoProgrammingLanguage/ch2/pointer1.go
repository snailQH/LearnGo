package main

import(
	"fmt"
)

func main(){
	v := 1
	incr(&v)
	fmt.Println(incr(&v))
	fmt.Println(incr(&v))
	fmt.Println("P")
	p := new(int)
	fmt.Println(*p)
	*p=2
	fmt.Println(*p)
	fmt.Println(p)
}

func incr(p *int) int {
	*p++
	fmt.Println(p)
	return *p
}