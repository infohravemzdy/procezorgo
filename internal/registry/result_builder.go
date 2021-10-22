package registry

import (
	legalios "github.com/mzdyhrave/payrollgo-legalios/pkg/service"
	factories "github.com/mzdyhrave/payrollgo-procezor/internal/registry_factories"
	providers "github.com/mzdyhrave/payrollgo-procezor/internal/registry_providers"
	"github.com/mzdyhrave/payrollgo-procezor/internal/types"
	"sort"
)

type IResultBuilder interface {
	GetVersion() types.VersionCode
	GetPeriod() legalios.IPeriod
	InitWithPeriod(version types.VersionCode, period legalios.IPeriod, articleFactory factories.IArticleSpecFactory, conceptFactory factories.IConceptSpecFactory) bool
	GetResults(targets []types.ITermTarget, finDefs types.IArticleDefine) providers.IBuilderResultList
}

type resultBuilder struct {
	version types.VersionCode
	period  legalios.IPeriod

	articlesModel []providers.IArticleSpec
	conceptsModel []providers.IConceptSpec

	articlesOrder []types.ArticleCode
	articlesPaths pathCodeMap
}

func (f resultBuilder) GetPeriod() legalios.IPeriod {
	return f.period
}

func (f resultBuilder) GetVersion() types.VersionCode {
	return f.version
}

func (f resultBuilder) GetResults(targets []types.ITermTarget, finDefs types.IArticleDefine) providers.IBuilderResultList {
	calculTargets := f.buildCalculsList(targets, finDefs)

	calculResults := f.buildResultsList(calculTargets)

	return calculResults
}

func (f *resultBuilder) InitWithPeriod(version types.VersionCode, period legalios.IPeriod, articleFactory factories.IArticleSpecFactory, conceptFactory factories.IConceptSpecFactory) bool {
	f.version = version
	f.period  = period

	f.articlesModel = articleFactory.GetSpecList(period, version)
	f.conceptsModel = conceptFactory.GetSpecList(period, version)

	initResult := false

	f.articlesOrder, f.articlesPaths, initResult = InitGraphModel(f.articlesModel, f.conceptsModel)

	return initResult
}

func NewResultBuilder() IResultBuilder {
	factory := resultBuilder{ version: types.VersionZero(), period: legalios.PeriodZero(),
		articlesOrder: make([]types.ArticleCode, 0), articlesPaths: make(pathCodeMap) }

	return &factory
}

func mergeResults(results providers.IBuilderResultList, resultTarget providers.IBuilderResultList) providers.IBuilderResultList {
	return append(results, resultTarget...)
}

func (f resultBuilder) buildCalculsList(targets []types.ITermTarget, finDefs types.IArticleDefine) []ITermCalcul {
	finCalc := types.CloneArticleDefine(finDefs)

	targetsSpec := f.addFinDefToTargets(f.period, toTargetList(targets), finCalc)

	targetsStep := f.addExternToTargets(f.articlesOrder, f.period, targetsSpec)

	calculsList := f.addTargetToCalculs(targetsStep)

	return calculsList
}

func (f resultBuilder) buildResultsList(calculs []ITermCalcul) providers.IBuilderResultList {
	resultsList := make(providers.IBuilderResultList, 0, len(calculs))

	for _, x := range calculs {
		resultsList = mergeResults(resultsList, x.GetResults(f.period, resultsList))
	}
	return resultsList
}

func (f resultBuilder) addFinDefToTargets(period legalios.IPeriod, targets []types.ITermTarget, finDef types.IArticleDefine) []types.ITermTarget {
	return f.mergeItemPendings(period, targets, finDef)
}

func (f resultBuilder) addExternToTargets(topoOrders []types.ArticleCode, period legalios.IPeriod, targets []types.ITermTarget)  []types.ITermTarget {
	targetList := toTargetList(targets)

	for _, item := range targets {
		targetList = f.mergePendings(period, targetList, item)
	}

	sort.Slice(targetList, func(i, j int) bool {
		return targetCompare(topoOrders, targetList[i], targetList[j]) < 0
	})
	return targetList
}

func (f resultBuilder) addTargetToCalculs(targets []types.ITermTarget)  []ITermCalcul {
	var calculList = make([]ITermCalcul, 0, len(targets))

	for _, term := range targets {
		calculList = append(calculList, NewTermCalcul(term, getCalculFunc(f.conceptsModel, term.Concept())))
	}
	return calculList
}

func (f resultBuilder) mergePendings(period legalios.IPeriod, init []types.ITermTarget, target types.ITermTarget) []types.ITermTarget {
	targetsInit := toTargetList(init)

	pendingsPath, exists := firstOrDefaultInPathMap(f.articlesPaths, findArticleCode(target.Article()))

	if exists == false {
		return targetsInit
	}

	targetsList := toTargetList(targetsInit)
	for _, item := range pendingsPath {
		targetsList = f.mergeItemPendings(period, targetsList, item)
	}
	return targetsList
}

func (f resultBuilder) mergeItemPendings(period legalios.IPeriod, init []types.ITermTarget, articleDefs types.IArticleDefine) []types.ITermTarget {
	monthCode := types.GetMonthCode(period.GetCode())

	contract := types.NewContractCode()
	position := types.NewPositionCode()

	resultList := toTargetList(init)

	_, exists := firstOrDefaultTermTarget(init, findTermTargetCode(articleDefs.Code()))

	if exists {
		return resultList
	}

	variant := types.GetVariantCode(1)

	resultItem := types.NewTermTarget(monthCode, contract, position, variant, articleDefs.Code(), articleDefs.Role())

	resultList = append(resultList, resultItem)

	return resultList
}

func getCalculFunc(conceptsModel []providers.IConceptSpec, concept types.ConceptCode) providers.ResultFunc {
	conceptSpec, _ := firstOrDefaultConceptSpec(conceptsModel, findConceptSpecCode(concept.Value()))
	if conceptSpec == nil {
		return NotFoundCalculFunc
	}
	return conceptSpec.ResultDelegate()
}

func NotFoundCalculFunc(target types.ITermTarget, period legalios.IPeriod, results providers.IBuilderResultList) providers.IBuilderResultList {
	return providers.IBuilderResultList{types.NewFailureResult(NewNoResultFuncError(period, target))}
}

func targetCompare(topoOrders []types.ArticleCode, x types.ITermTarget, y types.ITermTarget) int {
	var xIndex = -1
	for p, v := range topoOrders {
		if v == x.Article() {
			xIndex = p
			break
		}
	}
	var yIndex = -1
	for p, v := range topoOrders {
		if v == y.Article() {
			yIndex = p
			break
		}
	}

	if xIndex == -1 && yIndex == -1 {
		return 0
	}

	if xIndex == -1 && yIndex != -1 {
		return -1
	}

	if xIndex != -1 && yIndex == -1 {
		return 1
	}
	return cmpInt(xIndex, yIndex)
}

