package types

import "fmt"

type ITermSymbol interface {
	Contract() ContractCode
	Position() PositionCode
	Month() MonthCode
	Article() ArticleCode
	Variant() VariantCode
	ArticleDescr() string
}

type ITermTarget interface {
	ITermSymbol
	Concept() ConceptCode
	ConceptDescr() string
}

type ITermTargetList = []ITermTarget

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
}

func (t TermTarget) Concept() ConceptCode {
	return t.concept
}

func (t TermTarget) ArticleDescr() string {
	return fmt.Sprintf("ArticleCode for: %v", t.article.Value())
}

func (t TermTarget) ConceptDescr() string {
	return fmt.Sprintf("ConceptCode for: %v", t.concept.Value())
}

func NewTermTarget(mont MonthCode, cont ContractCode, post PositionCode, vars VariantCode, code ArticleCode, role ConceptCode) TermTarget {
	return TermTarget{ TermSymbol: NewTermSymbol(mont, cont, post, vars, code ), concept: role }
}

