package example

import (
	"context"
	"fmt"

	"github.com/Epritka/morpheus/config"
	"github.com/Epritka/morpheus/ogm"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

const (
	host     = "localhost"
	port     = 7687
	username = "neo4j"
	password = "password"
)

func DefaultCreateActor() error {
	driver, err := neo4j.NewDriverWithContext(
		fmt.Sprintf("neo4j://%s:%d", host, port),
		neo4j.BasicAuth(
			username,
			password,
			"",
		),
	)
	if err != nil {
		return err
	}

	session := driver.NewSession(context.Background(), neo4j.SessionConfig{})

	executer, err := session.BeginTransaction(context.Background())
	if err != nil {
		return err
	}

	result, err := executer.Run(
		context.Background(),
		"CREATE (n:person { name: \"Keanu Reeves\",job: \"actor\" }) RETURN n",
		map[string]any{},
	)
	if err != nil {
		return err
	}

	err = executer.Commit(context.Background())
	if err != nil {
		return err
	}

	if result.Next(context.Background()) {
		fmt.Println("not empty")
	} else {
		fmt.Println("empty")
	}
	return nil
}

func Test() error {
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

	executer := ogm.NewExecuter()

	// Создание автора
	executer.DoQuery("CREATE (n:person { name: \"Keanu Reeves\",job: \"actor\" }) RETURN n")

	return nil
}
