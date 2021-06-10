package tag

import (
	"context"
	"encoding/json"
	"fmt"
	"tag/domain/model"
	"tag/infra/roachdb"
	"tag/usecase/repository"
	"time"

	"github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgx"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type tagRepositoryPgx struct {
	db         *pgxpool.Pool
	collection string
}

func NewRepositoryWithPgx(host, database, user, password string) (repository.TagRepository, error) {
	db, err := roachdb.NewWithPgx(host, database, user, password)
	if err != nil {
		return nil, err
	}

	return &tagRepositoryPgx{
		db:         db,
		collection: "records",
	}, nil

}

func (u *tagRepositoryPgx) Create(ctx context.Context, record ...model.Record) error {
	str := "INSERT INTO " + u.collection + "(user_info , tag , createdat) VALUES "

	err := crdbpgx.ExecuteTx(ctx, u.db, pgx.TxOptions{}, func(p pgx.Tx) error {
		for _, v := range record {
			ujson, _ := json.Marshal(v.User)
			str += fmt.Sprintf("( '%s' , \"%s\" , \"%s\" ) ", ujson, v.Tag, time.Now().Local().Format(time.RFC3339))
		}

		str += ";"
		fmt.Println(str)

		comm, err := p.Exec(ctx, str)
		if err != nil {
			return err
		}

		fmt.Println(comm.String())
		return nil
	})

	return err
}

func (u *tagRepositoryPgx) Search(ctx context.Context, filter model.Filter) (total int64, records []*model.Record) {

	return 0, nil
}
