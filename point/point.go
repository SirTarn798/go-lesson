package point

import "fmt"

type Point struct {
	X int
	Y int
}

func (p *Point) Print() {
	fmt.Println(p.X)
}