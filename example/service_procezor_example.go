package example

import (
	"fmt"
	legalios "github.com/mzdyhrave/legaliosgo"
	procezor "github.com/mzdyhrave/procezorgo"
	"github.com/mzdyhrave/procezorgo/internal/types"
)

const (
	ARTICLE_DEFAULT_SEQUENS int16 = 0
)
type ExampleArticleFactory struct {
	procezor.IArticleSpecFactory
}

func NewExampleArticleFactory() procezor.IArticleSpecFactory{
	providersConfig := []procezor.ProviderRecord {
		procezor.NewProviderRecord(ARTICLE_TIMESHT_WORKING.Id(), ARTICLE_DEFAULT_SEQUENS, CONCEPT_TIMESHT_WORKING.Id(),
			[]int32{}),
		procezor.NewProviderRecord(ARTICLE_PAYMENT_SALARY.Id(), ARTICLE_DEFAULT_SEQUENS, CONCEPT_AMOUNT_BASIS.Id(),
			[]int32 {
				ARTICLE_INCOME_GROSS.Id(),
				ARTICLE_HEALTH_INSBASE.Id(),
				ARTICLE_SOCIAL_INSBASE.Id(),
				ARTICLE_TAXING_ADVBASE.Id(),
			}),
		procezor.NewProviderRecord(ARTICLE_PAYMENT_BONUS.Id(), ARTICLE_DEFAULT_SEQUENS, CONCEPT_AMOUNT_FIXED.Id(),
			[]int32 {
				ARTICLE_INCOME_GROSS.Id(),
				ARTICLE_HEALTH_INSBASE.Id(),
				ARTICLE_SOCIAL_INSBASE.Id(),
				ARTICLE_TAXING_ADVBASE.Id(),
			}),
		procezor.NewProviderRecord(ARTICLE_PAYMENT_BARTER.Id(), ARTICLE_DEFAULT_SEQUENS, CONCEPT_AMOUNT_FIXED.Id(),
			[]int32 {
				ARTICLE_HEALTH_INSBASE.Id(),
				ARTICLE_SOCIAL_INSBASE.Id(),
				ARTICLE_TAXING_ADVBASE.Id(),
			}),
		procezor.NewProviderRecord(ARTICLE_ALLOWCE_HOFFICE.Id(), ARTICLE_DEFAULT_SEQUENS, CONCEPT_AMOUNT_FIXED.Id(),
			[]int32 {
				ARTICLE_INCOME_NETTO.Id(),
			}),
		procezor.NewProviderRecord(ARTICLE_HEALTH_INSBASE.Id(), ARTICLE_DEFAULT_SEQUENS, CONCEPT_HEALTH_INSBASE.Id(),
			[]int32{}),
		procezor.NewProviderRecord(ARTICLE_SOCIAL_INSBASE.Id(), ARTICLE_DEFAULT_SEQUENS, CONCEPT_SOCIAL_INSBASE.Id(),
			[]int32{}),
		procezor.NewProviderRecord(ARTICLE_HEALTH_INSPAYM.Id(), ARTICLE_DEFAULT_SEQUENS, CONCEPT_HEALTH_INSPAYM.Id(),
			[]int32{}),
		procezor.NewProviderRecord(ARTICLE_SOCIAL_INSPAYM.Id(), ARTICLE_DEFAULT_SEQUENS, CONCEPT_SOCIAL_INSPAYM.Id(),
			[]int32{}),
		procezor.NewProviderRecord(ARTICLE_TAXING_ADVBASE.Id(), ARTICLE_DEFAULT_SEQUENS, CONCEPT_TAXING_ADVBASE.Id(),
			[]int32{}),
		procezor.NewProviderRecord(ARTICLE_TAXING_ADVPAYM.Id(), ARTICLE_DEFAULT_SEQUENS, CONCEPT_TAXING_ADVPAYM.Id(),
			[]int32{}),
		procezor.NewProviderRecord(ARTICLE_INCOME_GROSS.Id(), ARTICLE_DEFAULT_SEQUENS, CONCEPT_INCOME_GROSS.Id(),
			[]int32{}),
		procezor.NewProviderRecord(ARTICLE_INCOME_NETTO.Id(), ARTICLE_DEFAULT_SEQUENS, CONCEPT_INCOME_NETTO.Id(),
			[]int32{}),
	}
	return ExampleArticleFactory{procezor.NewArticleSpecFactoryWithRecords(providersConfig)}
}

type ExampleConceptFactory struct {
	procezor.IConceptSpecFactory
}

