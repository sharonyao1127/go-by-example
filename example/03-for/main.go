package main

import "fmt"

func main() {
	// for init; condition; post { }
	// for condition { }
	// for { }
	// init： 一般为赋值表达式，给控制变量赋初值；
	// condition： 关系表达式或逻辑表达式，循环控制条件；
	// post： 一般为赋值表达式，给控制变量增量或减量。
	// for语句执行过程如下：
	// ①先对表达式 init 赋初值；
	// ②判别赋值表达式 init 是否满足给定 condition 条件，若其值为真，满足循环条件，则执行循环体内语句，然后执行 post，进入第二次循环，再判别 condition；否则判断 condition 的值为假，不满足条件，就终止for循环，执行循环体外语句。

	// 在初始化语句中计算出全部结果是个好主意

	s := "abc"

	for i, n := 0, len(s); i < n; i++ { // 常见的 for 循环，支持初始化语句。
		println(s[i])
	}

	n := len(s)
	for n > 0 { // 替代 while (n > 0) {}
		println(s[n]) // 替代 for (; n > 0;) {}
		n--
	}

	for { // 替代 while (true) {}
		println(s) // 替代 for (;;) {}
	}

	i1 := 1
	for {
		fmt.Println("loop")
		break
	}
	for j := 7; j < 9; j++ {
		fmt.Println(j)
	}

	for n := 0; n < 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
	for i1 <= 3 {
		fmt.Println(i1)
		i1 = i1 + 1
	}

	// for 和 for range有什么区别?
	s8 := []int{1, 2, 3, 4, 5}

	for i, v := range s8 { // 复制 struct slice { pointer, len, cap }。

		if i == 0 {
			s8 = s8[:3] // 对 slice 的修改，不会影响 range。
			s8[2] = 100 // 对底层数据的修改。
		}

		println(i, v)
	}

	// 主要是使用场景不同

	// for可以
	// 遍历array和slice
	// 遍历key为整型递增的map
	// 遍历string

	// for range可以完成所有for可以做的事情，却能做到for不能做的，包括
	// 遍历key为string类型的map并同时获取key和value
	// 遍历channel
}
