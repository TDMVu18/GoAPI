package model

import (
	"time"
)

type GenerateItem struct {
	Id        string     `json:"id" gorm:"column:id"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at"`
}

type TodoItem struct {
	GenerateItem        //embedd strcut
	Title        string `json:"title" gorm:"column:title" form:"title"`
	Description  string `json:"description" gorm:"column:description" form:"description"`
	Status       string `json:"status" gorm:"column:status" form:"status"`
}

type TodoItemCreate struct {
	GenerateItem        //embedd struct
	Title        string `json:"title" gorm:"column:title"`
	Description  string `json:"description" gorm:"column:description"`
	Status       string `json:"status" gorm:"column:status"`
}

type TodoItemUpdate struct {
	GenerateItem        //embedd struct
	Title        string `json:"title" gorm:"column:title"`
	Description  string `json:"description" gorm:"column:description"`
	Status       string `json:"status" gorm:"column:status"`
}

//tao method table name cho cac struct

func (TodoItem) TableName() string { return "todo_items" }

func (TodoItemCreate) TableName() string { return TodoItem{}.TableName() }

func (TodoItemUpdate) TableName() string { return TodoItem{}.TableName() }