func NewExampleConceptFactory() procezor.IConceptSpecFactory {
	return ExampleConceptFactory{procezor.NewConceptSpecFactoryWithProviders(
		map[int32]procezor.IConceptSpecProvider {
			CONCEPT_TIMESHT_WORKING.Id(): NewTimeshtWorkingConProv(),
			CONCEPT_AMOUNT_BASIS.Id(): NewAmountBasisConProv(),
			CONCEPT_AMOUNT_FIXED.Id(): NewAmountFixedConProv(),
			CONCEPT_HEALTH_INSBASE.Id(): NewHealthInsbaseConProv(),
			CONCEPT_SOCIAL_INSBASE.Id(): NewSocialInsbaseConProv(),
			CONCEPT_HEALTH_INSPAYM.Id(): NewHealthInspaymConProv(),
			CONCEPT_SOCIAL_INSPAYM.Id(): NewSocialInspaymConProv(),
			CONCEPT_TAXING_ADVBASE.Id(): NewTaxingAdvbaseConProv(),
			CONCEPT_TAXING_ADVPAYM.Id(): NewTaxingAdvpaymConProv(),
			CONCEPT_INCOME_GROSS.Id(): NewIncomeGrossConProv(),
			CONCEPT_INCOME_NETTO.Id(): NewIncomeNettoConProv(),
		}),
	}
}

type serviceBuilder struct {
}

func (t serviceBuilder) BuildArticleFactory(s *procezor.ProcezorService) bool {
	s.ArticleFactory = NewExampleArticleFactory()
	if s.ArticleFactory == nil {
		return false
	}
	s.ArticleFactory.BuildFactory()
	return true
}

func (t serviceBuilder) BuildConceptFactory(s *procezor.ProcezorService) bool {
	s.ConceptFactory = NewExampleConceptFactory()
	if s.ConceptFactory == nil {
		return false
	}
	s.ConceptFactory.BuildFactory()
	return true
}

type contractBuilder struct {
}

func (b contractBuilder) GetContractTerms(period procezor.IPeriod, targets procezor.ITermTargetList) procezor.IContractTermList {
	return procezor.IContractTermList{}
}

type positionBuilder struct {
}

func (b positionBuilder) GetPositionTerms(period procezor.IPeriod, contracts procezor.IContractTermList, targets procezor.ITermTargetList) procezor.IPositionTermList {
	return procezor.IPositionTermList{}
}

type ExampleService struct {
	procezor.IProcezorService
}

func NewExampleServiceBuilder() procezor.IProcezorFactoryBuilder {
	return &serviceBuilder{}
}

func NewExampleContractBuilder() procezor.IProcezorContractBuilder {
	return &contractBuilder{}
}

func NewExamplePositionBuilder() procezor.IProcezorPositionBuilder {
	return &positionBuilder{}
}

func NewExampleService() procezor.IProcezorService{
	const (
		TestVersion      = TEST_VERSION
		TestFinalArticle = ARTICLE_INCOME_NETTO
	)
	var (
		TestCalcsArticle = []procezor.ArticleCode{procezor.GetArticleCode(TestFinalArticle.Id())}
	)

	service := ExampleService{
		procezor.NewProcezorService(TestVersion, TestCalcsArticle,
			NewExampleContractBuilder(), NewExamplePositionBuilder(), NewExampleServiceBuilder()),
	}
	buildSuccess := service.BuildFactories()
	if buildSuccess == false {
		println(fmt.Sprintf("Version: %d, build factories failed", service.Version().Value()))
	}
	return &service
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
			procezor.GetArticleCode(ARTICLE_TIMESHT_WORKING.Id()), procezor.GetConceptCode(CONCEPT_TIMESHT_WORKING.Id())),
		types.NewTermTarget(montCode, contract, position, variant1,
			procezor.GetArticleCode(ARTICLE_PAYMENT_SALARY.Id()),procezor.GetConceptCode(CONCEPT_AMOUNT_BASIS.Id())),
		types.NewTermTarget(montCode, contract, position, variant1,
			procezor.GetArticleCode(ARTICLE_ALLOWCE_HOFFICE.Id()),procezor.GetConceptCode(CONCEPT_AMOUNT_FIXED.Id())),
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
			procezor.GetArticleCode(ARTICLE_TIMESHT_WORKING.Id()), procezor.GetConceptCode(CONCEPT_TIMESHT_WORKING.Id())),
		types.NewTermTarget(montCode, contract, position, variant1,
			procezor.GetArticleCode(ARTICLE_PAYMENT_SALARY.Id()), procezor.GetConceptCode(CONCEPT_AMOUNT_BASIS.Id())),
		types.NewTermTarget(montCode, contract, position, variant1,
			procezor.GetArticleCode(ARTICLE_PAYMENT_BONUS.Id()), procezor.GetConceptCode(CONCEPT_AMOUNT_FIXED.Id())),
		types.NewTermTarget(montCode, contract, position, variant1,
			procezor.GetArticleCode(ARTICLE_PAYMENT_BARTER.Id()), procezor.GetConceptCode(CONCEPT_AMOUNT_FIXED.Id())),
	}
	return targets
}

