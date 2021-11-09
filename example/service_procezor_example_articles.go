package example

import procezor "github.com/mzdyhrave/procezorgo"

type TimeshtWorkingArtProv struct {
	procezor.ArticleSpecProvider
}

func (p TimeshtWorkingArtProv) GetSpec(period procezor.IPeriod, version procezor.VersionCode) procezor.IArticleSpec {
	return NewTimeshtWorkingArtSpec(p.Code().Value())
}

func NewTimeshtWorkingArtProv() procezor.IArticleSpecProvider {
	ArticleCode := ARTICLE_TIMESHT_WORKING.Id()

	return &TimeshtWorkingArtProv{ArticleSpecProvider: procezor.NewArticleProvider(ArticleCode)}
}

type TimeshtWorkingArtSpec struct {
	procezor.ArticleSpec
}

func NewTimeshtWorkingArtSpec(code int32) procezor.IArticleSpec {
	ConceptCode := CONCEPT_TIMESHT_WORKING.Id()

	return &TimeshtWorkingArtSpec{ArticleSpec: procezor.NewArticleSpec(code, ConceptCode)}
}

type PaymentSalaryArtProv struct {
	procezor.ArticleSpecProvider
}

func (p PaymentSalaryArtProv) GetSpec(period procezor.IPeriod, version procezor.VersionCode) procezor.IArticleSpec {
	return NewPaymentSalaryArtSpec(p.Code().Value())
}

func NewPaymentSalaryArtProv() procezor.IArticleSpecProvider {
	ArticleCode := ARTICLE_PAYMENT_SALARY.Id()

	return &PaymentSalaryArtProv{ArticleSpecProvider: procezor.NewArticleProvider(ArticleCode)}
}

type PaymentSalaryArtSpec struct {
	procezor.ArticleSpec
}

func NewPaymentSalaryArtSpec(code int32) procezor.IArticleSpec {
	ConceptCode := CONCEPT_AMOUNT_BASIS.Id()

	_sums := []int32{
		ARTICLE_INCOME_GROSS.Id(),
		ARTICLE_HEALTH_INSBASE.Id(),
		ARTICLE_SOCIAL_INSBASE.Id(),
		ARTICLE_TAXING_ADVBASE.Id(),
	}
	return &PaymentSalaryArtSpec{ArticleSpec: procezor.NewArticleSumIntSpec(code, ConceptCode, _sums)}
}

type PaymentBonusArtProv struct {
	procezor.ArticleSpecProvider
}

func (p PaymentBonusArtProv) GetSpec(period procezor.IPeriod, version procezor.VersionCode) procezor.IArticleSpec {
	return NewPaymentBonusArtSpec(p.Code().Value())
}

func NewPaymentBonusArtProv() procezor.IArticleSpecProvider {
	ArticleCode := ARTICLE_PAYMENT_BONUS.Id()

	return &PaymentBonusArtProv{ArticleSpecProvider: procezor.NewArticleProvider(ArticleCode)}
}

type PaymentBonusArtSpec struct {
	procezor.ArticleSpec
}

func NewPaymentBonusArtSpec(code int32) procezor.IArticleSpec {
	ConceptCode := CONCEPT_AMOUNT_FIXED.Id()

	_sums := []int32{
		ARTICLE_INCOME_GROSS.Id(),
		ARTICLE_HEALTH_INSBASE.Id(),
		ARTICLE_SOCIAL_INSBASE.Id(),
		ARTICLE_TAXING_ADVBASE.Id(),
	}
	return &PaymentBonusArtSpec{ArticleSpec: procezor.NewArticleSumIntSpec(code, ConceptCode, _sums)}
}

type PaymentBarterArtProv struct {
	procezor.ArticleSpecProvider
}

func (p PaymentBarterArtProv) GetSpec(period procezor.IPeriod, version procezor.VersionCode) procezor.IArticleSpec {
	return NewPaymentBarterArtSpec(p.Code().Value())
}

func NewPaymentBarterArtProv() procezor.IArticleSpecProvider {
	ArticleCode := ARTICLE_PAYMENT_BARTER.Id()

	return &PaymentBarterArtProv{ArticleSpecProvider: procezor.NewArticleProvider(ArticleCode)}
}

type PaymentBarterArtSpec struct {
	procezor.ArticleSpec
}

