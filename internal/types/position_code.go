package types

const (
	positionZeroCode int16 = 0
)

type PositionCode struct {
	value int16
}

func (v PositionCode) Value() int16 {
	return v.value
}

func (v PositionCode) IsZero() bool {
	return v.value == positionZeroCode
}

func PositionCodeZero() PositionCode {
	return NewPositionCode()
}

func NewPositionCode() PositionCode {
	return PositionCode{ value: positionZeroCode}
}

func GetPositionCode(id int16) PositionCode {
	return PositionCode{ value: id }
}

