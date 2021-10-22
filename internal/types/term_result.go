package types

type NameDescriptionFce func (int32) string

type ITermResult interface {
	ITermSymbol
	Concept() ConceptCode
	ResultDescr() string
	ResultBasis() int32
	ResultValue() int32
}

type TermResult struct {
	TermSymbol
	concept  ConceptCode
	target   ITermTarget
	resultDescr string
	resultBasis int32
	resultValue int32
	articleFunc NameDescriptionFce
	conceptFunc NameDescriptionFce
}

func (t TermResult) Concept() ConceptCode {
	return t.concept
}

func (t TermResult) Defs() IArticleDefine {
	return GetArticleDefine(t.Article().Value(), t.Concept().Value())
}

func (t TermResult) ResultDescr() string {
	return t.resultDescr
}

func (t TermResult) ResultBasis() int32 {
	return t.resultBasis
}

func (t TermResult) ResultValue() int32 {
	return t.resultValue
}

func NewTermResult(target ITermTarget, value int32, basis int32, descr string) TermResult {
	var mont = NewMonthCode()
	var cont = NewContractCode()
	var post = NewPositionCode()
	var code = NewArticleCode()
	var vars = NewVariantCode()
	var role = NewConceptCode()

	if target != nil {
		mont = target.Month()
		cont = target.Contract()
		post = target.Position()
		code = target.Article()
		vars = target.Variant()
		role = target.Concept()
	}

	return TermResult{ TermSymbol: NewTermSymbol(mont, cont, post, vars, code),
		concept: role, resultValue: value, resultBasis: basis, resultDescr: descr}
}

type ITermResultList []ITermResult

