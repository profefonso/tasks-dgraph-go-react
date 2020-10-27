// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Category struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type DataCategory struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type DataTask struct {
	Title      string `json:"title"`
	CategoryID string `json:"categoryId"`
	Completed  bool   `json:"completed"`
}

type Task struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Category  *Category `json:"category"`
	Completed bool      `json:"completed"`
}
