package main

func main() {
	var a int = 3
	var p *int
	p = &a
	println("Value of a: ", a)
	println("The memory address of a: ", p)
	println("Value of a with pointer: ", *p)
	println("Address of a: ", &a)
	println("Address of a: ", p)
}
