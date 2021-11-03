package registry

import (
	"errors"
	"fmt"
	legalios "github.com/mzdyhrave/legaliosgo"
	"github.com/mzdyhrave/procezorgo/internal/types"
)

type termResultError struct {
	period      legalios.IPeriod
	target      types.ITermTarget
	contract    types.ContractCode
	position    types.PositionCode
	article     types.ArticleCode
	articleName string
	concept     types.ConceptCode
	conceptName string
	variant     types.VariantCode
	innerResult types.ITermResultError
	err         error
}

func (e termResultError) Period() legalios.IPeriod {
	return e.period
}

func (e termResultError) Target() types.ITermTarget {
	return e.target
}

func (e termResultError) Contract() types.ContractCode {
	return e.contract
}

func (e termResultError) Position() types.PositionCode {
	return e.position
}

func (e termResultError) Article() types.ArticleCode {
	return e.article
}

func (e termResultError) ArticleName() string {
	return e.articleName
}

func (e termResultError) Concept() types.ConceptCode {
	return e.concept
}

func (e termResultError) ConceptName() string {
	return e.conceptName
}

func (e termResultError) Variant() types.VariantCode {
	return e.variant
}

func (e termResultError) InnerResult() types.ITermResultError {
	return e.innerResult
}

func (e termResultError) Err() error {
	return e.err
}

func (e termResultError) Error() string {
	return fmt.Sprintf("error: %v", e.Err())
}

func (e termResultError) ArticleDescr() string {
	if e.target == nil {
		return e.target.ArticleDescr()
	}
	return fmt.Sprintf("ArticleCode for: %v", e.article.Value())
}

func (e termResultError) ConceptDescr() string {
	if e.target == nil {
		return e.target.ConceptDescr()
	}
	return fmt.Sprintf("ConceptCode for: %v", e.concept.Value())
}

func newResultError(err string) types.ITermResultError {
	return termResultError{ err: errors.New(err) }
}

func NewNoResultFuncError(period legalios.IPeriod, target types.ITermTarget) types.ITermResultError {
	_contract := types.NewContractCode()
	_position := types.NewPositionCode()
	_article  := types.NewArticleCode()
	_concept  := types.NewConceptCode()
	_variant  := types.NewVariantCode()

	if target != nil {
		_contract = target.Contract()
		_position = target.Position()
		_article  = target.Article()
		_concept  = target.Concept()
		_variant  = target.Variant()
	}
	return termResultError{ period: period, target: target,
		contract: _contract,
		position: _position,
		article: _article,
		concept: _concept,
		variant: _variant,
		err: errors.New("no result calculation function") }
}

func NewEvalResultError(period legalios.IPeriod, target types.ITermTarget, inner types.ITermResultError, errorText string) types.ITermResultError {
	_contract := types.NewContractCode()
	_position := types.NewPositionCode()
	_article  := types.NewArticleCode()
	_concept  := types.NewConceptCode()
	_variant  := types.NewVariantCode()

	if target != nil {
		_contract = target.Contract()
		_position = target.Position()
		_article  = target.Article()
		_concept  = target.Concept()
		_variant  = target.Variant()
	}
	return termResultError{ period: period, target: target, innerResult: inner,
		contract: _contract,
		position: _position,
		article: _article,
		concept: _concept,
		variant: _variant,
		err: errors.New(errorText) }
}

func NewExtractResultError( period legalios.IPeriod, target types.ITermTarget, symbol types.ITermSymbol, inner types.ITermResultError, errorText string) types.ITermResultError {
	_contract := types.NewContractCode()
	_position := types.NewPositionCode()
	_article  := types.NewArticleCode()
	_concept  := types.NewConceptCode()
	_variant  := types.NewVariantCode()

	if target != nil {
		_contract = target.Contract()
		_position = target.Position()
		_article  = target.Article()
		_concept  = target.Concept()
		_variant  = target.Variant()
	}
	return termResultError{ period: period, target: target, innerResult: inner,
		contract: _contract,
		position: _position,
		article: _article,
		concept: _concept,
		variant: _variant,
		err: errors.New(errorText) }
}

func NewNoImplementationError( period legalios.IPeriod, target types.ITermTarget) types.ITermResultError {
	_contract := types.NewContractCode()
	_position := types.NewPositionCode()
	_article  := types.NewArticleCode()
	_concept  := types.NewConceptCode()
	_variant  := types.NewVariantCode()

	if target != nil {
		_contract = target.Contract()
		_position = target.Position()
		_article  = target.Article()
		_concept  = target.Concept()
		_variant  = target.Variant()
	}
	return termResultError{ period: period, target: target,
		contract: _contract,
		position: _position,
		article: _article,
		concept: _concept,
		variant: _variant,
		err: errors.New("no implementation") }
}

