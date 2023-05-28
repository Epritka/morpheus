package transaction

import (
	"context"

	"github.com/Epritka/morpheus/ogm/internal/cypher"
	"github.com/Epritka/morpheus/ogm/internal/executer"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Transaction struct {
	*executer.Executer
	tx neo4j.ExplicitTransaction
}

func Begin(
	ctx context.Context,
	session neo4j.SessionWithContext,
	cypher *cypher.Cypher,
) (*Transaction, error) {
	tx, err := session.BeginTransaction(ctx)
	if err != nil {
		session.Close(ctx)
		return nil, err
	}

	return &Transaction{
		Executer: executer.NewWithTx(tx, cypher),
		tx:       tx,
	}, nil
}

func (t *Transaction) Commit() error {
	return t.commit(context.Background())
}

func (t *Transaction) CommitWithContext(ctx context.Context) error {
	return t.commit(ctx)
}

func (t *Transaction) RollBack() error {
	return t.rollback(context.Background())
}

func (t *Transaction) RollBackWithContext(ctx context.Context) error {
	return t.rollback(ctx)
}

func (t *Transaction) Close() error {
	return t.close(context.Background())
}

func (t *Transaction) CloseWithContext(ctx context.Context) error {
	return t.close(ctx)
}

func (t *Transaction) commit(ctx context.Context) error {
	return t.tx.Commit(ctx)
}

func (t *Transaction) rollback(ctx context.Context) error {
	return t.tx.Rollback(ctx)
}

func (t *Transaction) close(ctx context.Context) error {
	return t.tx.Close(ctx)
}
