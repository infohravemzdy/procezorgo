package example

import procezor "github.com/mzdyhrave/procezorgo"

type TimeshtWorkingArtProv struct {
	procezor.ArticleSpecProvider
}

func (p TimeshtWorkingArtProv) GetSpec(period procezor.IPeriod, version procezor.VersionCode) procezor.IArticleSpec {
	return NewTimeshtWorkingArtSpec(p.Code().Value())
}

type TimeshtWorkingArtSpec struct {
	procezor.ArticleSpec
}

func NewTimeshtWorkingArtSpec(code int32) procezor.IArticleSpec {
	const (
		CONCEPT_CODE = CONCEPT_TIMESHT_WORKING
	)
	return &TimeshtWorkingArtSpec{ArticleSpec: procezor.NewArticleSpec(code, CONCEPT_CODE.Id())}
}

func NewTimeshtWorkingArtProv() procezor.IArticleSpecProvider {
	const (
		ARTICLE_CODE = ARTICLE_TIMESHT_WORKING
	)
	return &TimeshtWorkingArtProv{ArticleSpecProvider: procezor.NewArticleProvider(ARTICLE_CODE.Id())}
}

type PaymentSalaryArtProv struct {
	procezor.ArticleSpecProvider
}

func (p PaymentSalaryArtProv) GetSpec(period procezor.IPeriod, version procezor.VersionCode) procezor.IArticleSpec {
	return NewPaymentSalaryArtSpec(p.Code().Value())
}

type PaymentSalaryArtSpec struct {
	procezor.ArticleSpec
}

func NewPaymentSalaryArtSpec(code int32) procezor.IArticleSpec {
	const (
		CONCEPT_CODE = CONCEPT_AMOUNT_BASIS
	)
	_sums := []int32{
		ARTICLE_INCOME_GROSS.Id(),
		ARTICLE_HEALTH_INSBASE.Id(),
		ARTICLE_SOCIAL_INSBASE.Id(),
		ARTICLE_TAXING_ADVBASE.Id(),
	}
	return &PaymentSalaryArtSpec{ArticleSpec: procezor.NewArticleSumIntSpec(code, CONCEPT_CODE.Id(), _sums)}
}

func NewPaymentSalaryArtProv() procezor.IArticleSpecProvider {
	const (
		ARTICLE_CODE = ARTICLE_PAYMENT_SALARY
	)
	return &PaymentSalaryArtProv{ArticleSpecProvider: procezor.NewArticleProvider(ARTICLE_CODE.Id())}
}

type PaymentBonusArtProv struct {
	procezor.ArticleSpecProvider
}

func (p PaymentBonusArtProv) GetSpec(period procezor.IPeriod, version procezor.VersionCode) procezor.IArticleSpec {
	return NewPaymentBonusArtSpec(p.Code().Value())
}

type PaymentBonusArtSpec struct {
	procezor.ArticleSpec
}

func NewPaymentBonusArtSpec(code int32) procezor.IArticleSpec {
	const (
		CONCEPT_CODE = CONCEPT_AMOUNT_FIXED
	)
	_sums := []int32{
		ARTICLE_INCOME_GROSS.Id(),
		ARTICLE_HEALTH_INSBASE.Id(),
		ARTICLE_SOCIAL_INSBASE.Id(),
		ARTICLE_TAXING_ADVBASE.Id(),
	}
	return &PaymentBonusArtSpec{ArticleSpec: procezor.NewArticleSumIntSpec(code, CONCEPT_CODE.Id(), _sums)}
}

func NewPaymentBonusArtProv() procezor.IArticleSpecProvider {
	const (
		ARTICLE_CODE = ARTICLE_PAYMENT_BONUS
	)
	return &PaymentBonusArtProv{ArticleSpecProvider: procezor.NewArticleProvider(ARTICLE_CODE.Id())}
}

type PaymentBarterArtProv struct {
	procezor.ArticleSpecProvider
}

func (p PaymentBarterArtProv) GetSpec(period procezor.IPeriod, version procezor.VersionCode) procezor.IArticleSpec {
	return NewPaymentBarterArtSpec(p.Code().Value())
}

type PaymentBarterArtSpec struct {
	procezor.ArticleSpec
}

func NewPaymentBarterArtSpec(code int32) procezor.IArticleSpec {
	const (
		CONCEPT_CODE = CONCEPT_AMOUNT_FIXED
	)
	_sums := []int32{
		ARTICLE_HEALTH_INSBASE.Id(),
		ARTICLE_SOCIAL_INSBASE.Id(),
		ARTICLE_TAXING_ADVBASE.Id(),
	}
	return &PaymentBarterArtSpec{ArticleSpec: procezor.NewArticleSumIntSpec(code, CONCEPT_CODE.Id(), _sums)}
}

