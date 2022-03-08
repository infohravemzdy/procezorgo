package example

import (
	procezor "github.com/mzdyhrave/procezorgo"
)

type TimeshtWorkingConProv struct {
	procezor.ConceptSpecProvider
}

func NewTimeshtWorkingConProv() procezor.IConceptSpecProvider {
	ConceptCode := CONCEPT_TIMESHT_WORKING.Id()

	return &TimeshtWorkingConProv{ConceptSpecProvider: procezor.NewConceptProvider(ConceptCode)}
}

func (p TimeshtWorkingConProv) GetSpec(period procezor.IPeriod, version procezor.VersionCode) procezor.IConceptSpec {
	return NewTimeshtWorkingConSpec(p.Code().Value())
}

type TimeshtWorkingConSpec struct {
	ExampleConceptSpec
}

func NewTimeshtWorkingConSpec(role int32) procezor.IConceptSpec {
	return &TimeshtWorkingConSpec{
		ExampleConceptSpec: NewExampleConceptFuncSpec(role, TimeshtWorkingConceptEval)}
}

func TimeshtWorkingConceptEval(target procezor.ITermTarget, spec procezor.IArticleSpec, period procezor.IPeriod, ruleset procezor.IBundleProps, results procezor.IBuilderResultList) procezor.IBuilderResultList {
	resultsValues := NewTimeshtWorkingResult(target, spec)
	successResult := procezor.NewSuccessResult(resultsValues, period)
	return procezor.IBuilderResultList{successResult}
}

type AmountBasisConProv struct {
	procezor.ConceptSpecProvider
}

func NewAmountBasisConProv() procezor.IConceptSpecProvider {
	ConceptCode := CONCEPT_AMOUNT_BASIS.Id()

	return &AmountBasisConProv{ConceptSpecProvider: procezor.NewConceptProvider(ConceptCode)}
}

func (p AmountBasisConProv) GetSpec(period procezor.IPeriod, version procezor.VersionCode) procezor.IConceptSpec {
	return NewAmountBasisConSpec(p.Code().Value())
}

type AmountBasisConSpec struct {
	ExampleConceptSpec
}

func NewAmountBasisConSpec(role int32) procezor.IConceptSpec {
	_path := []int32{
		ARTICLE_TIMESHT_WORKING.Id(),
	}
	return &AmountBasisConSpec{
		ExampleConceptSpec: NewExampleConceptPathIntFuncSpec(role, _path, AmountBasisConceptEval)}
}

func AmountBasisConceptEval(target procezor.ITermTarget, spec procezor.IArticleSpec, period procezor.IPeriod, ruleset procezor.IBundleProps, results procezor.IBuilderResultList) procezor.IBuilderResultList {
	resultsValues := NewAmountBasisResult(target, spec)
	successResult := procezor.NewSuccessResult(resultsValues, period)
	return procezor.IBuilderResultList{successResult}
}

type AmountFixedConProv struct {
	procezor.ConceptSpecProvider
}

func NewAmountFixedConProv() procezor.IConceptSpecProvider {
	ConceptCode := CONCEPT_AMOUNT_FIXED.Id()

	return &AmountFixedConProv{ConceptSpecProvider: procezor.NewConceptProvider(ConceptCode)}
}

func (p AmountFixedConProv) GetSpec(period procezor.IPeriod, version procezor.VersionCode) procezor.IConceptSpec {
	return NewAmountFixedConSpec(p.Code().Value())
}

type AmountFixedConSpec struct {
	ExampleConceptSpec
}

func NewAmountFixedConSpec(role int32) procezor.IConceptSpec {
	return &AmountFixedConSpec{
		ExampleConceptSpec: NewExampleConceptFuncSpec(role, AmountFixedConceptEval)}
}

func AmountFixedConceptEval(target procezor.ITermTarget, spec procezor.IArticleSpec, period procezor.IPeriod, ruleset procezor.IBundleProps, results procezor.IBuilderResultList) procezor.IBuilderResultList {
	resultsValues := NewAmountFixedResult(target, spec)
	successResult := procezor.NewSuccessResult(resultsValues, period)
	return procezor.IBuilderResultList{successResult}
}

type HealthInsbaseConProv struct {
	procezor.ConceptSpecProvider
}

func NewHealthInsbaseConProv() procezor.IConceptSpecProvider {
	ConceptCode := CONCEPT_HEALTH_INSBASE.Id()

	return &HealthInsbaseConProv{ConceptSpecProvider: procezor.NewConceptProvider(ConceptCode)}
}

func (p HealthInsbaseConProv) GetSpec(period procezor.IPeriod, version procezor.VersionCode) procezor.IConceptSpec {
	return NewHealthInsbaseConSpec(p.Code().Value())
}

type HealthInsbaseConSpec struct {
	ExampleConceptSpec
}

func NewHealthInsbaseConSpec(role int32) procezor.IConceptSpec {
	return &HealthInsbaseConSpec{
		ExampleConceptSpec: NewExampleConceptFuncSpec(role, HealthInsbaseConceptEval)}
}

