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
		CONCEPT_CODE int32 = CONCEPT_TIMESHT_WORKING
	)
	return &TimeshtWorkingArtSpec{ArticleSpec: procezor.NewArticleSpec(code, CONCEPT_CODE)}
}

func NewTimeshtWorkingArtProv() procezor.IArticleSpecProvider {
	const (
		ARTICLE_CODE int32 = ARTICLE_TIMESHT_WORKING
	)
	return &TimeshtWorkingArtProv{ArticleSpecProvider: procezor.NewArticleProvider(ARTICLE_CODE)}
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
		CONCEPT_CODE int32 = CONCEPT_AMOUNT_BASIS
	)
	return &PaymentSalaryArtSpec{ArticleSpec: procezor.NewArticleSumSpec(code, CONCEPT_CODE,
		procezor.ArticleCodeList{
			procezor.GetArticleCode(ARTICLE_INCOME_GROSS),
			procezor.GetArticleCode(ARTICLE_HEALTH_INSBASE),
			procezor.GetArticleCode(ARTICLE_SOCIAL_INSBASE),
			procezor.GetArticleCode(ARTICLE_TAXING_ADVBASE),
		})}
}

func NewPaymentSalaryArtProv() procezor.IArticleSpecProvider {
	const (
		ARTICLE_CODE int32 = ARTICLE_PAYMENT_SALARY
	)
	return &PaymentSalaryArtProv{ArticleSpecProvider: procezor.NewArticleProvider(ARTICLE_CODE)}
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
		CONCEPT_CODE int32 = CONCEPT_AMOUNT_FIXED
	)
	return &PaymentBonusArtSpec{ArticleSpec: procezor.NewArticleSumSpec(code, CONCEPT_CODE,
		procezor.ArticleCodeList{
			procezor.GetArticleCode(ARTICLE_INCOME_GROSS),
			procezor.GetArticleCode(ARTICLE_HEALTH_INSBASE),
			procezor.GetArticleCode(ARTICLE_SOCIAL_INSBASE),
			procezor.GetArticleCode(ARTICLE_TAXING_ADVBASE),
		})}
}

func NewPaymentBonusArtProv() procezor.IArticleSpecProvider {
	const (
		ARTICLE_CODE int32 = ARTICLE_PAYMENT_BONUS
	)
	return &PaymentBonusArtProv{ArticleSpecProvider: procezor.NewArticleProvider(ARTICLE_CODE)}
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
		CONCEPT_CODE int32 = CONCEPT_AMOUNT_FIXED
	)
	return &PaymentBarterArtSpec{ArticleSpec: procezor.NewArticleSumSpec(code, CONCEPT_CODE,
		procezor.ArticleCodeList{
			procezor.GetArticleCode(ARTICLE_HEALTH_INSBASE),
			procezor.GetArticleCode(ARTICLE_SOCIAL_INSBASE),
			procezor.GetArticleCode(ARTICLE_TAXING_ADVBASE),
		})}
}

func NewPaymentBarterArtProv() procezor.IArticleSpecProvider {
	const (
		ARTICLE_CODE int32 = ARTICLE_PAYMENT_BARTER
	)
	return &PaymentBarterArtProv{ArticleSpecProvider: procezor.NewArticleProvider(ARTICLE_CODE)}
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
		CONCEPT_CODE int32 = CONCEPT_AMOUNT_FIXED
	)
	return &AllowceHofficeArtSpec{ArticleSpec: procezor.NewArticleSumSpec(code, CONCEPT_CODE,
		procezor.ArticleCodeList{
			procezor.GetArticleCode(ARTICLE_INCOME_NETTO),
		})}
}

func NewAllowceHofficeArtProv() procezor.IArticleSpecProvider {
	const (
		ARTICLE_CODE int32 = ARTICLE_ALLOWCE_HOFFICE
	)
	return &AllowceHofficeArtProv{ArticleSpecProvider: procezor.NewArticleProvider(ARTICLE_CODE)}
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
		CONCEPT_CODE int32 = CONCEPT_HEALTH_INSBASE
	)
	return &HealthInsbaseArtSpec{ArticleSpec: procezor.NewArticleSpec(code, CONCEPT_CODE)}
}

