# =====================================================================
# 1. CRUD DE CATEGORIAS
# =====================================================================

# crear una categoria
POST http://localhost:8080/categorias
Content-Type: application/json
{
  "nombre": "Perifericos Hurl",
  "descripcion": "Categoria creada desde Hurl"
}
HTTP 201
[Captures]
cat_id: jsonpath "$.id"
[Asserts]
jsonpath "$.id" exists
jsonpath "$.nombre" == "Perifericos Hurl"

# listar categorias
GET http://localhost:8080/categorias
HTTP 200
[Asserts]
jsonpath "$" count >= 1

# obtener categoria por id
GET http://localhost:8080/categorias/{{cat_id}}
HTTP 200
[Asserts]
jsonpath "$.id" == {{cat_id}}
jsonpath "$.nombre" == "Perifericos Hurl"

# actualizar categoria por id
PUT http://localhost:8080/categorias/{{cat_id}}
Content-Type: application/json
{
  "nombre": "Perifericos y Accesorios",
  "descripcion": "Categoria actualizada desde Hurl"
}
HTTP 200
[Asserts]
jsonpath "$.id" == {{cat_id}}
jsonpath "$.nombre" == "Perifericos y Accesorios"

# validacion: categoria con nombre vacio (400 Bad Request)
POST http://localhost:8080/categorias
Content-Type: application/json
{
  "nombre": "",
  "descripcion": "Invalida"
}
HTTP 400

# buscar categoria inexistente (404 Not Found)
GET http://localhost:8080/categorias/99999
HTTP 404


# =====================================================================
# 2. CRUD DE PRODUCTOS
# =====================================================================

# crear un producto asociado a la categoria creada
POST http://localhost:8080/productos
Content-Type: application/json
{
  "nombre": "Auriculares Gamer",
  "descripcion": "Auriculares 7.1 con microfono",
  "stock": 15,
  "precio": 12500.00,
  "id_categoria": {{cat_id}}
}
HTTP 201
[Captures]
prod_id: jsonpath "$.id"
[Asserts]
jsonpath "$.id" exists
jsonpath "$.nombre" == "Auriculares Gamer"
jsonpath "$.stock" == 15
jsonpath "$.precio" == 12500

# listar productos
GET http://localhost:8080/productos
HTTP 200
[Asserts]
jsonpath "$" count >= 1

# obtener producto por id
GET http://localhost:8080/productos/{{prod_id}}
HTTP 200
[Asserts]
jsonpath "$.id" == {{prod_id}}
jsonpath "$.nombre" == "Auriculares Gamer"

# actualizar producto por id
PUT http://localhost:8080/productos/{{prod_id}}
Content-Type: application/json
{
  "nombre": "Auriculares Gamer Pro",
  "descripcion": "Auriculares inalambricos",
  "stock": 8,
  "precio": 18000.50,
  "id_categoria": {{cat_id}}
}
HTTP 200
[Asserts]
jsonpath "$.id" == {{prod_id}}
jsonpath "$.nombre" == "Auriculares Gamer Pro"
jsonpath "$.stock" == 8

# validacion: producto con nombre vacio (400 Bad Request)
POST http://localhost:8080/productos
Content-Type: application/json
{
  "nombre": "",
  "stock": 10,
  "precio": 100.00,
  "id_categoria": {{cat_id}}
}
HTTP 400

# validacion: producto con categoria inexistente (400 Bad Request)
POST http://localhost:8080/productos
Content-Type: application/json
{
  "nombre": "Mouse",
  "stock": 10,
  "precio": 100.00,
  "id_categoria": 99999
}
HTTP 400

# buscar producto inexistente (404 Not Found)
GET http://localhost:8080/productos/99999
HTTP 404


# =====================================================================
# 3. CRUD DE CLIENTES
# =====================================================================

# crear un cliente
POST http://localhost:8080/clientes
Content-Type: application/json
{
  "nombre": "Esteban",
  "apellido": "Perez",
  "email": "esteban.hurl@example.com",
  "direccion": "Pinto 399, Tandil"
}
HTTP 201
[Captures]
cli_id: jsonpath "$.id"
[Asserts]
jsonpath "$.id" exists
jsonpath "$.nombre" == "Esteban"
jsonpath "$.email" == "esteban.hurl@example.com"

# listar clientes
GET http://localhost:8080/clientes
HTTP 200
[Asserts]
jsonpath "$" count >= 1

# obtener cliente por id
GET http://localhost:8080/clientes/{{cli_id}}
HTTP 200
[Asserts]
jsonpath "$.id" == {{cli_id}}
jsonpath "$.apellido" == "Perez"

# actualizar cliente por id
PUT http://localhost:8080/clientes/{{cli_id}}
Content-Type: application/json
{
  "nombre": "Esteban Actualizado",
  "apellido": "Perez",
  "email": "esteban.hurl@example.com",
  "direccion": "San Martin 850, Tandil"
}
HTTP 200
[Asserts]
jsonpath "$.id" == {{cli_id}}
jsonpath "$.nombre" == "Esteban Actualizado"

# validacion: cliente con email sin @ (400 Bad Request)
POST http://localhost:8080/clientes
Content-Type: application/json
{
  "nombre": "Juan",
  "apellido": "Gomez",
  "email": "email_invalido",
  "direccion": "Belgrano 100"
}
HTTP 400

# buscar cliente inexistente (404 Not Found)
GET http://localhost:8080/clientes/99999
HTTP 404


# =====================================================================
# 4. ELIMINACION (DELETE -> 204 No Content)
# =====================================================================

# eliminar producto creado
DELETE http://localhost:8080/productos/{{prod_id}}
HTTP 204

# verificar que el producto ya no existe
GET http://localhost:8080/productos/{{prod_id}}
HTTP 404

# eliminar cliente creado
DELETE http://localhost:8080/clientes/{{cli_id}}
HTTP 204

# verificar que el cliente ya no existe
GET http://localhost:8080/clientes/{{cli_id}}
HTTP 404

# eliminar categoria creada
DELETE http://localhost:8080/categorias/{{cat_id}}
HTTP 204

# verificar que la categoria ya no existe
GET http://localhost:8080/categorias/{{cat_id}}
HTTP 404
