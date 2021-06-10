package repository

import (
	"context"
	"tag/domain/model"
)

type TagRepository interface {
	Create(ctx context.Context, record ...model.Record) error
	Search(ctx context.Context, filter model.Filter) (total int64, records []*model.Record)
}
