package entity

// EntityRef is base for all storable objects
type Entity interface {
	Id() string
	SetId(string)

	Nature() string
}

type Provider interface {
	Accept(id string, nature string) Entity
}
