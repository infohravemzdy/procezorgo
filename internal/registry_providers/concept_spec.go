package registry_providers

import (
	legalios "github.com/mzdyhrave/legaliosgo"
	"github.com/mzdyhrave/procezorgo/internal/types"
)

type ResultFunc func (target types.ITermTarget, period legalios.IPeriod, ruleset legalios.IBundleProps, results IBuilderResultList) IBuilderResultList

type IConceptSpec interface {
	Code() types.ConceptCode
	Path() []types.ArticleCode
	ResultDelegate() ResultFunc
}

type ConceptSpec struct {
	code types.ConceptCode
	path           []types.ArticleCode
	resultDelegate ResultFunc
}

func (s ConceptSpec) Code() types.ConceptCode {
	return s.code
}

func (s ConceptSpec) Path() []types.ArticleCode {
	return s.path
}

func (s ConceptSpec) ResultDelegate() ResultFunc {
	return s.resultDelegate
}

