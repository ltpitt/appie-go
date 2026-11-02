package main

import (
	"testing"

	appie "github.com/gwillem/appie-go"
)

func TestFindList(t *testing.T) {
	lists := []appie.ShoppingList{
		{ID: "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee", Name: "Boodschappen"},
		{ID: "11111111-2222-3333-4444-555555555555", Name: "Weekmenu"},
	}

	t.Run("exact match", func(t *testing.T) {
		got, err := findList(lists, "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got.Name != "Boodschappen" {
			t.Fatalf("got %q, want %q", got.Name, "Boodschappen")
		}
	})

	t.Run("not found", func(t *testing.T) {
		_, err := findList(lists, "00000000-0000-0000-0000-000000000000")
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})

	t.Run("prefix does not match", func(t *testing.T) {
		_, err := findList(lists, "aaaaaaaa")
		if err == nil {
			t.Fatal("expected error for prefix match, got nil")
		}
	})
}

func TestMergeQuantities(t *testing.T) {
	listItems := []appie.ListItem{
		{ProductID: 1, Quantity: 2},
		{ProductID: 2, Quantity: 1},
		{ProductID: 3, Quantity: 1},
		{ProductID: 0, Quantity: 1}, // free-text item, skipped
	}
	basket := []appie.OrderItem{
		{ProductID: 1, Quantity: 2}, // already enough
		{ProductID: 2, Quantity: 5}, // more than the list wants, never reduced
	}

	got := mergeQuantities(listItems, basket)
	if len(got) != 1 || got[0].ProductID != 3 || got[0].Quantity != 1 {
		t.Fatalf("got %+v, want only product 3 x1", got)
	}
}
