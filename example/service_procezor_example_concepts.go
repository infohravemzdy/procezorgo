package example

import (
	legalios "github.com/mzdyhrave/legaliosgo"
	procezor "github.com/mzdyhrave/procezorgo"
)

type TimeshtWorkingConProv struct {
	procezor.ConceptSpecProvider
}

func (p TimeshtWorkingConProv) GetSpec(period procezor.IPeriod, version procezor.VersionCode) procezor.IConceptSpec {
	return NewTimeshtWorkingConSpec(p.Code().Value())
}

type TimeshtWorkingConSpec struct {
	TestConceptSpec
}

func TimeshtWorkingConceptEval(target procezor.ITermTarget, period procezor.IPeriod, ruleset legalios.IBundleProps, results procezor.IBuilderResultList) procezor.IBuilderResultList {
	resultsValues := NewTimeshtWorkingResult(target, RESULT_VALUE_ZERO, RESULT_BASIS_ZERO, RESULT_DESCRIPTION_EMPTY)
	successResult := procezor.NewSuccessResult(resultsValues, period)
	return procezor.IBuilderResultList{successResult}
}

func NewTimeshtWorkingConSpec(role int32) procezor.IConceptSpec {
	return &TimeshtWorkingConSpec{TestConceptSpec: NewTestConceptFuncSpec(role, TimeshtWorkingConceptEval) }
}

func NewTimeshtWorkingConProv() procezor.IConceptSpecProvider {
	const (
		CONCEPT_CODE int32 = CONCEPT_TIMESHT_WORKING
	)
	return &TimeshtWorkingConProv{ConceptSpecProvider: procezor.NewConceptProvider(CONCEPT_CODE)}
}

type AmountBasisConProv struct {
	procezor.ConceptSpecProvider
}

func (p AmountBasisConProv) GetSpec(period procezor.IPeriod, version procezor.VersionCode) procezor.IConceptSpec {
	return NewAmountBasisConSpec(p.Code().Value())
}

type AmountBasisConSpec struct {
	TestConceptSpec
}

func AmountBasisConceptEval(target procezor.ITermTarget, period procezor.IPeriod, ruleset legalios.IBundleProps, results procezor.IBuilderResultList) procezor.IBuilderResultList {
	resultsValues := NewAmountBasisResult(target, RESULT_VALUE_ZERO, RESULT_BASIS_ZERO, RESULT_DESCRIPTION_EMPTY)
	successResult := procezor.NewSuccessResult(resultsValues, period)
	return procezor.IBuilderResultList{successResult}
}

func NewAmountBasisConSpec(role int32) procezor.IConceptSpec {
	return &AmountBasisConSpec{TestConceptSpec: NewTestConceptPathFuncSpec(role,
		procezor.ArticleCodeList{
			procezor.GetArticleCode(ARTICLE_TIMESHT_WORKING),
		}, AmountBasisConceptEval) }
}

func NewAmountBasisConProv() procezor.IConceptSpecProvider {
	const (
		CONCEPT_CODE int32 = CONCEPT_AMOUNT_BASIS
	)
	return &AmountBasisConProv{ConceptSpecProvider: procezor.NewConceptProvider(CONCEPT_CODE)}
}

type AmountFixedConProv struct {
	procezor.ConceptSpecProvider
}

func (p AmountFixedConProv) GetSpec(period procezor.IPeriod, version procezor.VersionCode) procezor.IConceptSpec {
	return NewAmountFixedConSpec(p.Code().Value())
}

type AmountFixedConSpec struct {
	TestConceptSpec
}

func AmountFixedConceptEval(target procezor.ITermTarget, period procezor.IPeriod, ruleset legalios.IBundleProps, results procezor.IBuilderResultList) procezor.IBuilderResultList {
	resultsValues := NewAmountFixedResult(target, RESULT_VALUE_ZERO, RESULT_BASIS_ZERO, RESULT_DESCRIPTION_EMPTY)
	successResult := procezor.NewSuccessResult(resultsValues, period)
	return procezor.IBuilderResultList{successResult}
}

func NewAmountFixedConSpec(role int32) procezor.IConceptSpec {
	return &AmountFixedConSpec{TestConceptSpec: NewTestConceptFuncSpec(role, AmountFixedConceptEval) }
}

