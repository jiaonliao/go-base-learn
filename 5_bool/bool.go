package main

import "fmt"

func main() {
	var v1 bool         //布尔值声明采用bool关键字
	v1 = true           //布尔类型只可以赋值为预定义常量 true  false
	v2 := 1 == 2        //可以通过类型推导得到bool值
	fmt.Println(v1, v2) //Go为强类型语言 不可以通过弱类型语言中的假值：0,"",null等进行布尔判断 例如js中 a=0 if(a){}
}
