package db

import (
	"database/sql"
	"fmt"
	"kanbaru/config"

	_ "github.com/go-sql-driver/mysql"
)

var DB, _ = sql.Open(config.Driver, fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", config.Username, config.Password, config.Host, config.Port, config.Database))
