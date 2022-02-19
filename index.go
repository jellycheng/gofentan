package gofentan

import (
	"errors"
	"github.com/shopspring/decimal"
	"sort"
	"sync"
)

type FenTan struct {
	// 参与分摊的数据
	dataStore *sync.Map
	// 共同分摊值，如：共分摊优惠金额、共分摊的运费金额
	commonVal int64
	// 实际完成共同分摊值
	alreadyCommonVal int64
	// 是否完成分摊
	isFinish bool
}

// SetCommonVal 设置共同分摊值
func (m *FenTan)SetCommonVal(n int64) *FenTan {
	m.commonVal = n
	return m
}

// GetCommonVal 获取共同分摊值
func (m *FenTan)GetCommonVal() int64 {
	return m.commonVal
}

func (m *FenTan)GetAlreadyCommonVal() int64 {
	return m.alreadyCommonVal
}

// AddData 添加分摊数据
func (m *FenTan)AddData(k interface{}, d FentanDto) *FenTan {
	m.dataStore.Store(k, d)
	return m
}

// IsFinish 获取是否完成分摊计算状态
func (m *FenTan)IsFinish() bool {
	return m.isFinish
}

// StartFenTanV1 开始计算分摊
func (m *FenTan)StartFenTanV1() *FenTan {
	if m.isFinish {
		return m
	}
	// 有效记录数
	validNum := 0
	// 总值
	var totalVal int64 = 0
	var tmpSliceData FentanDtoSortV1
	m.dataStore.Range(func(key, value interface{}) bool {
		ftDto := value.(FentanDto)
		if ftDto.Num>0 && ftDto.Price>0 {
			validNum++
			totalVal += ftDto.Price * ftDto.Num
			ftDto.masterDataNum = key
			tmpSliceData = append(tmpSliceData, ftDto)
		}
		return true
	})
	if totalVal == 0 || m.commonVal == 0 {
		m.isFinish = true
		return m
	}
	// 比例： 共同分摊值/总值
	commonValDecimal := decimal.NewFromFloat(float64(m.commonVal))
	totalValDecimal := decimal.NewFromFloat(float64(totalVal))
	rateDecimal := commonValDecimal.DivRound(totalValDecimal, 2)
	//rate,_ := rateDecimal.Float64()
	//fmt.Println("rate=", rate, ";validNum=",validNum)
	// 已经分摊共同分摊值
	alreadyCommonValDecimal := decimal.NewFromFloat(0)
	// 排序
	sort.Sort(tmpSliceData)
	// 已分摊个数
	alreadyNum := 0
	// 开始分摊
	for _, v := range tmpSliceData {
		alreadyNum++
		if alreadyNum == validNum { // 最后一个分摊，用减法
			tmpFentanDecimal := commonValDecimal.Sub(alreadyCommonValDecimal)
			alreadyCommonValDecimal = alreadyCommonValDecimal.Add(tmpFentanDecimal)
			tmpFentan,_ := tmpFentanDecimal.Float64()
			v.fentanVal = int64(tmpFentan)
			m.dataStore.Store(v.masterDataNum, v)
		} else {
			selfTotal := v.Price * v.Num
			tmpFentanDecimal := rateDecimal.Mul(decimal.NewFromFloat(float64(selfTotal))).RoundCeil(0)
			tmpFentanDecimal = decimal.Min(tmpFentanDecimal, commonValDecimal.Sub(alreadyCommonValDecimal))
			tmpFentan,_ := tmpFentanDecimal.Float64()
			v.fentanVal = int64(tmpFentan)
			m.dataStore.Store(v.masterDataNum, v)
			alreadyCommonValDecimal = alreadyCommonValDecimal.Add(tmpFentanDecimal)
		}

	}

	alreadyCommonVal,_ := alreadyCommonValDecimal.Float64()
	m.alreadyCommonVal = int64(alreadyCommonVal)

	m.isFinish = true
	return m
}

// GetData 获取分摊结果
func (m *FenTan)GetData(k interface{}) (FentanDto,error) {
	if v, ok := m.dataStore.Load(k);ok {
		return v.(FentanDto), nil
	}
	return FentanDto{}, errors.New("分摊数据不存在")
}

// GetAllData 获取所有数据,仅克隆数据
func (m *FenTan)GetAllData() *sync.Map {
	var ret = &sync.Map{}
	m.dataStore.Range(func(key, value interface{}) bool {
		ret.Store(key, value)
		return true
	})
	return ret
}

func NewFenTan() *FenTan {
	ret := new(FenTan)
	ret.dataStore = &sync.Map{}
	return ret
}
