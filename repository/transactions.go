package repository

import (
	"context"
	"database/sql"
)

func BeginTransaction(userRep *UserRepository, todoRep *TodoItemRepository) error {
	ctx := context.Background()
	transaction, err := todoRep.dbHandler.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	userRep.transaction = transaction
	todoRep.transaction = transaction
	return nil
}

func RollBackTransaction(userRep *UserRepository, todoRep *TodoItemRepository) error {
	ctx := context.Background()
	transaction, err := todoRep.dbHandler.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	userRep.transaction = nil
	todoRep.transaction = nil
	return transaction.Rollback()
}

func CommitTransaction(userRep *UserRepository, todoRep *TodoItemRepository) error {
	ctx := context.Background()
	transaction, err := todoRep.dbHandler.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	userRep.transaction = nil
	todoRep.transaction = nil
	return transaction.Commit()
}

//TODO Add transaction logic in Service layer
