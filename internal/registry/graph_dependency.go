package registry

import (
	providers "github.com/mzdyhrave/payrollgo-procezor/internal/registry_providers"
	"github.com/mzdyhrave/payrollgo-procezor/internal/types"
	"sort"
)

func InitGraphModel(articlesModel []providers.IArticleSpec, conceptsModel []providers.IConceptSpec) ([]types.ArticleCode, pathCodeMap, bool) {
	vertModel := createVertModel(articlesModel)
	edgeModel := createEdgeModel(articlesModel, conceptsModel)

	order := createTopoModel(vertModel, edgeModel)

	paths := createPathModel(articlesModel, vertModel, edgeModel, order)

	return order, paths, true
}

func createVertModel(articlesModel []providers.IArticleSpec) []types.ArticleCode {
	result := make([]types.ArticleCode, 0, len(articlesModel))
	for _, article := range articlesModel {
		result = append(result, article.Code())
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Value() < result[j].Value()
	})
	return result
}

func createEdgeModel(articlesModel []providers.IArticleSpec, conceptsModel []providers.IConceptSpec) []articleEdge {
	result := articleEdgeSet{}
	for _, s := range articlesModel {
		result = mergeEdges(conceptsModel, result, s)
	}
	resultSlice := result.Sort()
	return resultSlice
}

func mergeEdges(conceptsModel []providers.IConceptSpec, init articleEdgeSet, article providers.IArticleSpec) articleEdgeSet {
	result := toEdgeHashSet(init)

	concept, _ := firstOrDefaultConceptSpec(conceptsModel, findConceptSpecCode(article.Role().Value()))

	for _, s := range article.Sums() {
		result.Add(newArticleEdge(article.Code(), s))
	}
	for _, p := range concept.Path() {
		result.Add(newArticleEdge(p, article.Code()))
	}
	return result
}

func createTopoModel(vertModel []types.ArticleCode, edgeModel []articleEdge) []types.ArticleCode {
	articlesOrder := make([]types.ArticleCode, 0, len(vertModel))

	degrees := make(map[types.ArticleCode]int32)
	for _, v := range vertModel {
		var count int32 = 0
		for _, e := range edgeModel {
			if e.stops == v {
				count++
			}
		}
		degrees[v] = count
	}
	zeroDegrees := make([]types.ArticleCode, 0)
	for d, v := range degrees {
		if v == 0 {
			zeroDegrees = append(zeroDegrees, d)
		}
	}
	sort.Slice(zeroDegrees, func(i, j int) bool {
		return zeroDegrees[i].Value() < zeroDegrees[j].Value()
	})

	queues := articleCodeQueue{}
	for _, v := range zeroDegrees {
		queues.Enqueue(v)
	}
	var index = 0
	for queues.Size() != 0 {
		index++
		article, err := queues.Dequeue()
		if err != nil {
			return make([]types.ArticleCode, 0)
		}
		articlesOrder = append(articlesOrder, article)
		paths := make([]types.ArticleCode, 0)
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
		return make([]types.ArticleCode, 0)
	}
	return articlesOrder
}

func createPathModel(articlesModel []providers.IArticleSpec, vertModel []types.ArticleCode, edgeModel []articleEdge, articlesOrder []types.ArticleCode) pathCodeMap {
	result := make(pathCodeMap)

	for _, v := range vertModel {
		result[v] = mergePaths(articlesModel, edgeModel, articlesOrder, v)
	}
	return result
}

func mergePaths(articlesModel []providers.IArticleSpec, edgeModel []articleEdge, articlesOrder []types.ArticleCode, article types.ArticleCode) pathCodeList {
	articleInit := make([]types.IArticleDefine, 0)
	for _, e := range edgeModel {
		if e.stops == article {
			articleInit = append(articleInit, getArticleDefs(e.start, articlesModel))
		}
	}
	articlePath := toDefsList(articleInit)
	for _, a := range articleInit {
		articlePath = mergeVert(articlesModel, edgeModel, articlePath, a)
	}

	articleDist := distinctDefsList(articlePath)

	sort.Slice(articleDist, func(i, j int) bool {
		return defineCompare(articlesOrder, articleDist[i], articleDist[j]) < 0
	})
	return pathCodeList{defs: getArticleDefs(article, articlesModel), list: articleDist }
}

func mergeVert(articlesModel []providers.IArticleSpec, edgeModel []articleEdge, agr []types.IArticleDefine, x types.IArticleDefine) []types.IArticleDefine {
	resultInit := make([]types.IArticleDefine, 0)
	for _, e := range edgeModel {
		if e.stops == x.Code() {
			resultInit = append(resultInit, getArticleDefs(e.start, articlesModel))
		}
	}
	resultList := toDefsList(resultInit)
	for _, d := range resultInit {
		resultList = mergeVert(articlesModel, edgeModel, resultList, d)
	}
	return append(append(agr, resultList...), x)
}

func defineCompare(topoOrders []types.ArticleCode, x types.IArticleDefine, y types.IArticleDefine) int {
	var xIndex = -1
	for p, v := range topoOrders {
		if v == x.Code() {
			xIndex = p
			break
		}
	}
	var yIndex = -1
	for p, v := range topoOrders {
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

