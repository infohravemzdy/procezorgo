package example

import (
	legalios "github.com/mzdyhrave/legaliosgo"
	procezor "github.com/mzdyhrave/procezorgo"
	"github.com/mzdyhrave/procezorgo/internal/types"
)

const TEST_VERSION int32 = 100

type TestConceptSpec struct {
	procezor.ConceptSpec
}

func NewTestConceptSpec(code int32) TestConceptSpec {
	return TestConceptSpec{procezor.NewConceptSpec(code)}
}

func NewTestConceptPathSpec(code int32, path []procezor.ArticleCode) TestConceptSpec {
	return TestConceptSpec{ procezor.NewConceptPathSpec(code, path)}
}

func NewTestConceptFuncSpec(code int32, resultFunc procezor.ResultFunc) TestConceptSpec {
	return TestConceptSpec{procezor.NewConceptFuncSpec(code, resultFunc)}
}

func NewTestConceptPathFuncSpec(code int32, path []procezor.ArticleCode, resultFunc procezor.ResultFunc) TestConceptSpec {
	return TestConceptSpec{procezor.NewConceptPathFuncSpec(code, path, resultFunc)}
}

type TestTermResult struct {
	procezor.TermResult
}

func NewTestTermResult(target procezor.ITermTarget, value int32, basis int32, descr string) TestTermResult {
	return TestTermResult{ TermResult: procezor.NewTermResult(target, value, basis, descr) }
}

type TestArticleFactory struct {
	procezor.IArticleSpecFactory
}

func NewTestArticleFactory() procezor.IArticleSpecFactory{
	providersConfig := []procezor.ProviderRecord {
		procezor.NewProviderRecord(ARTICLE_TIMESHT_WORKING, CONCEPT_TIMESHT_WORKING,
			[]int32{}),
		procezor.NewProviderRecord(ARTICLE_PAYMENT_SALARY, CONCEPT_AMOUNT_BASIS,
			[]int32 {
				ARTICLE_INCOME_GROSS,
				ARTICLE_HEALTH_INSBASE,
				ARTICLE_SOCIAL_INSBASE,
				ARTICLE_TAXING_ADVBASE,
			}),
		procezor.NewProviderRecord(ARTICLE_PAYMENT_BONUS, CONCEPT_AMOUNT_FIXED,
			[]int32 {
				ARTICLE_INCOME_GROSS,
				ARTICLE_HEALTH_INSBASE,
				ARTICLE_SOCIAL_INSBASE,
				ARTICLE_TAXING_ADVBASE,
			}),
		procezor.NewProviderRecord(ARTICLE_PAYMENT_BARTER, CONCEPT_AMOUNT_FIXED,
			[]int32 {
				ARTICLE_HEALTH_INSBASE,
				ARTICLE_SOCIAL_INSBASE,
				ARTICLE_TAXING_ADVBASE,
			}),
		procezor.NewProviderRecord(ARTICLE_ALLOWCE_HOFFICE, CONCEPT_AMOUNT_FIXED,
			[]int32 {
				ARTICLE_INCOME_NETTO,
			}),
		procezor.NewProviderRecord(ARTICLE_HEALTH_INSBASE, CONCEPT_HEALTH_INSBASE,
			[]int32{}),
		procezor.NewProviderRecord(ARTICLE_SOCIAL_INSBASE, CONCEPT_SOCIAL_INSBASE,
			[]int32{}),
		procezor.NewProviderRecord(ARTICLE_HEALTH_INSPAYM, CONCEPT_HEALTH_INSPAYM,
			[]int32{}),
		procezor.NewProviderRecord(ARTICLE_SOCIAL_INSPAYM, CONCEPT_SOCIAL_INSPAYM,
			[]int32{}),
		procezor.NewProviderRecord(ARTICLE_TAXING_ADVBASE, CONCEPT_TAXING_ADVBASE,
			[]int32{}),
		procezor.NewProviderRecord(ARTICLE_TAXING_ADVPAYM, CONCEPT_TAXING_ADVPAYM,
			[]int32{}),
		procezor.NewProviderRecord(ARTICLE_INCOME_GROSS, CONCEPT_INCOME_GROSS,
			[]int32{}),
		procezor.NewProviderRecord(ARTICLE_INCOME_NETTO, CONCEPT_INCOME_NETTO,
			[]int32{}),
	}
	return TestArticleFactory{procezor.NewArticleSpecFactoryWithRecords(providersConfig)}
}

type TestConceptFactory struct {
	procezor.IConceptSpecFactory
}

