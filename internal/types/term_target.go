package types

type ITermSymbol interface {
	Contract() ContractCode
	Position() PositionCode
	Month() MonthCode
	Article() ArticleCode
	Variant() VariantCode
}

type ITermTarget interface {
	ITermSymbol
	Concept() ConceptCode
	TarBasis() int32
	Defs() IArticleDefine
}

type ITermTargetList []ITermTarget

type TermSymbol struct {
	contract ContractCode
	position PositionCode
	month    MonthCode
	article  ArticleCode
	variant  VariantCode
}

func (t TermSymbol) Contract() ContractCode {
	return t.contract
}

func (t TermSymbol) Position() PositionCode {
	return t.position
}

func (t TermSymbol) Month() MonthCode {
	return t.month
}

func (t TermSymbol) Article() ArticleCode {
	return t.article
}

func (t TermSymbol) Variant() VariantCode {
	return t.variant
}

func NewTermSymbol(mont MonthCode, cont ContractCode, post PositionCode, vars VariantCode, code ArticleCode) TermSymbol {
	return TermSymbol{ month: mont, contract: cont, position: post, variant: vars, article: code }
}

type TermTarget struct {
	TermSymbol
	concept  ConceptCode
	tarBasis int32
}

func (t TermTarget) Concept() ConceptCode {
	return t.concept
}

func (t TermTarget) Defs() IArticleDefine {
	return GetArticleDefine(t.Article().Value(), t.Concept().Value())
}

func (t TermTarget) TarBasis()  int32 {
	return t.tarBasis
}

func NewTermTarget(mont MonthCode, cont ContractCode, post PositionCode, vars VariantCode, code ArticleCode, role ConceptCode) TermTarget {
	return TermTarget{ TermSymbol: NewTermSymbol(mont, cont, post, vars, code ), concept: role, tarBasis: int32(0) }
}

