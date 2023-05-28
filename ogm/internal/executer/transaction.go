package executer

import (
	"context"
)

type Transaction struct {
	*Executer
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