func NewTestConceptFactory() procezor.IConceptSpecFactory {
	return TestConceptFactory{procezor.NewConceptSpecFactoryWithProviders(
		map[int32]procezor.IConceptSpecProvider {
			CONCEPT_TIMESHT_WORKING: NewTimeshtWorkingConProv(),
			CONCEPT_AMOUNT_BASIS: NewAmountBasisConProv(),
			CONCEPT_AMOUNT_FIXED: NewAmountFixedConProv(),
			CONCEPT_HEALTH_INSBASE: NewHealthInsbaseConProv(),
			CONCEPT_SOCIAL_INSBASE: NewSocialInsbaseConProv(),
			CONCEPT_HEALTH_INSPAYM: NewHealthInspaymConProv(),
			CONCEPT_SOCIAL_INSPAYM: NewSocialInspaymConProv(),
			CONCEPT_TAXING_ADVBASE: NewTaxingAdvbaseConProv(),
			CONCEPT_TAXING_ADVPAYM: NewTaxingAdvpaymConProv(),
			CONCEPT_INCOME_GROSS: NewIncomeGrossConProv(),
			CONCEPT_INCOME_NETTO: NewIncomeNettoConProv(),
		}),
	}
}

type TestServiceBuilder struct {

}

func (t TestServiceBuilder) BuildArticleFactory(s *procezor.ProcezorService) bool {
	s.ArticleFactory = NewTestArticleFactory()
	if s.ArticleFactory == nil {
		return false
	}
	s.ArticleFactory.BuildFactory()
	return true
}

func (t TestServiceBuilder) BuildConceptFactory(s *procezor.ProcezorService) bool {
	s.ConceptFactory = NewTestConceptFactory()
	if s.ConceptFactory == nil {
		return false
	}
	s.ConceptFactory.BuildFactory()
	return true
}

type TestService struct {
	procezor.IProcezorService
}

func NewTestServiceBuilder() procezor.IProcezorFactoryBuilder {
	return &TestServiceBuilder{}
}

func NewTestService() procezor.IProcezorService{
	const (
		TestVersion      = TEST_VERSION
		TestFinalArticle = ARTICLE_INCOME_NETTO
		TestFinalConcept = CONCEPT_INCOME_NETTO
	)
	var (
		TestFinalDefs = procezor.GetArticleDefine(TestFinalArticle, TestFinalConcept)
	)

	return &TestService{
		procezor.NewProcezorService(TestVersion, TestFinalDefs, NewTestServiceBuilder()),
	}
}

func GetTargetsWithSalaryHomeOffice(period legalios.IPeriod) types.ITermTargetList {
	const CONTRACT_CODE = 0
	const POSITION_CODE = 0
	const VARIANT1_CODE = 1

	var montCode = types.GetMonthCode(period.GetCode())
	var contract = types.GetContractCode(CONTRACT_CODE)
	var position = types.GetPositionCode(POSITION_CODE)
	var variant1 = types.GetVariantCode(VARIANT1_CODE)

	var targets = []types.ITermTarget {
		types.NewTermTarget(montCode, contract, position, variant1,
			procezor.GetArticleCode(ARTICLE_TIMESHT_WORKING), procezor.GetConceptCode(CONCEPT_TIMESHT_WORKING)),
		types.NewTermTarget(montCode, contract, position, variant1,
			procezor.GetArticleCode(ARTICLE_PAYMENT_SALARY),procezor.GetConceptCode(CONCEPT_AMOUNT_BASIS)),
		types.NewTermTarget(montCode, contract, position, variant1,
			procezor.GetArticleCode(ARTICLE_ALLOWCE_HOFFICE),procezor.GetConceptCode(CONCEPT_AMOUNT_FIXED)),
	}
	return targets
}

func GetTargetsWithSalaryBonusBarter(period legalios.IPeriod) types.ITermTargetList {
	const CONTRACT_CODE = 0
	const POSITION_CODE = 0
	const VARIANT1_CODE = 1

	var montCode = types.GetMonthCode(period.GetCode())
	var contract = types.GetContractCode(CONTRACT_CODE)
	var position = types.GetPositionCode(POSITION_CODE)
	var variant1 = types.GetVariantCode(VARIANT1_CODE)

	var targets = []types.ITermTarget {
		types.NewTermTarget(montCode, contract, position, variant1,
			procezor.GetArticleCode(ARTICLE_TIMESHT_WORKING), procezor.GetConceptCode(CONCEPT_TIMESHT_WORKING)),
		types.NewTermTarget(montCode, contract, position, variant1,
			procezor.GetArticleCode(ARTICLE_PAYMENT_SALARY), procezor.GetConceptCode(CONCEPT_AMOUNT_BASIS)),
		types.NewTermTarget(montCode, contract, position, variant1,
			procezor.GetArticleCode(ARTICLE_PAYMENT_BONUS), procezor.GetConceptCode(CONCEPT_AMOUNT_FIXED)),
		types.NewTermTarget(montCode, contract, position, variant1,
			procezor.GetArticleCode(ARTICLE_PAYMENT_BARTER), procezor.GetConceptCode(CONCEPT_AMOUNT_FIXED)),
	}
	return targets
}

