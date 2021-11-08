package registry_providers

import (
	legalios "github.com/mzdyhrave/legaliosgo"
	"github.com/mzdyhrave/procezorgo/internal/types"
)

type IArticleCodeProvider interface {
	Code() types.ArticleCode
}

type IArticleSpecProvider interface {
	IArticleCodeProvider
	GetSpec(period legalios.IPeriod, version types.VersionCode) IArticleSpec
}

type ArticleSpecProvider struct {
	code types.ArticleCode
}

func (p ArticleSpecProvider) Code() types.ArticleCode {
	return p.code
}

func NewArticleProvider(code int32) ArticleSpecProvider {
	return ArticleSpecProvider{ code: types.GetArticleCode(code) }
}

func NewArticleSpec(code int32, role int32) ArticleSpec {
	return ArticleSpec{ code: types.GetArticleCode(code), role: types.GetConceptCode(role), sums: types.ArticleCodeList{} }
}

func NewArticleSumSpec(code int32, role int32, sums []types.ArticleCode) ArticleSpec {
	return ArticleSpec{ code: types.GetArticleCode(code), role: types.GetConceptCode(role), sums: sums }
}

func NewArticleSumIntSpec(code int32, role int32, sums []int32) ArticleSpec {
	return ArticleSpec{ code: types.GetArticleCode(code), role: types.GetConceptCode(role), sums: ConstToSumsArray(sums) }
}

func NewArticleCodeProvider(code int32) ArticleSpecProvider {
	return ArticleSpecProvider{ code: types.GetArticleCode(code) }
}

func ConstToSumsArray(_sums []int32) []types.ArticleCode {
	sumsArray := make([]types.ArticleCode, 0)
	for _, r := range _sums {
		sumsArray = append(sumsArray, types.GetArticleCode(r))
	}
	return sumsArray
}

