---------------------------------------------------------------------------------------------------------
---------------------------------------------------------------------------------------------------------

-- name: CrearCliente :one
INSERT INTO cliente (nombre, apellido, email, direccion) VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetCliente :one
SELECT * FROM cliente WHERE id = $1;

-- name: ListarClientes :many
SELECT * FROM cliente ORDER BY apellido;

-- name: CambiarDireccion :exec
UPDATE cliente SET direccion = $2 WHERE id = $1;

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

-- name: CambiarPrecio :exec
UPDATE producto SET precio = $2 WHERE id = $1;

-- name: GetStock :one
SELECT stock FROM producto WHERE id = $1;

-- name: ActualizarStock :exec
UPDATE producto SET stock = $2 WHERE id = $1;

-- name: ActualizarFoto :exec
UPDATE producto SET foto = $2 WHERE id = $1;

-- name: BorrarProducto :exec
DELETE FROM producto WHERE id = $1;

---------------------------------------------------------------------------------------------------------
---------------------------------------------------------------------------------------------------------

-- name: ListarCategorias :many
SELECT id, nombre, descripcion FROM categoria ORDER BY nombre;

-- name: GetCategoria :one
SELECT id, nombre, descripcion FROM categoria WHERE id = $1;

-- name: InsertarCategoria :one
INSERT INTO categoria (nombre, descripcion) VALUES ($1, $2) RETURNING *;

-- name: ActualizarCategoria :exec
UPDATE categoria SET nombre = $1, descripcion = $2 WHERE id = $3;

-- name: BorrarCategoria :exec
DELETE FROM categoria WHERE id = $1;
