package main

import "fmt"

type Base interface {
	Play()
}

type Name struct{}

// Play ...
func (n Name) Play() {
	fmt.Printf("%+v\n", "Play name") // output for debug
}

type Decorater struct {
	Base
}

// Play ...
func (a Decorater) Play() {
	fmt.Printf("%+v\n", "Decorate name") // output for debug
	a.Base.Play()
}

// Wrap ...
func Wrap(b Base) Base {
	return Decorater{Base: b}
}

//装饰器模式利用装饰器将接口包起来,并重写接口方法以达到对原结构体装饰的作用
//通过Wrap函数将实现接口的原有方法转变成实现接口的装饰器方法
//可以看出Decorater扩展了原有的Name的能力
//但并未对Name相关的代码进行修改,实现了对扩展开放,对修改封闭.
func main() {
	var a Base = Name{}
	a.Play()    //执行Name的play操作
	a = Wrap(a) //装饰器将Name转变成了Decorator
	a.Play()    //输出Decorator的play操作
}
