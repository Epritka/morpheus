package executer

import (
	"context"

	"github.com/Epritka/morpheus/v1/ogm/internal/cypher"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Executer struct {
	*cypher.Cypher
	tx         neo4j.ExplicitTransaction
	autoCommit bool
}

func New(session neo4j.SessionWithContext, cypher *cypher.Cypher) (*Executer, error) {
	tx, err := session.BeginTransaction(context.Background())
	if err != nil {
		return nil, err
	}

	return &Executer{
		Cypher:     cypher,
		tx:         tx,
		autoCommit: true,
	}, nil
}

func NewWithTx(tx neo4j.ExplicitTransaction, cypher *cypher.Cypher) *Executer {
	return &Executer{
		Cypher:     cypher,
		tx:         tx,
		autoCommit: false,
	}
}

func (b *Executer) Do() error {
	return b.do(context.Background())
}

func (b *Executer) DoWithContext(ctx context.Context) error {
	return b.do(ctx)
}

func (b *Executer) DoQuery(cypher string) error {
	return b.doQuery(context.Background(), cypher)
}

func (b *Executer) DoQueryWithContext(ctx context.Context, cypher string) error {
	return b.doQuery(ctx, cypher)
}

func (e *Executer) do(ctx context.Context) error {
	return e.doQuery(ctx, e.String())
}

func (b *Executer) doQuery(ctx context.Context, cypher string) error {
	err := func() error {
		result, err := b.tx.Run(ctx, cypher, map[string]any{})
		if err != nil {
			return err
		}

		if result.Err() != nil {
			return result.Err()
		}

		return nil
	}()

	if !b.autoCommit {
		return err
	}

	if err == nil {
		return b.tx.Commit(ctx)
	}

	{
		err := b.tx.Rollback(ctx)
		if err != nil {
			return err
		}
	}

	return err
}
