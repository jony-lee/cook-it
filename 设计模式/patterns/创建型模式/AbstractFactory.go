package main

import (
	"fmt"
)

//利用门店-工厂-手机产品来理解抽象工厂生产模式

//工厂接口
type Factory interface {
	ProducePhone(Product)
}

//产品接口
type Product interface {
	Create()
}

//华为手机
type HuaweiProduct struct{}

func (h HuaweiProduct) Create() {
	fmt.Println("Huawei P50")
}

//苹果手机
type AppleProduct struct{}

func (a AppleProduct) Create() {
	fmt.Println("IPhone 12")
}

//小米手机
type XiaomiProduct struct{}

func (x XiaomiProduct) Create() {
	fmt.Println("Xiaomi 9")
}

//上海工厂
type FactoryShanghai struct{}

func (sh FactoryShanghai) ProducePhone(p Product) {
	fmt.Println("shanghai factory")
	p.Create()
}

//深圳工厂
type FactoryShenzheng struct{}

func (sz FactoryShenzheng) ProducePhone(p Product) {
	fmt.Println("shenzhen factory")
	p.Create()
}

//新建武汉工厂生产线
type FactoryWuhan struct{}

func (wh FactoryWuhan) ProducePhone(p Product) {
	fmt.Println("wuhan factory")
	p.Create()
}

//抽象工厂类型的意思是把工厂也进行抽象吗?
//现在有很多销售门店都需要进货,所以工厂都要进行加工
//假设现在需要更换苹果公司的产品线,比如iphone11换成iphone12
//那只需要修改工厂的生产逻辑就可以了,门店代码不用换

//为什么要写4个门店,最初写这段代码的时候只有一个门店
//根本无法感受到使用抽象工厂模式带来的好处
//可以想象,一旦门店增加,从门店来更换生产线非常不划算
//这样就能感受到抽象工厂模式带来的好处了

//同理,增加工厂以及增加产品线都只需在工厂方面进行修改
//无需修改门店(即客户方)的代码即可完成

//那么如果需要增加比如小米产品线
//很简单,只需添加小米产品接口

//update 抽象工厂好像还有点问题,因为我想让产品和工厂都抽象,之前的工厂是抽象的,但是产品并不是抽象的

func main() {
	//门店1 选择生产工厂
	var f Factory = &FactoryShanghai{}
	f.ProducePhone(AppleProduct{})

	f = &FactoryShanghai{}
	f.ProducePhone(XiaomiProduct{})
	// //门店2 选择生产工厂
	// f = &FactoryShanghai{}
	// f.ProduceApple().Create()

	// //门店3 选择生产工厂
	// f = &FactoryShanghai{}
	// f.ProduceApple().Create()

	// //门店4 选择生产工厂
	// f = &FactoryShanghai{}
	// f.ProduceApple().Create()

	// //在武汉开始进行生产苹果手机
	// f = &FactoryWuhan{}
	// f.ProduceApple().Create()

	// //在武汉开始生产小米手机
	// f = &FactoryWuhan{}
	// f.ProduceXiaomi().Create()
}
