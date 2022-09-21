package models

import "github.com/uptrace/bun"

type User struct {
	bun.BaseModel `bun:"table:go_user,alias:u"`

	Name   string `json:"name"`
	Banana string `json:"banana"`
}
