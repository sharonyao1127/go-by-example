package main

import (
	"fmt"
	"sort"
	"strconv"
)

//输入aaabbcc，输出a3b2c2
//map
func main(){
	// 1、获取输入
	input1 := "aabbbcc"
	// 2、循环输入，并计数
	m1 := make(map[string]int)
	var counts = 1

	for _,k := range input1 {
		// 先看 map 里存不存在，不存在，赋值
		if _, ok := m1[string(k)] ; !ok {
			//添加
			m1[string(k)] = counts
			continue
		}
		// 存在，就加 1
		if _, ok := m1[string(k)] ; ok {
			m1[string(k)] ++
		}
	}
	// 5、最后拼接，map排序
	sorted_keys := make([]string, 0)
	for k,_ := range m1{
		sorted_keys = append(sorted_keys, k)
	}
	sort.Strings(sorted_keys)

	//拼接
	sum := ""
	for _, k := range sorted_keys {
		sum = sum + k + strconv.Itoa(m1[k])
	}
	fmt.Println(sum)

	//随机顺序
	//sum := ""
	//for key, value := range m1 {
	//	sum = sum + key + strconv.Itoa(value)
	//}
	//fmt.Println(sum)
}