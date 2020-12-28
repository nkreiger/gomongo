package mongo

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

// GetEvents executes a Find query with query parameters that don't include operations
func GetEvents(params ExecuteParams, queryInput QueryInputs) (interface{}, error) {
	sort := queryInput.buildOrderedSort()

	projection := queryInput.buildProjection()

	var match interface{}

	if queryInput.Params == nil {
		match = queryInput.buildParamsWithOperations()
	} else {
		match = queryInput.buildParams()
	}

	opts := options.Find()

	if sort != nil {
		opts.SetSort(sort)
	}

	if projection != nil {
		opts.SetProjection(projection)
	}

	if queryInput.Limit != 0 {
		opts.SetLimit(queryInput.Limit)
	}

	log.Printf("query: %v", match)

	c := params.Database.Collection(params.Collection)

	cursor, err := c.Find(params.Context, match, opts)
	if err != nil {
		return nil, fmt.Errorf("GetAll: error getting all events from db: %v", err)
	}

	var results []bson.M

	err = cursor.All(params.Context, &results)

	return results, err
}
