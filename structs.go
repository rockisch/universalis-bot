package main

import (
	"context"
	"log"
	"strconv"
	"universalis-bot/api"
)

type WorldData struct {
	Name       string
	DataCenter string
}

func FetchWorldData(ctx context.Context) (map[int]WorldData, error) {
	apiDataCenters, err := api.GetDataCenters(ctx)
	if err != nil {
		return nil, err
	}
	apiWorlds, err := api.GetWorlds(ctx)
	if err != nil {
		return nil, err
	}
	worldData := make(map[int]WorldData)
	for _, world := range apiWorlds {
		for _, dc := range apiDataCenters {
			found := false
			for _, worldID := range dc.Worlds {
				if world.ID == worldID {
					found = true
					break
				}
			}
			if found {
				worldData[world.ID] = WorldData{
					Name:       world.Name,
					DataCenter: dc.Name,
				}
			}
		}
	}
	return worldData, nil
}

type MinPriceItem struct {
	ID       int
	MinPrice int
	WorldID  int
}

func FetchMinPriceItems(ctx context.Context, items []int) ([]MinPriceItem, error) {
	mc, err := api.GetMarketCurrent(ctx, items)
	if err != nil {
		return nil, err
	}
	var minPrices []MinPriceItem
	for key, items := range mc.Items {
		id, err := strconv.Atoi(key)
		if err != nil {
			log.Fatalf("invalid key on market current: %s", key)
			continue
		}
		for _, listing := range items.Listings {
			if listing.PricePerUnit <= items.MinPriceHQ {
				minPrices = append(minPrices, MinPriceItem{id, listing.PricePerUnit, listing.WorldID})
				break
			}
		}
	}
	return minPrices, nil
}
