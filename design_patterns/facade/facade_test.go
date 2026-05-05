package facade

import "testing"

func TestFacade(t *testing.T) {
	api := NewAPI()
	out := api.Test()
	expected := "testA ... testB ... "

	if out != expected {
		t.Errorf("API Test() = %q, want %q", out, expected)
	}
}
