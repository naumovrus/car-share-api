package storage_postgres

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Tx struct {
	db pgx.Tx
}

type TxReq func(tx Tx) error

type TxRunner interface {
	Begin(ctx context.Context) (pgx.Tx, error)
}

func ExecTx(ctx context.Context, runner TxRunner, req TxReq) error {
	pgxTx, err := runner.Begin(ctx)
	if err != nil {
		return err
	}

	tx := Tx{
		db: pgxTx,
	}

	defer tx.Rollback(context.TODO())

	err = req(tx)
	if err != nil {
		return err
	}
	return tx.Commit(context.TODO())
}

func (p Tx) Stats() *pgxpool.Stat {
	return nil
}

func (p Tx) Begin(ctx context.Context) (pgx.Tx, error) {
	return p.db.Begin(ctx)
}

func (p Tx) Rollback(ctx context.Context) {
	_ = p.db.Rollback(ctx)
}

func (p Tx) Commit(ctx context.Context) error {
	return p.db.Commit(ctx)
}

func (p Tx) Query(query string, args ...any) (pgx.Rows, error) {
	return p.db.Query(context.Background(), query, args[:]...)
}

func (p Tx) Get(dest interface{}, query string, args ...interface{}) error {
	rows, err := p.db.Query(context.Background(), query, args[:]...)
	if err != nil {
		return err
	}
	return pgxscan.DefaultAPI.ScanOne(dest, rows)
}

func (p Tx) Select(dest interface{}, query string, args ...interface{}) error {
	rows, err := p.db.Query(context.Background(), query, args[:]...)
	if err != nil {
		return err
	}
	return pgxscan.DefaultAPI.ScanAll(dest, rows)
}

func (p Tx) Exec(query string, args ...any) (pgconn.CommandTag, error) {
	return p.db.Exec(context.Background(), query, args[:]...)
}

func (p Tx) QueryRow(query string, args ...interface{}) pgx.Row {
	return p.db.QueryRow(context.Background(), query, args[:]...)
}
