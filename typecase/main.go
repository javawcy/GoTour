package main

import (
	"fmt"
	"strconv"
)

var text string
var textInited string = "Hi"

const finalValue = "Final"

func main() {

	text = "Hello World"
	fmt.Println(text)
	fmt.Println(textInited)

	//AUTO VAR INITED
	textAutoInited := "AUTO"
	fmt.Println(textAutoInited)

	//比较难理解的基本类型

	//uint 无符号字符
	//byte utf8字节 = 8bit = unit8
	//rune unicode编码字符 = int32
	//complex 复数 （1 + 2i）

	//go零值问题
	//数值类的零值 = 0
	//字符串零值 = “”
	//布尔值零值 = false

	//go中数值类型转换需要显式标注
	var intValue int = 64
	fmt.Println(float64(intValue))
	fmt.Println(uint(intValue))

	fmt.Println(finalValue)

	//string类型需要使用strconv包进行转换
	fmt.Println(strconv.Itoa(1))
	fmt.Println(strconv.Atoi("1"))
	fmt.Println(strconv.ParseFloat("0.32", 64))
	fmt.Println(strconv.FormatFloat(0.32, 'f', -1, 64))
}
