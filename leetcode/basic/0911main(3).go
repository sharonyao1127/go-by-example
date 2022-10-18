package main

import (
	"log"
)

// 函数
// 函数是什么?
// 关键字是什么?
func main() {
	// 调用函数 一定要带()
	// 不带() go编译器认为这是个变量
	// 函数本身也是一个类型
	// 也可以作为成本变量使用
	Output()

	// 例如
	// 编译通过
	fun := Output
	_= fun

	// 调用
	fun()

	result := sum(1, 2)
	log.Println(result)

	log.Println(sub(10, 5.5))
	log.Println(sub1(10, 5.5))
	char()
}

// 定义一个函数
// 结构
// 关键字 [结构体成员函数] 函数名(入参) 出参 {
//
// }
// 带出入参的成员函数
/*func (this struct) name(入参) 出参 {

}*/

// 带出入参的全局函数
/*func name(入参) 出参 {

}*/

// 纯全局函数 遵循首字母大写 外包可调用的包规则
/*func name() {

}
*/

// 函数用来做什么?
// 实现一些功能算法过程
// 首先有需求 我要干什么?
// 其次是分解需求 拆分需求的步骤

// 需求1. 输出固定的字符串 "函数"
func Output() {
	// 固定的字符串
	// 定义字符串
	var str = "函数"
	// 输出字符串
	log.Println(str)
}

// 需求:
// 编写一个函数, 这个函数用来计算两个整数的和 并且返回
// a int, b int  与 a, b int 定义完全一样
func sum(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func div(a, b int) int {
	return a / b
}

// 我们一直在写 同一个类型的计算 int int64 int32 float32 float64...
// 不同类型怎么做计算?
// 例: 减法
// 需求: 给定入参为 int64 与 float64 相减的结果
func sub(a int64, b float64) float64 {
	// c := a - b
	// 基础类型转换  string int float bool
	// 将int系列转换到float系列
	// 转换后有什么损失?
	// 基础类型转换 最好是从小的 转换到大的
	// int8 int16 int32 int64  这些都标识这个类型能装多大的值
	// float32 float64 也是意味着它能装多大的值 以及多少位精度的小数
	// 不能将int64 转换到float32   如果int64的值 超过了float32的max 那么就会出现异常
	// 只能从小到大 最起码要等于 float32能被什么转换? int8 int16 int32
	aFloat64 := float64(a)
	c := aFloat64 - b
	return c
}

// string 基础类型是什么?
// 字符 串 一堆字符组合起来的值
// 什么是字符

var s string
// 字符演示
func char() {
	// 字符串类型定义
	// a := ""
	// 字符类型定义
	// 就是一堆字符组合起来值 字符就是 人所认识的最小单位 a, b, c, d... 0, 1, 2, 3...
	// 汉字  汉是一个字\
	// char := ' '
	// char1 := 'l'

	// 字符类型
	// 也就是byte
	// var k rune
	// 字符串 就是字符类型的数组
	// var i []rune
	// 字符串最终是什么
	// var str []byte
	str := []byte{97, 98, 99, 100, 101}
	str1 := string(str)
	log.Println("[]byte: ", str)
	log.Println("转换后的字符串; ", str1)

	// 真实的字符集 变成字符串
	str2 := "a"
	str3 := "b"

	log.Println("str2 []byte: ", []byte(str2))
	log.Println("str3 []byte: ", []byte(str3))
}
/*int16 也就2个字节
utf-8 字符集
[
	字符A: 二进制表示: 1000 0000  也就128
	字符B: 二进制表示: 1000 0001  也就129
...
]*/

// 小数转换整数 直接使用类型转换
// 转换后 小数位还有没有? 怎么处理的?
// 小数位肯定没有了, 处理的方式是 直接舍去
// 12.10 -> int64    12
// 23.9999 -> int64  23
// 所以使用类型直接转换 要注意这些细节
func sub1(a int64, b float64) int64 {
	// math.Floor()

	bInt64 := int64(b)
	c := a - bInt64
	return c
}

// 在出参直接定义返回值 叫有字返回
// 在出参只定义类型 叫无字返回
func sum1(a int, b int) int {
	c := a + b
	return c
}

func sum2(a int, b int) int {
	var c int
	c = a + b
	return c
}