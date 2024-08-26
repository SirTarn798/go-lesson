package point

import "fmt"

type Point struct {
	x int
	y int
}

func (p *Point) Print() {
	fmt.Println(p.x)
}