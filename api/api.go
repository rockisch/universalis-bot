package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"universalis-bot/utils"
)

const (
	API_ENDPOINT = "https://universalis.app/api/v2/"
)

var (
	client http.Client
)

func init() {
	client = http.Client{}
}

func apiRequest(ctx context.Context, path string, target interface{}) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, API_ENDPOINT+path, nil)
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode > 400 {
		return fmt.Errorf("invalid response code %d", resp.StatusCode)
	}
	err = json.NewDecoder(resp.Body).Decode(target)
	if err != nil {
		return err
	}
	return nil
}

func GetDataCenters(ctx context.Context) ([]DataCenter, error) {
	var result []DataCenter
	err := apiRequest(ctx, "data-centers", result)
	return result, err
}

func GetWorlds(ctx context.Context) ([]World, error) {
	var result []World
	err := apiRequest(ctx, "worlds", result)
	return result, err
}

func GetUserList(ctx context.Context, list string) (UserList, error) {
	var result UserList
	path := fmt.Sprintf("lists/%s", list)
	err := apiRequest(ctx, path, &result)
	return result, err
}

func GetMarketCurrent(ctx context.Context, items []int) (MarketCurrent, error) {
	var result MarketCurrent
	itemsString := utils.JoinInts(items, ",")
	path := fmt.Sprintf("North-America/%s?hq=1", itemsString)
	fmt.Println(path)
	err := apiRequest(ctx, path, &result)
	return result, err
}
