package usage

import (
	"fmt"

	"github.com/wuxiaoxiaoshen/go-version/common"
)

func NewHello() {
	var hello common.Hello
	fmt.Println(hello.Print())
	fmt.Println(hello.Println("Python"))
}
