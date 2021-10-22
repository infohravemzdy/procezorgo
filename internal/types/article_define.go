package types

type IArticleDefine interface {
	Code() ArticleCode
	Role() ConceptCode
}

type articleDefine struct {
	code ArticleCode
	role ConceptCode
}

func NewArticleDefine() IArticleDefine {
	return articleDefine{ code: ArticleZero(), role: ConceptZero() }
}

func GetArticleDefine(code int32, role int32) IArticleDefine {
	return articleDefine{ code: GetArticleCode(code), role: GetConceptCode(role) }
}

func CloneArticleDefine(defs IArticleDefine) IArticleDefine {
	return articleDefine{ code: CloneArticleCode(defs.Code()), role: CloneConceptCode(defs.Role()) }
}

func (d articleDefine) Code() ArticleCode {
	return d.code
}

func (d articleDefine) Role() ConceptCode {
	return d.role
}

