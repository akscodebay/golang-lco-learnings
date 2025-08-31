package model

type Netflix struct {
	ID      int    `json:"id,omitempty"`
	Title   string `json:"title,omitempty"`
	Watched bool   `json:"watched,omitempty"`
}
