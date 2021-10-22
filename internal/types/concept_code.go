package types

const (
	conceptZeroCode int32 = 0
)

type ConceptCode struct {
	value int32
}

func (v ConceptCode) Value() int32 {
	return v.value
}

func (v ConceptCode) IsZero() bool {
	return v.value == conceptZeroCode
}

func ConceptZero() ConceptCode {
	return NewConceptCode()
}

func NewConceptCode() ConceptCode {
	return ConceptCode{	value: conceptZeroCode}
}

func GetConceptCode(id int32) ConceptCode {
	return ConceptCode{	value: id }
}

func CloneConceptCode(code ConceptCode) ConceptCode {
	return ConceptCode{	value: code.value }
}

