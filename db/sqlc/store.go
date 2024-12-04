package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Store interface {
	Querier
	ExecTx(ctx context.Context, fn func(*Queries) error) error
}

type SQLStore struct {
	*Queries
	db *pgx.Conn
}

func NewStore(db *pgx.Conn) *SQLStore {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

func (store *SQLStore) ExecTx(ctx context.Context, fn func(*Queries) error) error {

	tx, err := store.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	q := New(tx)

	if err := fn(q); err != nil {
		return err
	}

	return tx.Commit(ctx)
}
