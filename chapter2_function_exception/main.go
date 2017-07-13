// chapter_function_exception project main.go
package main

import (
	"errors"
	"fmt"
)

//error 接口
// Go 语言引入
//type error interface{
//	Error() string
//}

type PathError struct {
	Op   string
	Path string
	Err  error
}

type MyError struct {
	Msg string
}

func (e *MyError) Error() string {
	return e.Msg
}

func (e *PathError) Error() string {
	return e.Op + " " + e.Path + ": " + e.Err.Error()
}

//关键字 defer
// panic recover关键字
// recover()函数用于终止错误处理流程。一般情况下，recover()应该在一个使用defer关键字的函数中执行
// 以有效截取错误处理流程。如果没有在发生异常的goroutine中明确调用恢复过程(使用recover关键字），会
// 导致该goroutine所属进程的进程打印异常信息后直接退出。

func main() {
	var err error

	err = &PathError{"stat", "abc.txt", &MyError{"unknow reason"}}

	fmt.Println(err.Error())

	err2 := &PathError{"stat2", "cde.txt", errors.New("just for test")}
	fmt.Println(err2.Error())

	var err3 error
	err3 = &PathError{"stat", "abc2.txt", &MyError{"unknow reason"}}
	if e, ok := err3.(*PathError); ok && e.Err != nil { //注意这里variable.(*type)中注意这里variable必须是一个接口类型
		fmt.Println("run here just for test")
	}

}
