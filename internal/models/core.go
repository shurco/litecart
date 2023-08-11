package models

// Core is ...
type Core struct {
	Created int64 `json:"created"`
	Updated int64 `json:"updated"`
}

// UpdateClause is ...
type UpdateClause struct {
	Field string
	Value string
}
