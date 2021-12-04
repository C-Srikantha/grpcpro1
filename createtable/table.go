package createtable

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type Information struct {
	Id    int32 `pg:",pk"`
	Name  string
	Age   int32
	Phone int64 `sql:",unique"`
}

func CreateTab(db *pg.DB) error {
	model := []interface{}{
		(*Information)(nil),
	}
	for _, val := range model {
		err := db.Model(val).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
			Varchar:     50,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
