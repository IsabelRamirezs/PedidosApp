# PedidosApp

PedidosApp es una aplicación web para gestionar pedidos de establecimientos. Permite crear nuevos pedidos, consultar el ranking de establecimientos por ventas y obtener la ubicación de un establecimiento.

## Características

- **Crear Pedido**: Permite crear nuevos pedidos para un establecimiento especificando los productos y sus cantidades.
- **Ranking de Establecimientos**: Genera un ranking de los establecimientos basado en el total vendido.
- **Ubicación del Establecimiento**: Obtiene la ubicación de un establecimiento por su ID.


## Comenzando

### Prerrequisitos

- [Go](https://golang.org/doc/install)
- [Python](https://www.python.org/downloads/)
- [MySQL](https://dev.mysql.com/downloads/mysql/)
- [Flask](https://flask.palletsprojects.com/en/2.0.x/installation/)

### Instalación

1. **Configuración de la base de datos**

    Crear la base de datos y ejecutar el archivo `schema.sql` ubicado en `database/`.

    ```sql
    CREATE DATABASE PedidoDB;
    USE PedidoDB;
    SOURCE database/schema.sql;
    ```

2. **Backend en Go**

    Navega al directorio `pedido_golang` y ejecuta el servidor Go.

    ```bash
    cd pedido_golang
    go run main.go
    ```

3. **Backend en Python**

    Navega al directorio `pedido_python`, instala las dependencias y ejecuta el servidor Flask.

    ```bash
    cd pedido_python
    pip install -r requirements.txt
    python app.py
    ```

4. **Frontend**

    Abre `frontend/index.html` en tu navegador.

### Configuración

Configura las variables de conexión a la base de datos en `config.py` para la aplicación Flask:

```python
MYSQL_HOST = 'localhost'
MYSQL_USER = 'root'
MYSQL_PASSWORD = ''
MYSQL_DB = 'PedidoDB'
