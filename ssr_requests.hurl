# =====================================================================
# PRUEBAS DE INTEGRACIÓN SSR Y PATRÓN PRG CON HURL (TP5)
# =====================================================================

# 1. Cargar la página principal con Server-Side Rendering
GET http://localhost:8080/
HTTP 200
[Asserts]
header "Content-Type" contains "text/html"
xpath "string(//h1)" contains "MarkeTandil"
xpath "//form[@id='form-producto']" exists

# 2. Crear un producto mediante formulario tradicional (POST application/x-www-form-urlencoded)
POST http://localhost:8080/productos
[FormParams]
nombre: Yerba Mate Especial Hurl
id_categoria: 1
descripcion: Paquete 1kg edicion especial Hurl
precio: 3500.00
stock: 25
HTTP 303
[Asserts]
header "Location" contains "/"

# 3. Seguir la redirección y comprobar que el nuevo producto aparece en la tabla
GET http://localhost:8080/
HTTP 200
[Asserts]
xpath "//table[@id='tabla-productos']" exists
xpath "//td[strong[text()='Yerba Mate Especial Hurl']]" exists
xpath "string(//body)" contains "Yerba Mate Especial Hurl"

# 4. Validación de negocio: Intentar crear con precio inválido (<= 0)
POST http://localhost:8080/productos
[FormParams]
nombre: Producto Invalido
id_categoria: 1
descripcion: No deberia crearse
precio: -10.00
stock: 5
HTTP 303
[Asserts]
header "Location" contains "error="

# 5. Validación de negocio: Intentar crear sin nombre
POST http://localhost:8080/productos
[FormParams]
nombre: 
id_categoria: 1
descripcion: Sin nombre
precio: 100.00
stock: 5
HTTP 303
[Asserts]
header "Location" contains "error="
