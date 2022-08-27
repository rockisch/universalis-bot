package api

type DataCenter struct {
	Name   string
	Region string
	Worlds []int
}

type World struct {
	ID   int
	Name string
}

type UserList struct {
	ID      string
	Created string
	Updated string
	Name    string
	ItemIDs []int
}

type MarketCurrentListing struct {
	PricePerUnit int
	WorldID      int
}

type MarketCurrentItem struct {
	MinPriceHQ int
	Listings   []MarketCurrentListing
}

type MarketCurrent struct {
	ItemIDs []int
	Items   map[string]MarketCurrentItem
}
