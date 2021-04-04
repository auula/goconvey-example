package testing

import(
	"errors"
)

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

func Add(x,y int) int {
	return x + y
}

func  Subtract(x,y int) int {
	return x - y
}

func Division(x,y int)(int,error){
	if y == 0 {
		return 0,errors.New("被除数不能为零")
	}
	return x / y,nil
}

func Multiply(x,y int) int{
	return x * y
}