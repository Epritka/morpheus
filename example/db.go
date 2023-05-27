package example

import (
	"context"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type DB struct {
	Driver neo4j.DriverWithContext
	tx     neo4j.ExplicitTransaction
}

func Connect(config *Config, configurers ...func(c *neo4j.Config)) (*DB, error) {
	driver, err := neo4j.NewDriverWithContext(
		config.URI,
		neo4j.BasicAuth(
			config.Username,
			config.Password,
			"",
		),
		configurers...,
	)
	if err != nil {
		return nil, err
	}

	return &DB{
		Driver: driver,
	}, nil
}

// func (db *DB) WithContext(ctx context.Context) *DB {
// 	db.ctx = ctx
// 	return db
// }

// func (db *DB) Context() context.Context {
// 	if db.ctx != nil {
// 		return db.ctx
// 	}
// 	return context.Background()
// }

func (db *DB) Close(ctx context.Context) error {
	return db.Driver.Close(ctx)
}
