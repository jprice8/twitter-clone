package model

type Category struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Categories struct {
	Categories []Category `json:"categories"`
}