func HealthInsbaseConceptEval(target procezor.ITermTarget, spec procezor.IArticleSpec, period procezor.IPeriod, ruleset procezor.IBundleProps, results procezor.IBuilderResultList) procezor.IBuilderResultList {
	resultsValues := NewHealthInsbaseResult(target, spec)
	successResult := procezor.NewSuccessResult(resultsValues, period)
	return procezor.IBuilderResultList{successResult}
}

type SocialInsbaseConProv struct {
	procezor.ConceptSpecProvider
}

func NewSocialInsbaseConProv() procezor.IConceptSpecProvider {
	ConceptCode := CONCEPT_SOCIAL_INSBASE.Id()

	return &SocialInsbaseConProv{ConceptSpecProvider: procezor.NewConceptProvider(ConceptCode)}
}

func (p SocialInsbaseConProv) GetSpec(period procezor.IPeriod, version procezor.VersionCode) procezor.IConceptSpec {
	return NewSocialInsbaseConSpec(p.Code().Value())
}

type SocialInsbaseConSpec struct {
	ExampleConceptSpec
}

func NewSocialInsbaseConSpec(role int32) procezor.IConceptSpec {
	return &SocialInsbaseConSpec{
		ExampleConceptSpec: NewExampleConceptFuncSpec(role, SocialInsbaseConceptEval)}
}

func SocialInsbaseConceptEval(target procezor.ITermTarget, spec procezor.IArticleSpec, period procezor.IPeriod, ruleset procezor.IBundleProps, results procezor.IBuilderResultList) procezor.IBuilderResultList {
	resultsValues := NewSocialInsbaseResult(target, spec)
	successResult := procezor.NewSuccessResult(resultsValues, period)
	return procezor.IBuilderResultList{successResult}
}

type HealthInspaymConProv struct {
	procezor.ConceptSpecProvider
}

func NewHealthInspaymConProv() procezor.IConceptSpecProvider {
	ConceptCode := CONCEPT_HEALTH_INSPAYM.Id()

	return &HealthInspaymConProv{ConceptSpecProvider: procezor.NewConceptProvider(ConceptCode)}
}

func (p HealthInspaymConProv) GetSpec(period procezor.IPeriod, version procezor.VersionCode) procezor.IConceptSpec {
	return NewHealthInspaymConSpec(p.Code().Value())
}

type HealthInspaymConSpec struct {
	ExampleConceptSpec
}

func NewHealthInspaymConSpec(role int32) procezor.IConceptSpec {
	_path := []int32{
		ARTICLE_HEALTH_INSBASE.Id(),
	}
	return &HealthInspaymConSpec{
		ExampleConceptSpec: NewExampleConceptPathIntFuncSpec(role, _path, HealthInspaymConceptEval)}
}

func HealthInspaymConceptEval(target procezor.ITermTarget, spec procezor.IArticleSpec, period procezor.IPeriod, ruleset procezor.IBundleProps, results procezor.IBuilderResultList) procezor.IBuilderResultList {
	resultsValues := NewHealthInspaymResult(target, spec)
	successResult := procezor.NewSuccessResult(resultsValues, period)
	return procezor.IBuilderResultList{successResult}
}

type SocialInspaymConProv struct {
	procezor.ConceptSpecProvider
}

func NewSocialInspaymConProv() procezor.IConceptSpecProvider {
	ConceptCode := CONCEPT_SOCIAL_INSPAYM.Id()

	return &SocialInspaymConProv{ConceptSpecProvider: procezor.NewConceptProvider(ConceptCode)}
}

func (p SocialInspaymConProv) GetSpec(period procezor.IPeriod, version procezor.VersionCode) procezor.IConceptSpec {
	return NewSocialInspaymConSpec(p.Code().Value())
}

type SocialInspaymConSpec struct {
	ExampleConceptSpec
}

func NewSocialInspaymConSpec(role int32) procezor.IConceptSpec {
	_path := []int32{
		ARTICLE_SOCIAL_INSBASE.Id(),
	}
	return &SocialInspaymConSpec{
		ExampleConceptSpec: NewExampleConceptPathIntFuncSpec(role, _path, SocialInspaymConceptEval)}
}

func SocialInspaymConceptEval(target procezor.ITermTarget, spec procezor.IArticleSpec, period procezor.IPeriod, ruleset procezor.IBundleProps, results procezor.IBuilderResultList) procezor.IBuilderResultList {
	resultsValues := NewSocialInspaymResult(target, spec)
	successResult := procezor.NewSuccessResult(resultsValues, period)
	return procezor.IBuilderResultList{successResult}
}

type TaxingAdvbaseConProv struct {
	procezor.ConceptSpecProvider
}

func NewTaxingAdvbaseConProv() procezor.IConceptSpecProvider {
	ConceptCode := CONCEPT_TAXING_ADVBASE.Id()

	return &TaxingAdvbaseConProv{ConceptSpecProvider: procezor.NewConceptProvider(ConceptCode)}
}

func (p TaxingAdvbaseConProv) GetSpec(period procezor.IPeriod, version procezor.VersionCode) procezor.IConceptSpec {
	return NewTaxingAdvbaseConSpec(p.Code().Value())
}

