package zuora

import (
	"bytes"
	"fmt"
	"strings"
)

type ZoqlComposer struct {
	table  string
	fields []string
	where  string
	and    string
	or     string
}

func NewZoqlComposer() ZoqlComposer {
	return ZoqlComposer{}
}

func (z ZoqlComposer) Fields(fields ...string) ZoqlComposer {
	z.fields = fields
	return z
}

func (z ZoqlComposer) From(table string) ZoqlComposer {
	if len(z.fields) == 0 {
		//TODO: An error here
	}

	z.table = table
	return z
}

func (z ZoqlComposer) Where(key, value string) ZoqlComposer {
	if z.table == "" {
		//TODO: An error here
	}

	z.where = fmt.Sprintf(` where %v = '%v'`, key, value)
	return z
}

func (z ZoqlComposer) And(key, value string) ZoqlComposer {
	if z.where == "" {
		//TODO: An error here
	}

	if z.and != "" {
		z.and = fmt.Sprintf("%v and %v = '%v'", z.and, key, value)
	} else {
		z.and = fmt.Sprintf(" and %v = '%v'", key, value)
	}

	return z
}

func (z ZoqlComposer) Or(key, value string) ZoqlComposer {
	if z.where == "" {
		//TODO: An error here
	}

	if z.or != "" {
		z.or = fmt.Sprintf("%v or %v = '%v'", z.and, key, value)
	} else {
		z.or = fmt.Sprintf(" or %v = '%v'", key, value)
	}

	return z
}

func (z ZoqlComposer) Build() string {
	//TODO: Check minimum strings
	var buffer bytes.Buffer
	buffer.WriteString(`{ "queryString" : "select %v from %v`)

	if z.where != "" {
		buffer.WriteString(z.where)
	}

	if z.and != "" {
		buffer.WriteString(z.and)
	}

	if z.or != "" {
		buffer.WriteString(z.or)
	}

	buffer.WriteString(`" }`)
	return fmt.Sprintf(buffer.String(), strings.Join(z.fields, ", "), z.table)
}
