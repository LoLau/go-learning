package src

func main() {
	//a:= max(2, 3)

	// fmt.Printf("%d\n", a)

	c, d := swap("str1", "str2")

	println(c, d)
}

func max(a, b int) int {
	var result int

	if a > b {
		result = a
	} else {
		result = b
	}
	return result
}

func swap(x, y string) (string, string)  {
	return y, x
}
