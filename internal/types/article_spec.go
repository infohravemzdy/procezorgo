package types

type IArticleSpec interface {
	Code() ArticleCode
	Seqs() ArticleSeqs
	Role() ConceptCode
	Term() ArticleTerm
	Defs() IArticleDefine
	Sums() []ArticleCode
}

type IArticleSpecList = []IArticleSpec

type ArticleSpec struct {
	sums []ArticleCode
	role ConceptCode
	code ArticleCode
	seqs ArticleSeqs
}

func (s ArticleSpec) Term() ArticleTerm {
	return GetArticleTerm(s.code.Value(), s.seqs.Value())
}

func (s ArticleSpec) Defs() IArticleDefine {
	return GetArticleDefine(s.code.Value(), s.seqs.Value(), s.role.Value())
}

func (s ArticleSpec) Sums() []ArticleCode {
	return s.sums
}

func (s ArticleSpec) Code() ArticleCode {
	return s.code
}

func (s ArticleSpec) Role() ConceptCode {
	return s.role
}

func (s ArticleSpec) Seqs() ArticleSeqs {
	return s.seqs
}

func NewArticleSpec(code int32, seqs int16, role int32) ArticleSpec {
	return ArticleSpec{ code: GetArticleCode(code), seqs: GetArticleSeqs(seqs), role: GetConceptCode(role), sums: ArticleCodeList{} }
}

func NewArticleSumSpec(code int32, seqs int16, role int32, sums []ArticleCode) ArticleSpec {
	return ArticleSpec{ code: GetArticleCode(code), seqs: GetArticleSeqs(seqs), role: GetConceptCode(role), sums: sums }
}

func NewArticleSumIntSpec(code int32, seqs int16, role int32, sums []int32) ArticleSpec {
	return ArticleSpec{ code: GetArticleCode(code), seqs: GetArticleSeqs(seqs), role: GetConceptCode(role), sums: ConstToSumsArray(sums) }
}

func ConstToSumsArray(_sums []int32) []ArticleCode {
	sumsArray := make([]ArticleCode, 0)
	for _, r := range _sums {
		sumsArray = append(sumsArray, GetArticleCode(r))
	}
	return sumsArray
}

