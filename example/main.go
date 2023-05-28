package example

import (
	"fmt"

	"github.com/Epritka/morpheus/v2/config"
	"github.com/Epritka/morpheus/v2/ogm"
)

func CreateActor() error {
	host := "localhost"
	port := 7687
	username := "neo4j"
	password := "password"

	ogm, err := ogm.Connect(
		&config.Config{
			URI:      fmt.Sprintf("neo4j://%s:%d", host, port),
			Username: username,
			Password: password,
		},
	)
	if err != nil {
		return err
	}

	query := "CREATE (a:actor {name: 'test'})"
	executer := ogm.NewExecuter()
	tx, err := executer.Begin()

	err = tx.DoQuery(query)
	if err != nil {
		return err
	}

	return tx.Commit()
}
