package mongo

import (
	"context"
	mgo "go.mongodb.org/mongo-driver/mongo"
)
// ExecuteParams defines the necessary values to execute a
// mongo query
type ExecuteParams struct {
	Context 	context.Context
	Database 	*mgo.Database
	Collection 	string
}


// QueryInputs handles the inputs that would go into a basic and advanced MQL
type QueryInputs struct {
	Limit					int64 `json:"limit"`
	Projection				[]string `json:"projection"`
	Sort					BasicParams `json:"sort"`
	Params					BasicParams `json:"params"`
	Update					map[string]BasicParams `json:"update"`
	ParamsWithOperations	map[string][]BasicParams `json:"paramsWithOperations"`
}

// BasicParams is a subset of query inputs
type BasicParams map[string]interface{}