package createtable

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"grpc.com/grpc/entity"
)

func CreateTab(db *pg.DB) error {
	model := []interface{}{
		(*entity.Information)(nil),
		(*entity.CollegeDetails)(nil),
	}
	for _, val := range model {
		err := db.Model(val).CreateTable(&orm.CreateTableOptions{
			IfNotExists:   true,
			Varchar:       50,
			FKConstraints: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
