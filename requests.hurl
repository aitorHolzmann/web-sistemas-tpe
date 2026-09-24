POST http://localhost:8080/categorias
Content-Type: application/json
{
  "nombre": "Perifericos Hurl",
  "descripcion": "Categoria creada desde Hurl"
}
HTTP 201
[Asserts]
jsonpath "$.id" exists
jsonpath "$.nombre" == "Perifericos Hurl"

GET http://localhost:8080/categorias
HTTP 200
[Asserts]
jsonpath "$" count >= 1

POST http://localhost:8080/productos
Content-Type: application/json
{
  "nombre": "Auriculares Gamer",
  "descripcion": "Auriculares 7.1 con microfono",
  "stock": "15",
  "precio": "12500.00",
  "foto": "https://example.com/auriculares.jpg",
  "id_categoria": 1
}
HTTP 201
[Asserts]
jsonpath "$.id" exists
jsonpath "$.nombre" == "Auriculares Gamer"
jsonpath "$.stock" == "15"

GET http://localhost:8080/productos
HTTP 200
[Asserts]
jsonpath "$" count >= 1

GET http://localhost:8080/productos/99999
HTTP 404

POST http://localhost:8080/productos
Content-Type: application/json
{
  "nombre": "",
  "stock": "10",
  "precio": "100.00",
  "id_categoria": 1
}
HTTP 400
