package registry

import (
	legalios "github.com/mzdyhrave/legaliosgo"
	providers "github.com/mzdyhrave/procezorgo/internal/registry_providers"
	"github.com/mzdyhrave/procezorgo/internal/types"
)

type ITermCalcul interface {
	Target() types.ITermTarget
	Spec() types.IArticleSpec
	ResultDelegate() providers.ResultFunc
	GetResults(period legalios.IPeriod, ruleset legalios.IBundleProps, results providers.IBuilderResultList) providers.IBuilderResultList
}

type termCalcul struct {
	target         types.ITermTarget
	spec           types.IArticleSpec
	resultDelegate providers.ResultFunc
}

func (t termCalcul) Target() types.ITermTarget {
	return t.target
}

func (t termCalcul) Spec() types.IArticleSpec {
	return t.spec
}

func (t termCalcul) ResultDelegate() providers.ResultFunc {
	return t.resultDelegate
}

func (t termCalcul) GetResults(period legalios.IPeriod, ruleset legalios.IBundleProps, results providers.IBuilderResultList) providers.IBuilderResultList {
	if t.resultDelegate == nil {
		resultErrors := types.NewFailureResult(NewNoResultFuncError(period, t.target))
		return providers.IBuilderResultList{resultErrors}
	}
	return t.resultDelegate(t.Target(), t.Spec(), period, ruleset, results)
}

func NewTermCalcul(target types.ITermTarget, spec types.IArticleSpec, resultDelegate providers.ResultFunc) ITermCalcul {
	return &termCalcul{ target: target, spec: spec, resultDelegate: resultDelegate }
}
