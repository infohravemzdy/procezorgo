package example

import procezor "github.com/mzdyhrave/procezorgo"

type TimeshtWorkingResult struct {
	ExampleTermResult
}

func NewTimeshtWorkingResult(target procezor.ITermTarget, spec procezor.IArticleSpec) procezor.ITermResult {
	return &TimeshtWorkingResult{ExampleTermResult: NewExampleTermResult(target, spec)}
}

type AmountBasisResult struct {
	ExampleTermResult
}

func NewAmountBasisResult(target procezor.ITermTarget, spec procezor.IArticleSpec) procezor.ITermResult {
	return &AmountBasisResult{ExampleTermResult: NewExampleTermResult(target, spec)}
}

type AmountFixedResult struct {
	ExampleTermResult
}

func NewAmountFixedResult(target procezor.ITermTarget, spec procezor.IArticleSpec) procezor.ITermResult {
	return &AmountFixedResult{ExampleTermResult: NewExampleTermResult(target, spec)}
}

type HealthInsbaseResult struct {
	ExampleTermResult
}

func NewHealthInsbaseResult(target procezor.ITermTarget, spec procezor.IArticleSpec) procezor.ITermResult {
	return &HealthInsbaseResult{ExampleTermResult: NewExampleTermResult(target, spec)}
}

type SocialInsbaseResult struct {
	ExampleTermResult
}

func NewSocialInsbaseResult(target procezor.ITermTarget, spec procezor.IArticleSpec) procezor.ITermResult {
	return &SocialInsbaseResult{ExampleTermResult: NewExampleTermResult(target, spec)}
}

type HealthInspaymResult struct {
	ExampleTermResult
}

func NewHealthInspaymResult(target procezor.ITermTarget, spec procezor.IArticleSpec) procezor.ITermResult {
	return &HealthInspaymResult{ExampleTermResult: NewExampleTermResult(target, spec)}
}

type SocialInspaymResult struct {
	ExampleTermResult
}

func NewSocialInspaymResult(target procezor.ITermTarget, spec procezor.IArticleSpec) procezor.ITermResult {
	return &SocialInspaymResult{ExampleTermResult: NewExampleTermResult(target, spec)}
}

type TaxingAdvbaseResult struct {
	ExampleTermResult
}

func NewTaxingAdvbaseResult(target procezor.ITermTarget, spec procezor.IArticleSpec) procezor.ITermResult {
	return &TaxingAdvbaseResult{ExampleTermResult: NewExampleTermResult(target, spec)}
}

type TaxingAdvpaymResult struct {
	ExampleTermResult
}

func NewTaxingAdvpaymResult(target procezor.ITermTarget, spec procezor.IArticleSpec) procezor.ITermResult {
	return &TaxingAdvpaymResult{ExampleTermResult: NewExampleTermResult(target, spec)}
}

type IncomeGrossResult struct {
	ExampleTermResult
}

func NewIncomeGrossResult(target procezor.ITermTarget, spec procezor.IArticleSpec) procezor.ITermResult {
	return &IncomeGrossResult{ExampleTermResult: NewExampleTermResult(target, spec)}
}

type IncomeNettoResult struct {
	ExampleTermResult
}

func NewIncomeNettoResult(target procezor.ITermTarget, spec procezor.IArticleSpec) procezor.ITermResult {
	return &IncomeNettoResult{ExampleTermResult: NewExampleTermResult(target, spec)}
}

