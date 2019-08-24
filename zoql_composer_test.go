package zuora_test

import (
	"testing"

	"github.com/hyeomans/zuora"
)

func TestComposerWithNoFilters(t *testing.T) {
	t.Parallel()
	expectedResult := `{ "queryString" : "select ID, Name from Product" }`

	result := zuora.NewZoqlComposer().
		Fields("ID", "Name").
		From("Product").
		Build()

	if result != expectedResult {
		t.Fatalf("Zoql Composer failed to build a simple query. Expected: `%v` but got `%v`", expectedResult, result)
	}
}

func TestComposerWithSingleFilter(t *testing.T) {
	t.Parallel()
	expectedResult := `{ "queryString" : "select ID, Name from Product where Name = 'productName'" }`

	result := zuora.NewZoqlComposer().Fields("ID", "Name").From("Product").Where("Name", "productName").Build()

	if result != expectedResult {
		t.Fatalf("Zoql Composer failed to build a simple query. Expected: `%v` but got `%v`", expectedResult, result)
	}
}

func TestComposerWithOrFilter(t *testing.T) {
	t.Parallel()
	expectedResult := `{ "queryString" : "select ID, Name from Product where Name = 'productName' or Name = 'productName'" }`

	result := zuora.NewZoqlComposer().Fields("ID", "Name").From("Product").Where("Name", "productName").Or("Name", "productName").Build()

	if result != expectedResult {
		t.Fatalf("Zoql Composer failed to build a simple query. Expected: `%v` but got `%v`", expectedResult, result)
	}
}

func TestComposerWithAndFilter(t *testing.T) {
	t.Parallel()
	expectedResult := `{ "queryString" : "select ID, Name from Product where Name = 'productName' and Name = 'productName'" }`

	result := zuora.NewZoqlComposer().Fields("ID", "Name").From("Product").Where("Name", "productName").And("Name", "productName").Build()

	if result != expectedResult {
		t.Fatalf("Zoql Composer failed to build a simple query. Expected: `%v` but got `%v`", expectedResult, result)
	}
}
