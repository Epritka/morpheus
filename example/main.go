package example

import (
	"context"
	"fmt"

	"github.com/Epritka/morpheus/builder/entity"
	"github.com/Epritka/morpheus/config"
	"github.com/Epritka/morpheus/ogm"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/dbtype"
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

	executer := ogm.NewExecuter()

	// Создание автора
	executer.
		Create(entity.
			NewNode("u").
			SetLables("User").SetProperties(
			map[string]any{
				"firstName":  "%s",
				"lastName":   "%s",
				"patronymic": "%s",
				"job":        "%s",
				"email":      "%s",
				"password":   "%s",
			}),
		)

	// Создание статьи
	executer.
		Create(entity.
			NewNode("a").
			SetLables("Article").SetProperties(
			map[string]any{
				"title":             "%s",
				"subject":           "%s",
				"keywords":          "%s",
				"annotation":        "%s",
				"yearOfPublication": "%s",
				"sourceLink":        "%s",
			}),
		)

	// Создание статьи
	executer.
		Match(
			entity.NewNode("a").SetLables("Article"),
			entity.NewNode("u").SetLables("User"),
		).
		Where(fmt.Sprintf("ID(u) = %d AND ID(a) = %d", 1, 1)).
		Create(entity.
			NewPatternList(entity.NewNode("u")).
			Related(entity.NewRelationship("").SetLables("WROTE")).
			To(entity.NewNode("a")),
		)

	// Поиск автора по id
	executer.
		Match(entity.NewNode("u").SetLables("User")).
		Where(fmt.Sprintf("ID(u) = %d", 1)).
		Return(entity.NewAlias("u"))

	// Поиск автора по email
	executer.
		Match(entity.NewNode("u").SetLables("User")).
		Where(fmt.Sprintf("u.email = %s", "")).
		Return(entity.NewAlias("u"))

	// Поиск статей по названию
	executer.
		Match(entity.NewNode("a").SetLables("Article")).
		Where(fmt.Sprintf("u.title CONTAINS %s", "")).
		Return(entity.NewAlias("a"))

	// Поиск статей по ФИО автора
	part1 := fmt.Sprintf("u.firstName CONTAINS \"%s\"", "")
	part2 := fmt.Sprintf("u.lastName CONTAINS \"%s\"", "")
	part3 := fmt.Sprintf("u.patronymic CONTAINS \"%s\"", "")

	executer.
		Match(entity.
			NewPatternList(entity.NewNode("u").SetLables("User")).
			Related(entity.NewRelationship("").SetLables("AUTHORSHIP")).
			To(entity.NewNode("a").SetLables("Article")),
		).
		Where(fmt.Sprintf("%s OR %s OR %s", part1, part2, part3)).
		Return(entity.NewAlias("a"))

	// Поиск статей по ключевым словам
	executer.
		Match(entity.NewNode("a").SetLables("Article")).
		Where(fmt.Sprintf("ANY(keyword IN a.keywords WHERE keyword CONTAINS \"%s\")", "")).
		Return(entity.NewAlias("a"))

	// Поиск статей по ключевым словам
	executer.
		Match(entity.NewNode("a").SetLables("Article")).
		Where(fmt.Sprintf("a.subject = \"%s\"", "")).
		Return(entity.NewAlias("a"))

	// Поиск статей по году публикации
	executer.
		Match(entity.NewNode("a").SetLables("Article")).
		Where(fmt.Sprintf("a.yearOfPublication = %d", 1)).
		Return(entity.NewAlias("a"))

	records, err := executer.Do()
	if err != nil {
		return err
	}

	for _, record := range records {
		for _, value := range record.Values {
			// тк как ты почти везде возвращаешь либо массив узлов
			// либо узел, кастуй всегда так
			node := value.(dbtype.Node)
			fmt.Println(node.Props)
		}
	}

	return nil
}