func NewPaymentBarterArtSpec(code int32) procezor.IArticleSpec {
	ConceptCode := CONCEPT_AMOUNT_FIXED.Id()

	_sums := []int32{
		ARTICLE_HEALTH_INSBASE.Id(),
		ARTICLE_SOCIAL_INSBASE.Id(),
		ARTICLE_TAXING_ADVBASE.Id(),
	}
	return &PaymentBarterArtSpec{ArticleSpec: procezor.NewArticleSumIntSpec(code, ConceptCode, _sums)}
}

type AllowceHofficeArtProv struct {
	procezor.ArticleSpecProvider
}

func (p AllowceHofficeArtProv) GetSpec(period procezor.IPeriod, version procezor.VersionCode) procezor.IArticleSpec {
	return NewAllowceHofficeArtSpec(p.Code().Value())
}

func NewAllowceHofficeArtProv() procezor.IArticleSpecProvider {
	ArticleCode := ARTICLE_ALLOWCE_HOFFICE.Id()

	return &AllowceHofficeArtProv{ArticleSpecProvider: procezor.NewArticleProvider(ArticleCode)}
}

type AllowceHofficeArtSpec struct {
	procezor.ArticleSpec
}

func NewAllowceHofficeArtSpec(code int32) procezor.IArticleSpec {
	ConceptCode := CONCEPT_AMOUNT_FIXED.Id()

	_sums := []int32{
		ARTICLE_INCOME_NETTO.Id(),
	}
	return &AllowceHofficeArtSpec{ArticleSpec: procezor.NewArticleSumIntSpec(code, ConceptCode, _sums)}
}

type HealthInsbaseArtProv struct {
	procezor.ArticleSpecProvider
}

func (p HealthInsbaseArtProv) GetSpec(period procezor.IPeriod, version procezor.VersionCode) procezor.IArticleSpec {
	return NewHealthInsbaseArtSpec(p.Code().Value())
}

func NewHealthInsbaseArtProv() procezor.IArticleSpecProvider {
	ArticleCode := ARTICLE_HEALTH_INSBASE.Id()

	return &HealthInsbaseArtProv{ArticleSpecProvider: procezor.NewArticleProvider(ArticleCode)}
}

type HealthInsbaseArtSpec struct {
	procezor.ArticleSpec
}

func NewHealthInsbaseArtSpec(code int32) procezor.IArticleSpec {
	ConceptCode := CONCEPT_HEALTH_INSBASE.Id()

	return &HealthInsbaseArtSpec{ArticleSpec: procezor.NewArticleSpec(code, ConceptCode)}
}

type SocialInsbaseArtProv struct {
	procezor.ArticleSpecProvider
}

func (p SocialInsbaseArtProv) GetSpec(period procezor.IPeriod, version procezor.VersionCode) procezor.IArticleSpec {
	return NewSocialInsbaseArtSpec(p.Code().Value())
}

func NewSocialInsbaseArtProv() procezor.IArticleSpecProvider {
	ArticleCode := ARTICLE_SOCIAL_INSBASE.Id()

	return &SocialInsbaseArtProv{ArticleSpecProvider: procezor.NewArticleProvider(ArticleCode)}
}

type SocialInsbaseArtSpec struct {
	procezor.ArticleSpec
}

func NewSocialInsbaseArtSpec(code int32) procezor.IArticleSpec {
	ConceptCode := CONCEPT_SOCIAL_INSBASE.Id()

	return &SocialInsbaseArtSpec{ArticleSpec: procezor.NewArticleSpec(code, ConceptCode)}
}

type HealthInspaymArtProv struct {
	procezor.ArticleSpecProvider
}

func (p HealthInspaymArtProv) GetSpec(period procezor.IPeriod, version procezor.VersionCode) procezor.IArticleSpec {
	return NewHealthInspaymArtSpec(p.Code().Value())
}

func NewHealthInspaymArtProv() procezor.IArticleSpecProvider {
	ArticleCode := ARTICLE_HEALTH_INSPAYM.Id()

	return &HealthInspaymArtProv{ArticleSpecProvider: procezor.NewArticleProvider(ArticleCode)}
}

type HealthInspaymArtSpec struct {
	procezor.ArticleSpec
}

func NewHealthInspaymArtSpec(code int32) procezor.IArticleSpec {
	ConceptCode := CONCEPT_HEALTH_INSPAYM.Id()

	return &HealthInspaymArtSpec{ArticleSpec: procezor.NewArticleSpec(code, ConceptCode)}
}

type SocialInspaymArtProv struct {
	procezor.ArticleSpecProvider
}

func (p SocialInspaymArtProv) GetSpec(period procezor.IPeriod, version procezor.VersionCode) procezor.IArticleSpec {
	return NewSocialInspaymArtSpec(p.Code().Value())
}

