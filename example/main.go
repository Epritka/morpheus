package example

import (
	"fmt"

	"github.com/Epritka/morpheus/config"
	"github.com/Epritka/morpheus/ogm"
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
	_, err = executer.DoQuery(query)
	if err != nil {
		return err
	}

	return nil
}
