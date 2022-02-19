package gofentan

// FentanDto 参与分摊的元数据
type FentanDto struct {
	Price int64  // 单价，单位是最小粒度，如：人民币分/积分个
	Num int64    // 数量，购买数量
	Weight int64 // 分摊权重，权重越高越优先分摊（值越大权重越高）
	fentanVal int64 // 分摊到的值
	masterDataNum interface{}
}

// GetFentanVal 获取分摊到的值
func (m FentanDto)GetFentanVal() int64 {
	return m.fentanVal
}

type FentanDtoSortV1 []FentanDto
func (m FentanDtoSortV1) Len() int {
	return len(m)
}

// Less 在同等权重下，最小优先
func (m FentanDtoSortV1) Less(i, j int) bool {
	if m[i].Weight != m[j].Weight {
		return m[i].Weight > m[j].Weight
	}
	iVal := m[i].Price * m[i].Num
	jVal := m[j].Price * m[j].Num
	return iVal < jVal
}

func (m FentanDtoSortV1) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