type TaxingAdvbaseConSpec struct {
	ExampleConceptSpec
}

func NewTaxingAdvbaseConSpec(role int32) procezor.IConceptSpec {
	return &TaxingAdvbaseConSpec{
		ExampleConceptSpec: NewExampleConceptFuncSpec(role, TaxingAdvbaseConceptEval)}
}

func TaxingAdvbaseConceptEval(target procezor.ITermTarget, spec procezor.IArticleSpec, period procezor.IPeriod, ruleset procezor.IBundleProps, results procezor.IBuilderResultList) procezor.IBuilderResultList {
	resultsValues := NewTaxingAdvbaseResult(target, spec)
	successResult := procezor.NewSuccessResult(resultsValues, period)
	return procezor.IBuilderResultList{successResult}
}

type TaxingAdvpaymConProv struct {
	procezor.ConceptSpecProvider
}

func NewTaxingAdvpaymConProv() procezor.IConceptSpecProvider {
	ConceptCode := CONCEPT_TAXING_ADVPAYM.Id()

	return &TaxingAdvpaymConProv{ConceptSpecProvider: procezor.NewConceptProvider(ConceptCode)}
}

func (p TaxingAdvpaymConProv) GetSpec(period procezor.IPeriod, version procezor.VersionCode) procezor.IConceptSpec {
	return NewTaxingAdvpaymConSpec(p.Code().Value())
}

type TaxingAdvpaymConSpec struct {
	ExampleConceptSpec
}

func NewTaxingAdvpaymConSpec(role int32) procezor.IConceptSpec {
	_path := []int32{
		ARTICLE_TAXING_ADVBASE.Id(),
	}
	return &TaxingAdvpaymConSpec{
		ExampleConceptSpec: NewExampleConceptPathIntFuncSpec(role, _path, TaxingAdvpaymConceptEval)}
}

func TaxingAdvpaymConceptEval(target procezor.ITermTarget, spec procezor.IArticleSpec, period procezor.IPeriod, ruleset procezor.IBundleProps, results procezor.IBuilderResultList) procezor.IBuilderResultList {
	resultsValues := NewTaxingAdvpaymResult(target, spec)
	successResult := procezor.NewSuccessResult(resultsValues, period)
	return procezor.IBuilderResultList{successResult}
}

type IncomeGrossConProv struct {
	procezor.ConceptSpecProvider
}

func NewIncomeGrossConProv() procezor.IConceptSpecProvider {
	ConceptCode := CONCEPT_INCOME_GROSS.Id()

	return &IncomeGrossConProv{ConceptSpecProvider: procezor.NewConceptProvider(ConceptCode)}
}

func (p IncomeGrossConProv) GetSpec(period procezor.IPeriod, version procezor.VersionCode) procezor.IConceptSpec {
	return NewIncomeGrossConSpec(p.Code().Value())
}

type IncomeGrossConSpec struct {
	ExampleConceptSpec
}

func NewIncomeGrossConSpec(role int32) procezor.IConceptSpec {
	return &IncomeGrossConSpec{
		ExampleConceptSpec: NewExampleConceptFuncSpec(role, IncomeGrossConceptEval)}
}

func IncomeGrossConceptEval(target procezor.ITermTarget, spec procezor.IArticleSpec, period procezor.IPeriod, ruleset procezor.IBundleProps, results procezor.IBuilderResultList) procezor.IBuilderResultList {
	resultsValues := NewIncomeGrossResult(target, spec)
	successResult := procezor.NewSuccessResult(resultsValues, period)
	return procezor.IBuilderResultList{successResult}
}

type IncomeNettoConProv struct {
	procezor.ConceptSpecProvider
}

func NewIncomeNettoConProv() procezor.IConceptSpecProvider {
	ConceptCode := CONCEPT_INCOME_NETTO.Id()

	return &IncomeNettoConProv{ConceptSpecProvider: procezor.NewConceptProvider(ConceptCode)}
}

func (p IncomeNettoConProv) GetSpec(period procezor.IPeriod, version procezor.VersionCode) procezor.IConceptSpec {
	return NewIncomeNettoConSpec(p.Code().Value())
}

type IncomeNettoConSpec struct {
	ExampleConceptSpec
}

func NewIncomeNettoConSpec(role int32) procezor.IConceptSpec {
	_path := []int32{
		ARTICLE_INCOME_GROSS.Id(),
		ARTICLE_HEALTH_INSPAYM.Id(),
		ARTICLE_SOCIAL_INSPAYM.Id(),
		ARTICLE_TAXING_ADVPAYM.Id(),
	}
	return &IncomeNettoConSpec{
		ExampleConceptSpec: NewExampleConceptPathIntFuncSpec(role, _path, IncomeNettoConceptEval)}
}

func IncomeNettoConceptEval(target procezor.ITermTarget, spec procezor.IArticleSpec, period procezor.IPeriod, ruleset procezor.IBundleProps, results procezor.IBuilderResultList) procezor.IBuilderResultList {
	resultsValues := NewIncomeNettoResult(target, spec)
	successResult := procezor.NewSuccessResult(resultsValues, period)
	return procezor.IBuilderResultList{successResult}
}

