package repository

import (
	"database/sql"
	"encurtador_url/src/internal/entity"
)

type CodeRepository struct {
	Db *sql.DB
}

func NewCodeRepository(db *sql.DB) *CodeRepository {
	return &CodeRepository{
		Db: db,
	}
}

func (r *CodeRepository) Save(code entity.Code) error {
	sql, err := r.Db.Prepare("insert into codes(url,code) values($1,$2)")
	if err != nil {
		return err
	}

	defer sql.Close()

	_, err = sql.Exec(code.Url, code.Code)
	if err != nil {
		return err
	}

	return nil
}

func (r *CodeRepository) Get(code string) (*entity.Code, error) {
	sql, err := r.Db.Prepare(`select * from codes where code = $1`)
	if err != nil {
		return nil, err
	}

	defer sql.Close()

	var c entity.Code

	err = sql.QueryRow(code).Scan(&c.Url, &c.Code)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}
		return nil, err
	}

	return &c, nil
}