func NewSocialInspaymArtProv() procezor.IArticleSpecProvider {
	ArticleCode := ARTICLE_SOCIAL_INSPAYM.Id()

	return &SocialInspaymArtProv{ArticleSpecProvider: procezor.NewArticleProvider(ArticleCode)}
}

type SocialInspaymArtSpec struct {
	procezor.ArticleSpec
}

func NewSocialInspaymArtSpec(code int32) procezor.IArticleSpec {
	ConceptCode := CONCEPT_SOCIAL_INSPAYM.Id()

	return &SocialInspaymArtSpec{ArticleSpec: procezor.NewArticleSpec(code, ConceptCode)}
}

type TaxingAdvbaseArtProv struct {
	procezor.ArticleSpecProvider
}

func (p TaxingAdvbaseArtProv) GetSpec(period procezor.IPeriod, version procezor.VersionCode) procezor.IArticleSpec {
	return NewTaxingAdvbaseArtSpec(p.Code().Value())
}

func NewTaxingAdvbaseArtProv() procezor.IArticleSpecProvider {
	ArticleCode := ARTICLE_TAXING_ADVBASE.Id()

	return &TaxingAdvbaseArtProv{ArticleSpecProvider: procezor.NewArticleProvider(ArticleCode)}
}

type TaxingAdvbaseArtSpec struct {
	procezor.ArticleSpec
}

func NewTaxingAdvbaseArtSpec(code int32) procezor.IArticleSpec {
	ConceptCode := CONCEPT_TAXING_ADVBASE.Id()

	return &TaxingAdvbaseArtSpec{ArticleSpec: procezor.NewArticleSpec(code, ConceptCode)}
}

type TaxingAdvpaymArtProv struct {
	procezor.ArticleSpecProvider
}

func (p TaxingAdvpaymArtProv) GetSpec(period procezor.IPeriod, version procezor.VersionCode) procezor.IArticleSpec {
	return NewTaxingAdvpaymArtSpec(p.Code().Value())
}

func NewTaxingAdvpaymArtProv() procezor.IArticleSpecProvider {
	ArticleCode := ARTICLE_TAXING_ADVPAYM.Id()

	return &TaxingAdvpaymArtProv{ArticleSpecProvider: procezor.NewArticleProvider(ArticleCode)}
}

type TaxingAdvpaymArtSpec struct {
	procezor.ArticleSpec
}

func NewTaxingAdvpaymArtSpec(code int32) procezor.IArticleSpec {
	ConceptCode := CONCEPT_TAXING_ADVPAYM.Id()

	return &TaxingAdvpaymArtSpec{ArticleSpec: procezor.NewArticleSpec(code, ConceptCode)}
}

type IncomeGrossArtProv struct {
	procezor.ArticleSpecProvider
}

func (p IncomeGrossArtProv) GetSpec(period procezor.IPeriod, version procezor.VersionCode) procezor.IArticleSpec {
	return NewIncomeGrossArtSpec(p.Code().Value())
}

func NewIncomeGrossArtProv() procezor.IArticleSpecProvider {
	ArticleCode := ARTICLE_INCOME_GROSS.Id()

	return &IncomeGrossArtProv{ArticleSpecProvider: procezor.NewArticleProvider(ArticleCode)}
}

type IncomeGrossArtSpec struct {
	procezor.ArticleSpec
}

func NewIncomeGrossArtSpec(code int32) procezor.IArticleSpec {
	ConceptCode := CONCEPT_INCOME_GROSS.Id()

	return &IncomeGrossArtSpec{ArticleSpec: procezor.NewArticleSpec(code, ConceptCode)}
}

type IncomeNettoArtProv struct {
	procezor.ArticleSpecProvider
}

func (p IncomeNettoArtProv) GetSpec(period procezor.IPeriod, version procezor.VersionCode) procezor.IArticleSpec {
	return NewIncomeNettoArtSpec(p.Code().Value())
}

func NewIncomeNettoArtProv() procezor.IArticleSpecProvider {
	ArticleCode := ARTICLE_INCOME_NETTO.Id()

	return &IncomeNettoArtProv{ArticleSpecProvider: procezor.NewArticleProvider(ArticleCode)}
}

type IncomeNettoArtSpec struct {
	procezor.ArticleSpec
}

func NewIncomeNettoArtSpec(code int32) procezor.IArticleSpec {
	ConceptCode := CONCEPT_INCOME_NETTO.Id()

	return &IncomeNettoArtSpec{ArticleSpec: procezor.NewArticleSpec(code, ConceptCode)}
}

