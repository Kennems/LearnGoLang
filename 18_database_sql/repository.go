package main

import "context"

type User struct {
	ID   int64
	Name string
}

type RowScanner interface {
	Scan(dest ...any) error
}

type QueryRower interface {
	QueryRowContext(ctx context.Context, query string, args ...any) RowScanner
}

func findUserByID(ctx context.Context, db QueryRower, id int64) (User, error) {
	row := db.QueryRowContext(ctx, "SELECT id, name FROM users WHERE id = ?", id)

	var user User
	if err := row.Scan(&user.ID, &user.Name); err != nil {
		return User{}, err
	}

	return user, nil
}
