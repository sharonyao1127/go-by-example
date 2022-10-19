package main

import (
	"log"
)

func main() {
	基础数据结构
	整数
	浮点数
	字符串 默认实现 C CPP 没有官方实现 char
	// 删除类函数 从xxx删xxx
	// map kv 键值对 数据结构 Go内部实现了  map
	// key: value
	// 小姚: 对应了你

	从 m(map) 删除 key 为 key的键值对
	delete(m map[key]value, key Type)

	泛型


	// 函数名 见名知意
	// 入参 函数实现的前提
	// 出参 函数实现的结果
	func 函数名(入参) 出参
}

// 算法
入参
过程中申请1G内存空间 去做计算 然后得到了结果
xxx
xx
xx
xxx
申请内存空间
xxxx
xxx


bool true false 8 4字节 结果

a := 0 // 他是什么类型 int 几个字节

int8 int16 int32 int64
int
b := 0 // int
func TestMap()  {
	存储设备中(内存)
	create(内存, 数据结构)
	xxx进程申请了 一块内存区域(1G , 100M) 二进制 存储到这个内存区域中
	名字叫 m -> 一块内存区域
	m是指针  指针也是一块内存区域 4字节
	这个指针存储一块区域的 内存地址
	m是指针 同时 指针存储了 map对象的内存区域地址  存储是起始内存区域 offset
	从内存起始地址 通过offset 扫描
	0x000001
	offset 1000
	0x000001 - > 0x001000

	array  arr := [8]int{}
	0x000010 -> 0x000018
	slice  申请连续内存空间

	m := map[string]string{
		"姚": "你",
		"余": "我",
	}
	// 作用
	delete(m, "余")
	m = map[string]string{
		"姚": "你",
	}
}

type Slice struct {
	Len int64 指当前slice 元素有多少个
	Cap int64 指当前slice 空间有多大(能容纳多少个元素)
	Array unsafe.Pointer 一段连续的内存空间 是一个数组
}

func GetLen() {
	sliece := make([]string, 0 , 0, 0, 0 ,0 ,0 ,0 )
	slice := []string{"1", "2", "3"}
	len := len(slice)
	// 输出了几?
	log.Println(len)
	r("1", "2", "3", "4", "5", "6")
	r()

	r2("2", "3", "4")
	r2()
	sli := []string{"2", "3"}
	r2(sli)

	start()
	d := end()

	div(d)
}

func r(str ...string) {
	for k := range str {
		log.Println(str[k])
	}
}

func r2(str []string) {

}

// div 除法
func div(d int) int {
	return 1 / d
}

func mapTest() {
	var m map[string]string // nil  map 零值 = nil
	m = map[string]string{} // 有内存地址
	m["k"] = "v"

	// 会不会panic()
}