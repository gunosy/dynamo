package dynamo

import (
	"time"

	"github.com/gunosy/aws-sdk-go/aws"
	"github.com/gunosy/aws-sdk-go/service/dynamodb"
)

type DB struct {
	client *dynamodb.DynamoDB
}

const retryTimeout = 1 * time.Minute // TODO: make this configurable

func New(cfg *aws.Config) *DB {
	db := &DB{
		dynamodb.New(cfg),
	}
	return db
}

// Iter is an iterator for query results.
type Iter interface {
	Next(out interface{}) bool
	Err() error
}
