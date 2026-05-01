package main

import (
	"github.com/stretchr/testify/mock"
	"testing"
)

type DB interface {
	Get(key string) (string, error)
}

type MockDB struct {
	mock.Mock
}

func (m *MockDB) Get(key string) (string, error) {
	args := m.Called(key)
	return args.String(0), args.Error(1)
}

func TestGetFromDb(t *testing.T) {
	mockDB := new(MockDB)
	mockDB.On("Get", "key").Return("value", nil)

	value, err := mockDB.Get("key")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if value != "value" {
		t.Errorf("Expected 'value', got '%s'", value)
	}

	mockDB.AssertExpectations(t)
}
