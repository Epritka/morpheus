package executer

import (
	"context"
	"log"

	"github.com/Epritka/morpheus/builder"
	"github.com/Epritka/morpheus/builder/entity"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

var logger = log.Default()

type Executer struct {
	entity.Builder

	session    neo4j.SessionWithContext
	tx         neo4j.ExplicitTransaction
	autoCommit bool
}

func New(session neo4j.SessionWithContext) *Executer {
	builder := builder.NewBuilder()
	return &Executer{
		Builder:    builder,
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
	logger.Print("transaction begining")
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
