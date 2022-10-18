package main

import "fmt"

//去除重复字符串
//ffdcvf
//利用map的key唯一性
func removedup(str string)string{
	//定义一个map
	targer := make(map[string]int)

	//循环原字符串
	var counts = 1
	for _,i:= range str{
		//fmt.Println(string(i))
		//判断i在不在map里，不在的话，添加，在的话，continue循环
		if _,ok := targer[string(i)]; !ok {
			targer[string(i)] = counts
		}else{
			counts += 1
		}
	}

	//最后循环map,输出key,连接起来，map是无序的
	result := ""
	for k,_ := range targer{
		result = result + k
	}
	return result
}


func main() {
	res := removedup("ffdcvf")
	fmt.Println(res)
}