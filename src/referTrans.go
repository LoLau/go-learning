package src

func main() {
	a := 10
	b := 20

	println(a)
	println(b)

	swap(&a, &b)

	println(a)
	println(b)
}

func swap(x , y * int) {
	tmp := *x
	*x = *y
	*y = tmp
}
