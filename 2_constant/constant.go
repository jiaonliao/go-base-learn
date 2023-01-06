package main

import "fmt"

func main() {
	/****************************常量定义******************************/

	const c1 int = 1 //常规命名方式 其中value必须为在编译期间确定的值 常量的作用域与变量一致
	const c2 = 1     //推导类型的命名方式
	const (          //批量定义
		c3 = 1
		c4 = 1
	)
	const c5, c6 = 1, "a"   //多重赋值定义
	const c7, c8 int = 2, 3 //多重赋值,如规定的类型是一样的可以只写一个类型
	fmt.Println(c1, c2, c3, c4, c5, c6, c7, c8)

	/****************************预定义常量******************************/
	const i = iota //iota是一个特殊的常量，是一个可以被编译器修改的常量，当每一个const出现时就会被重置为0,每出现一个iota就会自增1
	const (        //iota被重置为0 todo 我不认为iota可以被理解为常量，没找到相关资料
		i1 = iota //递增
		i2 = iota
		i3 = iota
	)
	const (
		i4 = iota //在声明变量和常量时，如果赋值语句后的行为是一样的可以只写一次，后面的省略
		i5
		i6
	)
	const (
		i7 = iota * 2 //i7=0
		i8            //i8=2
		i9            //i9=4
	)
	fmt.Println(true, false, i)
	fmt.Println(i1, i2, i3)
	fmt.Println(i4, i5, i6)
	fmt.Println(i7, i8, i9)
}
