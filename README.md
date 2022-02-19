# 分摊算法
```
按照价值比例均摊（即按占比分摊）,最后一个用减法

分摊顺序（也就是说抵扣顺序，抵扣优先级）：
    1.营销活动
    2.优惠券：抵扣适用的商品
    3.德分
    4.积分
    5.礼金卡（礼品卡）： 可抵扣适用的商品，抵扣适用商品的运费
    6.支付方式：使用第三方人民币支付

```

## 分摊类型
```
运费分摊：把运费分摊给需要运费的商品上
优惠券分摊：把优惠券金额分摊给适用的商品上
积分分摊：积分抵扣分摊
德分分摊：德分抵扣分摊
金币分摊：金币抵扣分摊
支付金额分摊：把实际支付的人民币分摊到购买的商品上
促销活动分摊： 
    满减（满件减、满元减）：把优惠的金额分摊到活动优惠的商品上
    满折（满件折、满元折）：把打折优惠的金额分摊到活动优惠的商品上
    满赠（满件赠、满元赠）：赠品不分摊，退主品时同时把赠品退回

礼金卡（礼品卡）分摊： 把消耗礼金卡的金额分摊到适用的商品和适用商品的运费上，礼金卡优先抵扣运费再抵扣商品

```

## Install下载依赖
```
go get -u github.com/jellycheng/gofentan
或者
GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go get -u github.com/jellycheng/gofentan

```

## 示例1 - map key为字符串
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
