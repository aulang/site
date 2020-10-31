package model

import "github.com/aulang/site/entity"

type ConfigRequest struct {
	Config entity.WebConfig `json:"config"`
	Menus  []entity.Menu    `json:"menus"`
}
