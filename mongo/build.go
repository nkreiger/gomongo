package mongo

import (
"go.mongodb.org/mongo-driver/bson"
"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	eventVersionKey 			= "event.version"
	eventTimestampKey 			= "event.timestamp"
	timestamp					= "timestamp"
	version						= "version"
)

var (
	sortOrderKeys = []string{eventTimestampKey, timestamp, eventVersionKey, version} // set order of sorting
)

// buildProjection builds the projection fields or returns nil if nothing is passed
func (q *QueryInputs) buildProjection() bson.M {
	p := make(bson.M)

	for _, v := range q.Projection {
		p[v] = 1
	}

	return p
}

// buildSort builds the sort fields or returns nil if nothing is passed
func (q *QueryInputs) buildOrderedSort() bson.D {
	var s bson.D

	if len(q.Sort) != 0 {
		s = buildOrderedD(q.Sort, sortOrderKeys)
	}

	return s
}

// buildParams builds the param fields or returns nil if nothing is passed
func (q *QueryInputs) buildParams() bson.D {
	p := bson.D{}

	if len(q.Params) != 0 {
		p = buildD(q.Params)
	}

	return p
}

// buildParamsWithOperations builds the param fields with the operation as key
func(q *QueryInputs) buildParamsWithOperations() bson.M {
	p := make(bson.M)

	for k, v := range q.ParamsWithOperations {
		p[k] = []bson.D{}

		for _, value := range v {
			param := buildD(value)
			p[k] = append(p[k].([]bson.D), param)
		}
	}

	return p
}

// builds ordered type bson.D
func buildOrderedD(values map[string]interface{}, order []string) bson.D {
	var e []primitive.E

	for _, k := range order {

		if _, ok := values[k]; !ok {
			continue
		}

		value := primitive.E{
			Key:   k,
			Value: values[k],
		}

		e = append(e, value)
	}

	return e
}

// builds type bson.D
func buildD(values map[string]interface{}) bson.D {
	var e []primitive.E

	for k, v := range values {
		// skip blank
		if k == "" || v == "" {
			continue
		}

		value := primitive.E{
			Key:   k,
			Value: v,
		}

		e = append(e, value)
	}

	return e
}
