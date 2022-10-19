package main

import (
	"fmt"
	"log"
	lift1 "projectZ/2021-08-23/lift"
)

// 零值 也就是只定义 不赋值的
// string  		零值  ""
// int类型 		零值 0
// float类型 	零值 0.0
// struct 		是指针就是nil  不是指针是有内存空间的但成员变量都是零值
// map          无论怎么定义不赋值都是nil
// slice 		同上

func main() {
	// 本质上 所有的类型 在内存中都是二进制 在编码层面 它会被定义成int 也就地址值


	// 类型初始化 有很多种方式
	// 暂时说struct 类型

	// struct 类型的零值为nil
	// 凡是遇到nil 然后进行调用的 都会发生panic() 空指针异常 nil.成员变量 nil.成员函数 进行了操作就会出现问题
	// struct类型指针必须初始化赋值, 否则是nil
	// struct类型object 可以不初始化复制 它是一个struct 基本类型
	// 任何类型只有定义 没有赋值都是默认取零值

	// 指针和对象初始化区别,
	// 		对象只定义是可以直接使用的
	// 		指针只定义是nil  不能直接使用

	// var i1 lift1.Lift // 这样定义是object  但是这样成员变量的值 都是零值(看具体类型)
	var i1 *lift1.Lift // 定义了指针 但是没初始化 所以它是什么  指针只是定义就是nil
	// & 取对象的指针地址值
	i1 = &lift1.Lift{}
	i1.Down()
	// 0xc0000ca018 0x开头代表16进制
	log.Println(fmt.Sprintf("%v", &i1))
	// 0xc000088200
	as := "" // 用了内存空间 那它一定有内存地址 0xc000088200 它是一个二进制 也可以转换成16进制 也可以转换成十进制...
	log.Println(&as)
	m := map[string]string{
		"k": "k1",
	}
	// 内置函数调用
	delete(m, "k")

	// 自己写模拟编译器
	// 编译器会这样修改你的代码
	// 指针本身就是一个int
	key := 1
	k1 := false
	m1 := map[Type]Type1{
		&key: &k1,
	}
	k := "k"
	delete1(m1, &k)
}

type Type int
type Type1 int

// m map[Type]Type1
// m map[int]int
// m map[指针地址]指针地址
func delete1(m map[Type]Type1, key Type)  {

}


// Dog 这个类型 只不过在Go里面 我们通常对一个事物有N多特征和行为的情况 使用struct
// struct 它有自己的成员变量(众多特征)
type Dog struct {
	Name string
	Age int64
	// ....
	xx struct {
		Name string
		kkk struct{
			name string
		}
	}
}

// ...
// func ...


// type 命名 类型
// 枚举 听过吗?
// 电梯 1 2 3 4 5 6 7....层

func (dog *Dog) eat()  {
	// 写法1
	lift1 := First
	lift1.UP(First)

	// 写法2
	lift2 := Second
	UP1(First, lift2)
}

// java
//
// class lift {
//   public static void UP(lift t, lift lift) {
// 		xxxx
// 	}
//
// 	public void UP(lift t) {
// 		xxxx
// 	}
// }

// 把现实事物 抽象成了  类型(名字)  特征(成员变量) 行为(成员函数)
// 最基础 并且是人为认知的最小单位
type lift int64

// 这是一个 lift类型独有的成员函数  只能通过lift.func()进行调用
// UP() 如果成员函数是大写 说明外部 初始化了lift 是可以直接调用成员函数-UP()
// up() 如果成员函数是小写 说明外部 初始化了lift 不可以直接调用成员函数-up()
func (lift lift) UP(t lift) {
	// 上去T层
	lift = t
}

// 这是一个不受保护的全局函数 所有人都可以调用
// U 首字母大写 外部包是可以调用的
// u 首字母小写 只能本包调用
func UP1(t lift, lift lift) {
	// 上去T层
	lift = t
}

func (lift lift) Down(t lift) {
	// 下去T层
}

const (
	First lift = iota + 1
	Second
)

func call() {
	// liftFirst := First
}

// 这是什么? 我不知道
// lift 我知道是什么
func up(t lift) {

}