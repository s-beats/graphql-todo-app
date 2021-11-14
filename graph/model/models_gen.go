// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type Task struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type User struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Tasks []*Task `json:"tasks"`
}

type CreateTaskInput struct {
	UserID string `json:"userID"`
	Title  string `json:"title"`
	Text   string `json:"text"`
}

type CreateTaskPayload struct {
	Task *Task `json:"task"`
}
