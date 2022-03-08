package registry

import (
	providers "github.com/mzdyhrave/procezorgo/internal/registry_providers"
	"github.com/mzdyhrave/procezorgo/internal/types"
)

func toDefsList(source []types.IArticleDefine) []types.IArticleDefine {
	result := make([]types.IArticleDefine, 0, len(source))
	for _, v := range source {
		result = append(result, v)
	}
	return result
}

func distinctDefsList(source []types.IArticleDefine) []types.IArticleDefine {
	checkMap := make(map[int32]types.IArticleDefine)
	for _, val := range source {
		checkMap[val.Code().Value()] = val
	}
	distinct := make([]types.IArticleDefine, 0, len(checkMap))
	for _, val := range checkMap {
		distinct = append(distinct, val)
	}
	return distinct
}

func toEdgeHashSet(edges articleEdgeSet) articleEdgeSet {
	result := articleEdgeSet{}
	for _, e := range edges.Items() {
		result.Add(e)
	}
	return result
}

type findArticleSpecFunc func (spec types.IArticleSpec) bool

func findArticleSpecCode(code int32) findArticleSpecFunc {
	return func (item types.IArticleSpec) bool { return item.Code().Value() == code	}
}

func firstOrDefaultArticleSpec(source []types.IArticleSpec, findFunc findArticleSpecFunc) (types.IArticleSpec, bool) {
	for _, v := range source {
		if findFunc(v) {
			return v, true
		}
	}
	return nil, false
}

func getArticleTerm(article types.ArticleCode, articlesModel []types.IArticleSpec) types.ArticleTerm {
	articleSpec, _ := firstOrDefaultArticleSpec(articlesModel, findArticleSpecCode(article.Value()))
	if articleSpec == nil {
		return types.NewArticleTerm()
	}
	return articleSpec.Term()
}

func getArticleDefs(article types.ArticleCode, articlesModel []types.IArticleSpec) types.IArticleDefine {
	articleSpec, _ := firstOrDefaultArticleSpec(articlesModel, findArticleSpecCode(article.Value()))
	if articleSpec == nil {
		return types.NewArticleDefine()
	}
	return articleSpec.Defs()
}

type findConceptSpecFunc func (spec providers.IConceptSpec) bool

func findConceptSpecCode(code int32) findConceptSpecFunc {
	return func (item providers.IConceptSpec) bool { return item.Code().Value() == code	}
}

func firstOrDefaultConceptSpec(source []providers.IConceptSpec, findFunc findConceptSpecFunc) (providers.IConceptSpec, bool) {
	for _, v := range source {
		if findFunc(v) {
			return v, true
		}
	}
	return nil, false
}

type findArticleCodeFunc func (spec types.ArticleCode) bool
type findArticleTermFunc func (spec types.ArticleTerm) bool

func findArticleCode(article types.ArticleCode) findArticleCodeFunc {
	return func (item types.ArticleCode) bool { return item == article }
}

func findArticleTern(article types.ArticleCode) findArticleTermFunc {
	return func (item types.ArticleTerm) bool { return item.Code() == article }
}

type findTermTargetFunc func (term types.ITermTarget) bool

func findTermTargetCode(article types.ArticleCode) findTermTargetFunc {
	return func (term types.ITermTarget) bool { return term.Article() == article }
}

func firstOrDefaultTermTarget(source []types.ITermTarget, findFunc findTermTargetFunc) (types.ITermTarget, bool) {
	for _, v := range source {
		if findFunc(v) {
			return v, true
		}
	}
	return nil, false
}

func toTargetList(source []types.ITermTarget) []types.ITermTarget {
	result := make([]types.ITermTarget, 0, len(source))
	for _, v := range source {
		result = append(result, v)
	}
	return result
}

func firstOrDefaultInPathMap(source PathsTermsMap, findFunc findArticleTermFunc) ([]types.IArticleDefine, bool) {
	for k, v := range source {
		if findFunc(k) {
			return v.list, true
		}
	}
	return nil, false
}

func cmpInt32(a, b int32) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return +1
}

func cmpInt16(a, b int16) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return +1
}

func cmpInt(a, b int) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return +1
}

func cmpInterface(x, y interface{}) int {
	if x == nil && y != nil	{
		return -1
	}
	if x != nil && y == nil	{
		return 1
	}
	return 0
}

