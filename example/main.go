package example

import (
	"fmt"

	"github.com/Epritka/morpheus/v1/config"
	"github.com/Epritka/morpheus/v1/ogm"
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

	session := ogm.NewSession()

	err = session.DoQuery("CREATE (a:actor {name: 'name'})")
	if err != nil {
		return err
	}

	return nil
}
