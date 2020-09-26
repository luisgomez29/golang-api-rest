package auto

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/luisgomez29/golang-api-rest/models"
)

type DBLogger struct{}

func (d DBLogger) BeforeQuery(ctx context.Context, q *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}

func (d DBLogger) AfterQuery(ctx context.Context, q *pg.QueryEvent) error {
	query, _ := q.FormattedQuery()
	fmt.Printf("%s\n\n", query)
	return nil
}

func Load(db *pg.DB) {
	modelsSlice := []interface{}{
		(*models.User)(nil),
		(*models.Product)(nil),
	}

	for _, model := range modelsSlice {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists:   true,
			FKConstraints: true,
		})
		if err != nil {
			panic(err)
		}
	}
}
