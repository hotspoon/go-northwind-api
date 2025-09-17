package models

type Territory struct {
	TerritoryID   string `json:"territory_id"`
	TerritoryDesc string `json:"territory_description"`
	RegionID      int    `json:"region_id"`
}
