package main

import (
	"database/sql"
	"fmt"
	"log"
	"pedido_golang/pedido"
	"net/http"
	"os"
	"encoding/json"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	var err error
	// Conectar a la base de datos MySQL
	dsn := "root:@tcp(127.0.0.1:3306)/pedidodb"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Verificar la conexión a la base de datos
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexión exitosa a la base de datos")

	// Configurar los manejadores HTTP para las diferentes rutas

	// Ruta para obtener el ranking de establecimientos
	http.HandleFunc("/ranking", func(w http.ResponseWriter, r *http.Request) {
		pedido.RankingEstablecimientosSinFW(db, w, r)
	})

	// Ruta para obtener la ubicación de un establecimiento por su ID
	http.HandleFunc("/ubicacion", func(w http.ResponseWriter, r *http.Request) {
		pedido.UbicacionEstablecimientoSinFW(db, w, r)
	})

	// Ruta para crear un nuevo pedido
	http.HandleFunc("/crearPedido", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		// Estructura para decodificar los datos JSON del cuerpo de la solicitud
		var datos struct {
			Nombre            string `json:"Nombre"`
			IDEstablecimiento int    `json:"ID_Establecimiento"`
		}
		// Decodifica los datos JSON del cuerpo de la solicitud en la estructura datos
		if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Llama a la función para crear un pedido en la base de datos
		id, err := pedido.CrearPedidoSinFW(db, datos.Nombre, datos.IDEstablecimiento)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Prepara la respuesta JSON con el ID del pedido creado
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"ID_Pedido": id,
		})
	})

	// Inicia el servidor HTTP en el puerto 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

	//fmt.Println(pedido.CrearPedidoSinFW(db, "OtroPedidoCodigo", 1))
}
