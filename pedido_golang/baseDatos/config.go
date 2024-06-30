package baseDatos

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() *sql.DB {
	var err error
	DB, err := sql.Open("mysql", "root:@tcp(localhost:3306)/PedidoDB")
	if err != nil {
		panic(err)
	}

	return DB
}
