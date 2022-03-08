package registry

import (
	legalios "github.com/mzdyhrave/legaliosgo"
	factories "github.com/mzdyhrave/procezorgo/internal/registry_factories"
	providers "github.com/mzdyhrave/procezorgo/internal/registry_providers"
	"github.com/mzdyhrave/procezorgo/internal/types"
	"sort"
)

type IResultBuilder interface {
	GetVersion() types.VersionCode
	GetPeriod() legalios.IPeriod

	InitWithPeriod(version types.VersionCode, period legalios.IPeriod, articleFactory factories.IArticleSpecFactory, conceptFactory factories.IConceptSpecFactory) bool
	GetResults(ruleset legalios.IBundleProps,
		contractTerms types.IContractTermList, positionTerms types.IPositionTermList,
		targets types.ITermTargetList, calcArts types.ArticleCodeList) providers.IBuilderResultList
	GetOrder() OrderCodeList
	GetPaths() PathsTermsMap
}

type resultBuilder struct {
	version types.VersionCode
	period  legalios.IPeriod

	articlesModel []types.IArticleSpec
	conceptsModel []providers.IConceptSpec

	articlesOrder OrderCodeList
	articlesPaths PathsTermsMap
}

func (f resultBuilder) GetPeriod() legalios.IPeriod {
	return f.period
}

func (f resultBuilder) GetVersion() types.VersionCode {
	return f.version
}

