package types

const (
	variantZeroCode int16 = 0
)

type VariantCode struct {
	value int16
}

func (v VariantCode) Value() int16 {
	return v.value
}

func (v VariantCode) IsZero() bool {
	return v.value == variantZeroCode
}

func VariantZero() VariantCode {
	return NewVariantCode()
}

func NewVariantCode() VariantCode {
	return VariantCode{ value: variantZeroCode}
}

func GetVariantCode(id int16) VariantCode {
	return VariantCode{ value: id }
}


