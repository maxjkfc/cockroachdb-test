package tag

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
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

func (u *tagRepositoryWithUpper) Create(ctx context.Context, record ...model.Record) error {
	r := ConvertToDB(record...)

	err := u.db.Tx(func(sess db.Session) error {
		for _, v := range r {
			_, err := sess.Collection(u.collection).Insert(v)
			if err != nil {
				return err
			}
		}
		return nil
	})

	return err
}

func (u *tagRepositoryWithUpper) Search(ctx context.Context, filter model.Filter) (total int64, records []*model.Record) {

	sql := u.db.SQL().Select("id", "user_info", "tag", "createdat").From(u.collection)
	sql2 := `SELECT count(id) FROM ` + u.collection + " WHERE "

	var filterlist []string

	if filter.User.Name != "" {
		temp := fmt.Sprintf("user_name = '%s'", filter.User.Name)
		sql = sql.And(temp)
		filterlist = append(filterlist, temp)
	}
	if filter.User.ID != "" {
		temp := fmt.Sprintf("user_id = '%s'", filter.User.ID)
		sql = sql.And(temp)
		filterlist = append(filterlist, temp)
	}

	if len(filter.Tags) > 0 {
		temp := "("
		for i, v := range filter.Tags {
			if i > 0 {
				temp += fmt.Sprintf(" OR tag = '%s'", v)
			} else {
				temp += fmt.Sprintf("tag = '%s'", v)
			}
		}
		temp += ")"

		sql = sql.And(temp)
		filterlist = append(filterlist, temp)
	}

	//---------------

	if filter.Page > 1 {
		sql = sql.Offset(int(filter.Page * filter.Size))
	}

	if filter.Size > 0 {
		sql = sql.Limit(int(filter.Size))
	}
	ans := make([]Record, 0)

	if err := sql.All(&ans); err != nil {
		log.Fatal(err)
	}

	result, err := u.db.SQL().QueryRow(sql2 + strings.Join(filterlist, " AND "))
	if err != nil {
		log.Fatal(err)
	}

	if err := result.Scan(&total); err != nil {
		log.Fatal(err)
	}

	records = make([]*model.Record, len(ans))
	for i, v := range ans {
		uJSON, _ := v.UserInfo.MarshalJSON()
		u := model.User{}
		json.Unmarshal(uJSON, &u)

		records[i] = &model.Record{User: u, Tag: v.Tag, CreatedAt: v.CreatedAt}
	}

	return
}
