package example

import (
	"github.com/kr/pretty"
	legalios "github.com/mzdyhrave/legaliosgo"
	procezor "github.com/mzdyhrave/procezorgo"
	"github.com/mzdyhrave/procezorgo/internal/types"
	"testing"
)

func TestServiceProcezorExampleWithSalaryHomeOffice(t *testing.T) {
	var testPeriod = legalios.NewPeriodWithYearMonth(2021, 1)
	var testVersion = types.GetVersionCode(TEST_VERSION)

	var service procezor.IProcezorService = NewExampleService()

	var gotVersion = service.Version().Value()
	var expVersion int32 = 100
	if gotVersion != expVersion {
		t.Errorf("Error getting version from service expected: %d; got %d",
			expVersion, gotVersion)
	}

	var gotPeriod = testPeriod.GetCode()
	var expPeriod int32 = 202101
	if gotPeriod != expPeriod {
		t.Errorf("Error getting period from service expected: %d; got %d",
			expPeriod, gotPeriod)
	}
	var testLegal legalios.IBundleProps = legalios.EmptyBundleProps(testPeriod)

	var factoryArticleCode = procezor.GetArticleCode(ARTICLE_TIMESHT_WORKING.Id())
	var factoryConceptCode = procezor.GetConceptCode(CONCEPT_TIMESHT_WORKING.Id())

	var factoryArticle = service.GetArticleSpec(factoryArticleCode, testPeriod, testVersion)
	if factoryArticle == nil {
		t.Errorf("Error getting article from service expected: %d; got nil",
			ARTICLE_TIMESHT_WORKING)
	}
	if factoryArticle != nil && factoryArticle.Code().Value() != CONCEPT_TIMESHT_WORKING.Id() {
		t.Errorf("Error getting article from service expected: %d; got: %d",
			CONCEPT_TIMESHT_WORKING, factoryArticle.Code().Value())
	}
	var factoryConcept = service.GetConceptSpec(factoryConceptCode, testPeriod, testVersion)
	if factoryConcept == nil {
		t.Errorf("Error getting concept from service expected: %d; got nil",
			CONCEPT_TIMESHT_WORKING)
	}
	if factoryConcept != nil && factoryConcept.Code().Value() != CONCEPT_TIMESHT_WORKING.Id() {
		t.Errorf("Error getting concept from service expected: %d; got: %d",
			CONCEPT_TIMESHT_WORKING, factoryConcept.Code().Value())
	}
	var initService = service.InitWithPeriod(testPeriod)
	if initService == false {
		t.Errorf("Error initializating service got false")
	}
	var restTargets = GetTargetsWithSalaryHomeOffice(testPeriod)
	var restService = service.GetResults(testPeriod, testLegal, restTargets)
	if len(restService) == 0 {
		t.Errorf("Error gerring result from service got empty list")
	}
	var restArticles = make([]int32, 0)
	for _, res := range restService {
		if res.IsSuccess() {
			restArticles = append(restArticles, res.Value().Article().Value())
		}
	}
	for index, res := range restService {
		if res.IsSuccess() {
			resultValue := res.Value()
			articleSymbol := resultValue.ArticleDescr()
			conceptSymbol := resultValue.ConceptDescr()
			t.Logf("Index: %d, ART: %v, CON: %v", index, articleSymbol, conceptSymbol)
		}
		if res.IsFailure() {
			errorsValue := res.Error()
			resultValue := res.ResultError()
			articleSymbol := resultValue.ArticleDescr()
			conceptSymbol := resultValue.ConceptDescr()
			t.Logf("Index: %d, ART: %v, CON: %v, Error: %v", index, articleSymbol, conceptSymbol, errorsValue)
		}
	}

	var testArticles = []int32 { 80001, 80005, 80002, 80006, 80007, 80010, 80012, 80008, 80009, 80011, 80013 }
	var articlesDiff = false
	if len(restArticles) != len(testArticles) {
		t.Errorf("Error gerring result from service result len don't match, expected %d; got: %d",
			len(testArticles), len(restArticles))
	}
	if len(restArticles) == len(testArticles) {
		for i, v := range restArticles {
			if v != testArticles[i] {
				articlesDiff = true
			}
		}
	}
	if articlesDiff {
		t.Errorf("Error gerring result from service result article code don't match,\n expected %s;\n got: %s\n",
			pretty.Sprint(testArticles), pretty.Sprint(restArticles))
	}
}

