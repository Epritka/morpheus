package executer

import (
	"context"

	"github.com/Epritka/morpheus/builder"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func (e *Executer) Do() ([]*neo4j.Record, error) {
	return e.doResult(context.Background())
}

func (e *Executer) DoWithContext(ctx context.Context) ([]*neo4j.Record, error) {
	return e.doResult(ctx)
}

func (e *Executer) DoQuery(cypher string) ([]*neo4j.Record, error) {
	return e.doQueryResult(context.Background(), cypher)
}

func (e *Executer) DoQueryWithContext(ctx context.Context, cypher string) ([]*neo4j.Record, error) {
	return e.doQueryResult(ctx, cypher)
}

func (e *Executer) doResult(ctx context.Context) ([]*neo4j.Record, error) {
	return e.doQueryResult(ctx, e.Build())
}

func (e *Executer) doQueryResult(ctx context.Context, cypher string) ([]*neo4j.Record, error) {
	var records []*neo4j.Record
	return records, e.do(ctx, cypher, func(rwc neo4j.ResultWithContext) error {
		dbRecords, err := rwc.Collect(context.Background())
		if err != nil {
			return err
		}

		records = dbRecords
		return nil
	})
}

func (e *Executer) doParse(ctx context.Context, resultParser func(neo4j.ResultWithContext) error) error {
	return e.doQueryParse(ctx, e.Build(), resultParser)
}

func (e *Executer) doQueryParse(ctx context.Context, cypher string, resultParser func(neo4j.ResultWithContext) error) error {
	return e.do(ctx, cypher, resultParser)
}

func (e *Executer) do(ctx context.Context, cypher string, resultParser func(neo4j.ResultWithContext) error) error {
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

	if err == nil {
		err = resultParser(result)
	}

	if !e.autoCommit {
		return err
	}

	if err == nil {
		{
			err := e.tx.Commit(ctx)
			if err != nil {
				return err
			}
		}
		return nil
	}

	{
		err := e.tx.Rollback(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}
