package fema

import (
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"
)

type QueryOptions struct {
	// OrderByFields specifies which fields are important
	//
	// e.g., $orderby=state desc, declaredCountyArea
	OrderByFields string

	// MaxCount determines the number of returned results == "$top"
	MaxCount *int

	// InitialOffset determines the number of 'skipped' values from the query == "$skip"
	InitialOffset *int

	// SelectedFields specifies which fields should be returned in the query
	SelectedFields string

	// Filters determine logical operations on the query to get specific data.
	//
	// e.g. $filter=declarationDate ge '2010-01-01T04:00:00.000z' and state eq 'VA'
	Filters string
}

func newDefaultQueryOptions() *QueryOptions {
	return &QueryOptions{}
}

func (queryOpt QueryOptions) getURI() string {
	// generate our string builder that we'll use to store our variables
	var builder strings.Builder

	// -----
	// Populate the query
	// -----

	if queryOpt.OrderByFields != "" {
		builder.WriteString(fmt.Sprintf("$orderby=%s", url.QueryEscape(queryOpt.OrderByFields)))
	} else {
		builder.WriteString(fmt.Sprintf("$orderby=%s", url.QueryEscape("declarationDate desc")))
	}

	// Write out our max results count if specified
	if queryOpt.MaxCount != nil {
		builder.WriteString(fmt.Sprintf("&$top=%d", *queryOpt.MaxCount))
	}

	if queryOpt.Filters != "" {
		builder.WriteString(fmt.Sprintf("&$filter=%s", queryOpt.Filters))
	}

	// TODO: Add other query option parsing here!
	log.Println(builder.String())

	return builder.String()
}

type QueryFunc func(*QueryOptions)

func WithMaxCount(count int) QueryFunc {
	return func(q *QueryOptions) {
		q.MaxCount = &count
	}
}

func WithStateFilter(state string) QueryFunc {
	return func (q *QueryOptions) {
		q.addFilterAND()
		q.Filters += url.QueryEscape(fmt.Sprintf("state eq '%s'", state))
	}
}

func getDateFilterString(year, month int) string {
	// base message that needs to be URL escaped
	baseMsg := url.QueryEscape("declarationDate gt %s")
	date := fmt.Sprintf("'%d-%d-01T00:00:01.000Z'", year, month)
	return fmt.Sprintf(baseMsg, date)
}

func WithAfterDateFilter(year, month int) QueryFunc {
	return func (q *QueryOptions) {
		q.addFilterAND()
		q.Filters += getDateFilterString(year, month)
	}
}

func WithCurrentMonthFilter() QueryFunc {
	return func (q *QueryOptions) {
		q.addFilterAND()
		now := time.Now()
		year, month, _ := now.Date()
		q.Filters += getDateFilterString(year, int(month))
	}
}

// addFilterAND will only add a URL escapted ' and ' string if there's something in the string
func (q *QueryOptions) addFilterAND() {
	if q.Filters != "" {
		q.Filters += url.QueryEscape(" and ")
	}
}

func resolveQueryOptions(options []QueryFunc) string {
	// Create a new default query option
	opts := newDefaultQueryOptions()
	// iterate  over the given funcs to mutate the default queries
	for _, optFunc := range options {
		optFunc(opts)
	}

	// from the mutated query options, generate the URI
	// and return that from this function
	query := opts.getURI()
	return query
}
