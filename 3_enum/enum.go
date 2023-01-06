package main

// GO语言中不支持enum的关键字
// 此时可以采用如下方式定义枚举

func main() {
	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		numberOfDays
	)
	const (
		Sunday1 = "Sunday1"
		Monday2 = "Monday2"
	)
}
