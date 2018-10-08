package src

import "fmt"

type Circle struct {
	radius float64
}

// 此方法属于Circle类型的对象
func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}

func main() {
	var c1 Circle
	c1.radius = 10.00
	fmt.Println(c1.getArea())
}
