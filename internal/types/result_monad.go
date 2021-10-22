package types

import (
	legalios "github.com/mzdyhrave/legaliosgo"
)

type ITermResultError interface {
	Period() legalios.IPeriod
	Target() ITermTarget
	Contract() ContractCode
	Position() PositionCode
	Article() ArticleCode
	ArticleName() string
	Concept() ConceptCode
	ConceptName() string
	Variant() VariantCode
	InnerResult() ITermResultError
	Err()         error
	Error()       string
}

type FailureResult struct {
	error ITermResultError
}

func (f FailureResult) IsFailure() bool {
	return true
}

func (f FailureResult) IsSuccess() bool {
	return false
}

func (f FailureResult) Error() error {
	return f.error
}

func (f FailureResult) ResultError() ITermResultError {
	return f.error
}

func (f FailureResult) Value() ITermResult {
	return nil
}

type SuccessResult struct {
	value ITermResult
}

func (s SuccessResult) IsFailure() bool {
	return false
}

func (s SuccessResult) IsSuccess() bool {
	return true
}

func (s SuccessResult) Error() error {
	return nil
}

func (s SuccessResult) ResultError() ITermResultError {
	return nil
}

func (s SuccessResult) Value() ITermResult {
	return s.value
}

func NewFailureResult(error ITermResultError) FailureResult {
	return FailureResult{error }
}

func NewSuccessResult(result ITermResult, period legalios.IPeriod) SuccessResult {
	return SuccessResult{result }
}

