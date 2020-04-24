package main

//创建者模式:将一个复杂对象的构建与它的表示分离,使得同样的创建过程可以创建不同的表示.
//以领导(director)-员工(productBuilder)-计算机(product)组装来描述建造者模式
import "fmt"

//生成器抽象接口
type Builder interface {
	New()
	BuildHost()
	BuildMonitor()
	BuildMouse()
	BuildKeyboard()
	GetResult() interface{}
}

//director通过接口函数调用屏蔽了具体的对象实现细节,只需要调用接口对应的函数即可
type Director struct {
	Builder Builder
}

func (d *Director) SetBuilder(builder Builder) {
	d.Builder = builder
}

//director在指挥builder时不需要知道Builder内部的具体细节
//director像一个领导,下属有多个员工,领导安排不同员工干不同的活
//但是具体怎么干都交给员工,领导统筹,以下就是统筹的方法
func (d *Director) Construct() *Computer {
	d.Builder.New()
	d.Builder.BuildHost()
	d.Builder.BuildMonitor()
	d.Builder.BuildMouse()
	d.Builder.BuildKeyboard()
	return d.Builder.GetResult().(*Computer)
}

//具体建造的产品
type Computer struct {
	Host     string //主机
	Monitor  string //显示器
	Mouse    string //鼠标
	Keyboard string //键盘
}

func (c *Computer) Show() {
	fmt.Printf("\n%s \n%s \n%s \n%s", c.Host, c.Monitor, c.Mouse, c.Keyboard)
}

//为什么要使用Computer和ComputerBuilder两层结构体,只有Computer层结构体不够吗?
//构造代码,实现Builder接口
type ComputerBuilder struct {
	Computer *Computer
}

func (cb *ComputerBuilder) New() {
	cb.Computer = &Computer{}
}

func (cb *ComputerBuilder) BuildHost() {
	cb.Computer.Host = "build host"
}
func (cb *ComputerBuilder) BuildMonitor() {
	cb.Computer.Monitor = "build Monitor"
}
func (cb *ComputerBuilder) BuildMouse() {
	cb.Computer.Mouse = "build Mouse"
}
func (cb *ComputerBuilder) BuildKeyboard() {
	cb.Computer.Keyboard = "build keyboard"
}

func (cb *ComputerBuilder) GetResult() interface{} {
	return cb.Computer
}

//那么建造者的优势在哪里呢?
//可扩展性在哪里?
//可扩展性源于接口,在接口不改变的前提下(产品组成不变),生产细节交给具体的实现者
//可以使用相同的步骤生产不同的产品(只要这两者产品生产步骤相同)
//如何理解创建者模式中的构建与表示分离
//我理解的是电脑产品和组装电脑产品的过程分离
func main() {
	//指挥者
	director := &Director{}
	//建造者
	builder := &ComputerBuilder{}
	director.SetBuilder(builder)
	//组装产品
	computer := director.Construct()
	//显示产品
	computer.Show()
}
