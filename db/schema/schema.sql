CREATE TABLE categoria (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(255) NOT NULL,
    descripcion TEXT
);

CREATE TABLE producto (
    id SERIAL PRIMARY KEY,
    id_categoria INTEGER NOT NULL,
    nombre VARCHAR(255) NOT NULL,
    descripcion VARCHAR(255),
    stock NUMERIC NOT NULL,
    precio NUMERIC(7,2) NOT NULL,
    foto VARCHAR(500)
);

CREATE TABLE cliente (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(255) NOT NULL,
    apellido VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    direccion VARCHAR(255) NOT NULL
);

ALTER TABLE producto
ADD CONSTRAINT fk_categoria_producto FOREIGN KEY (id_categoria) REFERENCES categoria(id) ON DELETE CASCADE;

