from flask import Flask, request, jsonify
from flask_mysqldb import MySQL
import MySQLdb.cursors

app = Flask(__name__)
app.config.from_pyfile('config.py') # Carga la configuración desde el archivo config.py

mysql = MySQL(app) # Inicializa la extensión MySQL para Flask


# Handler para crear un pedido
@app.route('/pedido', methods=['POST'])
def crear_pedido():
    datos = request.get_json() # Obtiene los datos JSON de la solicitud
    establecimiento_id = datos['ID_Establecimiento']
    productos = datos['productos']
    
    cursor = mysql.connection.cursor() # Obtiene un cursor para interactuar con la base de datos
    # Inserta un nuevo pedido en la tabla Pedido
    cursor.execute('INSERT INTO Pedido (Nombre, ID_Establecimiento) VALUES (%s, %s)', (datos['Nombre'], establecimiento_id))
    pedido_id = cursor.lastrowid # Obtiene el ID del último pedido insertado
    
    # Inserta los productos del pedido en la tabla Pedido_Producto
    for producto in productos:
        cursor.execute('INSERT INTO Pedido_Producto (ID_Pedido, ID_Producto, Cantidad) VALUES (%s, %s, %s)', 
                       (pedido_id, producto['ID_Producto'], producto['Cantidad']))
    
    mysql.connection.commit()
    cursor.close()
    
    # Retorna una respuesta JSON indicando el éxito de la operación y el ID del pedido creado
    return jsonify({'mensaje': 'Pedido creado exitosamente', 'ID_Pedido': pedido_id})


# Handler para obtener el ranking de establecimientos por ventas
@app.route('/ranking', methods=['GET'])
def ranking_establecimientos():
    cursor = mysql.connection.cursor(MySQLdb.cursors.DictCursor)
    # Consulta SQL para obtener el ranking de establecimientos por ventas
    cursor.execute('''
        SELECT e.RazonSocial, SUM(pp.Cantidad * pr.Precio) AS Total_Vendido
        FROM Pedido_Producto pp
        JOIN Pedido p ON pp.ID_Pedido = p.ID_Pedido
        JOIN Producto pr ON pp.ID_Producto = pr.ID_Producto
        JOIN Establecimiento e ON p.ID_Establecimiento = e.ID_Establecimiento
        GROUP BY e.ID_Establecimiento
        ORDER BY Total_Vendido DESC
    ''')
    resultados = cursor.fetchall() # Se obtienen los resultados de da consulta
    cursor.close()
    return jsonify(resultados)

# Handler para obtener ubicación
@app.route('/ubicacion/<int:id>', methods=['GET'])
def ubicacion_establecimiento(id):
    cursor = mysql.connection.cursor(MySQLdb.cursors.DictCursor)
    # Consulta SQL para obtener la ubicación de un establecimiento por su ID
    cursor.execute('SELECT Ubicacion_Establecimiento FROM Establecimiento WHERE ID_Establecimiento = %s', [id])
    resultado = cursor.fetchone()
    cursor.close()
    # Retorna el resultado como JSON
    return jsonify(resultado)

if __name__ == '__main__':
    app.run(debug=True, host='127.0.0.1', port=8000)
