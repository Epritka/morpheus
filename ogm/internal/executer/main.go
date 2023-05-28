package executer

import (
	"context"

	"github.com/Epritka/morpheus/v2/ogm/internal/cypher"
	"github.com/Epritka/morpheus/v2/ogm/types"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Executer struct {
	types.Cypher

	cypher     *cypher.Cypher
	session    neo4j.SessionWithContext
	tx         neo4j.ExplicitTransaction
	autoCommit bool
}

func New(session neo4j.SessionWithContext) *Executer {
	cypher := cypher.New()

	return &Executer{
		Cypher:     cypher,
		cypher:     cypher,
		session:    session,
		autoCommit: false,
	}
}

func (e *Executer) Begin() (*Transaction, error) {
	return e.begin(context.Background())
}

func (e *Executer) BeginWithContext(ctx context.Context) (*Transaction, error) {
	return e.begin(ctx)
}

func (e *Executer) Close() error {
	return e.close(context.Background())
}

func (e *Executer) CloseWithContext(ctx context.Context) error {
	return e.close(ctx)
}

func (e *Executer) begin(ctx context.Context) (*Transaction, error) {
	tx, err := e.session.BeginTransaction(ctx)
	if err != nil {
		return nil, err
	}

	e.tx = tx

	return &Transaction{
		Executer: e,
	}, nil
}

func (e *Executer) close(ctx context.Context) error {
	if e.tx != nil {
		err := e.tx.Close(ctx)
		if err != nil {
			return err
		}
	}

	return e.session.Close(ctx)
}
