package pedido

import (
	"database/sql"
	"net/http"

	"encoding/json"
	//"github.com/gin-gonic/gin"
	
)

func CrearPedidoSinFW(db *sql.DB, name string, establishmentID int) (int64, error) {
	result, err := db.Exec("INSERT INTO pedidodb.pedido (Nombre, ID_Establecimiento) VALUES (?, ?)", name, establishmentID)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()

}


// func CrearPedido(c *gin.Context, db *sql.DB) {
// 	var datos struct {
// 		Nombre            string `json:"Nombre"`
// 		IDEstablecimiento int    `json:"ID_Establecimiento"`
// 		Productos         []struct {
// 			IDProducto int `json:"ID_Producto"`
// 			Cantidad   int `json:"Cantidad"`
// 		} `json:"productos"`
// 	}

// 	if err := c.ShouldBindJSON(&datos); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	tx, err := db.Begin()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	res, err := tx.Exec("INSERT INTO Pedido (Nombre, ID_Establecimiento) VALUES (?, ?)", datos.Nombre, datos.IDEstablecimiento)
// 	if err != nil {
// 		tx.Rollback()
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	pedidoID, err := res.LastInsertId()
// 	if err != nil {
// 		tx.Rollback()
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	for _, producto := range datos.Productos {
// 		_, err := tx.Exec("INSERT INTO Pedido_Producto (ID_Pedido, ID_Producto, Cantidad) VALUES (?, ?, ?)", pedidoID, producto.IDProducto, producto.Cantidad)
// 		if err != nil {
// 			tx.Rollback()
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}
// 	}

// 	tx.Commit()
// 	c.JSON(http.StatusOK, gin.H{"mensaje": "Pedido creado exitosamente", "ID_Pedido": pedidoID})
// }

func RankingEstablecimientosSinFW(db *sql.DB, w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var resultados []map[string]interface{}
	for rows.Next() {
		var razonSocial string
		var totalVendido float64
		if err := rows.Scan(&razonSocial, &totalVendido); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		resultados = append(resultados, map[string]interface{}{
			"RazonSocial":   razonSocial,
			"Total_Vendido": totalVendido,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resultados); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// func RankingEstablecimientos(c *gin.Context, db *sql.DB) {
// 	rows, err := db.Query(`
//         SELECT e.RazonSocial, SUM(pp.Cantidad * p.Precio) AS Total_Vendido
//         FROM Pedido_Producto pp
//         JOIN Pedido p ON pp.ID_Pedido = p.ID_Pedido
//         JOIN Producto pr ON pp.ID_Producto = pr.ID_Producto
//         JOIN Establecimiento e ON p.ID_Establecimiento = e.ID_Establecimiento
//         GROUP BY e.ID_Establecimiento
//         ORDER BY Total_Vendido DESC
//     `)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	defer rows.Close()

// 	var resultados []map[string]interface{}
// 	for rows.Next() {
// 		var razonSocial string
// 		var totalVendido float64
// 		rows.Scan(&razonSocial, &totalVendido)
// 		resultados = append(resultados, map[string]interface{}{
// 			"RazonSocial":   razonSocial,
// 			"Total_Vendido": totalVendido,
// 		})
// 	}

// 	c.JSON(http.StatusOK, resultados)
// }

func UbicacionEstablecimientoSinFW(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	row := db.QueryRow("SELECT Ubicacion_Establecimiento FROM Establecimiento WHERE ID_Establecimiento = ?", id)

	var ubicacion string
	if err := row.Scan(&ubicacion); err != nil {
		http.Error(w, "Establecimiento no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"Ubicacion_Establecimiento": ubicacion,
	})
}


// func UbicacionEstablecimiento(c *gin.Context, db *sql.DB) {
// 	id := c.Param("id")
// 	row := db.QueryRow("SELECT Ubicacion_Establecimiento FROM Establecimiento WHERE ID_Establecimiento = ?", id)

// 	var ubicacion string
// 	if err := row.Scan(&ubicacion); err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Establecimiento no encontrado"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"Ubicacion_Establecimiento": ubicacion})
// }