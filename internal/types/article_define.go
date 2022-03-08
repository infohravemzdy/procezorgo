package types

type IArticleDefine interface {
	Code() ArticleCode
	Seqs() ArticleSeqs
	Role() ConceptCode
	Term() ArticleTerm
}

type IArticleDefineList = []IArticleDefine

type articleDefine struct {
	code ArticleCode
	seqs ArticleSeqs
	role ConceptCode
}

func NewArticleDefine() IArticleDefine {
	return articleDefine{ code: ArticleCodeZero(), seqs: ArticleSeqsZero(), role: ConceptCodeZero() }
}

func GetArticleDefine(code int32, seqs int16, role int32) IArticleDefine {
	return articleDefine{ code: GetArticleCode(code), seqs: GetArticleSeqs(seqs), role: GetConceptCode(role) }
}

func CloneArticleDefine(defs IArticleDefine) IArticleDefine {
	return articleDefine{ code: CloneArticleCode(defs.Code()), seqs: CloneArticleSeqs(defs.Seqs()), role: CloneConceptCode(defs.Role()) }
}

func (d articleDefine) Code() ArticleCode {
	return d.code
}

func (d articleDefine) Role() ConceptCode {
	return d.role
}

func (d articleDefine) Seqs() ArticleSeqs {
	return d.seqs
}

func (d articleDefine) Term() ArticleTerm {
	return GetArticleTerm(d.code.Value(), d.seqs.Value())
}

