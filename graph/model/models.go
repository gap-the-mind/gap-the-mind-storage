package model

type Node interface {
	IsNode()
}

type Note struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

func (Note) IsNode() {}

type PageInfo struct {
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *string `json:"startCursor"`
	EndCursor       *string `json:"endCursor"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (User) IsNode() {}

type UserNoteEdge struct {
	Cursor string `json:"cursor"`
	Node   *Note  `json:"node"`
}

type UserNotesConnection struct {
	Edges      []*UserNoteEdge `json:"edges"`
	PageInfo   *PageInfo       `json:"pageInfo"`
	TotalCount int             `json:"totalCount"`
}
