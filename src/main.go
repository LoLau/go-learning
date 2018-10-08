package src

import (
	"fmt"
	"math"
	"strconv"
)

type newType int

type gopher struct {
}

type goInterface interface {
}

func main() {
	var a [1]int

	// var b uint8

	// var c byte

	fmt.Println(a)
	fmt.Println(math.MaxInt32)

	var c, _, e, f int = 1, 2, 3, 4

	fmt.Println(c)
	fmt.Println(e)
	fmt.Println(f)

	var x = 100.1
	var y = int(x)

	fmt.Println(x, y)

	// var t int = 65
	// q := string(t)
	// fmt.Println(q)

	// string 和 int 进行转换时要用strconv包
	var t int = 65
	q := strconv.Itoa(t)   //整数转字符串
	t, _ = strconv.Atoi(q) //字符串转整数
	fmt.Println(a)

	const (
		aa, bb = 1, 2
		cc, dd
	)

	fmt.Println(aa, bb, cc, dd)

	const (
		M = iota + 1
		T
		S
	)

	fmt.Println(M, T, S)

	const (
		B = 1 << (10 * iota)
		KB
		MB
		GB
		TB
	)

	fmt.Println(B, KB, MB, GB, TB)

	xx := 1
	xx++
	yy := xx
	fmt.Println(yy)
}
