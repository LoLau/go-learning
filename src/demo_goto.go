package src

import (
	"fmt"
)

func main() {

	for i := 0; i < 10; i++ {
		for {
			fmt.Println(i)
			goto LABEL2
		}
	}
LABEL2:
}
