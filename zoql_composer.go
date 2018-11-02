package zuora

import (
	"fmt"
	"strings"
)

type ZoqlComposer struct {
	Table     string
	Fields    []string
	Filter    QueryFilter
	OrFilter  []QueryFilter
	AndFilter []QueryFilter
}

type QueryFilter struct {
	Key   string
	Value string
}

type ZoqlComposerOption func(*ZoqlComposer)

func NewZoqlComposer(table string, fields []string, zoqlComposerOption ...ZoqlComposerOption) *ZoqlComposer {
	//TODO: Validate table and fields
	zoqlComposer := &ZoqlComposer{
		Table:  table,
		Fields: fields,
	}

	for _, option := range zoqlComposerOption {
		option(zoqlComposer)
	}

	return zoqlComposer
}

func QueryWithFilter(filter QueryFilter) ZoqlComposerOption {
	return func(zoqlComposer *ZoqlComposer) {
		zoqlComposer.Filter = filter
	}
}

func QueryWithOrFilter(orFilter []QueryFilter) ZoqlComposerOption {
	return func(zoqlComposer *ZoqlComposer) {
		zoqlComposer.OrFilter = orFilter
	}
}

func QueryWithAndFilter(andFilter []QueryFilter) ZoqlComposerOption {
	return func(zoqlComposer *ZoqlComposer) {
		zoqlComposer.AndFilter = andFilter
	}
}

func (q *ZoqlComposer) Build() string {
	if len(q.AndFilter) > 0 || len(q.OrFilter) > 0 || q.Filter.Key != "" {
		andFilter := buildFilter(q.AndFilter, " and ")
		orFilter := buildFilter(q.OrFilter, " or ")
		singleFilter := fmt.Sprintf("%v = '%v'", q.Filter.Key, q.Filter.Value)
		combinedFilter := fmt.Sprintf("%v %v %v", singleFilter, andFilter, orFilter)
		return fmt.Sprintf(`{ "queryString" : "select %v from %v where %v" }`, strings.Join(q.Fields, ", "), q.Table, combinedFilter)
	}

	return fmt.Sprintf(`{ "queryString" : "select %v from %v" }`, strings.Join(q.Fields, ", "), q.Table)
}

func buildFilter(filters []QueryFilter, separator string) string {
	if len(filters) == 0 {
		return ""
	}

	var stringBuilder []string
	for _, filter := range filters {
		stringBuilder = append(stringBuilder, fmt.Sprintf("%v = '%v'", filter.Key, filter.Value))
	}

	return fmt.Sprintf("%v %v", separator, strings.Join(stringBuilder, separator))
}

func (q *ZoqlComposer) String() string {
	return q.Build()
}
