#!/usr/bin/env bash

set -e

BASE_URL="http://localhost:8080"

echo "=== 1. Probar servidor de archivos estaticos / ==="
curl -s -o /dev/null -w "%{http_code}\n" "$BASE_URL/"

echo "=== 2. Crear Categoria (POST /categorias) ==="
CAT_RESP=$(curl -s -i -X POST "$BASE_URL/categorias" \
  -H "Content-Type: application/json" \
  -d '{"nombre": "Hardware", "descripcion": "Componentes de computacion"}')
echo "$CAT_RESP" | head -n 1
CAT_ID=$(echo "$CAT_RESP" | grep -o '"id":[0-9]*' | head -n 1 | cut -d: -f2)
echo "Categoria creada con ID: $CAT_ID"

echo "=== 3. Listar Categorias (GET /categorias) ==="
curl -s -X GET "$BASE_URL/categorias"
echo ""

echo "=== 4. Obtener Categoria por ID (GET /categorias/$CAT_ID) ==="
curl -s -X GET "$BASE_URL/categorias/$CAT_ID"
echo ""

echo "=== 5. Actualizar Categoria (PUT /categorias/$CAT_ID) ==="
curl -s -X PUT "$BASE_URL/categorias/$CAT_ID" \
  -H "Content-Type: application/json" \
  -d '{"nombre": "Hardware y Perifericos", "descripcion": "Componentes y accesorios"}'
echo ""

echo "=== 6. Crear Producto Valido (POST /productos) ==="
PROD_RESP=$(curl -s -i -X POST "$BASE_URL/productos" \
  -H "Content-Type: application/json" \
  -d "{\"nombre\": \"Teclado Mecanico\", \"descripcion\": \"RGB switch blue\", \"stock\": \"25\", \"precio\": \"18500.50\", \"foto\": \"https://example.com/teclado.jpg\", \"id_categoria\": $CAT_ID}")
echo "$PROD_RESP" | head -n 1
PROD_ID=$(echo "$PROD_RESP" | grep -o '"id":[0-9]*' | head -n 1 | cut -d: -f2)
echo "Producto creado con ID: $PROD_ID"

echo "=== 7. Listar Productos (GET /productos) ==="
curl -s -X GET "$BASE_URL/productos"
echo ""

echo "=== 8. Obtener Producto por ID (GET /productos/$PROD_ID) ==="
curl -s -X GET "$BASE_URL/productos/$PROD_ID"
echo ""

echo "=== 9. Actualizar Producto (PUT /productos/$PROD_ID) ==="
curl -s -X PUT "$BASE_URL/productos/$PROD_ID" \
  -H "Content-Type: application/json" \
  -d "{\"nombre\": \"Teclado Mecanico Pro\", \"descripcion\": \"RGB switch red silencioso\", \"stock\": \"10\", \"precio\": \"22000.00\", \"foto\": \"https://example.com/teclado_red.jpg\", \"id_categoria\": $CAT_ID}"
echo ""

echo "=== 10. Validacion: POST /productos con nombre vacio (esperado: 400 Bad Request) ==="
curl -s -i -X POST "$BASE_URL/productos" \
  -H "Content-Type: application/json" \
  -d "{\"nombre\": \"\", \"stock\": \"5\", \"precio\": \"100\", \"id_categoria\": $CAT_ID}" | head -n 1

echo "=== 11. Validacion: POST /productos con precio negativo (esperado: 400 Bad Request) ==="
curl -s -i -X POST "$BASE_URL/productos" \
  -H "Content-Type: application/json" \
  -d "{\"nombre\": \"Mouse\", \"stock\": \"5\", \"precio\": \"-100\", \"id_categoria\": $CAT_ID}" | head -n 1

echo "=== 12. Validacion: POST /productos con categoria inexistente (esperado: 400 Bad Request) ==="
curl -s -i -X POST "$BASE_URL/productos" \
  -H "Content-Type: application/json" \
  -d '{"nombre": "Mouse", "stock": "5", "precio": "5000", "id_categoria": 99999}' | head -n 1

echo "=== 13. Buscar Producto Inexistente (esperado: 404 Not Found) ==="
curl -s -i -X GET "$BASE_URL/productos/99999" | head -n 1

echo "=== 14. Metodo No Permitido (esperado: 405 Method Not Allowed) ==="
curl -s -i -X PATCH "$BASE_URL/productos" | head -n 1

echo "=== 15. Crear Cliente (POST /clientes) ==="
CLI_RESP=$(curl -s -i -X POST "$BASE_URL/clientes" \
  -H "Content-Type: application/json" \
  -d "{\"nombre\": \"Carlos\", \"apellido\": \"Gomez\", \"email\": \"carlos.$RANDOM@example.com\", \"direccion\": \"Av. Colon 1234\"}")
echo "$CLI_RESP" | head -n 1
CLI_ID=$(echo "$CLI_RESP" | grep -o '"id":[0-9]*' | head -n 1 | cut -d: -f2)
echo "Cliente creado con ID: $CLI_ID"

echo "=== 16. Listar Clientes (GET /clientes) ==="
curl -s -X GET "$BASE_URL/clientes"
echo ""

echo "=== 17. Eliminar Producto (DELETE /productos/$PROD_ID) (esperado: 204 No Content) ==="
curl -s -i -X DELETE "$BASE_URL/productos/$PROD_ID" | head -n 1

echo "=== 18. Verificar Producto Eliminado (GET /productos/$PROD_ID) (esperado: 404 Not Found) ==="
curl -s -i -X GET "$BASE_URL/productos/$PROD_ID" | head -n 1

echo "=== 19. Eliminar Cliente (DELETE /clientes/$CLI_ID) (esperado: 204 No Content) ==="
curl -s -i -X DELETE "$BASE_URL/clientes/$CLI_ID" | head -n 1

echo "=== 20. Eliminar Categoria (DELETE /categorias/$CAT_ID) (esperado: 204 No Content) ==="
curl -s -i -X DELETE "$BASE_URL/categorias/$CAT_ID" | head -n 1

echo "=== Todas las pruebas completadas exitosamente ==="