func NewAmountFixedConProv() procezor.IConceptSpecProvider {
	const (
		CONCEPT_CODE int32 = CONCEPT_AMOUNT_FIXED
	)
	return &AmountFixedConProv{ConceptSpecProvider: procezor.NewConceptProvider(CONCEPT_CODE)}
}

type HealthInsbaseConProv struct {
	procezor.ConceptSpecProvider
}

func (p HealthInsbaseConProv) GetSpec(period procezor.IPeriod, version procezor.VersionCode) procezor.IConceptSpec {
	return NewHealthInsbaseConSpec(p.Code().Value())
}

type HealthInsbaseConSpec struct {
	TestConceptSpec
}

func HealthInsbaseConceptEval(target procezor.ITermTarget, period procezor.IPeriod, ruleset legalios.IBundleProps, results procezor.IBuilderResultList) procezor.IBuilderResultList {
	resultsValues := NewHealthInsbaseResult(target, RESULT_VALUE_ZERO, RESULT_BASIS_ZERO, RESULT_DESCRIPTION_EMPTY)
	successResult := procezor.NewSuccessResult(resultsValues, period)
	return procezor.IBuilderResultList{successResult}
}

func NewHealthInsbaseConSpec(role int32) procezor.IConceptSpec {
	return &HealthInsbaseConSpec{TestConceptSpec: NewTestConceptFuncSpec(role, HealthInsbaseConceptEval) }
}

func NewHealthInsbaseConProv() procezor.IConceptSpecProvider {
	const (
		CONCEPT_CODE int32 = CONCEPT_HEALTH_INSBASE
	)
	return &HealthInsbaseConProv{ConceptSpecProvider: procezor.NewConceptProvider(CONCEPT_CODE)}
}

type SocialInsbaseConProv struct {
	procezor.ConceptSpecProvider
}

func (p SocialInsbaseConProv) GetSpec(period procezor.IPeriod, version procezor.VersionCode) procezor.IConceptSpec {
	return NewSocialInsbaseConSpec(p.Code().Value())
}

type SocialInsbaseConSpec struct {
	TestConceptSpec
}

func SocialInsbaseConceptEval(target procezor.ITermTarget, period procezor.IPeriod, ruleset legalios.IBundleProps, results procezor.IBuilderResultList) procezor.IBuilderResultList {
	resultsValues := NewSocialInsbaseResult(target, RESULT_VALUE_ZERO, RESULT_BASIS_ZERO, RESULT_DESCRIPTION_EMPTY)
	successResult := procezor.NewSuccessResult(resultsValues, period)
	return procezor.IBuilderResultList{successResult}
}

func NewSocialInsbaseConSpec(role int32) procezor.IConceptSpec {
	return &SocialInsbaseConSpec{TestConceptSpec: NewTestConceptFuncSpec(role, SocialInsbaseConceptEval) }
}

func NewSocialInsbaseConProv() procezor.IConceptSpecProvider {
	const (
		CONCEPT_CODE int32 = CONCEPT_SOCIAL_INSBASE
	)
	return &SocialInsbaseConProv{ConceptSpecProvider: procezor.NewConceptProvider(CONCEPT_CODE)}
}

type TaxingAdvbaseConProv struct {
	procezor.ConceptSpecProvider
}

func (p TaxingAdvbaseConProv) GetSpec(period procezor.IPeriod, version procezor.VersionCode) procezor.IConceptSpec {
	return NewTaxingAdvbaseConSpec(p.Code().Value())
}

type TaxingAdvbaseConSpec struct {
	TestConceptSpec
}

func TaxingAdvbaseConceptEval(target procezor.ITermTarget, period procezor.IPeriod, ruleset legalios.IBundleProps, results procezor.IBuilderResultList) procezor.IBuilderResultList {
	resultsValues := NewTaxingAdvbaseResult(target, RESULT_VALUE_ZERO, RESULT_BASIS_ZERO, RESULT_DESCRIPTION_EMPTY)
	successResult := procezor.NewSuccessResult(resultsValues, period)
	return procezor.IBuilderResultList{successResult}
}

