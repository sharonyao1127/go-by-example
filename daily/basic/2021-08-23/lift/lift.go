package lift

import "log"

// type(关键字) name type

type Lift struct {
	Name string    // string ""
	Age  int64     // int64 0
	XX   *struct{} // *struct nil
}

func NewLift() *Lift {
	return &Lift{}
}

// * 是指针意思
// Lift 是object
// * 和 object
// 暂时可以认为 指针和对象 是对等的

// UP() 如果成员函数是大写 说明外部 初始化了lift 是可以直接调用成员函数-UP()
// up() 如果成员函数是小写 说明外部 初始化了lift 不可以直接调用成员函数-up()
func (lift Lift) UP() {
	log.Println("call UP()")
}

func (lift Lift) Down() {
	log.Println("call Down()")
}
