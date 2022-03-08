package example

import (
	procezor "github.com/mzdyhrave/procezorgo"
	"github.com/mzdyhrave/procezorgo/internal/types"
)

const TEST_VERSION int32 = 100

type ExampleConceptSpec struct {
	procezor.ConceptSpec
}

func NewExampleConceptSpec(code int32) ExampleConceptSpec {
	return ExampleConceptSpec{procezor.NewConceptSpec(code)}
}

func NewExampleConceptPathSpec(code int32, path []procezor.ArticleCode) ExampleConceptSpec {
	return ExampleConceptSpec{procezor.NewConceptPathSpec(code, path)}
}

func NewExampleConceptPathIntSpec(code int32, path []int32) ExampleConceptSpec {
	return ExampleConceptSpec{procezor.NewConceptPathIntSpec(code, path)}
}

func NewExampleConceptFuncSpec(code int32, resultFunc procezor.ResultFunc) ExampleConceptSpec {
	return ExampleConceptSpec{procezor.NewConceptFuncSpec(code, resultFunc)}
}

func NewExampleConceptPathFuncSpec(code int32, path []procezor.ArticleCode, resultFunc procezor.ResultFunc) ExampleConceptSpec {
	return ExampleConceptSpec{procezor.NewConceptPathFuncSpec(code, path, resultFunc)}
}

func NewExampleConceptPathIntFuncSpec(code int32, path []int32, resultFunc procezor.ResultFunc) ExampleConceptSpec {
	return ExampleConceptSpec{procezor.NewConceptPathIntFuncSpec(code, path, resultFunc)}
}

type ExampleTermTarget struct {
	procezor.TermTarget
}

func (t ExampleTermTarget) ArticleDescr() string {
	return ExampleArticleConst(t.Article().Value()).String()
}

func (t ExampleTermTarget) ConceptDescr() string {
	return ExampleConceptConst(t.Concept().Value()).String()
}

func NewExampleTermTarget(mont types.MonthCode, cont types.ContractCode, post types.PositionCode, vars types.VariantCode, code types.ArticleCode, role types.ConceptCode) ExampleTermTarget {
	return ExampleTermTarget{ TermTarget: procezor.NewTermTarget(mont, cont, post, vars, code, role) }
}

type ExampleTermResult struct {
	procezor.TermResult
}

func (t ExampleTermResult) ArticleDescr() string {
	return ExampleArticleConst(t.Article().Value()).String()
}

func (t ExampleTermResult) ConceptDescr() string {
	return ExampleConceptConst(t.Concept().Value()).String()
}


func NewExampleTermResult(target procezor.ITermTarget, spec procezor.IArticleSpec) ExampleTermResult {
	return ExampleTermResult{ TermResult: procezor.NewTermResult(target, spec) }
}

