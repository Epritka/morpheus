package session

import (
	"context"

	"github.com/Epritka/morpheus/v1/ogm/internal/cypher"
	"github.com/Epritka/morpheus/v1/ogm/internal/executer"
	"github.com/Epritka/morpheus/v1/ogm/internal/transaction"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Session struct {
	*cypher.Cypher
	session neo4j.SessionWithContext
}

func New(session neo4j.SessionWithContext) *Session {
	return &Session{
		Cypher:  cypher.New(),
		session: session,
	}
}

func (s *Session) Begin() (*transaction.Transaction, error) {
	return transaction.Begin(context.Background(), s.session, cypher.New())
}

func (s *Session) BeginWithContext(ctx context.Context) (*transaction.Transaction, error) {
	return transaction.Begin(ctx, s.session, cypher.New())
}

func (s *Session) Close(ctx context.Context) error {
	return s.session.Close(ctx)
}

func (s *Session) Do() error {
	return s.do(context.Background())
}

func (s *Session) DoWithContext(ctx context.Context) error {
	return s.do(ctx)
}

func (s *Session) DoQuery(cypher string) error {
	return s.doQuery(context.Background(), cypher)
}

func (s *Session) DoQueryWithContext(ctx context.Context, cypher string) error {
	return s.doQuery(ctx, cypher)
}

func (s *Session) do(ctx context.Context) error {
	executer, err := executer.New(s.session, s.Cypher)
	if err != nil {
		return err
	}

	return s.doQuery(ctx, executer.String())
}

func (s *Session) doQuery(ctx context.Context, cypher string) error {
	executer, err := executer.New(s.session, s.Cypher)
	if err != nil {
		return err
	}

	return executer.DoQueryWithContext(ctx, cypher)
}
