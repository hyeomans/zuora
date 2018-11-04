package zuora

import (
	"bytes"
	"fmt"
	"strings"
)

//ZoqlComposer helper struct to build a zoql query
type ZoqlComposer struct {
	Table     string
	Fields    []string
	Filter    QueryFilter
	OrFilter  []QueryFilter
	AndFilter []QueryFilter
}

//QueryFilter key/value combination that represent filters.
type QueryFilter struct {
	Key   string
	Value string
}

//ZoqlComposerOption using functional options to construct a query
type ZoqlComposerOption func(*ZoqlComposer)

//NewZoqlComposer helper function to get a ready ZoqlComposer struct
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

//QueryWithFilter add a single filter to query
func QueryWithFilter(filter QueryFilter) ZoqlComposerOption {
	return func(zoqlComposer *ZoqlComposer) {
		zoqlComposer.Filter = filter
	}
}

//QueryWithOrFilter adds an OR filter to query
func QueryWithOrFilter(orFilter []QueryFilter) ZoqlComposerOption {
	return func(zoqlComposer *ZoqlComposer) {
		zoqlComposer.OrFilter = orFilter
	}
}

//QueryWithAndFilter adds an AND filter to final query
func QueryWithAndFilter(andFilter []QueryFilter) ZoqlComposerOption {
	return func(zoqlComposer *ZoqlComposer) {
		zoqlComposer.AndFilter = andFilter
	}
}

//Build going away
func (q *ZoqlComposer) Build() string {
	var buffer bytes.Buffer
	buffer.WriteString(`{ "queryString" : "select %v from %v`)

	if q.Filter.Key != "" {
		buffer.WriteString(fmt.Sprintf(" where %v = '%v'", q.Filter.Key, q.Filter.Value))
	}

	if len(q.AndFilter) > 0 {
		andFilter := buildFilter(q.AndFilter, "and")
		buffer.WriteString(andFilter)
	}

	if len(q.OrFilter) > 0 {
		orFilter := buildFilter(q.OrFilter, "or")
		buffer.WriteString(orFilter)
	}

	buffer.WriteString(`" }`)
	return fmt.Sprintf(buffer.String(), strings.Join(q.Fields, ", "), q.Table)
}

func buildFilter(filters []QueryFilter, separator string) string {
	var buffer bytes.Buffer
	for _, filter := range filters {
		buffer.WriteString(fmt.Sprintf(" %v %v = '%v'", separator, filter.Key, filter.Value))
	}

	return buffer.String()
}

func (q *ZoqlComposer) String() string {
	return q.Build()
}
