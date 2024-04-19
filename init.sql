-- Crear la base de datos
CREATE DATABASE inventario;

-- Conectar a la base de datos
\c inventario;

-- Crear la tabla Marca
CREATE TABLE Marca (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL
);

-- Crear la tabla Categoría
CREATE TABLE Categoria (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL
);

-- Crear la tabla Producto
CREATE TABLE Producto (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    id_marca INT REFERENCES Marca(id),
    id_categoria INT REFERENCES Categoria(id),
    precio DECIMAL(10, 2) NOT NULL,
    stock INT NOT NULL
);
