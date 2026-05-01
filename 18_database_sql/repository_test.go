package main

import (
	"context"
	"errors"
	"testing"
)

type stubRow struct {
	user User
	err  error
}

func (r stubRow) Scan(dest ...any) error {
	if r.err != nil {
		return r.err
	}
	*(dest[0].(*int64)) = r.user.ID
	*(dest[1].(*string)) = r.user.Name
	return nil
}

type stubDB struct {
	row stubRow
}

func (db stubDB) QueryRowContext(_ context.Context, _ string, _ ...any) RowScanner {
	return db.row
}

func TestFindUserByID(t *testing.T) {
	db := stubDB{
		row: stubRow{user: User{ID: 1, Name: "Ken"}},
	}

	got, err := findUserByID(context.Background(), db, 1)
	if err != nil {
		t.Fatalf("findUserByID() error = %v", err)
	}
	if got.Name != "Ken" {
		t.Fatalf("got.Name = %q, want %q", got.Name, "Ken")
	}
}

func TestFindUserByIDScanError(t *testing.T) {
	wantErr := errors.New("scan failed")
	db := stubDB{
		row: stubRow{err: wantErr},
	}

	_, err := findUserByID(context.Background(), db, 1)
	if !errors.Is(err, wantErr) {
		t.Fatalf("expected %v, got %v", wantErr, err)
	}
}
