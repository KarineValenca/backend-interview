package sql

import (
	"context"
	"fmt"
	"strings"

	"github.com/gustvision/backend-interview/pkg/account"
)

func (s *Store) Fetch(ctx context.Context, f account.Filter) (account.Account, error) {
	b := strings.Builder{}
	b.WriteString(`SELECT id, user_id, total `)
	b.WriteString(`FROM account `)
	b.WriteString(`WHERE id = $1 ;`)

	row := s.QueryRowContext(ctx, b.String(), []interface{}{
		f.ID,
	}...)

	var a account.Account

	if err := row.Scan(
		&a.ID,
		&a.UserID,
		&a.Total,
	); err != nil {
		return account.Account{}, err
	}

	return a, nil
}

func (s *Store) FetchMany(ctx context.Context, f account.Filter, callback func(account.Account) error) error {
	b := strings.Builder{}
	b.WriteString(`SELECT id, user_id, total `)
	b.WriteString(`FROM account `)
	b.WriteString(`WHERE user_id = $1 ;`)

	rows, err := s.QueryContext(ctx, b.String(), []interface{}{
		f.UserID,
	}...)
	if err != nil {
		return err
	}

	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var a account.Account

		if err := rows.Scan(
			&a.ID,
			&a.UserID,
			&a.Total,
		); err != nil {
			return err
		}

		if err := callback(a); err != nil {
			return err
		}
	}

	return rows.Err()
}

func (s *Store) UpdateAccountTotal(ctx context.Context, f account.Filter) error {
	fmt.Println("updating total", f.Total)
	b := strings.Builder{}
	b.WriteString(`UPDATE account `)
	b.WriteString(`SET total = $1 `)
	b.WriteString(`WHERE id = $2 ;`)

	if _, err := s.ExecContext(ctx, b.String(), []interface{}{
		f.Total,
		f.ID,
	}...); err != nil {
		return err
	}

	return nil
}
