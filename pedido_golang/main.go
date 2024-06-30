package main

import (
	"database/sql"
	"fmt"
	"log"
	"pedido_golang/pedido"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {

	//db = baseDatos.InitDB()

	var err error
	// Conectar a la base de datos
	dsn := "root:@tcp(127.0.0.1:3306)/pedidodb"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Verificar la conexión
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexión exitosa a la base de datos")

	fmt.Println(pedido.CrearPedidoSinFW(db, "OtroPedidoCodigo", 1))

	// router := gin.Default()
	// router.POST("/pedido", pedido.CrearPedido)
	// router.GET("/ranking", pedido.RankingEstablecimientos)
	// router.GET("/ubicacion/:id", pedido.UbicacionEstablecimiento)
	// router.Run(":8080")
}