func NewPaymentBarterArtProv() procezor.IArticleSpecProvider {
	const (
		ARTICLE_CODE = ARTICLE_PAYMENT_BARTER
	)
	return &PaymentBarterArtProv{ArticleSpecProvider: procezor.NewArticleProvider(ARTICLE_CODE.Id())}
}

type AllowceHofficeArtProv struct {
	procezor.ArticleSpecProvider
}

func (p AllowceHofficeArtProv) GetSpec(period procezor.IPeriod, version procezor.VersionCode) procezor.IArticleSpec {
	return NewAllowceHofficeArtSpec(p.Code().Value())
}

type AllowceHofficeArtSpec struct {
	procezor.ArticleSpec
}

func NewAllowceHofficeArtSpec(code int32) procezor.IArticleSpec {
	const (
		CONCEPT_CODE = CONCEPT_AMOUNT_FIXED
	)
	_sums := []int32{
		ARTICLE_INCOME_NETTO.Id(),
	}
	return &AllowceHofficeArtSpec{ArticleSpec: procezor.NewArticleSumIntSpec(code, CONCEPT_CODE.Id(), _sums)}
}

func NewAllowceHofficeArtProv() procezor.IArticleSpecProvider {
	const (
		ARTICLE_CODE = ARTICLE_ALLOWCE_HOFFICE
	)
	return &AllowceHofficeArtProv{ArticleSpecProvider: procezor.NewArticleProvider(ARTICLE_CODE.Id())}
}

type HealthInsbaseArtProv struct {
	procezor.ArticleSpecProvider
}

func (p HealthInsbaseArtProv) GetSpec(period procezor.IPeriod, version procezor.VersionCode) procezor.IArticleSpec {
	return NewHealthInsbaseArtSpec(p.Code().Value())
}

type HealthInsbaseArtSpec struct {
	procezor.ArticleSpec
}

func NewHealthInsbaseArtSpec(code int32) procezor.IArticleSpec {
	const (
		CONCEPT_CODE = CONCEPT_HEALTH_INSBASE
	)
	return &HealthInsbaseArtSpec{ArticleSpec: procezor.NewArticleSpec(code, CONCEPT_CODE.Id())}
}

func NewHealthInsbaseArtProv() procezor.IArticleSpecProvider {
	const (
		ARTICLE_CODE = ARTICLE_HEALTH_INSBASE
	)
	return &HealthInsbaseArtProv{ArticleSpecProvider: procezor.NewArticleProvider(ARTICLE_CODE.Id())}
}

type SocialInsbaseArtProv struct {
	procezor.ArticleSpecProvider
}

func (p SocialInsbaseArtProv) GetSpec(period procezor.IPeriod, version procezor.VersionCode) procezor.IArticleSpec {
	return NewSocialInsbaseArtSpec(p.Code().Value())
}

type SocialInsbaseArtSpec struct {
	procezor.ArticleSpec
}

func NewSocialInsbaseArtSpec(code int32) procezor.IArticleSpec {
	const (
		CONCEPT_CODE = CONCEPT_SOCIAL_INSBASE
	)
	return &SocialInsbaseArtSpec{ArticleSpec: procezor.NewArticleSpec(code, CONCEPT_CODE.Id())}
}

func NewSocialInsbaseArtProv() procezor.IArticleSpecProvider {
	const (
		ARTICLE_CODE = ARTICLE_SOCIAL_INSBASE
	)
	return &SocialInsbaseArtProv{ArticleSpecProvider: procezor.NewArticleProvider(ARTICLE_CODE.Id())}
}

type TaxingAdvbaseArtProv struct {
	procezor.ArticleSpecProvider
}

func (p TaxingAdvbaseArtProv) GetSpec(period procezor.IPeriod, version procezor.VersionCode) procezor.IArticleSpec {
	return NewTaxingAdvbaseArtSpec(p.Code().Value())
}

type TaxingAdvbaseArtSpec struct {
	procezor.ArticleSpec
}

func NewTaxingAdvbaseArtSpec(code int32) procezor.IArticleSpec {
	const (
		CONCEPT_CODE = CONCEPT_TAXING_ADVBASE
	)
	return &TaxingAdvbaseArtSpec{ArticleSpec: procezor.NewArticleSpec(code, CONCEPT_CODE.Id())}
}

func NewTaxingAdvbaseArtProv() procezor.IArticleSpecProvider {
	const (
		ARTICLE_CODE = ARTICLE_TAXING_ADVBASE
	)
	return &TaxingAdvbaseArtProv{ArticleSpecProvider: procezor.NewArticleProvider(ARTICLE_CODE.Id())}
}

type HealthInspaymArtProv struct {
	procezor.ArticleSpecProvider
}

func (p HealthInspaymArtProv) GetSpec(period procezor.IPeriod, version procezor.VersionCode) procezor.IArticleSpec {
	return NewHealthInspaymArtSpec(p.Code().Value())
}

