package registry

import (
	legalios "github.com/mzdyhrave/payrollgo-legalios/pkg/service"
	providers "github.com/mzdyhrave/payrollgo-procezor/internal/registry_providers"
	"github.com/mzdyhrave/payrollgo-procezor/internal/types"
)

type ITermCalcul interface {
	Target() types.ITermTarget
	ResultDelegate() providers.ResultFunc
	GetResults(period legalios.IPeriod, results providers.IBuilderResultList) providers.IBuilderResultList
}

type termCalcul struct {
	target         types.ITermTarget
	resultDelegate providers.ResultFunc
}

func (t termCalcul) Target() types.ITermTarget {
	return t.target
}

func (t termCalcul) ResultDelegate() providers.ResultFunc {
	return t.resultDelegate
}

func (t termCalcul) CallResultDelegate(target types.ITermTarget, period legalios.IPeriod, results providers.IBuilderResultList) providers.IBuilderResultList {
	if t.resultDelegate == nil {
		resultErrors := types.NewFailureResult(NewNoResultFuncError(period, t.target))
		return providers.IBuilderResultList{resultErrors}
	}
	return t.resultDelegate(target, period, results)
}

func (t termCalcul) GetResults(period legalios.IPeriod, results providers.IBuilderResultList) providers.IBuilderResultList {
	resultTarget := t.CallResultDelegate(t.Target(), period, results)
	return resultTarget
}

func NewTermCalcul(target types.ITermTarget, resultDelegate providers.ResultFunc) ITermCalcul {
	return &termCalcul{ target: target, resultDelegate: resultDelegate }
}
