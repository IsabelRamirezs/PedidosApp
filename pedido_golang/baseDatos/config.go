package baseDatos

import (
	"database/sql"

	// Importa el controlador MySQL para utilizarlo con database/sql
	_ "github.com/go-sql-driver/mysql"
)


// InitDB inicializa y devuelve una conexión a la base de datos MySQL.
func InitDB() *sql.DB {
	var err error
	// Abre una conexión a la base de datos utilizando el controlador MySQL
	DB, err := sql.Open("mysql", "root:@tcp(localhost:3306)/PedidoDB")
	if err != nil {
		panic(err)
	}

	// Devuelve la conexión a la base de datos
	return DB
}
