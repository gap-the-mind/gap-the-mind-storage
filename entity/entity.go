//go:generate easyjson
package entity

import "github.com/mailru/easyjson"

// EntityRef is base for all storable objects
type Entity interface {
	easyjson.MarshalerUnmarshaler
	Id() string
	SetId(string)

	Nature() string
}

type Provider interface {
	Accept(e EntityPick) (bool, Entity)
}

//easyjson:json
// EntityPick is a minial implementation of Entity
type EntityPick struct {
	EntityId     string `json:"id"`
	EntityNature string `json:"nature"`
}

func (e EntityPick) Id() string {
	return e.EntityId
}

func (e EntityPick) SetId(s string) {
	e.EntityId = s
}

func (e EntityPick) Nature() string {
	return e.EntityNature
}
