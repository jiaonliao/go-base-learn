package main

import "fmt"

// go支持复数 即x=y+zi格式的复数
// x称之为实部
// zi称之为虚部
// 支持complex64 ： 32位实部和32位虚部
// 支持complex128 ： 64位实部与虚部
func main() {
	var com1 complex64 //定义复数
	com1 = 1 + 2i
	fmt.Println(com1)
	r1 := real(com1)                                    //获取实部
	i1 := imag(com1)                                    //获取虚假部
	fmt.Printf("r1 type of %T ,value of %v \n", r1, r1) //实部类型float32
	fmt.Printf("i1 type of %T ,value of %v", i1, i1)    //虚部类型float32
	//从上述打印可知，实部和虚部的类型也是浮点型，进行计算和比较时也需要注意精度的问题
}
