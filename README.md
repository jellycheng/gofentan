# 分摊算法
```
按照价值比例均摊（即按占比分摊）,最后一个用减法

分摊顺序（也就是说抵扣顺序，抵扣优先级）：
    1.营销活动：活动优惠，如满减、满折
    2.优惠券：抵扣适用的商品
    3.德分：一种虚拟货币
    4.积分：一种虚拟货币，不可提现
    5.金币：一种虚拟货币，可提现
    6.礼金卡（礼品卡）：可抵扣适用的商品，但抵扣适用商品的运费或者不能抵扣运费取决于产品需求,一般礼金卡不能抵扣运费
    7.支付方式：使用第三方人民币支付

```

## 分摊类型
```
运费分摊：把运费分摊给需要运费的商品上，运费是否可以被虚拟货币抵扣取决于产品需求（如：运费券、德分、积分、金币、礼金卡等,一般虚拟货币不能抵扣运费）
优惠券分摊：把优惠券金额分摊给适用的商品上，如果运费优惠券则抵扣运费
积分分摊：积分抵扣分摊,抵扣商品和运费
德分分摊：德分抵扣分摊,抵扣商品和运费
金币分摊：金币抵扣分摊,抵扣商品和运费
促销活动分摊： 
    满减（满件减、满元减）：把优惠的金额分摊到活动优惠的商品上
    满折（满件折、满元折）：把打折优惠的金额分摊到活动优惠的商品上
    满赠（满件赠、满元赠）：赠品不分摊，退主品时同时把赠品退回

礼金卡（礼品卡）分摊： 把消耗礼金卡的金额分摊到适用的商品，但不能抵扣适用商品的运费上
支付金额分摊：把实际支付的人民币分摊到购买的商品和运费上；
    公式： 商品分摊人民币 = 商品剩余第三方支付金额 + 商品分摊剩余实付运费 = 商品单价*数量 - 所有抵扣商品的金额（不含运费） + 商品分摊剩余实付运费，分摊值为负数时强制存0
        商品剩余第三方支付金额 = 商品单价*数量 - 所有抵扣商品的金额（不含运费）
        商品分摊剩余实付运费 = 商品分摊运费 - 优惠券抵扣运费 - 虚拟货币抵扣运费 - 其它抵扣运费

```

## Install下载依赖
```
go get -u github.com/jellycheng/gofentan
或者
GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go get -u github.com/jellycheng/gofentan

```

## 示例1 - map key为字符串
```
备注： 原则上 共同分摊值 要<= 参与分摊的数据值总和，但程序本身不控制（由业务系统控制），为了适用于参与分摊数据做膨胀分摊
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
