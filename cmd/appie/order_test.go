package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	appie "github.com/gwillem/appie-go"
)

func TestEnsureOrderOpenActiveOrder(t *testing.T) {
	newClient := func(state string) *appie.Client {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(`{"id":123,"state":"` + state + `","orderedProducts":[]}`))
		}))
		t.Cleanup(srv.Close)
		return appie.New(appie.WithBaseURL(srv.URL))
	}

	t.Run("unconfirmed active order", func(t *testing.T) {
		if err := ensureOrderOpen(context.Background(), newClient("UNCONFIRMED"), nil, 123); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("other order id", func(t *testing.T) {
		if err := ensureOrderOpen(context.Background(), newClient("UNCONFIRMED"), nil, 456); err == nil {
			t.Fatal("expected error")
		}
	})

	t.Run("submitted active order is not modified", func(t *testing.T) {
		if err := ensureOrderOpen(context.Background(), newClient("SUBMITTED"), nil, 123); err == nil {
			t.Fatal("expected error")
		}
	})
}
