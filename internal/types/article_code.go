package types

const (
	articleZeroCode int32 = 0
)

type ArticleCode struct {
	value int32
}

type ArticleCodeList []ArticleCode

func (v ArticleCode) Value() int32 {
	return v.value
}

func (v ArticleCode) IsZero() bool {
	return v.value == articleZeroCode
}

func ArticleCodeZero() ArticleCode {
	return NewArticleCode()
}

func NewArticleCode() ArticleCode {
	return ArticleCode{	value: articleZeroCode}
}

func GetArticleCode(id int32) ArticleCode {
	return ArticleCode{	value: id }
}

func CloneArticleCode(code ArticleCode) ArticleCode {
	return ArticleCode{	value: code.value }
}

