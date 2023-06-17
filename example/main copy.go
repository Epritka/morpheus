package example

// import (
// 	"context"
// 	"fmt"

// 	"github.com/Epritka/morpheus/builder/entity"
// 	"github.com/Epritka/morpheus/config"
// 	"github.com/Epritka/morpheus/ogm"
// 	"github.com/Epritka/morpheus/ogm/types"
// 	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
// )

// const (
// 	host     = "localhost"
// 	port     = 7687
// 	username = "neo4j"
// 	password = "password"
// )

// func DefaultCreateActor() error {
// 	driver, err := neo4j.NewDriverWithContext(
// 		fmt.Sprintf("neo4j://%s:%d", host, port),
// 		neo4j.BasicAuth(
// 			username,
// 			password,
// 			"",
// 		),
// 	)
// 	if err != nil {
// 		return err
// 	}

// 	session := driver.NewSession(context.Background(), neo4j.SessionConfig{})

// 	tx, err := session.BeginTransaction(context.Background())
// 	if err != nil {
// 		return err
// 	}

// 	result, err := tx.Run(
// 		context.Background(),
// 		"CREATE (n:person { name: \"Keanu Reeves\",job: \"actor\" }) RETURN n",
// 		map[string]any{},
// 	)
// 	if err != nil {
// 		return err
// 	}

// 	err = tx.Commit(context.Background())
// 	if err != nil {
// 		return err
// 	}

// 	if result.Next(context.Background()) {
// 		fmt.Println("not empty")
// 	} else {
// 		fmt.Println("empty")
// 	}
// 	return nil
// }

// func Test() error {
// 	ogm, err := ogm.Connect(
// 		&config.Config{
// 			URI:      fmt.Sprintf("neo4j://%s:%d", host, port),
// 			Username: username,
// 			Password: password,
// 		},
// 	)
// 	if err != nil {
// 		return err
// 	}

// 	person := &Person{
// 		Node: types.NewNode(),
// 		Name: "Keanu Reeves",
// 		Job:  "actor",
// 	}

// 	movie := &Movie{
// 		Node:     types.NewNode(),
// 		Title:    "The Matrix",
// 		Released: 1999,
// 	}

// 	fmt.Println("\n-------------------------")
// 	fmt.Println("Person до выполнения запроса")
// 	fmt.Println("-------------------------")
// 	fmt.Println("id: ", person.Id)
// 	fmt.Println("element id: ", person.ElementId)
// 	fmt.Println("name: ", person.Name)
// 	fmt.Println("job: ", person.Job)
// 	fmt.Println("-------------------------\n")

// 	fmt.Println("\n-------------------------")
// 	fmt.Println("Movie до выполнения запроса")
// 	fmt.Println("-------------------------")
// 	fmt.Println("id: ", movie.Id)
// 	fmt.Println("element id: ", movie.ElementId)
// 	fmt.Println("title: ", movie.Title)
// 	fmt.Println("released: ", movie.Released)
// 	fmt.Println("-------------------------\n")

// 	executer := ogm.NewExecuter()

// 	tx, err := executer.Begin()
// 	if err != nil {
// 		return err
// 	}

// 	err = tx.Create(person)
// 	if err != nil {
// 		return err
// 	}

// 	err = tx.Create(movie)
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Println("\n-------------------------")
// 	fmt.Println("Person после выполнения запроса")
// 	fmt.Println("-------------------------")
// 	fmt.Println("id: ", person.Id)
// 	fmt.Println("element id: ", person.ElementId)
// 	fmt.Println("name: ", person.Name)
// 	fmt.Println("job: ", person.Job)
// 	fmt.Println("-------------------------")

// 	fmt.Println("\n-------------------------")
// 	fmt.Println("Movie после выполнения запроса")
// 	fmt.Println("-------------------------")
// 	fmt.Println("id: ", movie.Id)
// 	fmt.Println("element id: ", movie.ElementId)
// 	fmt.Println("title: ", movie.Title)
// 	fmt.Println("released: ", movie.Released)
// 	fmt.Println("-------------------------\n")

// 	actedIn := ActedIn{
// 		Relationship: types.NewRelationship(),
// 		Role:         "neo",
// 	}

// 	tx.Match(
// 		entity.NewNode("p"),
// 		entity.NewNode("m"),
// 	).
// 		Where(fmt.Sprintf("ID(p) = %d AND ID(m) = %d", person.Id, movie.Id)).
// 		Create(entity.NewPatternList(entity.NewNode("p")).
// 			Related(entity.NewRelationship("a").
// 				SetLables("actedIn").
// 				SetProperties(map[string]any{
// 					"role": actedIn.Role,
// 				}),
// 			).
// 			To(entity.NewNode("m")),
// 		)

// 	_, err = tx.Do()
// 	if err != nil {
// 		return err
// 	}

// 	err = tx.Commit()
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
