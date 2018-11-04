package zuora_test

import (
	"fmt"
	"testing"

	"github.com/hyeomans/zuora"
)

func TestComposerWithNoFilters(t *testing.T) {
	t.Parallel()
	expectedResult := `{ "queryString" : "select ID, Name from Product" }`

	fields := []string{"ID", "Name"}
	zoqlComposer := zuora.NewZoqlComposer("Product", fields)
	result := fmt.Sprint(zoqlComposer)

	if result != expectedResult {
		t.Fatalf("Zoql Composer failed to build a simple query. Expected: `%v` but got `%v`", expectedResult, result)
	}
}

func TestComposerWithSingleFilter(t *testing.T) {
	t.Parallel()
	expectedResult := `{ "queryString" : "select ID, Name from Product where Name = 'productName'" }`

	fields := []string{"ID", "Name"}
	queryFilter := zuora.QueryFilter{Key: "Name", Value: "productName"}
	singleFilter := zuora.QueryWithFilter(queryFilter)
	zoqlComposer := zuora.NewZoqlComposer("Product", fields)
	singleFilter(zoqlComposer)
	result := fmt.Sprint(zoqlComposer)

	if result != expectedResult {
		t.Fatalf("Zoql Composer failed to build a simple query. Expected: `%v` but got `%v`", expectedResult, result)
	}
}

func TestComposerWithOrFilter(t *testing.T) {
	t.Parallel()
	expectedResult := `{ "queryString" : "select ID, Name from Product where Name = 'productName' or Name = 'productName'" }`

	fields := []string{"ID", "Name"}
	singleFilter := zuora.QueryFilter{Key: "Name", Value: "productName"}
	queryFilter := []zuora.QueryFilter{singleFilter}
	addSingleFilter := zuora.QueryWithFilter(singleFilter)
	addOrFilter := zuora.QueryWithOrFilter(queryFilter)

	zoqlComposer := zuora.NewZoqlComposer("Product", fields)
	addSingleFilter(zoqlComposer)
	addOrFilter(zoqlComposer)
	result := fmt.Sprint(zoqlComposer)

	if result != expectedResult {
		t.Fatalf("Zoql Composer failed to build a simple query. Expected: `%v` but got `%v`", expectedResult, result)
	}
}

func TestComposerWithAndFilter(t *testing.T) {
	t.Parallel()
	expectedResult := `{ "queryString" : "select ID, Name from Product where Name = 'productName' and Name = 'productName'" }`

	fields := []string{"ID", "Name"}
	singleFilter := zuora.QueryFilter{Key: "Name", Value: "productName"}
	queryFilter := []zuora.QueryFilter{singleFilter}
	addSingleFilter := zuora.QueryWithFilter(singleFilter)
	addAndFilter := zuora.QueryWithAndFilter(queryFilter)

	zoqlComposer := zuora.NewZoqlComposer("Product", fields)
	addSingleFilter(zoqlComposer)
	addAndFilter(zoqlComposer)
	result := fmt.Sprint(zoqlComposer)

	if result != expectedResult {
		t.Fatalf("Zoql Composer failed to build a simple query. Expected: `%v` but got `%v`", expectedResult, result)
	}
}

func TestComposerWithAllFilters(t *testing.T) {
	t.Parallel()
	expectedResult := `{ "queryString" : "select ID, Name from Product where Name = 'productName' and Name = 'productName' and Name = 'productName' and Name = 'productName' and Name = 'productName' or Name = 'productName' or Name = 'productName' or Name = 'productName' or Name = 'productName'" }`

	fields := []string{"ID", "Name"}
	singleFilter := zuora.QueryFilter{Key: "Name", Value: "productName"}
	filters := []zuora.QueryFilter{singleFilter, singleFilter, singleFilter, singleFilter}
	addSingleFilter := zuora.QueryWithFilter(singleFilter)
	addAndFilter := zuora.QueryWithAndFilter(filters)
	addOrFilter := zuora.QueryWithOrFilter(filters)
	zoqlComposer := zuora.NewZoqlComposer("Product", fields)
	addSingleFilter(zoqlComposer)
	addAndFilter(zoqlComposer)
	addOrFilter(zoqlComposer)

	result := fmt.Sprint(zoqlComposer)

	if result != expectedResult {
		t.Fatalf("Zoql Composer failed to build a simple query. Expected: `%v` but got `%v`", expectedResult, result)
	}
}
