CREATE DATABASE IF NOT EXISTS PedidoDB;

USE PedidoDB;

CREATE TABLE IF NOT EXISTS Establecimiento (
    ID_Establecimiento INT PRIMARY KEY AUTO_INCREMENT,
    RazonSocial VARCHAR(20) NOT NULL,
    Ubicacion_Establecimiento VARCHAR(20) NOT NULL
);

CREATE TABLE IF NOT EXISTS Producto (
    ID_Producto INT PRIMARY KEY AUTO_INCREMENT,
    Nombre VARCHAR(20) NOT NULL,
    Precio FLOAT NOT NULL
);

CREATE TABLE IF NOT EXISTS Pedido (
    ID_Pedido INT PRIMARY KEY AUTO_INCREMENT,
    Nombre VARCHAR(20) NOT NULL,
    ID_Establecimiento INT NOT NULL,
    FOREIGN KEY (ID_Establecimiento) REFERENCES Establecimiento(ID_Establecimiento)
);

CREATE TABLE IF NOT EXISTS Pedido_Producto (
    ID_Pedido INT NOT NULL,
    ID_Producto INT NOT NULL,
    Cantidad INT NOT NULL,
    FOREIGN KEY (ID_Pedido) REFERENCES Pedido(ID_Pedido),
    FOREIGN KEY (ID_Producto) REFERENCES Producto(ID_Producto),
    PRIMARY KEY (ID_Pedido, ID_Producto)
);
