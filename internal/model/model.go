package model

type Cat struct {
	ID    int64  `json:"id"`
	Age   int64  `json:"age"`
	Name  string `json:"name"`
	Color string `json:"color"`
}