func NewTaxingAdvbaseConSpec(role int32) procezor.IConceptSpec {
	return &TaxingAdvbaseConSpec{TestConceptSpec: NewTestConceptFuncSpec(role, TaxingAdvbaseConceptEval) }
}

func NewTaxingAdvbaseConProv() procezor.IConceptSpecProvider {
	const (
		CONCEPT_CODE int32 = CONCEPT_TAXING_ADVBASE
	)
	return &TaxingAdvbaseConProv{ConceptSpecProvider: procezor.NewConceptProvider(CONCEPT_CODE)}
}

type HealthInspaymConProv struct {
	procezor.ConceptSpecProvider
}

func (p HealthInspaymConProv) GetSpec(period procezor.IPeriod, version procezor.VersionCode) procezor.IConceptSpec {
	return NewHealthInspaymConSpec(p.Code().Value())
}

type HealthInspaymConSpec struct {
	TestConceptSpec
}

func HealthInspaymConceptEval(target procezor.ITermTarget, period procezor.IPeriod, ruleset legalios.IBundleProps, results procezor.IBuilderResultList) procezor.IBuilderResultList {
	resultsValues := NewHealthInspaymResult(target, RESULT_VALUE_ZERO, RESULT_BASIS_ZERO, RESULT_DESCRIPTION_EMPTY)
	successResult := procezor.NewSuccessResult(resultsValues, period)
	return procezor.IBuilderResultList{successResult}
}

func NewHealthInspaymConSpec(role int32) procezor.IConceptSpec {
	return &HealthInspaymConSpec{TestConceptSpec: NewTestConceptPathFuncSpec(role,
		procezor.ArticleCodeList{
			procezor.GetArticleCode(ARTICLE_HEALTH_INSBASE),
		}, HealthInspaymConceptEval) }
}

func NewHealthInspaymConProv() procezor.IConceptSpecProvider {
	const (
		CONCEPT_CODE int32 = CONCEPT_HEALTH_INSPAYM
	)
	return &HealthInspaymConProv{ConceptSpecProvider: procezor.NewConceptProvider(CONCEPT_CODE)}
}

type SocialInspaymConProv struct {
	procezor.ConceptSpecProvider
}

func (p SocialInspaymConProv) GetSpec(period procezor.IPeriod, version procezor.VersionCode) procezor.IConceptSpec {
	return NewSocialInspaymConSpec(p.Code().Value())
}

type SocialInspaymConSpec struct {
	TestConceptSpec
}

func SocialInspaymConceptEval(target procezor.ITermTarget, period procezor.IPeriod, ruleset legalios.IBundleProps, results procezor.IBuilderResultList) procezor.IBuilderResultList {
	resultsValues := NewSocialInspaymResult(target, RESULT_VALUE_ZERO, RESULT_BASIS_ZERO, RESULT_DESCRIPTION_EMPTY)
	successResult := procezor.NewSuccessResult(resultsValues, period)
	return procezor.IBuilderResultList{successResult}
}

func NewSocialInspaymConSpec(role int32) procezor.IConceptSpec {
	return &SocialInspaymConSpec{TestConceptSpec: NewTestConceptPathFuncSpec(role,
		procezor.ArticleCodeList{
			procezor.GetArticleCode(ARTICLE_SOCIAL_INSBASE),
		}, SocialInspaymConceptEval) }
}

func NewSocialInspaymConProv() procezor.IConceptSpecProvider {
	const (
		CONCEPT_CODE int32 = CONCEPT_SOCIAL_INSPAYM
	)
	return &SocialInspaymConProv{ConceptSpecProvider: procezor.NewConceptProvider(CONCEPT_CODE)}
}

type TaxingAdvpaymConProv struct {
	procezor.ConceptSpecProvider
}

func (p TaxingAdvpaymConProv) GetSpec(period procezor.IPeriod, version procezor.VersionCode) procezor.IConceptSpec {
	return NewTaxingAdvpaymConSpec(p.Code().Value())
}

type TaxingAdvpaymConSpec struct {
	TestConceptSpec
}

