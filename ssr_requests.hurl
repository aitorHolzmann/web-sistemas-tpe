# 1. cargar la pagina principal (estado vacio inicial)
GET http://localhost:8080/
HTTP 200
[Asserts]
header "Content-Type" contains "text/html"
xpath "string(//h1)" contains "MarkeTandil"
body contains "htmx.org@4.0.0"
xpath "//form[@id='form-producto']" exists
xpath "string(//form[@id='form-producto']/@hx-post)" == "/productos"
xpath "string(//form[@id='form-producto']/@hx-target)" == "#lista-productos"
xpath "string(//form[@id='form-producto']/@hx-swap)" == "outerHTML"
xpath "//div[@id='mensajes-feedback']" exists
xpath "//p[@class='vacio']" exists

# 2. crear un producto mediante htmx sin recarga de pagina
POST http://localhost:8080/productos
HX-Request: true
[FormParams]
nombre: Yerba Mate Especial Hurl
id_categoria: 1
descripcion: Paquete 1kg edicion especial Hurl
precio: 3500.00
stock: 25
HTTP 200
[Asserts]
header "Content-Type" contains "text/html"
xpath "//section[@id='lista-productos']" exists
xpath "string(//body)" contains "Yerba Mate Especial Hurl"
xpath "//hx-partial" exists
body not contains "<!DOCTYPE html>"
[Captures]
producto_id: regex "hx-delete=\"/productos/([0-9]+)\""

# 3. obtener solo el fragmento de la lista usando htmx
GET http://localhost:8080/productos
HX-Request: true
HTTP 200
[Asserts]
xpath "//section[@id='lista-productos']" exists
body not contains "<!DOCTYPE html>"

# 4. validacion de negocio: precio invalido
POST http://localhost:8080/productos
HX-Request: true
[FormParams]
nombre: Producto Invalido
id_categoria: 1
descripcion: No deberia crearse
precio: -10.00
stock: 5
HTTP 400
[Asserts]
xpath "//hx-partial" exists
xpath "//div[contains(@class, 'mensaje error')]" exists
xpath "//section[@id='lista-productos']" not exists

# 5. validacion de negocio: nombre vacio
POST http://localhost:8080/productos
HX-Request: true
[FormParams]
nombre: 
id_categoria: 1
descripcion: Sin nombre
precio: 100.00
stock: 5
HTTP 400
[Asserts]
xpath "//hx-partial" exists
xpath "//div[contains(@class, 'mensaje error')]" exists

# 6. eliminar el producto creado sin recarga
DELETE http://localhost:8080/productos/{{producto_id}}
HX-Request: true
HTTP 200
[Asserts]
body == ""

# 7. eliminar producto inexistente
DELETE http://localhost:8080/productos/999999
HX-Request: true
HTTP 404
[Asserts]
xpath "//hx-partial" exists
xpath "//div[contains(@class, 'mensaje error')]" exists

# 8. comprobar en la pagina principal que el producto ya no esta
GET http://localhost:8080/
HTTP 200
[Asserts]
body not contains "Yerba Mate Especial Hurl"
xpath "//p[@class='vacio']" exists
