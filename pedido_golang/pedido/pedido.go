package pedido

import (
	"database/sql"	// Importa el paquete sql para manejar bases de datos relacionales
	"net/http"		// Importa el paquete http para manejar solicitudes y respuestas HTTP
	"encoding/json"	// Importa el paquete encoding/json para manejar codificación y decodificación JSON
	
)

// CrearPedidoSinFW inserta un nuevo pedido en la base de datos.
// Recibe una conexión a la base de datos `db`, el nombre del pedido `name`
// y el ID del establecimiento `establishmentID`.
// Devuelve el ID del nuevo pedido y un error en caso de que ocurra.
func CrearPedidoSinFW(db *sql.DB, name string, establishmentID int) (int64, error) {
	// Se ejecuta la instrucción SQL para insertar un nuevo pedido en la tabla pedido
	result, err := db.Exec("INSERT INTO pedidodb.pedido (Nombre, ID_Establecimiento) VALUES (?, ?)", name, establishmentID)
	if err != nil {
		return 0, err
	}
	// Retorna el ID del último pedido insertado
	return result.LastInsertId()

}

// RankingEstablecimientosSinFW genera un ranking de establecimientos basado en el total vendido.
// Recibe una conexión a la base de datos `db`, un ResponseWriter `w` y una solicitud `r`.
// Devuelve un JSON con los resultados del ranking.
func RankingEstablecimientosSinFW(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// Ejecuta la consulta SQL para obtener el ranking de establecimientos
	rows, err := db.Query(`
        SELECT e.RazonSocial, SUM(pp.Cantidad * p.Precio) AS Total_Vendido
        FROM Pedido_Producto pp
        JOIN Pedido p ON pp.ID_Pedido = p.ID_Pedido
        JOIN Producto pr ON pp.ID_Producto = pr.ID_Producto
        JOIN Establecimiento e ON p.ID_Establecimiento = e.ID_Establecimiento
        GROUP BY e.ID_Establecimiento
        ORDER BY Total_Vendido DESC
    `)
	if err != nil {
		// Si ocurre un error durante la consulta, devuelve un error HTTP 
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var resultados []map[string]interface{}
	// Itera sobre los resultados de la consulta
	for rows.Next() {
		var razonSocial string
		var totalVendido float64
		// Escanea los resultados de la fila actual en las variables correspondientes
		if err := rows.Scan(&razonSocial, &totalVendido); err != nil {
			// Si ocurre un error durante el escaneo, devuelve un error HTTP
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Añade el resultado actual al slice de resultados
		resultados = append(resultados, map[string]interface{}{
			"RazonSocial":   razonSocial,
			"Total_Vendido": totalVendido,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	// Codifica los resultados en JSON y los envía en la respuesta
	if err := json.NewEncoder(w).Encode(resultados); err != nil {
		// Si ocurre un error durante la codificación JSON, devuelve un error HTTP 500
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


// UbicacionEstablecimientoSinFW obtiene la ubicación de un establecimiento por su ID.
// Recibe una conexión a la base de datos `db`, un ResponseWriter `w` y una solicitud `r`.
// Devuelve un JSON con la ubicación del establecimiento
func UbicacionEstablecimientoSinFW(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// Obtiene el ID del establecimiento de los parámetros de la URL
	id := r.URL.Query().Get("id")
	// Ejecuta la consulta SQL para obtener la ubicación del establecimiento
	row := db.QueryRow("SELECT Ubicacion_Establecimiento FROM Establecimiento WHERE ID_Establecimiento = ?", id)

	var ubicacion string
	// Escanea el resultado de la consulta en la variable ubicación
	if err := row.Scan(&ubicacion); err != nil {
		http.Error(w, "Establecimiento no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	// Codifica la ubicación en JSON y la envía en la respuesta
	json.NewEncoder(w).Encode(map[string]interface{}{
		"Ubicacion_Establecimiento": ubicacion,
	})
}
