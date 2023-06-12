package example

import (
	"fmt"

	"github.com/Epritka/morpheus/config"
	"github.com/Epritka/morpheus/ogm"
	"github.com/Epritka/morpheus/ogm/types"
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

	p := &Person{
		PersonSchema: &PersonSchema{
			Node: types.NewNode(),
			Name: "Keanu Reeves",
			Job:  "actor",
		},
		Movies: []MovieSchema{
			{
				Node:     types.NewNode(),
				Title:    "The Matrix",
				Released: 1999,
			},
		},
	}

	// query := "CREATE (a:actor {name: 'test'})"
	executer := ogm.NewExecuter()
	// tx, err := executer.Begin()
	// executer.
	// executer.DoQuery(query)
	// tx.Cypher.Create(p)
	err = executer.Save(p)
	if err != nil {
		return err
	}

	// return tx.Commit()
	return nil
}
