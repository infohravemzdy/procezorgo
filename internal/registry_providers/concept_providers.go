package registry_providers

import (
	legalios "github.com/mzdyhrave/legaliosgo"
	"github.com/mzdyhrave/procezorgo/internal/types"
)

type IBuilderResult interface {
	IsFailure() bool
	IsSuccess() bool
	Error() error
	ResultError() types.ITermResultError
	Value() types.ITermResult
}

type IBuilderResultList []IBuilderResult

type IConceptCodeProvider interface {
	Code() types.ConceptCode
}

type IConceptSpecProvider interface {
	IConceptCodeProvider
	GetSpec(period legalios.IPeriod, version types.VersionCode) IConceptSpec
}

type ConceptSpecProvider struct {
	code types.ConceptCode
}

func (p ConceptSpecProvider) Code() types.ConceptCode {
	return p.code
}

func EvalEmptyResultList(target types.ITermTarget, period legalios.IPeriod, ruleset legalios.IBundleProps, results IBuilderResultList) IBuilderResultList {
	return IBuilderResultList{}
}

func NewConceptProvider(code int32) ConceptSpecProvider {
	return ConceptSpecProvider{ code: types.GetConceptCode(code) }
}

func NewConceptSpec(code int32) ConceptSpec {
	return ConceptSpec{ code: types.GetConceptCode(code), path: []types.ArticleCode{}, resultDelegate: EvalEmptyResultList }
}

func NewConceptPathSpec(code int32, path []types.ArticleCode) ConceptSpec {
	return ConceptSpec{ code: types.GetConceptCode(code), path: path, resultDelegate: EvalEmptyResultList }
}

func NewConceptFuncSpec(code int32, resultFunc ResultFunc) ConceptSpec {
	return ConceptSpec{ code: types.GetConceptCode(code), path: []types.ArticleCode{}, resultDelegate: resultFunc }
}

func NewConceptPathFuncSpec(code int32, path []types.ArticleCode, resultFunc ResultFunc) ConceptSpec {
	return ConceptSpec{ code: types.GetConceptCode(code), path: path, resultDelegate: resultFunc }
}

func NewConceptCodeProvider(code int32) IConceptCodeProvider {
	return ConceptSpecProvider{ code: types.GetConceptCode(code) }
}

