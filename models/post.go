package models

// PostApi struct
type PostApi struct {
	ID     int    `json:"id"`
	UserId int    `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
