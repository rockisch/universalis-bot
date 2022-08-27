package api

import (
	"context"
	"testing"
)

func TestGetUserList(t *testing.T) {
	ctx := context.Background()
	list := "26989876-49b9-46c3-bc80-3f1aca98fa5e"
	result, err := GetUserList(ctx, list)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", result)
}

func TestGetMarketCurrent(t *testing.T) {
	ctx := context.Background()
	items := []int{37813, 37812}
	result, err := GetMarketCurrent(ctx, items)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", result)
}
