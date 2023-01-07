package main

import (
	"fmt"
	"math"
)

// go语言中的浮点数采用IEEE-754标准的表达方式，定义了两种类型：float32 & float64
// float32 为单精度浮点数，可精确到小数点后7位
// float64 位双精度浮点数，可精确到小数点后15位
func main() {
	var f1 float32 //定义单精度浮点数
	var f2 float64 //定义双精度浮点数
	f1 = 10.21
	f2 = 10.1
	fmt.Println(f1, f2)
	f3 := 10.2 //浮点数自动推导时会推导为float64
	//f1=f3 //此时类型不同，会编译异常
	f1 = float32(f3) //可类型强转
	fmt.Println(f3)  //在实际开发中，应该尽可能地使用 float64 类型，因为 math 包中所有有关数学运算的函数都会要求接收这个类型。
	//浮点数不是一种精确的表达方式，因为二进制无法表示所有十进制小数，
	floatValue4 := 0.1
	floatValue5 := 0.7
	floatValue6 := floatValue4 + floatValue5 //0.7999999999999999
	fmt.Println(floatValue6)
	///@split -> 浮点数比较 ,因为浮点数来表达数字是不准确的，一般的实现方式是设置一个可以接受的差异值进行比较
	p := 0.1 //可接受的比较差异
	// 判断 f2 与 f3 是否相等
	//这个判断语句在math包内的语句为 f3-f2 < p
	if math.Dim(f3, f2) < p {
		fmt.Println("f2 和 f3 相等")
	}
}
