package types

import (
	"fmt"
)

type NameDescriptionFce func (int32) string

type ITermResult interface {
	ITermSymbol
	Spec() IArticleSpec
	Concept() ConceptCode
	ConceptDescr() string
}

type TermResult struct {
	TermSymbol
	spec    IArticleSpec
	concept ConceptCode
	target   ITermTarget
	articleFunc NameDescriptionFce
	conceptFunc NameDescriptionFce
}

func (t TermResult) Spec() IArticleSpec {
	return t.spec
}

func (t TermResult) Concept() ConceptCode {
	return t.concept
}

func (t TermResult) ArticleDescr() string {
	return fmt.Sprintf("ArticleCode for: %v", t.article.Value())
}

func (t TermResult) ConceptDescr() string {
	return fmt.Sprintf("ConceptCode for: %v", t.concept.Value())
}

func NewTermResult(target ITermTarget, spec IArticleSpec) TermResult {
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
		concept: role, spec: spec}
}

type ITermResultList = []ITermResult

