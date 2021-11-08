package example

import procezor "github.com/mzdyhrave/procezorgo"

type TimeshtWorkingResult struct {
	ExampleTermResult
}

func NewTimeshtWorkingResult(target procezor.ITermTarget, value int32, basis int32, descr string) procezor.ITermResult {
	return &TimeshtWorkingResult{ExampleTermResult: NewExampleTermResult(target, value, basis, descr)}
}

func NewEmptyTimeshtWorkingResult(target procezor.ITermTarget) procezor.ITermResult {
	return &TimeshtWorkingResult{ExampleTermResult: NewExampleTermResult(target, RESULT_VALUE_ZERO, RESULT_BASIS_ZERO, RESULT_DESCRIPTION_EMPTY)}
}

type AmountBasisResult struct {
	ExampleTermResult
}

func NewAmountBasisResult(target procezor.ITermTarget, value int32, basis int32, descr string) procezor.ITermResult {
	return &AmountBasisResult{ExampleTermResult: NewExampleTermResult(target, value, basis, descr)}
}

func NewEmptyAmountBasisResult(target procezor.ITermTarget) procezor.ITermResult {
	return &AmountBasisResult{ExampleTermResult: NewExampleTermResult(target, RESULT_VALUE_ZERO, RESULT_BASIS_ZERO, RESULT_DESCRIPTION_EMPTY)}
}

type AmountFixedResult struct {
	ExampleTermResult
}

func NewAmountFixedResult(target procezor.ITermTarget, value int32, basis int32, descr string) procezor.ITermResult {
	return &AmountFixedResult{ExampleTermResult: NewExampleTermResult(target, value, basis, descr)}
}

func NewEmptyAmountFixedResult(target procezor.ITermTarget) procezor.ITermResult {
	return &AmountFixedResult{ExampleTermResult: NewExampleTermResult(target, RESULT_VALUE_ZERO, RESULT_BASIS_ZERO, RESULT_DESCRIPTION_EMPTY)}
}

type HealthInsbaseResult struct {
	ExampleTermResult
}

func NewHealthInsbaseResult(target procezor.ITermTarget, value int32, basis int32, descr string) procezor.ITermResult {
	return &HealthInsbaseResult{ExampleTermResult: NewExampleTermResult(target, value, basis, descr)}
}

func NewEmptyHealthInsbaseResult(target procezor.ITermTarget) procezor.ITermResult {
	return &HealthInsbaseResult{ExampleTermResult: NewExampleTermResult(target, RESULT_VALUE_ZERO, RESULT_BASIS_ZERO, RESULT_DESCRIPTION_EMPTY)}
}

type SocialInsbaseResult struct {
	ExampleTermResult
}

func NewSocialInsbaseResult(target procezor.ITermTarget, value int32, basis int32, descr string) procezor.ITermResult {
	return &SocialInsbaseResult{ExampleTermResult: NewExampleTermResult(target, value, basis, descr)}
}

func NewEmptySocialInsbaseResult(target procezor.ITermTarget) procezor.ITermResult {
	return &SocialInsbaseResult{ExampleTermResult: NewExampleTermResult(target, RESULT_VALUE_ZERO, RESULT_BASIS_ZERO, RESULT_DESCRIPTION_EMPTY)}
}

type HealthInspaymResult struct {
	ExampleTermResult
}

func NewHealthInspaymResult(target procezor.ITermTarget, value int32, basis int32, descr string) procezor.ITermResult {
	return &HealthInspaymResult{ExampleTermResult: NewExampleTermResult(target, value, basis, descr)}
}

func NewEmptyHealthInspaymResult(target procezor.ITermTarget) procezor.ITermResult {
	return &HealthInspaymResult{ExampleTermResult: NewExampleTermResult(target, RESULT_VALUE_ZERO, RESULT_BASIS_ZERO, RESULT_DESCRIPTION_EMPTY)}
}

type SocialInspaymResult struct {
	ExampleTermResult
}

func NewSocialInspaymResult(target procezor.ITermTarget, value int32, basis int32, descr string) procezor.ITermResult {
	return &SocialInspaymResult{ExampleTermResult: NewExampleTermResult(target, value, basis, descr)}
}

func NewEmptySocialInspaymResult(target procezor.ITermTarget) procezor.ITermResult {
	return &SocialInspaymResult{ExampleTermResult: NewExampleTermResult(target, RESULT_VALUE_ZERO, RESULT_BASIS_ZERO, RESULT_DESCRIPTION_EMPTY)}
}

type TaxingAdvbaseResult struct {
	ExampleTermResult
}

func NewTaxingAdvbaseResult(target procezor.ITermTarget, value int32, basis int32, descr string) procezor.ITermResult {
	return &TaxingAdvbaseResult{ExampleTermResult: NewExampleTermResult(target, value, basis, descr)}
}

func NewEmptyTaxingAdvbaseResult(target procezor.ITermTarget) procezor.ITermResult {
	return &TaxingAdvbaseResult{ExampleTermResult: NewExampleTermResult(target, RESULT_VALUE_ZERO, RESULT_BASIS_ZERO, RESULT_DESCRIPTION_EMPTY)}
}

type TaxingAdvpaymResult struct {
	ExampleTermResult
}

func NewTaxingAdvpaymResult(target procezor.ITermTarget, value int32, basis int32, descr string) procezor.ITermResult {
	return &TaxingAdvpaymResult{ExampleTermResult: NewExampleTermResult(target, value, basis, descr)}
}

func NewEmptyTaxingAdvpaymResult(target procezor.ITermTarget) procezor.ITermResult {
	return &TaxingAdvpaymResult{ExampleTermResult: NewExampleTermResult(target, RESULT_VALUE_ZERO, RESULT_BASIS_ZERO, RESULT_DESCRIPTION_EMPTY)}
}

type IncomeGrossResult struct {
	ExampleTermResult
}

func NewIncomeGrossResult(target procezor.ITermTarget, value int32, basis int32, descr string) procezor.ITermResult {
	return &IncomeGrossResult{ExampleTermResult: NewExampleTermResult(target, value, basis, descr)}
}

func NewEmptyIncomeGrossResult(target procezor.ITermTarget) procezor.ITermResult {
	return &IncomeGrossResult{ExampleTermResult: NewExampleTermResult(target, RESULT_VALUE_ZERO, RESULT_BASIS_ZERO, RESULT_DESCRIPTION_EMPTY)}
}

type IncomeNettoResult struct {
	ExampleTermResult
}

func NewIncomeNettoResult(target procezor.ITermTarget, value int32, basis int32, descr string) procezor.ITermResult {
	return &IncomeNettoResult{ExampleTermResult: NewExampleTermResult(target, value, basis, descr)}
}

func NewEmptyIncomeNettoResult(target procezor.ITermTarget) procezor.ITermResult {
	return &IncomeNettoResult{ExampleTermResult: NewExampleTermResult(target, RESULT_VALUE_ZERO, RESULT_BASIS_ZERO, RESULT_DESCRIPTION_EMPTY)}
}

