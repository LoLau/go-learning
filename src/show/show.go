package show

import (
	"fmt"
	learn_pkg "../learn"  //为包设置别名
)

func init() {
	fmt.Println("show init ")
}

func Show()  {
	learn_pkg.Learn()
}
