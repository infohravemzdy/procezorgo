package example

import procezor "github.com/mzdyhrave/procezorgo"

type TimeshtWorkingResult struct {
	TestTermResult
}

func NewTimeshtWorkingResult(target procezor.ITermTarget, value int32, basis int32, descr string) procezor.ITermResult {
	return &TimeshtWorkingResult{TestTermResult: NewTestTermResult(target, value, basis, descr)}
}

func NewEmptyTimeshtWorkingResult(target procezor.ITermTarget) procezor.ITermResult {
	return &TimeshtWorkingResult{TestTermResult: NewTestTermResult(target, RESULT_VALUE_ZERO, RESULT_BASIS_ZERO, RESULT_DESCRIPTION_EMPTY)}
}

type AmountBasisResult struct {
	TestTermResult
}

func NewAmountBasisResult(target procezor.ITermTarget, value int32, basis int32, descr string) procezor.ITermResult {
	return &AmountBasisResult{TestTermResult: NewTestTermResult(target, value, basis, descr)}
}

func NewEmptyAmountBasisResult(target procezor.ITermTarget) procezor.ITermResult {
	return &AmountBasisResult{TestTermResult: NewTestTermResult(target, RESULT_VALUE_ZERO, RESULT_BASIS_ZERO, RESULT_DESCRIPTION_EMPTY)}
}

type AmountFixedResult struct {
	TestTermResult
}

func NewAmountFixedResult(target procezor.ITermTarget, value int32, basis int32, descr string) procezor.ITermResult {
	return &AmountFixedResult{TestTermResult: NewTestTermResult(target, value, basis, descr)}
}

func NewEmptyAmountFixedResult(target procezor.ITermTarget) procezor.ITermResult {
	return &AmountFixedResult{TestTermResult: NewTestTermResult(target, RESULT_VALUE_ZERO, RESULT_BASIS_ZERO, RESULT_DESCRIPTION_EMPTY)}
}

type HealthInsbaseResult struct {
	TestTermResult
}

func NewHealthInsbaseResult(target procezor.ITermTarget, value int32, basis int32, descr string) procezor.ITermResult {
	return &HealthInsbaseResult{TestTermResult: NewTestTermResult(target, value, basis, descr)}
}

func NewEmptyHealthInsbaseResult(target procezor.ITermTarget) procezor.ITermResult {
	return &HealthInsbaseResult{TestTermResult: NewTestTermResult(target, RESULT_VALUE_ZERO, RESULT_BASIS_ZERO, RESULT_DESCRIPTION_EMPTY)}
}

type SocialInsbaseResult struct {
	TestTermResult
}

func NewSocialInsbaseResult(target procezor.ITermTarget, value int32, basis int32, descr string) procezor.ITermResult {
	return &SocialInsbaseResult{TestTermResult: NewTestTermResult(target, value, basis, descr)}
}

func NewEmptySocialInsbaseResult(target procezor.ITermTarget) procezor.ITermResult {
	return &SocialInsbaseResult{TestTermResult: NewTestTermResult(target, RESULT_VALUE_ZERO, RESULT_BASIS_ZERO, RESULT_DESCRIPTION_EMPTY)}
}

type TaxingAdvbaseResult struct {
	TestTermResult
}

func NewTaxingAdvbaseResult(target procezor.ITermTarget, value int32, basis int32, descr string) procezor.ITermResult {
	return &TaxingAdvbaseResult{TestTermResult: NewTestTermResult(target, value, basis, descr)}
}

func NewEmptyTaxingAdvbaseResult(target procezor.ITermTarget) procezor.ITermResult {
	return &TaxingAdvbaseResult{TestTermResult: NewTestTermResult(target, RESULT_VALUE_ZERO, RESULT_BASIS_ZERO, RESULT_DESCRIPTION_EMPTY)}
}

type HealthInspaymResult struct {
	TestTermResult
}

func NewHealthInspaymResult(target procezor.ITermTarget, value int32, basis int32, descr string) procezor.ITermResult {
	return &HealthInspaymResult{TestTermResult: NewTestTermResult(target, value, basis, descr)}
}

func NewEmptyHealthInspaymResult(target procezor.ITermTarget) procezor.ITermResult {
	return &HealthInspaymResult{TestTermResult: NewTestTermResult(target, RESULT_VALUE_ZERO, RESULT_BASIS_ZERO, RESULT_DESCRIPTION_EMPTY)}
}

type SocialInspaymResult struct {
	TestTermResult
}

func NewSocialInspaymResult(target procezor.ITermTarget, value int32, basis int32, descr string) procezor.ITermResult {
	return &SocialInspaymResult{TestTermResult: NewTestTermResult(target, value, basis, descr)}
}

func NewEmptySocialInspaymResult(target procezor.ITermTarget) procezor.ITermResult {
	return &SocialInspaymResult{TestTermResult: NewTestTermResult(target, RESULT_VALUE_ZERO, RESULT_BASIS_ZERO, RESULT_DESCRIPTION_EMPTY)}
}

type TaxingAdvpaymResult struct {
	TestTermResult
}

func NewTaxingAdvpaymResult(target procezor.ITermTarget, value int32, basis int32, descr string) procezor.ITermResult {
	return &TaxingAdvpaymResult{TestTermResult: NewTestTermResult(target, value, basis, descr)}
}

func NewEmptyTaxingAdvpaymResult(target procezor.ITermTarget) procezor.ITermResult {
	return &TaxingAdvpaymResult{TestTermResult: NewTestTermResult(target, RESULT_VALUE_ZERO, RESULT_BASIS_ZERO, RESULT_DESCRIPTION_EMPTY)}
}

type IncomeGrossResult struct {
	TestTermResult
}

func NewIncomeGrossResult(target procezor.ITermTarget, value int32, basis int32, descr string) procezor.ITermResult {
	return &IncomeGrossResult{TestTermResult: NewTestTermResult(target, value, basis, descr)}
}

func NewEmptyIncomeGrossResult(target procezor.ITermTarget) procezor.ITermResult {
	return &IncomeGrossResult{TestTermResult: NewTestTermResult(target, RESULT_VALUE_ZERO, RESULT_BASIS_ZERO, RESULT_DESCRIPTION_EMPTY)}
}

type IncomeNettoResult struct {
	TestTermResult
}

func NewIncomeNettoResult(target procezor.ITermTarget, value int32, basis int32, descr string) procezor.ITermResult {
	return &IncomeNettoResult{TestTermResult: NewTestTermResult(target, value, basis, descr)}
}

func NewEmptyIncomeNettoResult(target procezor.ITermTarget) procezor.ITermResult {
	return &IncomeNettoResult{TestTermResult: NewTestTermResult(target, RESULT_VALUE_ZERO, RESULT_BASIS_ZERO, RESULT_DESCRIPTION_EMPTY)}
}
