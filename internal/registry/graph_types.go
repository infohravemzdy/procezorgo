package registry

import (
	"fmt"
	"github.com/mzdyhrave/procezorgo/internal/types"
	"sort"
)

type pathCodeList struct {
	defs types.IArticleDefine
	list []types.IArticleDefine
}

type OrderCodeList = types.ArticleTermList
type PathsTermsMap = map[types.ArticleTerm]pathCodeList

type articleEdge struct {
	start types.ArticleTerm
	stops types.ArticleTerm
}

func newArticleEdge(codeStart types.ArticleTerm, codeStops types.ArticleTerm) articleEdge {
	return 	articleEdge{ start: codeStart, stops: codeStops }
}

type mapEdge = map[articleEdge]bool

// articleEdgeSet the set of ArticleEdge
type articleEdgeSet struct {
	items mapEdge
}

// Add adds a new element to the Set. Returns a pointer to the Set.
func (s *articleEdgeSet) Add(t articleEdge) *articleEdgeSet {
	if s.items == nil {
		s.items = make(mapEdge)
	}
	_, ok := s.items[t]
	if !ok {
		s.items[t] = true
	}
	return s
}

// Clear removes all elements from the Set
func (s *articleEdgeSet) Clear() {
	s.items = make(mapEdge)
}

// Delete removes the Item from the Set and returns Has(Item)
func (s *articleEdgeSet) Delete(item articleEdge) bool {
	_, ok := s.items[item]
	if ok {
		delete(s.items, item)
	}
	return ok
}

// Has returns true if the Set contains the Item
func (s *articleEdgeSet) Has(item articleEdge) bool {
	_, ok := s.items[item]
	return ok
}

// Items returns the Item(s) stored
func (s *articleEdgeSet) Items() []articleEdge {
	var items []articleEdge
	for i := range s.items {
		items = append(items, i)
	}
	return items
}

// Size returns the size of the set
func (s *articleEdgeSet) Size() int {
	return len(s.items)
}

// Sort returns the array of set sorted
func (s *articleEdgeSet) Sort() []articleEdge {
	itemsSlice := make([]articleEdge, 0)
	for k, _ := range s.items {
		itemsSlice = append(itemsSlice, k)
	}
	sort.Slice(itemsSlice, func(i, j int) bool {
		if types.ArticleTermCompare(itemsSlice[i].start, itemsSlice[j].start)==0 {
			return types.ArticleTermCompare(itemsSlice[i].stops, itemsSlice[j].stops) < 0
		}
		return types.ArticleTermCompare(itemsSlice[i].start, itemsSlice[j].start) < 0
	})
	return itemsSlice
}

type articleTermQueue struct {
	queue types.ArticleTermList
}

func (c *articleTermQueue) Enqueue(article types.ArticleTerm) {
	c.queue = append(c.queue, article)
}

func (c *articleTermQueue) Dequeue() (types.ArticleTerm, error) {
	if len(c.queue) > 0 {
		code := c.queue[0]
		c.queue = c.queue[1:]
		return code, nil
	}
	return types.NewArticleTerm(), fmt.Errorf("pop Error: Queue is empty")
}

func (c *articleTermQueue) Front() (types.ArticleTerm, error) {
	if len(c.queue) > 0 {
		return c.queue[0], nil
	}
	return types.NewArticleTerm(), fmt.Errorf("peep Error: Queue is empty")
}

func (c *articleTermQueue) Size() int {
	return len(c.queue)
}

func (c *articleTermQueue) Empty() bool {
	return len(c.queue) == 0
}

