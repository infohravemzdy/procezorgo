package types

type ArticleTerm struct {
	code ArticleCode
	seqs ArticleSeqs
}

type ArticleTermList []ArticleTerm

func (v ArticleTerm) Code() ArticleCode {
	return v.code
}

func (v ArticleTerm) CodeValue() int32 {
	return v.code.Value()
}

func (v ArticleTerm) Seqs() ArticleSeqs {
	return v.seqs
}

func (v ArticleTerm) SeqsValue() int16 {
	return v.seqs.Value()
}

func (v ArticleTerm) IsZero() bool {
	return v.code.IsZero()
}

func ArticleTermZero() ArticleTerm {
	return NewArticleTerm()
}

func NewArticleTerm() ArticleTerm {
	return ArticleTerm{	code: ArticleCodeZero(), seqs: ArticleSeqsZero() }
}

func GetArticleTerm(code int32, seqs int16) ArticleTerm {
	return ArticleTerm{	code: GetArticleCode(code), seqs: GetArticleSeqs(seqs) }
}

func CloneArticleTerm(term ArticleTerm) ArticleTerm {
	return ArticleTerm{	code: CloneArticleCode(term.Code()), seqs: CloneArticleSeqs(term.Seqs()) }
}

func ArticleTermCompare(x ArticleTerm, y ArticleTerm) int {
	if x.SeqsValue() > y.SeqsValue() {
		return 1
	}

	if x.SeqsValue() < y.SeqsValue() {
		return -1
	}

	if x.CodeValue() > y.CodeValue() {
		return 1
	}

	if x.CodeValue() < y.CodeValue() {
		return -1
	}
	return 0
}

