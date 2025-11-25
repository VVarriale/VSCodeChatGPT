package main

import "fmt"

type Vertex struct {
	X int
	Y int
	z int
}

func (v *Vertex) String() string {
	return fmt.Sprintf("(%d,%d)", v.X, v.Y)
}

func printSomething(s fmt.Stringer) {
	fmt.Println(s.String())
}

func main() {
	v := &Vertex{X: 10, Y: 20}
	printSomething(v)
}
