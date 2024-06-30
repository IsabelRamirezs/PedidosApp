from flask import Flask, request, jsonify
from flask_mysqldb import MySQL
import MySQLdb.cursors

app = Flask(__name__)
app.config.from_pyfile('config.py')

mysql = MySQL(app)

@app.route('/pedido', methods=['POST'])
def crear_pedido():
    datos = request.get_json()
    establecimiento_id = datos['ID_Establecimiento']
    productos = datos['productos']
    
    cursor = mysql.connection.cursor()
    cursor.execute('INSERT INTO Pedido (Nombre, ID_Establecimiento) VALUES (%s, %s)', (datos['Nombre'], establecimiento_id))
    pedido_id = cursor.lastrowid
    
    for producto in productos:
        cursor.execute('INSERT INTO Pedido_Producto (ID_Pedido, ID_Producto, Cantidad) VALUES (%s, %s, %s)', 
                       (pedido_id, producto['ID_Producto'], producto['Cantidad']))
    
    mysql.connection.commit()
    cursor.close()
    
    return jsonify({'mensaje': 'Pedido creado exitosamente', 'ID_Pedido': pedido_id})

@app.route('/ranking', methods=['GET'])
def ranking_establecimientos():
    cursor = mysql.connection.cursor(MySQLdb.cursors.DictCursor)
    cursor.execute('''
        SELECT e.RazonSocial, SUM(pp.Cantidad * pr.Precio) AS Total_Vendido
        FROM Pedido_Producto pp
        JOIN Pedido p ON pp.ID_Pedido = p.ID_Pedido
        JOIN Producto pr ON pp.ID_Producto = pr.ID_Producto
        JOIN Establecimiento e ON p.ID_Establecimiento = e.ID_Establecimiento
        GROUP BY e.ID_Establecimiento
        ORDER BY Total_Vendido DESC
    ''')
    resultados = cursor.fetchall()
    cursor.close()
    return jsonify(resultados)

@app.route('/ubicacion/<int:id>', methods=['GET'])
def ubicacion_establecimiento(id):
    cursor = mysql.connection.cursor(MySQLdb.cursors.DictCursor)
    cursor.execute('SELECT Ubicacion_Establecimiento FROM Establecimiento WHERE ID_Establecimiento = %s', [id])
    resultado = cursor.fetchone()
    cursor.close()
    return jsonify(resultado)

if __name__ == '__main__':
    app.run(debug=True, host='127.0.0.1', port=8000)
