package gbank

import "testing"

func TestNestedGroup(t *testing.T) {
	r := NewEngine()
	v1 := r.NewGroup("/v1")
	v2 := v1.NewGroup("/v2")
	v3 := v2.NewGroup("/v3")
	if v2.prefix != "/v1/v2" {
		t.Fatal("v2 prefix should be /v1/v2")
	}
	if v3.prefix != "/v1/v2/v3" {
		t.Fatal("v2 prefix should be /v1/v2")
	}
}
