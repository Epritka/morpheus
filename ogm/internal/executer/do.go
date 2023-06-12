package executer

import (
	"context"

	"github.com/Epritka/morpheus/builder"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func (e *Executer) Do() (neo4j.ResultWithContext, error) {
	return e.do(context.Background())
}

func (e *Executer) DoWithContext(ctx context.Context) (neo4j.ResultWithContext, error) {
	return e.do(ctx)
}

func (e *Executer) DoQuery(cypher string) (neo4j.ResultWithContext, error) {
	return e.doQuery(context.Background(), cypher)
}

func (e *Executer) DoQueryWithContext(ctx context.Context, cypher string) (neo4j.ResultWithContext, error) {
	return e.doQuery(ctx, cypher)
}

func (e *Executer) do(ctx context.Context) (neo4j.ResultWithContext, error) {
	return e.doQuery(ctx, e.Build())
}

func (e *Executer) doQuery(ctx context.Context, cypher string) (neo4j.ResultWithContext, error) {
	result, err := func() (neo4j.ResultWithContext, error) {
		if e.tx == nil {
			tx, err := e.session.BeginTransaction(ctx)
			if err != nil {
				return nil, err
			}

			e.tx = tx
			e.autoCommit = true
		}

		result, err := e.tx.Run(ctx, cypher, map[string]any{})
		if err != nil {
			return nil, err
		}

		if result.Err() != nil {
			return nil, result.Err()
		}

		return result, nil
	}()

	e.Builder = builder.NewBuilder()

	if !e.autoCommit {
		return result, err
	}

	if err == nil {
		{
			err := e.tx.Commit(ctx)
			if err != nil {
				return nil, err
			}
		}

		return result, nil
	}

	{
		err := e.tx.Rollback(ctx)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}
