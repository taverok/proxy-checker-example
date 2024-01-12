package mysql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql" // mysql driver
	"github.com/taverok/proxy-checker-example/pkg/db"
)

func NewMysql(ds db.Datasource) (*sql.DB, error) {
	url := fmt.Sprintf("%s:%s@(%s:%d)/%s", ds.User, ds.Pass, ds.Host, ds.Port, ds.Name)
	DB, err := sql.Open("mysql", url)
	if err != nil {
		return nil, err
	}

	DB.SetConnMaxLifetime(time.Minute * 3)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)

	err = DB.Ping()
	if err != nil {
		return nil, err
	}

	return DB, nil
}
