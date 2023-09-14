package model

import (
	"GoAPI/common"
	"errors"
)

var (
	ErrTitleIsBlank = errors.New("title can not be empty")
)

// tao struct TodoItem va xac dinh truong tuong ung trong database bang gorm
type TodoItem struct {
	common.SQLModel        //embedding struct
	Title           string `json:"title" gorm:"column:title;"`
	Description     string `json:"description" gorm:"column:description;"`
	Status          string `json:"status" gorm:"column:status;"`
}

// tao method TableName cho Struct TodoItem, tra ve ten bang "todo_items"
func (TodoItem) TableName() string { return "todo_items" }

// tao struct TodoItemCreation va xac dinh truong tuong ung trong database bang gorm
type TodoItemCreation struct {
	common.SQLModel        //embedd struct
	Title           string `json:"title" gorm:"column:title;"`
	Description     string `json:"description" gorm:"column:description;"`
	Status          string `json:"status" gorm:"column:status;"`
}

// tao method TableName cho Struct TodoItemCreation, tra ve ten bang "todo_items"
func (TodoItemCreation) TableName() string { return TodoItem{}.TableName() }

// tao struct TodoItemUpdate va xac dinh truong tuong ung trong database bang gorm
type TodoItemUpdate struct {
	//Dùng con trỏ, xử lý trường hợp truyền vào chuỗi rỗng thì bị bỏ qua
	common.SQLModel         //embedd struct
	Title           *string `json:"title" gorm:"column:title;"`
	Description     *string `json:"description" gorm:"column:description;"`
	Status          *string `json:"status" gorm:"column:status;"`
}

// tao method TableName cho Struct TodoItemUpdate, tra ve ten bang "todo_items"
func (TodoItemUpdate) TableName() string { return TodoItem{}.TableName() }
