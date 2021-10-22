package registry

import (
	"fmt"
	"github.com/mzdyhrave/payrollgo-procezor/internal/types"
	"sort"
)

type pathCodeList struct {
	defs types.IArticleDefine
	list []types.IArticleDefine
}

type pathCodeMap map[types.ArticleCode]pathCodeList

type articleEdge struct {
	start types.ArticleCode
	stops types.ArticleCode
}

func newArticleEdge(codeStart types.ArticleCode, codeStops types.ArticleCode) articleEdge {
	return 	articleEdge{ start: codeStart, stops: codeStops }
}

type mapEdge map[articleEdge]bool

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
		if itemsSlice[i].start.Value() == itemsSlice[j].start.Value() {
			return itemsSlice[i].stops.Value() < itemsSlice[j].stops.Value()
		}
		return itemsSlice[i].start.Value() < itemsSlice[j].start.Value()
	})
	return itemsSlice
}

type articleCodeQueue struct {
	queue []types.ArticleCode
}

func (c *articleCodeQueue) Enqueue(article types.ArticleCode) {
	c.queue = append(c.queue, article)
}

func (c *articleCodeQueue) Dequeue() (types.ArticleCode, error) {
	if len(c.queue) > 0 {
		code := c.queue[0]
		c.queue = c.queue[1:]
		return code, nil
	}
	return types.NewArticleCode(), fmt.Errorf("pop Error: Queue is empty")
}

func (c *articleCodeQueue) Front() (types.ArticleCode, error) {
	if len(c.queue) > 0 {
		return c.queue[0], nil
	}
	return types.NewArticleCode(), fmt.Errorf("peep Error: Queue is empty")
}

func (c *articleCodeQueue) Size() int {
	return len(c.queue)
}

func (c *articleCodeQueue) Empty() bool {
	return len(c.queue) == 0
}