type HealthInspaymArtSpec struct {
	procezor.ArticleSpec
}

func NewHealthInspaymArtSpec(code int32) procezor.IArticleSpec {
	const (
		CONCEPT_CODE = CONCEPT_HEALTH_INSPAYM
	)
	return &HealthInspaymArtSpec{ArticleSpec: procezor.NewArticleSpec(code, CONCEPT_CODE.Id())}
}

func NewHealthInspaymArtProv() procezor.IArticleSpecProvider {
	const (
		ARTICLE_CODE = ARTICLE_HEALTH_INSPAYM
	)
	return &HealthInspaymArtProv{ArticleSpecProvider: procezor.NewArticleProvider(ARTICLE_CODE.Id())}
}

type SocialInspaymArtProv struct {
	procezor.ArticleSpecProvider
}

func (p SocialInspaymArtProv) GetSpec(period procezor.IPeriod, version procezor.VersionCode) procezor.IArticleSpec {
	return NewSocialInspaymArtSpec(p.Code().Value())
}

type SocialInspaymArtSpec struct {
	procezor.ArticleSpec
}

func NewSocialInspaymArtSpec(code int32) procezor.IArticleSpec {
	const (
		CONCEPT_CODE = CONCEPT_SOCIAL_INSPAYM
	)
	return &SocialInspaymArtSpec{ArticleSpec: procezor.NewArticleSpec(code, CONCEPT_CODE.Id())}
}

func NewSocialInspaymArtProv() procezor.IArticleSpecProvider {
	const (
		ARTICLE_CODE = ARTICLE_SOCIAL_INSPAYM
	)
	return &SocialInspaymArtProv{ArticleSpecProvider: procezor.NewArticleProvider(ARTICLE_CODE.Id())}
}

type TaxingAdvpaymArtProv struct {
	procezor.ArticleSpecProvider
}

func (p TaxingAdvpaymArtProv) GetSpec(period procezor.IPeriod, version procezor.VersionCode) procezor.IArticleSpec {
	return NewTaxingAdvpaymArtSpec(p.Code().Value())
}

type TaxingAdvpaymArtSpec struct {
	procezor.ArticleSpec
}

func NewTaxingAdvpaymArtSpec(code int32) procezor.IArticleSpec {
	const (
		CONCEPT_CODE = CONCEPT_TAXING_ADVPAYM
	)
	return &TaxingAdvpaymArtSpec{ArticleSpec: procezor.NewArticleSpec(code, CONCEPT_CODE.Id())}
}

func NewTaxingAdvpaymArtProv() procezor.IArticleSpecProvider {
	const (
		ARTICLE_CODE = ARTICLE_TAXING_ADVPAYM
	)
	return &TaxingAdvpaymArtProv{ArticleSpecProvider: procezor.NewArticleProvider(ARTICLE_CODE.Id())}
}

type IncomeGrossArtProv struct {
	procezor.ArticleSpecProvider
}

func (p IncomeGrossArtProv) GetSpec(period procezor.IPeriod, version procezor.VersionCode) procezor.IArticleSpec {
	return NewIncomeGrossArtSpec(p.Code().Value())
}

type IncomeGrossArtSpec struct {
	procezor.ArticleSpec
}

func NewIncomeGrossArtSpec(code int32) procezor.IArticleSpec {
	const (
		CONCEPT_CODE = CONCEPT_INCOME_GROSS
	)
	return &IncomeGrossArtSpec{ArticleSpec: procezor.NewArticleSpec(code, CONCEPT_CODE.Id())}
}

func NewIncomeGrossArtProv() procezor.IArticleSpecProvider {
	const (
		ARTICLE_CODE = ARTICLE_INCOME_GROSS
	)
	return &IncomeGrossArtProv{ArticleSpecProvider: procezor.NewArticleProvider(ARTICLE_CODE.Id())}
}

type IncomeNettoArtProv struct {
	procezor.ArticleSpecProvider
}

func (p IncomeNettoArtProv) GetSpec(period procezor.IPeriod, version procezor.VersionCode) procezor.IArticleSpec {
	return NewIncomeNettoArtSpec(p.Code().Value())
}

type IncomeNettoArtSpec struct {
	procezor.ArticleSpec
}

func NewIncomeNettoArtSpec(code int32) procezor.IArticleSpec {
	const (
		CONCEPT_CODE = CONCEPT_INCOME_NETTO
	)
	return &IncomeNettoArtSpec{ArticleSpec: procezor.NewArticleSpec(code, CONCEPT_CODE.Id())}
}

func NewIncomeNettoArtProv() procezor.IArticleSpecProvider {
	const (
		ARTICLE_CODE = ARTICLE_INCOME_NETTO
	)
	return &IncomeNettoArtProv{ArticleSpecProvider: procezor.NewArticleProvider(ARTICLE_CODE.Id())}
}
