package procezorgo

import (
	"fmt"
	legalios "github.com/mzdyhrave/legaliosgo"
	"github.com/mzdyhrave/procezorgo/internal/registry"
	factories "github.com/mzdyhrave/procezorgo/internal/registry_factories"
	providers "github.com/mzdyhrave/procezorgo/internal/registry_providers"
	"github.com/mzdyhrave/procezorgo/internal/types"
)

type IProcezorService interface {
	Version() types.VersionCode
	CalcArticles() types.ArticleCodeList

	InitWithPeriod(period IPeriod) bool
	BuildFactories() bool

	BuilderOrder() OrderCodeList
	BuilderPaths() PathsTermsMap

	GetResults(period IPeriod, ruleset legalios.IBundleProps, targets types.ITermTargetList) providers.IBuilderResultList
	GetArticleSpec(code types.ArticleCode, period IPeriod, version types.VersionCode) types.IArticleSpec
	GetConceptSpec(code types.ConceptCode, period IPeriod, version types.VersionCode) providers.IConceptSpec
}

type IProcezorFactoryBuilder interface {
	BuildArticleFactory(s *ProcezorService) bool
	BuildConceptFactory(s *ProcezorService) bool
}

type IProcezorContractBuilder interface {
	GetContractTerms(period IPeriod, targets ITermTargetList) IContractTermList
}

type IProcezorPositionBuilder interface {
	GetPositionTerms(period IPeriod, contracts IContractTermList, targets ITermTargetList) IPositionTermList
}

type ProcezorService struct {
	factoryBuilder IProcezorFactoryBuilder
	contractBuilder IProcezorContractBuilder
	positionBuilder IProcezorPositionBuilder

	ArticleFactory factories.IArticleSpecFactory
	ConceptFactory factories.IConceptSpecFactory

	resultsBuilder registry.IResultBuilder

	version types.VersionCode
	calcArticles types.ArticleCodeList
}

func (s ProcezorService) Version() types.VersionCode {
	return s.version
}

func (s ProcezorService) CalcArticles() types.ArticleCodeList {
	return s.calcArticles
}

func (s ProcezorService) BuilderOrder() OrderCodeList {
	return s.resultsBuilder.GetOrder()
}
func (s ProcezorService) BuilderPaths() PathsTermsMap {
	return s.resultsBuilder.GetPaths()
}
func (s *ProcezorService) InitWithPeriod(period IPeriod) bool {
	initResult := false

	if s.resultsBuilder != nil {
		initResult = true
	}

	initBuilder := false

	if s.resultsBuilder != nil {
		initBuilder = s.resultsBuilder.GetPeriod().Equals(period)==false
	}

	if initBuilder && s.ArticleFactory != nil && s.ConceptFactory != nil {
		initResult = s.resultsBuilder.InitWithPeriod(s.version, period, s.ArticleFactory, s.ConceptFactory)
	}
	if initResult == false {
		println(fmt.Sprintf("Period: %d, init with period failed", period.GetCode()))
	}
	return initResult
}

func (s *ProcezorService) BuildFactories() bool {
	articleFactorySuccess := false
	conceptFactorySuccess := false

	if s.factoryBuilder != nil {
		articleFactorySuccess = s.factoryBuilder.BuildArticleFactory(s)
		conceptFactorySuccess = s.factoryBuilder.BuildConceptFactory(s)
	}
	if !(articleFactorySuccess && conceptFactorySuccess) {
		println(fmt.Sprintf("ServiceProcezor::BuildFactories(): Version: %d, build factories failed", s.version.Value()))
	}
	return articleFactorySuccess && conceptFactorySuccess
}

func (s ProcezorService) GetResults(period IPeriod, ruleset legalios.IBundleProps, targets types.ITermTargetList) providers.IBuilderResultList {
	success := s.InitWithPeriod(period)

	if success == false	{
		return make(providers.IBuilderResultList, 0)
	}

	contractTerms := s.contractBuilder.GetContractTerms(period, targets)
	positionTerms := s.positionBuilder.GetPositionTerms(period, contractTerms, targets)

	if s.factoryBuilder != nil {
	}

	results := s.resultsBuilder.GetResults(ruleset, contractTerms, positionTerms, targets, s.calcArticles)

	return results
}

func (s ProcezorService) GetArticleSpec(code types.ArticleCode, period IPeriod, version types.VersionCode) types.IArticleSpec {
	if  s.ArticleFactory == nil {
		return nil
	}
	return s.ArticleFactory.GetSpec(code, period, version)
}

func (s ProcezorService) GetConceptSpec(code types.ConceptCode, period IPeriod, version types.VersionCode) providers.IConceptSpec {
	if  s.ConceptFactory == nil {
		return nil
	}
	return s.ConceptFactory.GetSpec(code, period, version)
}

func NewProcezorService(version int32, calcArts types.ArticleCodeList,
	contractSpec IProcezorContractBuilder, positionSpec IProcezorPositionBuilder,
	serviceSpec IProcezorFactoryBuilder) IProcezorService {
	service := ProcezorService{ version: types.GetVersionCode(version), calcArticles: calcArts,
		resultsBuilder: registry.NewResultBuilder(),
		contractBuilder: contractSpec, positionBuilder: positionSpec, factoryBuilder: serviceSpec }
	if service.BuildFactories() == false {
		println(fmt.Sprintf("Version: %d, build factories failed", version))
	}
	return &service
}

