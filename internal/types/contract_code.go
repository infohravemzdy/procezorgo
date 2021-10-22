package types

const (
	contractZeroCode int16 = 0
)

type ContractCode struct {
	value int16
}

func (v ContractCode) Value() int16 {
	return v.value
}

func (v ContractCode) IsZero() bool {
	return v.value == contractZeroCode
}

func ContractZero() ContractCode {
	return NewContractCode()
}

func NewContractCode() ContractCode {
	return ContractCode{ value: contractZeroCode}
}

func GetContractCode(id int16) ContractCode {
	return ContractCode{ value: id }
}

