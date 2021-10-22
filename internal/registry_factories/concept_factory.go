package registry_factories

import (
	legalios "github.com/mzdyhrave/payrollgo-legalios/pkg/service"
	providers "github.com/mzdyhrave/payrollgo-procezor/internal/registry_providers"
	"github.com/mzdyhrave/payrollgo-procezor/internal/types"
)

type IConceptSpecFactory interface {
	GetSpec(code types.ConceptCode, period legalios.IPeriod, version types.VersionCode) providers.IConceptSpec
	GetSpecList(period legalios.IPeriod, version types.VersionCode) []providers.IConceptSpec
	GetProviders() []providers.IConceptSpecProvider
	BuildFactory() bool
}

func NotFoundConceptSpec() providers.IConceptSpec {
	ConceptCode := conceptNotfound

	return providers.NewConceptSpec(ConceptCode)
}

type NotFoundConceptProvider struct {
	providers.IConceptCodeProvider
}

func (p NotFoundConceptProvider) GetSpec(period legalios.IPeriod, version types.VersionCode) providers.IConceptSpec {
	return NotFoundConceptSpec()
}

func NewNotFoundConceptProvider() providers.IConceptSpecProvider {
	return NotFoundConceptProvider{providers.NewConceptCodeProvider(conceptNotfound) }
}

type BuildConceptProvidersFunc func (f *ConceptSpecFactory) bool

type ConceptSpecFactory struct {
	Providers          map[int32]providers.IConceptSpecProvider
	NotFoundProvider   providers.IConceptSpecProvider
	NotFoundSpec       providers.IConceptSpec
	BuildProvidersFunc BuildConceptProvidersFunc
}

func (f ConceptSpecFactory) GetProvider(code types.ConceptCode, defProvider providers.IConceptSpecProvider) providers.IConceptSpecProvider {
	provider, exists := f.Providers[code.Value()]
	if exists == false {
		return defProvider
	}
	return provider
}

func (f *ConceptSpecFactory) BuildFactory() bool {
	return f.BuildProvidersFunc(f)
}

func (f ConceptSpecFactory) GetProviders() []providers.IConceptSpecProvider {
	conceptProviders := make([]providers.IConceptSpecProvider, 0, len(f.Providers))
	for _, p := range f.Providers {
		conceptProviders = append(conceptProviders, p)
	}
	return conceptProviders
}

func (f ConceptSpecFactory) GetSpec(code types.ConceptCode, period legalios.IPeriod, version types.VersionCode) providers.IConceptSpec {
	provider := f.GetProvider(code, f.NotFoundProvider)
	if provider == nil {
		return f.NotFoundSpec
	}
	return provider.GetSpec(period, version)
}

func (f ConceptSpecFactory) GetSpecList(period legalios.IPeriod, version types.VersionCode) []providers.IConceptSpec {
	conceptSpecs := make([]providers.IConceptSpec, 0, len(f.Providers))
	for _, p := range f.Providers {
		conceptSpecs = append(conceptSpecs, p.GetSpec(period, version))
	}
	return conceptSpecs
}
