package procezorgo

import (
	legalios "github.com/mzdyhrave/legaliosgo"
	"github.com/mzdyhrave/procezorgo/internal/registry"
	factories "github.com/mzdyhrave/procezorgo/internal/registry_factories"
	providers "github.com/mzdyhrave/procezorgo/internal/registry_providers"
	"github.com/mzdyhrave/procezorgo/internal/types"
)

type IProcezorService interface {
	Version() types.VersionCode
	FinDefs() types.IArticleDefine

	InitWithPeriod(period IPeriod) bool
	BuildFactories() bool
	GetResults(period IPeriod, ruleset legalios.IBundleProps, targets types.ITermTargetList) providers.IBuilderResultList
	GetArticleSpec(code types.ArticleCode, period IPeriod, version types.VersionCode) providers.IArticleSpec
	GetConceptSpec(code types.ConceptCode, period IPeriod, version types.VersionCode) providers.IConceptSpec
}

type IProcezorFactoryBuilder interface {
	BuildArticleFactory(s *ProcezorService) bool
	BuildConceptFactory(s *ProcezorService) bool
}

type ProcezorService struct {
	factoryBuilder IProcezorFactoryBuilder

	ArticleFactory factories.IArticleSpecFactory
	ConceptFactory factories.IConceptSpecFactory

	resultsBuilder registry.IResultBuilder

	version types.VersionCode
	finDefs types.IArticleDefine
}

func (s ProcezorService) Version() types.VersionCode {
	return s.version
}

func (s ProcezorService) FinDefs() types.IArticleDefine {
	return s.finDefs
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
	return initResult
}

func (s *ProcezorService) BuildFactories() bool {
	articleFactorySuccess := false
	conceptFactorySuccess := false

	if s.factoryBuilder != nil {
		articleFactorySuccess = s.factoryBuilder.BuildArticleFactory(s)
		conceptFactorySuccess = s.factoryBuilder.BuildConceptFactory(s)
	}
	return articleFactorySuccess && conceptFactorySuccess
}

func (s ProcezorService) GetResults(period IPeriod, ruleset legalios.IBundleProps, targets types.ITermTargetList) providers.IBuilderResultList {
	success := s.InitWithPeriod(period)

	if success == false	{
		return make(providers.IBuilderResultList, 0)
	}

	results := s.resultsBuilder.GetResults(ruleset, targets, s.finDefs)

	return results
}

func (s ProcezorService) GetArticleSpec(code types.ArticleCode, period IPeriod, version types.VersionCode) providers.IArticleSpec {
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

func NewProcezorService(version int32, finDefs types.IArticleDefine, serviceSpec IProcezorFactoryBuilder) IProcezorService {
	service := ProcezorService{ version: types.GetVersionCode(version), finDefs: finDefs,
		resultsBuilder: registry.NewResultBuilder(), factoryBuilder: serviceSpec }
	service.BuildFactories()
	return &service
}

