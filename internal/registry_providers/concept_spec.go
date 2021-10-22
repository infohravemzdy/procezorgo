package registry_providers

import (
	legalios "github.com/mzdyhrave/payrollgo-legalios/pkg/service"
	"github.com/mzdyhrave/payrollgo-procezor/internal/types"
)

type ResultFunc func (target types.ITermTarget, period legalios.IPeriod, results IBuilderResultList) IBuilderResultList

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

