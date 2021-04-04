package testing

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// 单元测试函数和普通go测试命名一样
func TestAdd(t *testing.T) {

	//  Convey 函数 第一个参数是测试名称
	//  t = *testing.T
	// 	func = 自定义测试逻辑
	Convey("两数相加测试，11 + 11 = 22 ？", t, func() {
		x, y := 11, 11
		// So 函数比较结果 ShouldEqual 是相等的意思
		So(Operator["add"](x, y), ShouldEqual, 22)
	})
}

func TestSubtract(t *testing.T) {
	Convey("测试两数相减，22 - 11 != 22 ？", t, func() {
		x, y := 22, 11
		// So 函数比较结果 ShouldEqual 是相等的意思
		So(Operator["subtract"](x, y), ShouldNotEqual, 22)
	})
}

func TestMultiply(t *testing.T) {
	Convey("将两数相乘，11 * 2 = 22 ？", t, func() {
		x, y := 11, 2
		So(Operator["multiply"](x, y), ShouldEqual, 22)
	})
}

func TestDivision(t *testing.T) {
	Convey("将两数相除", t, func() {
		x, y := 10, 2
		Convey("除以非 0 数", func() {
			n,_ := Division(x, y)
			So(n, ShouldEqual, 5)
		})

		Convey("除以 0", func() {
			y = 0
			_,err := Division(x, y)
			So(err, ShouldNotBeNil)
		})
	})
}
