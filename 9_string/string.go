package main

import "fmt"

// 在GO语言中，字符串是一种基本类型，默认是使用UTF-8编码，当字符位ASCLL时,占用一个字节，其他字符需要占用2-4个字节
// 中文编码通常需要占用3个字节
func main() {
	var u1 string //字符类型声明
	u1 = "小明"     //赋值
	//len()内置函数可以取字符长度，此处的字符长度为6就是因为其占用的字节数为6
	//关于Printf的具体用法，可查阅Go语言SDK的api文档
	fmt.Printf("The length of \"%s\"  is %d \n", u1, len(u1))

	//Go语言中有如下转义字符
	fmt.Print("换行 \n")
	fmt.Print("回车 \r")
	fmt.Print("tab \t")
	fmt.Print("unicode 字符 \u662f")
	fmt.Print("反斜杠 \\")

	//多行字符串,类似es6中的模板字符串，但是没有填入变量的功能
	template := `
			strings 
				-str1
				-str2
				-str3`
	fmt.Println(template)

	//值不可变，字符串一旦被赋值后，其中字符的值是不能更换的
	str := "bbbbbbb"
	//str[0] = "a" 编译异常
	fmt.Println(str)

	//字符串拼接操作
	str1 := str + "str1"
	str1 = str1 + "str1"

	fmt.Println(str1)

	//字符串切片
	//在Go语言中，可以通过字符串切面来实现获取字串的功能 str=str[x:y]
	//切片功能是一个左闭右开的操作
	//有中文字符时需要注意
	longStr := "I am a very long string"
	shortStr := longStr[:5]   //取索引5（不包含5）之前的字符串
	shortStr1 := longStr[0:5] //取索引0-5（不包含5）的字符串
	shortStr2 := longStr[2:5] //取索引2-5（不包含5的字符串）
	shortStr3 := longStr[5:]  //取索引5之后的字符串
	fmt.Println(shortStr)
	fmt.Println(shortStr1)
	fmt.Println(shortStr2)
	fmt.Println(shortStr3)
}
