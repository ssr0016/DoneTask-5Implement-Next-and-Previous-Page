// functionality for store
package main

import (
	"context"
	"database/sql"
	"errors"
)

func (s *store) getNextPageBankByID(ctx context.Context, bankID int64) (*bank.Bank, error) {
	var result bank.Bank

	err := s.db.WithTransaction(ctx, func(tx postgres.Tx) error {
		rawSQL := `
			SELECT
				id,
				UPPER(code) as code,
				name,
				currency,
				logo,
				url,
				create_by,
				create_at,
				update_by,
				update_at
			FROM bank
			WHERE update_at < (SELECT update_at FROM bank WHERE id = ?)
			ORDER BY update_at DESC
			LIMIT 1
		`

		err := tx.Get(ctx, &result, rawSQL, bankID)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, bank.ErrBankNotFound
		} else {
			return nil, err
		}
	}

	return &result, nil
}

func (s *store) getPreviousPageBankByID(ctx context.Context, bankID int64) (*bank.Bank, error) {
	var result bank.Bank

	err := s.db.WithTransaction(ctx, func(tx postgres.Tx) error {
		rawSQL := `
			SELECT
				id,
				UPPER(code) as code,
				name,
				currency,
				logo,
				url,
				create_by,
				create_at,
				update_by,
				update_at
			FROM bank
			WHERE (update_at > (SELECT update_at FROM bank WHERE id = ?) OR (update_at = (SELECT update_at FROM bank WHERE id = ?) AND id < ?))
			ORDER BY update_at ASC, id ASC
			 LIMIT 1
		`

		err := tx.Get(ctx, &result, rawSQL, bankID, bankID, bankID)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, bank.ErrBankNotFound
		} else {
			return nil, err
		}
	}

	return &result, nil
}
