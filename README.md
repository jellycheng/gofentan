# 分摊算法及抵扣顺序
```
按照价值比例均摊（即按占比分摊）,且最后一个（一项）用减法

分摊抵扣顺序（即分摊抵扣优先级）：
    1.营销活动：活动优惠，如满减、满折
    2.优惠券：抵扣适用的商品
    3.德分：一种虚拟货币，不可提现，用于抵扣商品金额
    4.积分：一种虚拟货币，不可提现，用于积分商城兑换商品
    5.金币：一种虚拟货币，可提现，不能抵扣商品金额
    6.礼金卡（礼品卡）：可抵扣适用的商品，但抵扣适用商品的运费或者不能抵扣运费取决于产品需求,一般礼金卡不能抵扣运费
    7.支付方式：使用第三方人民币支付（现金部分）

```

## 分摊类型
```
运费分摊：把运费分摊给需要运费的商品上，运费是否可以被虚拟货币抵扣取决于产品需求（如：运费券、德分、积分、金币、礼金卡等,一般虚拟货币不能抵扣运费）
优惠券分摊：把优惠券金额分摊给适用的商品上，如果是运费优惠券则仅抵扣运费
积分分摊：积分抵扣分摊,抵扣商品和运费
德分分摊：德分抵扣分摊,抵扣商品和运费
金币分摊：金币抵扣分摊,抵扣商品和运费，一般情况金币不能抵扣商品金额，仅用于提现；
        正向：根据购买商品进行拨付金币
        逆向：部分商品的数量进行退货，则按分摊比例追回相应金币
促销活动分摊： 
    满减（满件减、满元减）：把优惠的金额分摊到活动优惠的商品上
    满折（满件折、满元折）：把打折优惠的金额分摊到活动优惠的商品上
    满赠（满件赠、满元赠）：赠品不分摊，退主品时同时把赠品退回
组合品分摊（套装优惠）： 一个组合品可以组合多个sku商品，组合品不能嵌套组合品；
        组合品优惠金额 = 组合品销售金额 - 子品sku金额（即子品单价*数量）之和；
        然后把 组合品优惠金额 按比例分摊到各子品上；
礼金卡（礼品卡）分摊： 把消耗礼金卡的金额分摊到适用的商品，但不能抵扣适用商品的运费上
支付金额分摊：把实际要支付的人民币分摊到购买的商品和运费上；
    公式： 
        商品剩余第三方支付金额A（未抵扣金额） = 商品单价*数量 - 所有抵扣商品的金额（不含运费）
        商品分摊剩余实付运费B（未抵扣运费） = 商品分摊运费 - 优惠券抵扣运费 - 虚拟货币抵扣运费 - 其它抵扣运费
        商品实付人民币 = 商品剩余第三方支付金额A + 商品分摊剩余实付运费B = 商品单价*数量 - 所有抵扣商品的金额（不含运费） + 商品分摊剩余实付运费B，
        分摊值为负数时强制存0
        

```

## 举例

商品 | 是否参与活动 | 价格（元）| 购买数量 | 总价（元）
---|--- | --- | --- | ---
A | 是 | 24 | 3 | 72
B | 是 |20 | 2 | 40
C | 否 | 10 | 3 | 30


```
综上：假如A、B商品适用于满减活动（满100减20）；
根据分摊算法推理逻辑如下：
    优惠金额：Y=20元
    需参与分摊金额：商品A + 商品B = 24*3 + 20 * 2= 112元
    不参与分摊金额： 10*3 = 30元
    商品实付金额 = 需参与分摊金额+不参与分摊金额-优惠金额，（即112 + 30 - 20 = 122）
    商品A分摊：商品A金额/实付金额 * 优惠金额 = 24*3/112*20 = 12.86 （12.8571429结果小数点保留2位且第三位四舍五入）
    商品B分摊：20-12.86=7.14
    商品C不分摊
    结论： 商品A分摊满减活动12.86元、商品B分摊分摊满减活动7.14元、商品C不分摊
    
```

## Install下载依赖
```
go get -u github.com/jellycheng/gofentan
或者
GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go get -u github.com/jellycheng/gofentan

```

