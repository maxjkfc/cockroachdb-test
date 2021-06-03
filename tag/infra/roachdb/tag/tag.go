package tag

import (
	"context"
	"fmt"
	"tag/domain/model"
	"tag/infra/roachdb"
	"tag/usecase/repository"

	"github.com/upper/db/v4"
)

type tagRepositoryWithUpper struct {
	db         db.Session
	collection string
}

func NewRepositoryWithUpper(host, database, user, password string) (repository.TagRepository, error) {
	db, err := roachdb.NewWithUpper(host, database, user, password)
	if err != nil {
		return nil, err
	}

	return &tagRepositoryWithUpper{
		db:         db,
		collection: "records",
	}, nil
}

func (u *tagRepositoryWithUpper) Create(ctx context.Context, record model.Record) error {
	r := ConvertToDB(record)

	result, err := u.db.Collection(u.collection).Insert(r)
	if err != nil {
		return err
	}

	fmt.Println(result.ID())

	return nil
}

func (u *tagRepositoryWithUpper) Search(ctx context.Context, filter model.Filter) (total int64, records []*model.Record) {

	return 0, nil
}
