package example

import (
	"context"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Transaction struct {
	tx neo4j.ExplicitTransaction
}

func (db *DB) Begin() (*Transaction, error) {
	return db.begin(context.Background())
}

func (db *DB) BeginWithContext(ctx context.Context) (*Transaction, error) {
	return db.begin(ctx)
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

func (db *DB) begin(ctx context.Context) (*Transaction, error) {
	session := db.Driver.NewSession(ctx, neo4j.SessionConfig{})
	tx, err := session.BeginTransaction(ctx)
	if err != nil {
		tx.Close(ctx)
		return nil, err
	}

	db.tx = tx
	return &Transaction{
		tx: tx,
	}, nil
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