## 示例1 - map key为字符串
```
备注： 
    原则上 共同分摊值 要<= 参与分摊的数据值总和，但程序本身不控制（由业务系统控制），为了适用于参与分摊数据做膨胀分摊；
    金额和数量必须大于0；
    
package main

import (
	"fmt"
	"github.com/jellycheng/gofentan"
)

func main() {
	fenTanObj := gofentan.NewFenTan()
	// 共同分摊值
	fenTanObj.SetCommonVal(2)
	// 需要参与分摊的数据，如果Price或Num值<=0则分摊算法给该项分摊值为0
	data := map[string]gofentan.FentanDto{
		"sku_100":{Price: 1, Num: 1},
		"sku_200":{Price: 2, Num: 1},
		"sku_400":{Price: 1, Num: 3},
	}
	for k, v := range data {
		fenTanObj.AddData(k, v)
	}
	fenTanObj.StartFenTanV1() // 开始分摊
	// 获取某一个值的分摊结果
	if v,err := fenTanObj.GetData("sku_200");err == nil {
		fmt.Println("sku_200的分摊结果：", v.GetFentanVal())
	}
	// 获取所有分摊结果
	fenTanObj.GetAllData().Range(func(key, value interface{}) bool {
		fmt.Println("key=", key, fmt.Sprintf(";分摊结果:%+v", value))

		return true
	})
	fmt.Println("获取共同分摊值：", fenTanObj.GetCommonVal())
	fmt.Println("完成共同分摊值：", fenTanObj.GetAlreadyCommonVal())

}

```

## 示例2 - map key为int64
```
package main

import (
	"fmt"
	"github.com/jellycheng/gofentan"
)

func main() {
	fenTanObj := gofentan.NewFenTan()
	// 共同分摊值
	fenTanObj.SetCommonVal(2)
	// 需要参与分摊的数据，如果Price或Num值<=0则分摊算法给该项分摊值为0
	data := map[int64]gofentan.FentanDto{
		123:{Price: 1, Num: 1},
		200:{Price: 2, Num: 1},
		987654321:{Price: 1, Num: 3},
	}
	for k, v := range data {
		fenTanObj.AddData(k, v)
	}
	fenTanObj.StartFenTanV1() // 开始分摊
	// 获取某一个值的分摊结果
	if v,err := fenTanObj.GetData(int64(200));err == nil {
		fmt.Println("200的分摊结果：", v.GetFentanVal())
	}
	// 获取所有分摊结果
	fenTanObj.GetAllData().Range(func(key, value interface{}) bool {
		fmt.Println("key=", key, fmt.Sprintf(";分摊结果:%+v", value))

		return true
	})
	fmt.Println("获取共同分摊值：", fenTanObj.GetCommonVal())
	fmt.Println("完成共同分摊值：", fenTanObj.GetAlreadyCommonVal())

}

```

## 示例3 - map key为int
```
package main

import (
	"fmt"
	"github.com/jellycheng/gofentan"
)

func main() {
	fenTanObj := gofentan.NewFenTan()
	// 共同分摊值
	fenTanObj.SetCommonVal(2)
	// 需要参与分摊的数据，如果Price或Num值<=0则分摊算法给该项分摊值为0
	data := map[int]gofentan.FentanDto{
		123:{Price: 1, Num: 1},
		200:{Price: 2, Num: 1},
		987654321:{Price: 1, Num: 3},
	}
	for k, v := range data {
		fenTanObj.AddData(k, v)
	}
	fenTanObj.StartFenTanV1() // 开始分摊
	// 获取某一个值的分摊结果
	if v,err := fenTanObj.GetData(200);err == nil {
		fmt.Println("200的分摊结果：", v.GetFentanVal())
	}
	// 获取所有分摊结果
	fenTanObj.GetAllData().Range(func(key, value interface{}) bool {
		fmt.Println("key=", key, fmt.Sprintf(";分摊结果:%+v", value))

		return true
	})
	fmt.Println("获取共同分摊值：", fenTanObj.GetCommonVal())
	fmt.Println("完成共同分摊值：", fenTanObj.GetAlreadyCommonVal())

}

```
