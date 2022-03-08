package types

const (
	articleZeroSeqs int16 = 0
)

type ArticleSeqs struct {
	value int16
}

type ArticleSeqsList []ArticleSeqs

func (v ArticleSeqs) Value() int16 {
	return v.value
}

func (v ArticleSeqs) IsZero() bool {
	return v.value == articleZeroSeqs
}

func ArticleSeqsZero() ArticleSeqs {
	return NewArticleSeqs()
}

func NewArticleSeqs() ArticleSeqs {
	return ArticleSeqs{	value: articleZeroSeqs}
}

func GetArticleSeqs(id int16) ArticleSeqs {
	return ArticleSeqs{	value: id }
}

func CloneArticleSeqs(seqs ArticleSeqs) ArticleSeqs {
	return ArticleSeqs{	value: seqs.value }
}

