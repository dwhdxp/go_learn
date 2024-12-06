package lib1

import (
	"fmt"
	_ "project/lib2"
)

func init() {
	fmt.Println("lib1 init ...")
}
