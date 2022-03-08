package procezorgo

import (
	legalios "github.com/mzdyhrave/legaliosgo"
	"github.com/mzdyhrave/procezorgo/internal/registry"
	factories "github.com/mzdyhrave/procezorgo/internal/registry_factories"
	providers "github.com/mzdyhrave/procezorgo/internal/registry_providers"
	"github.com/mzdyhrave/procezorgo/internal/types"
	"time"
)

type IPeriod = legalios.IPeriod
type IVersionId = legalios.IVersionId

func NewVersionId() IVersionId {
	return legalios.NewVersionId()
}

func GetVersionId(id int16) IVersionId {
	return legalios.GetVersionId(id)
}

func PeriodZero() IPeriod {
	return legalios.PeriodZero()
}

func NewPeriodWithCode(code int32) IPeriod {
	return legalios.NewPeriodWithCode(code)
}

func NewPeriodWithYearMonth(year int16, month int16) IPeriod {
	return legalios.NewPeriodWithYearMonth(year, month)
}

type VersionCode = types.VersionCode

type ArticleSeqs = types.ArticleSeqs

type MonthCode = types.MonthCode

type ContractCode = types.ContractCode

type PositionCode = types.PositionCode

type VariantCode = types.VariantCode

func ArticleSeqsZero() ArticleSeqs {
	return types.ArticleSeqsZero()
}

func NewArticleSeqs() ArticleSeqs {
	return types.NewArticleSeqs()
}

func GetArticleSeqs(id int16) ArticleSeqs {
	return types.GetArticleSeqs(id)
}

func CloneArticleSeqs(seqs ArticleSeqs) ArticleSeqs {
	return types.CloneArticleSeqs(seqs)
}

func GetMonthCode(id int32) MonthCode {
	return types.GetMonthCode(id)
}

func GetPeriodMonthCode(period IPeriod) MonthCode {
	return types.GetMonthCode(period.GetCode())
}

func GetContractCode(id int16) ContractCode {
	return types.GetContractCode(id)
}

func GetPositionCode(id int16) PositionCode {
	return types.GetPositionCode(id)
}

func GetVariantCode(id int16) VariantCode {
	return types.GetVariantCode(id)
}

func NewMonthCode() MonthCode {
	return types.NewMonthCode()
}

func NewPeriodMonthCode() MonthCode {
	return types.NewMonthCode()
}

func NewContractCode() ContractCode {
	return types.NewContractCode()
}

func NewPositionCode() PositionCode {
	return types.NewPositionCode()
}

func NewVariantCode() VariantCode {
	return types.NewVariantCode()
}

type ArticleCode = types.ArticleCode
type ArticleCodeList = types.ArticleCodeList

func GetArticleCode(id int32) ArticleCode {
	return types.GetArticleCode(id)
}

func NewArticleCode() ArticleCode {
	return types.NewArticleCode()
}

type ArticleTerm = types.ArticleTerm
type ArticleTermList = types.ArticleTermList

func ArticleTermZero() ArticleTerm {
	return types.ArticleTermZero()
}

func NewArticleTerm() ArticleTerm {
	return types.NewArticleTerm()
}

func GetArticleTerm(code int32, seqs int16) ArticleTerm {
	return types.GetArticleTerm(code, seqs)
}

func CloneArticleTerm(term ArticleTerm) ArticleTerm {
	return types.CloneArticleTerm(term)
}

type ConceptCode = types.ConceptCode

func GetConceptCode(id int32) ConceptCode {
	return types.GetConceptCode(id)
}

func NewConceptCode() ConceptCode {
	return types.NewConceptCode()
}

type IContractTerm = types.IContractTerm
type ContractTerm = types.ContractTerm

func NewContractTerm(_contract ContractCode,
	_dateFrom time.Time, _dateStop time.Time,
	_termDayFrom byte, _termDayStop byte) IContractTerm {
	return types.NewContractTerm(_contract, _dateFrom, _dateStop, _termDayFrom, _termDayStop)
}

type IContractTermList = types.IContractTermList

type IPositionTerm = types.IPositionTerm
type PositionTerm = types.PositionTerm

func NewPositionTerm(_contract ContractCode, _position PositionCode, _baseTerm IContractTerm,
	_dateFrom time.Time, _dateStop time.Time,
	_termDayFrom byte, _termDayStop byte) IPositionTerm {
	return NewPositionTerm(_contract, _position, _baseTerm,
		_dateFrom, _dateStop, _termDayFrom, _termDayStop)
}

type IPositionTermList = types.IPositionTermList

type IArticleSpec = types.IArticleSpec
type IArticleCodeProvider = providers.IArticleCodeProvider
type IArticleSpecProvider = providers.IArticleSpecProvider
type ArticleSpec = types.ArticleSpec
type ArticleSpecProvider = providers.ArticleSpecProvider

func NewArticleProvider(code int32) ArticleSpecProvider {
	return providers.NewArticleProvider(code)
}

func NewArticleSpec(code int32, seqs int16, role int32) ArticleSpec {
	return types.NewArticleSpec(code, seqs, role)
}

func NewArticleSumSpec(code int32, seqs int16, role int32, sums []types.ArticleCode) ArticleSpec {
	return types.NewArticleSumSpec(code, seqs, role, sums)
}

