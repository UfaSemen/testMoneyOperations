package server

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type PostgreBalansController struct {
	connection *sql.DB
	timeout    time.Duration
}

type ExecContexter interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

func updateAndCheck(ec ExecContexter, ctx context.Context, userId, amount int) (bool, error) {
	result, err := ec.ExecContext(ctx, "UPDATE accaunts SET balans = balans + $1 WHERE id = $2", amount, userId)
	if err != nil {
		return false, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	if rows == 0 {
		return false, err
	}
	return true, nil
}

func NewPostgreBalansController(pc PostgreConfig) (PostgreBalansController, error) {
	constr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", pc.Host, pc.Port, pc.User, pc.Password, pc.DBname)
	db, err := sql.Open("postgres", constr)
	if err != nil {
		return PostgreBalansController{}, err
	}
	return PostgreBalansController{connection: db, timeout: time.Duration(time.Second * time.Duration(pc.Timeout))}, nil
}

func (pbc PostgreBalansController) withdraw(userId, amount int) error {
	ctx, cansel := context.WithTimeout(context.Background(), pbc.timeout)
	defer cansel()
	success, err := updateAndCheck(pbc.connection, ctx, userId, -1*amount)
	if err != nil {
		return err
	}
	if !success {
		return errors.New("can't withdraw from non existent account")
	}
	return nil
}

func (pbc PostgreBalansController) deposit(userId, amount int) error {
	ctx, cansel := context.WithTimeout(context.Background(), pbc.timeout)
	defer cansel()

	tx, err := pbc.connection.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}
	defer tx.Rollback()

	success, err := updateAndCheck(tx, ctx, userId, amount)
	if err != nil {
		return err
	}
	if !success {
		_, err = tx.ExecContext(ctx, "INSERT INTO accaunts VAlUES ($1,$2)", userId, amount)
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (pbc PostgreBalansController) transfer(senderId, receiverId, amount int) error {
	ctx, cansel := context.WithTimeout(context.Background(), pbc.timeout)
	defer cansel()

	tx, err := pbc.connection.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}
	defer tx.Rollback()

	success, err := updateAndCheck(tx, ctx, senderId, -1*amount)
	if err != nil {
		return err
	}
	if !success {
		return errors.New("can't withdraw from non existent account")
	}

	success, err = updateAndCheck(tx, ctx, receiverId, amount)
	if err != nil {
		return err
	}
	if !success {
		_, err = tx.ExecContext(ctx, "INSERT INTO accaunts VAlUES ($1,$2)", receiverId, amount)
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (pbc PostgreBalansController) Close() {
	pbc.connection.Close()
}
