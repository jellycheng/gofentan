package gofentan

import (
	"fmt"
	"sort"
	"testing"
)

// go test -run="TestNewFenTan"
func TestNewFenTan(t *testing.T) {
	fenTanObj := NewFenTan()
	// 共同分摊值
	fenTanObj.SetCommonVal(123)
	// 需要参与分摊的数据，如果Price或Num值<=0则分摊算法给该项分摊值为0
	data := map[string]FentanDto{
			"sku_100":{Price: 1, Num: 1},
			"sku_200":{Price: 1, Num: 1},
			"sku_300":{Price: 3, Num: 0},
			"sku_400":{Price: 1, Num: 2},
	}
	data["sku_500"] = FentanDto{Price: 2, Num: 2}
	for k, v := range data {
		fenTanObj.AddData(k, v)
	}
	fenTanObj.StartFenTanV1() // 开始分摊
	// 获取某一个值的分摊结果
	if v,err := fenTanObj.GetData("sku_500");err == nil {
		fmt.Println("sku_500的分摊结果：", v.GetFentanVal())
	}
	// 获取所有分摊结果
	fenTanObj.GetAllData().Range(func(key, value interface{}) bool {
		fmt.Println("key=", key, fmt.Sprintf(";分摊结果:%+v", value))

		return true
	})
	fmt.Println("设置共同分摊值：", fenTanObj.GetCommonVal())
	fmt.Println("完成共同分摊值：", fenTanObj.GetAlreadyCommonVal())
}

// go test -run="TestFentanDtoSortV1"
func TestFentanDtoSortV1(t *testing.T) {
	d := FentanDtoSortV1 {
		{Price: 1, Num: 1},
		{Price: 1, Num: 1},
		{Price: 3, Num: 0},
		{Price: 1, Num: 500},
		{Price: 1, Num: 500,Weight: 9},
		{Price: 1, Num: 1,Weight: -1},
		{Price: 1, Num: 2},
	}
	d = append(d, FentanDto{Price: 100, Num: 0},)
	sort.Sort(d)
	for _, v := range d {
		fmt.Printf("%+v\n", v)
	}

}

