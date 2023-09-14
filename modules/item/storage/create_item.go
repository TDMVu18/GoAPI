package storage

import (
	"GoAPI/modules/item/model"
	"context"
)

func (s *sqlStorage) CreateItem(ctx context.Context, data *model.TodoItemCreation) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
