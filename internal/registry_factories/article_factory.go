package registry_factories

import (
	legalios "github.com/mzdyhrave/payrollgo-legalios/pkg/service"
	providers "github.com/mzdyhrave/payrollgo-procezor/internal/registry_providers"
	"github.com/mzdyhrave/payrollgo-procezor/internal/types"
)

const (
	articleNotfound int32 = 0
	conceptNotfound int32 = 0
)

type IArticleSpecFactory interface {
	GetSpec(code types.ArticleCode, period legalios.IPeriod, version types.VersionCode) providers.IArticleSpec
	GetSpecList(period legalios.IPeriod, version types.VersionCode) []providers.IArticleSpec
	GetProviders() []providers.IArticleSpecProvider
	BuildFactory() bool
}

func NotFoundArticleSpec() providers.IArticleSpec {
	ArticleCode := articleNotfound
	ConceptCode := conceptNotfound

	return providers.NewArticleSpec(ArticleCode, ConceptCode)
}

type NotFoundArticleProvider struct  {
	providers.IArticleCodeProvider
}

func (p NotFoundArticleProvider) GetSpec(period legalios.IPeriod, version types.VersionCode) providers.IArticleSpec {
	return NotFoundArticleSpec()
}

func NewNotFoundArticleProvider() providers.IArticleSpecProvider {
	return NotFoundArticleProvider{providers.NewArticleCodeProvider(articleNotfound) }
}

type BuildArticleProvidersFunc func (f *ArticleSpecFactory) bool

type ProviderRecord struct {
	Article int32
	Concept int32
	Sums    []int32
}

func NewProviderRecord(_article int32, _concept int32, _sums []int32) ProviderRecord {
	return ProviderRecord{Article: _article, Concept: _concept, Sums: _sums}
}

type ArticleSpecConfig struct {
	providers.ArticleSpec
}

func NewArticleSpecConfig(article int32, concept int32, sums []int32) providers.IArticleSpec {
	return ArticleSpecConfig{providers.NewArticleSumSpec(article, concept, providers.ConstToSumsArray(sums))}
}

type ArticleProviderConfig struct {
	providers.ArticleSpecProvider
	articleSpec providers.IArticleSpec
}

func NewArticleProviderConfig(article int32, concept int32, sums []int32) providers.IArticleSpecProvider {
	return ArticleProviderConfig{providers.NewArticleProvider(article), NewArticleSpecConfig(article, concept, sums)}
}

func (p ArticleProviderConfig) GetSpec(period legalios.IPeriod, version types.VersionCode) providers.IArticleSpec {
	return p.articleSpec
}

type ArticleSpecFactory struct {
	Providers          map[int32]providers.IArticleSpecProvider
	NotFoundProvider   providers.IArticleSpecProvider
	NotFoundSpec       providers.IArticleSpec
	BuildProvidersFunc BuildArticleProvidersFunc
}

func (f ArticleSpecFactory) GetProvider(code types.ArticleCode, defProvider providers.IArticleSpecProvider) providers.IArticleSpecProvider {
	provider, exists := f.Providers[code.Value()]
	if exists == false {
		return defProvider
	}
	return provider
}

func (f *ArticleSpecFactory) BuildFactory() bool {
	return f.BuildProvidersFunc(f)
}

func (f ArticleSpecFactory) GetProviders() []providers.IArticleSpecProvider {
	articleProviders := make([]providers.IArticleSpecProvider, 0, len(f.Providers))
	for _, p := range f.Providers {
		articleProviders = append(articleProviders, p)
	}
	return articleProviders
}

func (f ArticleSpecFactory) GetSpec(code types.ArticleCode, period legalios.IPeriod, version types.VersionCode) providers.IArticleSpec {
	provider := f.GetProvider(code, f.NotFoundProvider)
	if provider == nil {
		return f.NotFoundSpec
	}
	return provider.GetSpec(period, version)
}

func (f ArticleSpecFactory) GetSpecList(period legalios.IPeriod, version types.VersionCode) []providers.IArticleSpec {
	articleSpecs := make([]providers.IArticleSpec, 0, len(f.Providers))
	for _, p := range f.Providers {
		articleSpecs = append(articleSpecs, p.GetSpec(period, version))
	}
	return articleSpecs
}
