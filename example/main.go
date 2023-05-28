package example

import (
	"fmt"

	"github.com/Epritka/morpheus/config"
	"github.com/Epritka/morpheus/ogm"
)

func Create() error {
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

	session := ogm.NewSession()

	session.Save(map[string]any{})
	// tx, err := session.Begin()
	// if err != nil {
	// 	return err
	// }

	// query := "CREATE (a:actor {name: 'name'})"
	err = session.Do()
	if err != nil {
		return err
	}

	// return tx.Commit()
	return nil
}
