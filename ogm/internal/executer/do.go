package executer

import "context"

func (e *Executer) Do() error {
	return e.do(context.Background())
}

func (e *Executer) DoWithContext(ctx context.Context) error {
	return e.do(ctx)
}

func (e *Executer) DoQuery(cypher string) error {
	return e.doQuery(context.Background(), cypher)
}

func (e *Executer) DoQueryWithContext(ctx context.Context, cypher string) error {
	return e.doQuery(ctx, cypher)
}

func (e *Executer) do(ctx context.Context) error {
	return e.doQuery(ctx, e.cypher.Build())
}

func (e *Executer) doQuery(ctx context.Context, cypher string) error {
	err := func() error {
		if e.tx == nil {
			tx, err := e.session.BeginTransaction(ctx)
			if err != nil {
				return err
			}

			e.tx = tx
			e.autoCommit = true
		}

		result, err := e.tx.Run(ctx, cypher, map[string]any{})
		if err != nil {
			return err
		}

		if result.Err() != nil {
			return result.Err()
		}

		return nil
	}()

	if !e.autoCommit {
		return err
	}

	if err == nil {
		return e.tx.Commit(ctx)
	}

	{
		err := e.tx.Rollback(ctx)
		if err != nil {
			return err
		}
	}

	return err
}
