package main

import "fmt"

// 设计模式

//// 饿汉模式
//type singleton struct {
//}
//
//var ins *singleton = &singleton{}
//
//func GetInsOr() *singleton {
//	return ins
//}

//// 懒汉模式
//type singleton struct {
//}
//
//var ins *singleton
//
//func GetInsOr() *singleton {
//	if ins == nil {
//		ins = &singleton{}
//	}
//	return ins
//}

//// 加锁
//import "sync"
//
//type singleton struct {
//}
//
//var ins *singleton
//var mu sync.Mutex
//
//func GetInsOr() *singleton {
//	if ins == nil {
//		mu.Lock()
//		if ins == nil {
//			ins = &singleton{}
//		}
//		mu.Unlock()
//	}
//	return ins
//}
//

//type singleton struct {
//}
//
//var ins *singleton
//var once sync.Once
//
//func GetInsOr() *singleton {
//	// 使用once.Do可以确保 ins 实例全局只被创建一次，once.Do 函数还可以确保当同时有多个创建动作时，只有一个创建动作在被执行。
//	once.Do(func() {
//		ins = &singleton{}
//	})
//	return ins
//}

// 工厂模式
type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() {
	fmt.Printf("name:")
}
