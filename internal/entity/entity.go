package entity

import "github.com/google/uuid" // dependencies = command: go mod tidy

type Category struct {
	ID   string
	Name string
}

func NewCategory(name string) *Category {
	return &Category{
		ID:   uuid.New().String(),
		Name: name,
	}
}
