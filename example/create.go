package example

import (
	"context"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func Create() error {
	host := "localhost"
	port := 7687
	username := "neo4j"
	password := "password"
	driver, err := neo4j.NewDriverWithContext(
		fmt.Sprintf("neo4j://%s:%d", host, port),
		neo4j.BasicAuth(
			username,
			password,
			"",
		))

	if err != nil {
		return err
	}

	session := driver.NewSession(context.Background(), neo4j.SessionConfig{})
	tx, err := session.BeginTransaction(context.Background())

	if err != nil {
		return err
	}

	query := "CREATE (a:actor {name: 'name'})"
	result, err := tx.Run(context.Background(), query, map[string]any{})
	if err != nil {
		return err
	}

	if result.Err() != nil {
		return result.Err()
	}

	return tx.Commit(context.Background())
}
