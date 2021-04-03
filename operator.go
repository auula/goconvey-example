package testing

import "fmt"

type OperatorType map[string]func(x, y int) (result int)

var (
	Operator = make(OperatorType, 4)
)

// 为了测试编写的加减乘除函数
func init() {

	Operator["add"] = func(x, y int) int {
		return x + y
	}
	Operator["subtract"] = func(x, y int) int {
		return x - y
	}
	Operator["division"] = func(x, y int) int {
		return x / y
	}
	Operator["multiply"] = func(x, y int) int {
		return x * y
	}
}

func main() {
	fmt.Println(Operator["add"](1, 2))
}
