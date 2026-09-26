---------------------------------------------------------------------------------------------------------
---------------------------------------------------------------------------------------------------------

-- name: CrearCliente :one
INSERT INTO cliente (nombre, apellido, email, direccion) VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetCliente :one
SELECT * FROM cliente WHERE id = $1;

-- name: ListarClientes :many
SELECT * FROM cliente ORDER BY apellido;

-- name: ActualizarCliente :exec
UPDATE cliente SET nombre = $2, apellido = $3, email = $4, direccion = $5 WHERE id = $1;

-- name: BorrarCliente :exec
DELETE FROM cliente WHERE id = $1;

---------------------------------------------------------------------------------------------------------
---------------------------------------------------------------------------------------------------------

-- name: CrearProducto :one
INSERT INTO producto (nombre, descripcion, stock, precio, foto, id_categoria) VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetProducto :one
SELECT * FROM producto WHERE id = $1;

-- name: ListarProductos :many
SELECT * FROM producto ORDER BY nombre;

-- name: ActualizarProducto :one
UPDATE producto SET nombre = $2, descripcion = $3, stock = $4, precio = $5, foto = $6, id_categoria = $7 WHERE id = $1
RETURNING *;

-- name: BorrarProducto :exec
DELETE FROM producto WHERE id = $1;

---------------------------------------------------------------------------------------------------------
---------------------------------------------------------------------------------------------------------

-- name: ListarCategorias :many
SELECT * FROM categoria ORDER BY nombre;

-- name: GetCategoria :one
SELECT * FROM categoria WHERE id = $1;

-- name: CrearCategoria :one
INSERT INTO categoria (nombre, descripcion) VALUES ($1, $2) RETURNING *;

-- name: ActualizarCategoria :exec
UPDATE categoria SET nombre = $1, descripcion = $2 WHERE id = $3;

-- name: BorrarCategoria :exec
DELETE FROM categoria WHERE id = $1;
