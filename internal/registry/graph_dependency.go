package registry

import (
	providers "github.com/mzdyhrave/procezorgo/internal/registry_providers"
	"github.com/mzdyhrave/procezorgo/internal/types"
	"sort"
)

func InitGraphModel(articlesModel []types.IArticleSpec, conceptsModel []providers.IConceptSpec) (OrderCodeList, PathsTermsMap, bool) {
	vertModel := createVertModel(articlesModel)
	edgeModel := createEdgeModel(articlesModel, conceptsModel)
	depsModel := createPendModel(articlesModel, conceptsModel)

	order := createTopoModel(vertModel, edgeModel)

	paths := createPathModel(articlesModel, vertModel, depsModel, order)

	return order, paths, true
}

func createVertModel(articlesModel []types.IArticleSpec) OrderCodeList {
	result := make(OrderCodeList, 0, len(articlesModel))
	for _, article := range articlesModel {
		result = append(result, article.Term())
	}
	sort.Slice(result, func(i, j int) bool {
		return types.ArticleTermCompare(result[i], result[j]) < 0
	})
	return result
}

func createEdgeModel(articlesModel []types.IArticleSpec, conceptsModel []providers.IConceptSpec) []articleEdge {
	result := articleEdgeSet{}
	for _, s := range articlesModel {
		result = mergeEdges(articlesModel, conceptsModel, result, s)
	}
	resultSlice := result.Sort()

	return resultSlice
}

func createPendModel(articlesModel []types.IArticleSpec, conceptsModel []providers.IConceptSpec) []articleEdge {
	result := articleEdgeSet{}
	for _, s := range articlesModel {
		result = mergePends(articlesModel, conceptsModel, result, s)
	}
	resultSlice := result.Sort()
	return resultSlice
}

func mergeEdges(articlesModel []types.IArticleSpec, conceptsModel []providers.IConceptSpec, init articleEdgeSet, article types.IArticleSpec) articleEdgeSet {
	result := toEdgeHashSet(init)

	concept, _ := firstOrDefaultConceptSpec(conceptsModel, findConceptSpecCode(article.Role().Value()))

	for _, s := range article.Sums() {
		result.Add(newArticleEdge(article.Term(), getArticleTerm(s, articlesModel)))
	}
	for _, p := range concept.Path() {
		result.Add(newArticleEdge(getArticleTerm(p, articlesModel), article.Term()))
	}
	return result
}

func mergePends(articlesModel []types.IArticleSpec, conceptsModel []providers.IConceptSpec, init articleEdgeSet, article types.IArticleSpec) articleEdgeSet {
	result := toEdgeHashSet(init)

	concept, _ := firstOrDefaultConceptSpec(conceptsModel, findConceptSpecCode(article.Role().Value()))

	for _, p := range concept.Path() {
		result.Add(newArticleEdge(getArticleTerm(p, articlesModel), article.Term()))
	}
	return result
}

func createTopoModel(vertModel OrderCodeList, edgeModel []articleEdge) types.ArticleTermList {
	articlesOrder := make(OrderCodeList, 0, len(vertModel))

	degrees := make(map[types.ArticleTerm]int32)
	for _, v := range vertModel {
		var count int32 = 0
		for _, e := range edgeModel {
			if e.stops == v {
				count++
			}
		}
		degrees[v] = count
	}
	zeroDegrees := make(OrderCodeList, 0)
	for d, v := range degrees {
		if v == 0 {
			zeroDegrees = append(zeroDegrees, d)
		}
	}
	sort.Slice(zeroDegrees, func(i, j int) bool {
		return types.ArticleTermCompare(zeroDegrees[i], zeroDegrees[j]) < 0
	})

	queues := articleTermQueue{}
	for _, v := range zeroDegrees {
		queues.Enqueue(v)
	}
	var index = 0
	for queues.Size() != 0 {
		index++
		article, err := queues.Dequeue()
		if err != nil {
			return make([]types.ArticleTerm, 0)
		}
		articlesOrder = append(articlesOrder, article)
		paths := make([]types.ArticleTerm, 0)
		for _, e := range edgeModel {
			if e.start == article {
				paths = append(paths, e.stops)
			}
		}
		for _, p := range paths {
			degrees[p] -= 1
			if degrees[p] == 0 {
				queues.Enqueue(p)
			}
		}
	}
	if index != len(vertModel) {
		return make([]types.ArticleTerm, 0)
	}
	return articlesOrder
}

func createPathModel(articlesModel []types.IArticleSpec, vertModel OrderCodeList, edgeModel []articleEdge, vertOrder OrderCodeList) PathsTermsMap {
	result := make(PathsTermsMap)

	for _, v := range vertModel {
		result[v] = mergePaths(articlesModel, edgeModel, vertOrder, v)
	}
	return result
}

func mergePaths(articlesModel []types.IArticleSpec, edgeModel []articleEdge, vertOrder OrderCodeList, article types.ArticleTerm) pathCodeList {
	articleInit := make([]types.IArticleDefine, 0)
	for _, e := range edgeModel {
		if e.stops == article {
			articleInit = append(articleInit, getArticleDefs(e.start.Code(), articlesModel))
		}
	}
	articlePath := toDefsList(articleInit)
	for _, a := range articleInit {
		articlePath = mergeVert(articlesModel, edgeModel, articlePath, a)
	}

	articleDist := distinctDefsList(articlePath)

	sort.Slice(articleDist, func(i, j int) bool {
		return defineCompare(vertOrder, articleDist[i], articleDist[j]) < 0
	})
	return pathCodeList{defs: getArticleDefs(article.Code(), articlesModel), list: articleDist }
}

func mergeVert(articlesModel []types.IArticleSpec, edgeModel []articleEdge, agr []types.IArticleDefine, x types.IArticleDefine) []types.IArticleDefine {
	resultInit := make([]types.IArticleDefine, 0)
	for _, e := range edgeModel {
		if e.stops == x.Term() {
			resultInit = append(resultInit, getArticleDefs(e.start.Code(), articlesModel))
		}
	}
	resultList := toDefsList(resultInit)
	for _, d := range resultInit {
		resultList = mergeVert(articlesModel, edgeModel, resultList, d)
	}
	return append(append(agr, resultList...), x)
}

func defineCompare(topoOrders []types.ArticleTerm, x types.IArticleDefine, y types.IArticleDefine) int {
	var codeOrders []types.ArticleCode = make([]types.ArticleCode, 0, len(topoOrders))
	for _, t := range topoOrders {
		codeOrders = append(codeOrders, t.Code())
	}
	var xIndex = -1
	for p, v := range codeOrders {
		if v == x.Code() {
			xIndex = p
			break
		}
	}
	var yIndex = -1
	for p, v := range codeOrders {
		if v == y.Code() {
			yIndex = p
			break
		}
	}

	if xIndex == -1 && yIndex == -1 {
		return 0
	}

	if xIndex == -1 && yIndex != -1 {
		return -1
	}

	if xIndex != -1 && yIndex == -1 {
		return 1
	}
	return cmpInt(xIndex, yIndex)
}

