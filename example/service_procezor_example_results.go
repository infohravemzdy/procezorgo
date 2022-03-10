package example

import procezor "github.com/mzdyhrave/procezorgo"

// TimeshtWorking		TIMESHT_WORKING
type TimeshtWorkingResult struct {
    ExampleTermResult
}

func NewTimeshtWorkingResult(target procezor.ITermTarget, spec procezor.IArticleSpec) procezor.ITermResult {
    return &TimeshtWorkingResult{ExampleTermResult: NewExampleTermResult(target, spec)}
}

// AmountBasis		AMOUNT_BASIS
type AmountBasisResult struct {
    ExampleTermResult
}

func NewAmountBasisResult(target procezor.ITermTarget, spec procezor.IArticleSpec) procezor.ITermResult {
    return &AmountBasisResult{ExampleTermResult: NewExampleTermResult(target, spec)}
}

// AmountFixed		AMOUNT_FIXED
type AmountFixedResult struct {
    ExampleTermResult
}

func NewAmountFixedResult(target procezor.ITermTarget, spec procezor.IArticleSpec) procezor.ITermResult {
    return &AmountFixedResult{ExampleTermResult: NewExampleTermResult(target, spec)}
}

// HealthInsbase		HEALTH_INSBASE
type HealthInsbaseResult struct {
    ExampleTermResult
}

func NewHealthInsbaseResult(target procezor.ITermTarget, spec procezor.IArticleSpec) procezor.ITermResult {
    return &HealthInsbaseResult{ExampleTermResult: NewExampleTermResult(target, spec)}
}

// SocialInsbase		SOCIAL_INSBASE
type SocialInsbaseResult struct {
    ExampleTermResult
}

func NewSocialInsbaseResult(target procezor.ITermTarget, spec procezor.IArticleSpec) procezor.ITermResult {
    return &SocialInsbaseResult{ExampleTermResult: NewExampleTermResult(target, spec)}
}

// HealthInspaym		HEALTH_INSPAYM
type HealthInspaymResult struct {
    ExampleTermResult
}

func NewHealthInspaymResult(target procezor.ITermTarget, spec procezor.IArticleSpec) procezor.ITermResult {
    return &HealthInspaymResult{ExampleTermResult: NewExampleTermResult(target, spec)}
}

// SocialInspaym		SOCIAL_INSPAYM
type SocialInspaymResult struct {
    ExampleTermResult
}

func NewSocialInspaymResult(target procezor.ITermTarget, spec procezor.IArticleSpec) procezor.ITermResult {
    return &SocialInspaymResult{ExampleTermResult: NewExampleTermResult(target, spec)}
}

// TaxingAdvbase		TAXING_ADVBASE
type TaxingAdvbaseResult struct {
    ExampleTermResult
}

func NewTaxingAdvbaseResult(target procezor.ITermTarget, spec procezor.IArticleSpec) procezor.ITermResult {
    return &TaxingAdvbaseResult{ExampleTermResult: NewExampleTermResult(target, spec)}
}

// TaxingAdvpaym		TAXING_ADVPAYM
type TaxingAdvpaymResult struct {
    ExampleTermResult
}

func NewTaxingAdvpaymResult(target procezor.ITermTarget, spec procezor.IArticleSpec) procezor.ITermResult {
    return &TaxingAdvpaymResult{ExampleTermResult: NewExampleTermResult(target, spec)}
}

// IncomeGross		INCOME_GROSS
type IncomeGrossResult struct {
    ExampleTermResult
}

func NewIncomeGrossResult(target procezor.ITermTarget, spec procezor.IArticleSpec) procezor.ITermResult {
    return &IncomeGrossResult{ExampleTermResult: NewExampleTermResult(target, spec)}
}

// IncomeNetto		INCOME_NETTO
type IncomeNettoResult struct {
    ExampleTermResult
}

func NewIncomeNettoResult(target procezor.ITermTarget, spec procezor.IArticleSpec) procezor.ITermResult {
    return &IncomeNettoResult{ExampleTermResult: NewExampleTermResult(target, spec)}
}

