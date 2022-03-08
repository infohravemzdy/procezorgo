package registry_factories

import (
	legalios "github.com/mzdyhrave/legaliosgo"
	providers "github.com/mzdyhrave/procezorgo/internal/registry_providers"
	"github.com/mzdyhrave/procezorgo/internal/types"
)

const (
	articleCodeNotfound int32 = 0
	articleSeqsNotfound int16 = 0
	conceptCodeNotfound int32 = 0
)

type IArticleSpecFactory interface {
	GetSpec(code types.ArticleCode, period legalios.IPeriod, version types.VersionCode) types.IArticleSpec
	GetSpecList(period legalios.IPeriod, version types.VersionCode) []types.IArticleSpec
	GetProviders() []providers.IArticleSpecProvider
	BuildFactory() bool
}

func NotFoundArticleSpec() types.IArticleSpec {
	ArticleCode := articleCodeNotfound
	ArticleSeqs := articleSeqsNotfound
	ConceptCode := conceptCodeNotfound

	return types.NewArticleSpec(ArticleCode, ArticleSeqs, ConceptCode)
}

type NotFoundArticleProvider struct  {
	providers.IArticleCodeProvider
}

func (p NotFoundArticleProvider) GetSpec(period legalios.IPeriod, version types.VersionCode) types.IArticleSpec {
	return NotFoundArticleSpec()
}

func NewNotFoundArticleProvider() providers.IArticleSpecProvider {
	return NotFoundArticleProvider{providers.NewArticleCodeProvider(articleCodeNotfound) }
}

type BuildArticleProvidersFunc func (f *ArticleSpecFactory) bool

type ProviderRecord struct {
	Article int32
	Sequens int16
	Concept int32
	Sums    []int32
}

func NewProviderRecord(_article int32, _sequens int16, _concept int32, _sums []int32) ProviderRecord {
	return ProviderRecord{Article: _article, Sequens: _sequens, Concept: _concept, Sums: _sums}
}

type ArticleSpecConfig struct {
	types.ArticleSpec
}

func NewArticleSpecConfig(article int32, sequens int16, concept int32, sums []int32) types.IArticleSpec {
	return ArticleSpecConfig{types.NewArticleSumSpec(article, sequens, concept, types.ConstToSumsArray(sums))}
}

type ArticleProviderConfig struct {
	providers.ArticleSpecProvider
	articleSpec types.IArticleSpec
}

func NewArticleProviderConfig(article int32, sequens int16, concept int32, sums []int32) providers.IArticleSpecProvider {
	return ArticleProviderConfig{providers.NewArticleProvider(article), NewArticleSpecConfig(article, sequens, concept, sums)}
}

func (p ArticleProviderConfig) GetSpec(period legalios.IPeriod, version types.VersionCode) types.IArticleSpec {
	return p.articleSpec
}

type ArticleSpecFactory struct {
	Providers          map[int32]providers.IArticleSpecProvider
	NotFoundProvider   providers.IArticleSpecProvider
	NotFoundSpec       types.IArticleSpec
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

func (f ArticleSpecFactory) GetSpec(code types.ArticleCode, period legalios.IPeriod, version types.VersionCode) types.IArticleSpec {
	provider := f.GetProvider(code, f.NotFoundProvider)
	if provider == nil {
		return f.NotFoundSpec
	}
	return provider.GetSpec(period, version)
}

func (f ArticleSpecFactory) GetSpecList(period legalios.IPeriod, version types.VersionCode) []types.IArticleSpec {
	articleSpecs := make([]types.IArticleSpec, 0, len(f.Providers))
	for _, p := range f.Providers {
		articleSpecs = append(articleSpecs, p.GetSpec(period, version))
	}
	return articleSpecs
}