func (f resultBuilder) GetOrder() OrderCodeList {
	return f.articlesOrder
}
func (f resultBuilder) GetPaths() PathsTermsMap {
	return f.articlesPaths
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

func (f resultBuilder) GetResults(ruleset legalios.IBundleProps,
	contractTerms types.IContractTermList, positionTerms types.IPositionTermList,
	targets types.ITermTargetList, calcArts types.ArticleCodeList) providers.IBuilderResultList {

	calculTargets := f.buildCalculsList(f.period, ruleset,
		contractTerms, positionTerms, targets, calcArts)

	calculResults := f.buildResultsList(f.period, ruleset,
		calculTargets)

	return calculResults
}

func NewResultBuilder() IResultBuilder {
	factory := resultBuilder{ version: types.VersionCodeZero(), period: legalios.PeriodZero(),
		articlesOrder: make([]types.ArticleTerm, 0), articlesPaths: make(PathsTermsMap) }

	return &factory
}

func mergeResults(results providers.IBuilderResultList, resultTarget providers.IBuilderResultList) providers.IBuilderResultList {
	return append(results, resultTarget...)
}

func (f resultBuilder) buildCalculsList(period legalios.IPeriod, ruleset legalios.IBundleProps,
	contractTerms types.IContractTermList, positionTerms types.IPositionTermList,
	targets types.ITermTargetList, calcArts types.ArticleCodeList) []ITermCalcul {
	specDefines := make(types.IArticleSpecList, 0)
	for _, s := range calcArts {
		articleSpec, _ := firstOrDefaultArticleSpec(f.articlesModel, findArticleSpecCode(s.Value()))
		if articleSpec != nil {
			specDefines = append(specDefines, articleSpec)
		}
	}
	calcDefines := make(types.IArticleDefineList, 0)
	for _, s := range specDefines {
		calcDefines = append(calcDefines, types.CloneArticleDefine(s.Defs()))
	}

	targetsSpec := f.addFinDefToTargets(period, ruleset,
		contractTerms, positionTerms, toTargetList(targets), calcDefines)

	targetsStep := f.addExternToTargets(f.articlesOrder, period, ruleset,
		contractTerms, positionTerms, targetsSpec)

	calculsList := f.addTargetToCalculs(targetsStep)

	return calculsList
}

func (f resultBuilder) buildResultsList(period legalios.IPeriod, ruleset legalios.IBundleProps,
	calculs []ITermCalcul) providers.IBuilderResultList {
	resultsList := make(providers.IBuilderResultList, 0, len(calculs))

	for _, x := range calculs {
		resultsList = mergeResults(resultsList, x.GetResults(period, ruleset, resultsList))
	}
	return resultsList
}

func (f resultBuilder) addFinDefToTargets(period legalios.IPeriod, ruleset legalios.IBundleProps,
	contractTerms types.IContractTermList, positionTerms types.IPositionTermList,
	targets types.ITermTargetList, calcDefs types.IArticleDefineList) types.ITermTargetList {
	return f.mergeListPendings(period, ruleset,
		contractTerms, positionTerms, targets, calcDefs)
}

func (f resultBuilder) addExternToTargets(topoOrders OrderCodeList,
	period legalios.IPeriod, ruleset legalios.IBundleProps,
	contractTerms types.IContractTermList, positionTerms types.IPositionTermList,
	targets types.ITermTargetList) types.ITermTargetList {
	targetList := toTargetList(targets)

	for _, item := range targets {
		targetList = f.mergePendings(period, ruleset,
			contractTerms, positionTerms, targetList, item)
	}

	sort.Slice(targetList, func(i, j int) bool {
		return targetCompare(topoOrders, targetList[i], targetList[j]) < 0
	})
	return targetList
}

func (f resultBuilder) addDefinesToTargets(period legalios.IPeriod, ruleset legalios.IBundleProps,
	contractTerms types.IContractTermList, positionTerms types.IPositionTermList,
	targets types.ITermTargetList, defines types.IArticleDefineList) types.ITermTargetList {
	targetList := make(types.ITermTargetList, 0)
	for _, x := range defines {
		targetCodes := make(types.ITermTargetList, 0)
		for _, t := range targets {
			if  t.Article() == x.Code() {
				targetCodes = append(targetCodes, t)
			}
		}
		targetList = append(targetList,
			getTargetList(period, ruleset, f.conceptsModel,
				contractTerms, positionTerms, targetCodes, x.Code(), x.Role())...)

	}
	return targetList
}

func (f resultBuilder) addTargetToCalculs(targets types.ITermTargetList) []ITermCalcul {
	var calculList = make([]ITermCalcul, 0, len(targets))

	for _, term := range targets {
		articleSpec, _ := firstOrDefaultArticleSpec(f.articlesModel, findArticleSpecCode(term.Article().Value()))
		if articleSpec != nil {
			calculList = append(calculList, NewTermCalcul(term, articleSpec, getCalculFunc(f.conceptsModel, term.Concept())))
		}
	}
	return calculList
}

func (f resultBuilder) mergePendings(period legalios.IPeriod, ruleset legalios.IBundleProps,
	contractTerms types.IContractTermList, positionTerms types.IPositionTermList,
	targets types.ITermTargetList, target types.ITermTarget) types.ITermTargetList {
	targetsInit := toTargetList(targets)

	pendingsPath, exists := firstOrDefaultInPathMap(f.articlesPaths, findArticleTern(target.Article()))

	if exists == false {
		return targetsInit
	}

	targetsList := toTargetList(targetsInit)
	for _, item := range pendingsPath {
		targetsList = f.mergeItemPendings(period, ruleset, contractTerms, positionTerms, targetsList, item)
	}
	return targetsList
}

func (f resultBuilder) mergeItemPendings(period legalios.IPeriod, ruleset legalios.IBundleProps,
	contractTerms types.IContractTermList, positionTerms types.IPositionTermList,
	targets types.ITermTargetList, articleDefs types.IArticleDefine) types.ITermTargetList {
	resultList := toTargetList(targets)

	initTargets := make(types.ITermTargetList, 0)
	for _, t := range targets {
		if t.Article() == articleDefs.Code() {
			initTargets = append(initTargets, t)
		}
	}

	targetList := getTargetList(period, ruleset, f.conceptsModel,
		contractTerms, positionTerms, initTargets, articleDefs.Code(), articleDefs.Role())

	resultList = append(resultList, targetList...)

	return resultList
}

func (f resultBuilder) mergeListPendings(period legalios.IPeriod, ruleset legalios.IBundleProps,
	contractTerms types.IContractTermList, positionTerms types.IPositionTermList,
	targets types.ITermTargetList, calcDefs types.IArticleDefineList) types.ITermTargetList {
	resultList := toTargetList(targets)

	defineList := make(types.IArticleDefineList, 0)
	for _, x := range calcDefs {
		_, exists := firstOrDefaultTermTarget(targets, findTermTargetCode(x.Code()))
		if exists == false {
			defineList = append(defineList, x)
		}
	}

	targetList := f.addDefinesToTargets(period, ruleset, contractTerms, positionTerms, targets, defineList)

	resultList = append(resultList, targetList...)

	return resultList
}

func getTargetList(period legalios.IPeriod, ruleset legalios.IBundleProps, conceptsModel []providers.IConceptSpec,
	contractTerms types.IContractTermList, positionTerms types.IPositionTermList,
	targets types.ITermTargetList, article types.ArticleCode, concept types.ConceptCode) types.ITermTargetList {
	monthCode := types.GetMonthCode(period.GetCode())
	variant := types.GetVariantCode(1)

	conceptSpec, _ := firstOrDefaultConceptSpec(conceptsModel, findConceptSpecCode(concept.Value()))
	if conceptSpec == nil {
		contract := types.NewContractCode()
		position := types.NewPositionCode()
		return types.ITermTargetList{types.NewTermTarget(monthCode, contract, position, variant, article, concept)}
	}
	return conceptSpec.DefaultTargetList(article, period, ruleset, monthCode, contractTerms, positionTerms, targets, variant)
}

func getCalculFunc(conceptsModel []providers.IConceptSpec, concept types.ConceptCode) providers.ResultFunc {
	conceptSpec, _ := firstOrDefaultConceptSpec(conceptsModel, findConceptSpecCode(concept.Value()))
	if conceptSpec == nil {
		return NotFoundCalculFunc
	}
	return conceptSpec.ResultDelegate()
}

func NotFoundCalculFunc(target types.ITermTarget, spec types.IArticleSpec, period legalios.IPeriod, ruleset legalios.IBundleProps, results providers.IBuilderResultList) providers.IBuilderResultList {
	return providers.IBuilderResultList{types.NewFailureResult(NewNoResultFuncError(period, target))}
}

func targetCompare(topoOrders types.ArticleTermList, x types.ITermTarget, y types.ITermTarget) int {
	var codeOrders = make(types.ArticleCodeList, 0, len(topoOrders))
	for _, t := range topoOrders {
		codeOrders = append(codeOrders, t.Code())
	}
	var xIndex = -1
	for p, v := range codeOrders {
		if v == x.Article() {
			xIndex = p
			break
		}
	}
	var yIndex = -1
	for p, v := range codeOrders {
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

