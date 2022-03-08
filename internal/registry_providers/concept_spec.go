package registry_providers

import (
	legalios "github.com/mzdyhrave/legaliosgo"
	"github.com/mzdyhrave/procezorgo/internal/types"
)

type ResultFunc func (target types.ITermTarget, spec types.IArticleSpec, period legalios.IPeriod, ruleset legalios.IBundleProps, results IBuilderResultList) IBuilderResultList

type IConceptSpec interface {
	Code() types.ConceptCode
	Path() []types.ArticleCode
	ResultDelegate() ResultFunc
	DefaultTargetList(article types.ArticleCode, period legalios.IPeriod, ruleset legalios.IBundleProps, month types.MonthCode,
		contractTerms types.IContractTermList, positionTerms types.IPositionTermList,
		targets types.ITermTargetList, vars types.VariantCode) types.ITermTargetList
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

func (s ConceptSpec) DefaultTargetList(article types.ArticleCode, period legalios.IPeriod, ruleset legalios.IBundleProps, month types.MonthCode,
	contractTerms types.IContractTermList, positionTerms types.IPositionTermList,
	targets types.ITermTargetList, vars types.VariantCode) types.ITermTargetList {
	con := types.ContractCodeZero()
	pos := types.PositionCodeZero()

	if len(targets)!=0 {
		return types.ITermTargetList{}
	}
	return types.ITermTargetList{ types.NewTermTarget(month, con, pos, vars, article, s.Code()) }
}
