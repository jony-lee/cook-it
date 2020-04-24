package main

import "fmt"

//USB接口,传输USB数据
type Usb interface {
	UsbData() string
}

//TypeC接口,传输TypeC数据
type TypeC interface {
	TypeCData() string
}

type TypeCImp struct {
}

// TypeCData ...
func (TypeCImp) TypeCData() string {
	return "TypeC data"
}

// NewTypeC ...
func NewTypeC() TypeC {
	return TypeCImp{}
}

//转换器包含TypeC接口
type adapter struct {
	TypeC
}

// NewAdapter ...
func NewAdapter(t TypeC) Usb {
	return adapter{
		TypeC: t,
	}
}

// (a adapter)UsbData ...
//这是个USB转换器,实现了USB传输数据的方法,转换器重写USB方法,使得USB可以读取TypeC接口数据
func (a adapter) UsbData() string {
	return a.TypeCData()
}

//其实我想实现USB和TypeC接口之间的智能传输,即不论是USB还是TypeC只要插入之后
//我都能通过出口的接口读取到对应进口的数据
//这样就实现了 智能转换器模式 但是目前还没有成功
func main() {
	//新建一个TypeC设备
	t := NewTypeC()
	//将Type设备转成USB设备
	u := NewAdapter(t)
	//u通过USBDate数据传输成功读取到了TypeC设备t 的数据
	res := u.UsbData()
	//USB设备获取到了TypeC设备的传输数据
	fmt.Printf("%+v\n", res) // output for debug

}
