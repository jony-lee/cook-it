package main

import (
	"fmt"
	"sync"
)

type example struct {
	name string
}

var instance *example

// // GetExp1 ...
// //懒汉模式的单例设计模式
// func GetExp1() *example {
// 	if instance == nil {
// 		instance = &example{}
// 	}
// 	return instance
// }

// // GetExp2 ...
// //饿汉模式的单例设计模式
// func init() {
// 	instance = &example{}
// 	instance.name = "init"
// }
// func GetExp2() *example {
// 	return instance
// }

var once sync.Once

// GetExp3 ...
//并发安全的懒汉单例设计模式
func GetExp3() *example {
	once.Do(func() {
		instance = &example{}
		instance.name = "once"
	})
	return instance
}

//单例设计模式比较简单,就是为了避免创建多个对象实例
//每次获取对象实例的时候都返回全局的对象实例
//懒汉是在第一次调用时创建实例并返回,延迟分配提高启动加载速度
//饿汉是在软件启动的时候创建实例,什么时候需要什么时候返回
func main() {
	a := GetExp3()
	fmt.Printf("%+v\n", a) // output for debug
}
