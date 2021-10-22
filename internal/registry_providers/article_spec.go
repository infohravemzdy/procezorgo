package registry_providers

import "github.com/mzdyhrave/procezorgo/internal/types"

type IArticleSpec interface {
	Code() types.ArticleCode
	Role() types.ConceptCode
	Defs() types.IArticleDefine
	Sums() []types.ArticleCode
}

type ArticleSpec struct {
	sums []types.ArticleCode
	role types.ConceptCode
	code types.ArticleCode
}

func (s ArticleSpec) Defs() types.IArticleDefine {
	return types.GetArticleDefine(s.code.Value(), s.role.Value())
}

func (s ArticleSpec) Sums() []types.ArticleCode {
	return s.sums
}

func (s ArticleSpec) Code() types.ArticleCode {
	return s.code
}

func (s ArticleSpec) Role() types.ConceptCode {
	return s.role
}