func NewArticleSumIntSpec(code int32, seqs int16, role int32, sums []int32) ArticleSpec {
	return types.NewArticleSumIntSpec(code, seqs, role, sums)
}

func NewArticleCodeProvider(code int32) IArticleCodeProvider {
	return providers.NewArticleCodeProvider(code)
}

type IConceptSpec = providers.IConceptSpec
type IConceptCodeProvider = providers.IConceptCodeProvider
type IConceptSpecProvider = providers.IConceptSpecProvider
type ConceptSpec = providers.ConceptSpec
type ConceptSpecProvider = providers.ConceptSpecProvider
type ResultFunc = providers.ResultFunc

func NewConceptProvider(code int32) ConceptSpecProvider {
	return providers.NewConceptProvider(code)
}

func NewConceptSpec(code int32) ConceptSpec {
	return providers.NewConceptSpec(code)
}

func NewConceptPathSpec(code int32, path []types.ArticleCode) ConceptSpec {
	return providers.NewConceptPathSpec(code, path)
}

func NewConceptPathIntSpec(code int32, path []int32) ConceptSpec {
	return providers.NewConceptPathIntSpec(code, path)
}

func NewConceptFuncSpec(code int32, resultFunc providers.ResultFunc) ConceptSpec {
	return providers.NewConceptFuncSpec(code, resultFunc)
}

func NewConceptPathFuncSpec(code int32, path []types.ArticleCode, resultFunc providers.ResultFunc) ConceptSpec {
	return providers.NewConceptPathFuncSpec(code, path, resultFunc)
}

func NewConceptPathIntFuncSpec(code int32, path []int32, resultFunc providers.ResultFunc) ConceptSpec {
	return providers.NewConceptPathIntFuncSpec(code, path, resultFunc)
}

func NewConceptCodeProvider(code int32) IConceptCodeProvider {
	return providers.NewConceptCodeProvider(code)
}

type IArticleDefine = types.IArticleDefine
type IArticleDefineList = types.IArticleDefineList

func GetArticleDefine(code int32, seqs int16, role int32) IArticleDefine {
	return types.GetArticleDefine(code, seqs, role)
}

type ProviderRecord = factories.ProviderRecord

func NewProviderRecord(_article int32, _sequens int16, _concept int32, _sums []int32) factories.ProviderRecord {
	return factories.NewProviderRecord(_article, _sequens, _concept, _sums)
}

type IArticleSpecFactory = factories.IArticleSpecFactory

func NewArticleSpecFactoryWithProviders(providersMap map[int32]IArticleSpecProvider) IArticleSpecFactory {
	return &factories.ArticleSpecFactory{ BuildProvidersFunc: func (f *factories.ArticleSpecFactory) bool {
			f.NotFoundSpec = factories.NotFoundArticleSpec()
			f.NotFoundProvider = factories.NewNotFoundArticleProvider()
			f.Providers = providersMap
			return true
		},
	}
}

func NewArticleSpecFactoryWithRecords(records []factories.ProviderRecord) IArticleSpecFactory {
	return &factories.ArticleSpecFactory{ BuildProvidersFunc: func (f *factories.ArticleSpecFactory) bool {
		f.NotFoundSpec = factories.NotFoundArticleSpec()
		f.NotFoundProvider = factories.NewNotFoundArticleProvider()
		f.Providers = make(map[int32]IArticleSpecProvider)
		for _, r := range records {
			f.Providers[r.Article] = factories.NewArticleProviderConfig(r.Article, r.Sequens, r.Concept, r.Sums)
		}
		return true
	},
	}
}

type IConceptSpecFactory = factories.IConceptSpecFactory

func NewConceptSpecFactoryWithProviders(providersMap map[int32]IConceptSpecProvider) IConceptSpecFactory {
	return &factories.ConceptSpecFactory{ BuildProvidersFunc: func (f *factories.ConceptSpecFactory) bool {
			f.NotFoundSpec = factories.NotFoundConceptSpec()
			f.NotFoundProvider = factories.NewNotFoundConceptProvider()
			f.Providers = providersMap
			return true
		},
	}
}

type IBundleProps = legalios.IBundleProps

type ITermTarget = types.ITermTarget

type ITermTargetList = types.ITermTargetList

type TermTarget = types.TermTarget

func NewTermTarget(mont MonthCode, cont ContractCode, post PositionCode, vars VariantCode, code ArticleCode, role ConceptCode) TermTarget {
	return types.NewTermTarget(mont, cont, post, vars, code, role)
}

type ITermResult = types.ITermResult

type TermResult = types.TermResult

type ITermResultError = types.ITermResultError

type IBuilderResult = providers.IBuilderResult

type IBuilderResultList = providers.IBuilderResultList

func NewTermResult(target ITermTarget, spec IArticleSpec) TermResult {
	return types.NewTermResult(target, spec)
}

func NewFailureResult(error ITermResultError) types.FailureResult {
	return types.NewFailureResult(error)
}

type OrderCodeList = registry.OrderCodeList
type PathsTermsMap = registry.PathsTermsMap

func NewSuccessResult(result ITermResult, period legalios.IPeriod) types.SuccessResult {
	return types.NewSuccessResult(result, period)
}
