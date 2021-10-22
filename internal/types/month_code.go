package types

const (
	monthZeroCode int32 = 0
)

type MonthCode struct {
	value int32
}

func (v MonthCode) Value() int32 {
	return v.value
}

func (v MonthCode) IsZero() bool {
	return v.value == monthZeroCode
}

func MonthZero() MonthCode {
	return NewMonthCode()
}

func NewMonthCode() MonthCode {
	return MonthCode{ value: monthZeroCode}
}

func GetMonthCode(id int32) MonthCode {
	return MonthCode{ value: id }
}