func TaxingAdvpaymConceptEval(target procezor.ITermTarget, period procezor.IPeriod, ruleset legalios.IBundleProps, results procezor.IBuilderResultList) procezor.IBuilderResultList {
	resultsValues := NewTaxingAdvpaymResult(target, RESULT_VALUE_ZERO, RESULT_BASIS_ZERO, RESULT_DESCRIPTION_EMPTY)
	successResult := procezor.NewSuccessResult(resultsValues, period)
	return procezor.IBuilderResultList{successResult}
}

func NewTaxingAdvpaymConSpec(role int32) procezor.IConceptSpec {
	return &TaxingAdvpaymConSpec{TestConceptSpec: NewTestConceptPathFuncSpec(role,
		procezor.ArticleCodeList{
			procezor.GetArticleCode(ARTICLE_TAXING_ADVBASE),
		}, TaxingAdvpaymConceptEval) }
}

func NewTaxingAdvpaymConProv() procezor.IConceptSpecProvider {
	const (
		CONCEPT_CODE int32 = CONCEPT_TAXING_ADVPAYM
	)
	return &TaxingAdvpaymConProv{ConceptSpecProvider: procezor.NewConceptProvider(CONCEPT_CODE)}
}

type IncomeGrossConProv struct {
	procezor.ConceptSpecProvider
}

func (p IncomeGrossConProv) GetSpec(period procezor.IPeriod, version procezor.VersionCode) procezor.IConceptSpec {
	return NewIncomeGrossConSpec(p.Code().Value())
}

type IncomeGrossConSpec struct {
	TestConceptSpec
}

func IncomeGrossConceptEval(target procezor.ITermTarget, period procezor.IPeriod, ruleset legalios.IBundleProps, results procezor.IBuilderResultList) procezor.IBuilderResultList {
	resultsValues := NewIncomeGrossResult(target, RESULT_VALUE_ZERO, RESULT_BASIS_ZERO, RESULT_DESCRIPTION_EMPTY)
	successResult := procezor.NewSuccessResult(resultsValues, period)
	return procezor.IBuilderResultList{successResult}
}

func NewIncomeGrossConSpec(role int32) procezor.IConceptSpec {
	return &IncomeGrossConSpec{TestConceptSpec: NewTestConceptFuncSpec(role, IncomeGrossConceptEval) }
}

func NewIncomeGrossConProv() procezor.IConceptSpecProvider {
	const (
		CONCEPT_CODE int32 = CONCEPT_INCOME_GROSS
	)
	return &IncomeGrossConProv{ConceptSpecProvider: procezor.NewConceptProvider(CONCEPT_CODE)}
}

type IncomeNettoConProv struct {
	procezor.ConceptSpecProvider
}

func (p IncomeNettoConProv) GetSpec(period procezor.IPeriod, version procezor.VersionCode) procezor.IConceptSpec {
	return NewIncomeNettoConSpec(p.Code().Value())
}

type IncomeNettoConSpec struct {
	TestConceptSpec
}

func IncomeNettoConceptEval(target procezor.ITermTarget, period procezor.IPeriod, ruleset legalios.IBundleProps, results procezor.IBuilderResultList) procezor.IBuilderResultList {
	resultsValues := NewIncomeNettoResult(target, RESULT_VALUE_ZERO, RESULT_BASIS_ZERO, RESULT_DESCRIPTION_EMPTY)
	successResult := procezor.NewSuccessResult(resultsValues, period)
	return procezor.IBuilderResultList{successResult}
}

func NewIncomeNettoConSpec(role int32) procezor.IConceptSpec {
	return &IncomeNettoConSpec{TestConceptSpec: NewTestConceptPathFuncSpec(role,
		procezor.ArticleCodeList{
			procezor.GetArticleCode(ARTICLE_INCOME_GROSS),
			procezor.GetArticleCode(ARTICLE_HEALTH_INSPAYM),
			procezor.GetArticleCode(ARTICLE_SOCIAL_INSPAYM),
			procezor.GetArticleCode(ARTICLE_TAXING_ADVPAYM),
		}, IncomeNettoConceptEval) }
}

func NewIncomeNettoConProv() procezor.IConceptSpecProvider {
	const (
		CONCEPT_CODE int32 = CONCEPT_INCOME_NETTO
	)
	return &IncomeNettoConProv{ConceptSpecProvider: procezor.NewConceptProvider(CONCEPT_CODE)}
}