func NewHealthInsbaseArtProv() procezor.IArticleSpecProvider {
	const (
		ARTICLE_CODE int32 = ARTICLE_HEALTH_INSBASE
	)
	return &HealthInsbaseArtProv{ArticleSpecProvider: procezor.NewArticleProvider(ARTICLE_CODE)}
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
		CONCEPT_CODE int32 = CONCEPT_SOCIAL_INSBASE
	)
	return &SocialInsbaseArtSpec{ArticleSpec: procezor.NewArticleSpec(code, CONCEPT_CODE)}
}

func NewSocialInsbaseArtProv() procezor.IArticleSpecProvider {
	const (
		ARTICLE_CODE int32 = ARTICLE_SOCIAL_INSBASE
	)
	return &SocialInsbaseArtProv{ArticleSpecProvider: procezor.NewArticleProvider(ARTICLE_CODE)}
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
		CONCEPT_CODE int32 = CONCEPT_TAXING_ADVBASE
	)
	return &TaxingAdvbaseArtSpec{ArticleSpec: procezor.NewArticleSpec(code, CONCEPT_CODE)}
}

func NewTaxingAdvbaseArtProv() procezor.IArticleSpecProvider {
	const (
		ARTICLE_CODE int32 = ARTICLE_TAXING_ADVBASE
	)
	return &TaxingAdvbaseArtProv{ArticleSpecProvider: procezor.NewArticleProvider(ARTICLE_CODE)}
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
		CONCEPT_CODE int32 = CONCEPT_HEALTH_INSPAYM
	)
	return &HealthInspaymArtSpec{ArticleSpec: procezor.NewArticleSpec(code, CONCEPT_CODE)}
}

func NewHealthInspaymArtProv() procezor.IArticleSpecProvider {
	const (
		ARTICLE_CODE int32 = ARTICLE_HEALTH_INSPAYM
	)
	return &HealthInspaymArtProv{ArticleSpecProvider: procezor.NewArticleProvider(ARTICLE_CODE)}
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
		CONCEPT_CODE int32 = CONCEPT_SOCIAL_INSPAYM
	)
	return &SocialInspaymArtSpec{ArticleSpec: procezor.NewArticleSpec(code, CONCEPT_CODE)}
}

func NewSocialInspaymArtProv() procezor.IArticleSpecProvider {
	const (
		ARTICLE_CODE int32 = ARTICLE_SOCIAL_INSPAYM
	)
	return &SocialInspaymArtProv{ArticleSpecProvider: procezor.NewArticleProvider(ARTICLE_CODE)}
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
		CONCEPT_CODE int32 = CONCEPT_TAXING_ADVPAYM
	)
	return &TaxingAdvpaymArtSpec{ArticleSpec: procezor.NewArticleSpec(code, CONCEPT_CODE)}
}

func NewTaxingAdvpaymArtProv() procezor.IArticleSpecProvider {
	const (
		ARTICLE_CODE int32 = ARTICLE_TAXING_ADVPAYM
	)
	return &TaxingAdvpaymArtProv{ArticleSpecProvider: procezor.NewArticleProvider(ARTICLE_CODE)}
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
		CONCEPT_CODE int32 = CONCEPT_INCOME_GROSS
	)
	return &IncomeGrossArtSpec{ArticleSpec: procezor.NewArticleSpec(code, CONCEPT_CODE)}
}

func NewIncomeGrossArtProv() procezor.IArticleSpecProvider {
	const (
		ARTICLE_CODE int32 = ARTICLE_INCOME_GROSS
	)
	return &IncomeGrossArtProv{ArticleSpecProvider: procezor.NewArticleProvider(ARTICLE_CODE)}
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
		CONCEPT_CODE int32 = CONCEPT_INCOME_NETTO
	)
	return &IncomeNettoArtSpec{ArticleSpec: procezor.NewArticleSpec(code, CONCEPT_CODE)}
}

func NewIncomeNettoArtProv() procezor.IArticleSpecProvider {
	const (
		ARTICLE_CODE int32 = ARTICLE_INCOME_NETTO
	)
	return &IncomeNettoArtProv{ArticleSpecProvider: procezor.NewArticleProvider(ARTICLE_CODE)}
}