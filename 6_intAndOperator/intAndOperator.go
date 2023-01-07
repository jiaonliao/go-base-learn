package main

import "fmt"

// go中支持如下整形
// +─────────+───────────────────+─────────────────────────────────+───────────────────────────────────────────+───────+
// | 类型      | 长度（单位：字节）  | 说明                            | 值范围                                     | 默认值  |
// +─────────+───────────────────+─────────────────────────────────+───────────────────────────────────────────+───────+
// | int8    | 1                 | 带符号8位整型                     | -128~127                                  | 0     |
// +─────────+───────────────────+─────────────────────────────────+───────────────────────────────────────────+───────+
// | uint8   | 1                 | 无符号8位整型，与 byte 类型等价     | 0~255                                     | 0     |
// +─────────+───────────────── ─+─────────────────────────────────+───────────────────────────────────────────+───────+
// | int16   | 2                 | 带符号16位整型                    | -32768~32767                              | 0     |
// +─────────+───────────────────+─────────────────────────────────+───────────────────────────────────────────+───────+
// | uint16  | 2                 | 无符号16位整型                    | 0~65535                                   | 0     |
// +─────────+───────────────────+─────────────────────────────────+───────────────────────────────────────────+───────+
// | int32   | 4                 | 带符号32位整型，与 rune 类型等价    | -2147483648~2147483647                    | 0     |
// +─────────+───────────────────+─────────────────────────────────+───────────────────────────────────────────+───────+
// | uint32  | 4                 | 无符号32位整型                    | 0~4294967295                              | 0     |
// +─────────+───────────────────+─────────────────────────────────+───────────────────────────────────────────+───────+
// | int64   | 8                 | 带符号64位整型                    | -9223372036854775808~9223372036854775807  | 0     |
// +─────────+───────────────────+─────────────────────────────────+───────────────────────────────────────────+───────+
// | uint64  | 8                 | 无符号64位整型                    | 0~18446744073709551615                    | 0     |
// +─────────+───────────────────+─────────────────────────────────+───────────────────────────────────────────+───────+
// | int     | 32位或64位         | 与具体平台相关                    | 与具体平台相关                               | 0     |
// +─────────+───────────────────+─────────────────────────────────+───────────────────────────────────────────+───────+
// | uint    | 32位或64位         | 与具体平台相关                    | 与具体平台相关                               | 0     |
// +─────────+───────────────────+─────────────────────────────────+───────────────────────────────────────────+───────+
// | uintptr | 与对应指针相同      | 无符号整型，足以存储指针值的未解释位  | 32位平台下为4字节，64位平台下为8字节            | 0     |
// +─────────+───────────────────+─────────────────────────────────+───────────────────────────────────────────+───────+
// Go支持的整形类型非常丰富，可以根据需要设置合适的类型，以节省内存空间
// 此外不同的整形类型之间会被认为是不同的类型，不能做自动转换
func main() {
	var v1 int8 = 8
	v2 := 1 //此处会被推导为int类型
	//v1=v2  此处因为类型不同，会报编译异常
	v1 = int8(v2) //类型强转可以完成类型转换和赋值操作
	fmt.Println(v1, v2)
	v3 := 010  //8进制赋值
	v4 := 0xff //16进制赋值
	v5 := 1e3  //E连乘赋值
	fmt.Println(v3, v4, v5)

	/**************************算术运算符***********************/
	v1 = v1 + 1  //加法运算
	v1 = v1 - 1  //减法运算
	v1 = v1 * 1  //乘法运算
	v1 = v1 % v1 //取余(取余运算只能用于整数)
	//v1 = v1 + v2 不同类型之间不能直接进行算术运算，此处会编译异常
	v1 = v1 + int8(v2) //强转类型后，可进行计算
	v1++
	v1-- //自增(++)/自减(--) 只能作为语句，不能作为表达式，且只能放置于变量后 --v1 会编译异常
	//var v6 int = v1++  编译异常，++不能作为表达式

	/**************************比较运算符***********************/

	//支持 > < == >= <= 和 != ,运算结果为bool值，不同类型之间不可以一起比较，go为强类型语言，无弱类型语言中 === !== 的机制

	var b1 = v1 > 1 //变量可以与字面量进行比较
	fmt.Println(b1)

	/**************************位运算符***********************/
	//位运算符以二进制的方式对数值进行运算，效率更高，性能更好，Go 语言支持以下这几种位运算符：
	// +────────+─────────+─────────────────────────────+
	// | 运算符    | 含义   | 结果
	// +────────+─────────+─────────────────────────────+
	// | x & y  | 按位与   | 把 x 和 y 都为 1 的位设为 1
	// +────────+─────────+─────────────────────────────+
	// | x | y  | 按位或   | 把 x 或 y 为 1 的位设为 1
	// +────────+─────────+─────────────────────────────+
	// | x ^ y  | 按位异或  | 把 x 和 y 一个为 1 一个为 0 的位设为 1
	// +────────+─────────+─────────────────────────────+
	// | ^x     | 按位取反  | 把 x 中为 0 的位设为 1，为 1 的位设为 0
	// +────────+─────────+─────────────────────────────+
	// | x << y | 左移     | 把 x 中的位向左移动 y 次，每次移动相当于乘以 2
	// +────────+─────────+─────────────────────────────+
	// | x >> y | 右移     | 把 x 中的位向右移动 y 次，每次移动相当于除以 2
	// +────────+─────────+─────────────────────────────+
	var intValueBit uint8
	intValueBit = 255
	intValueBit = ^intValueBit // 按位取反
	fmt.Println(intValueBit)   // 0
	intValueBit = 1
	intValueBit = intValueBit << 3 // 左移 3 位，相当于乘以 2^3 = 8
	fmt.Println(intValueBit)       // 8

	/**************************逻辑运算符***********************/
	//go支持如下运算符
	// +─────────+──────────────+─────────────────────────────────────────+
	// | 运算符     | 含义           | 结果
	// +─────────+──────────────+─────────────────────────────────────────+
	// | x && y  | 逻辑与运算符（AND）  | 如果 x 和 y 都是 true，则结果为 true，否则结果为 false
	// +─────────+──────────────+─────────────────────────────────────────+
	// | x || y  | 逻辑或运算符（OR）   | 如果 x 或 y 是 true，则结果为 true，否则结果为 false
	// +─────────+──────────────+─────────────────────────────────────────+
	// | !x      | 逻辑非运算符（NOT）  | 如果 x 为 true，则结果为 false，否则结果为 true
	// +─────────+──────────────+─────────────────────────────────────────+

	/**************************运算符优先级***********************/
	// 优先级从上到下排列为

	// ^（按位取反） !
	// *  /  %  <<  >>  &  &^
	// +  -  |  ^（按位异或）
	// ==  !=  <  <=  >  >=
	// &&
	// ||
}
