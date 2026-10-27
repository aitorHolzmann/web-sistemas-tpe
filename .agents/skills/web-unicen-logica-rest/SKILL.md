---
name: web-unicen-logica-rest
description: >-
  Guía y patrones de la cátedra de Programación Web (UNICEN) para el diseño e implementación
  de la capa de Lógica de Negocio y APIs REST en Go. Usar para TP3, creación de endpoints HTTP,
  funciones puras de negocio, middleware, concurrencia (sync.RWMutex) y tests con cURL/Hurl.
---

# Lógica de Negocio y APIs REST (Cátedra Web UNICEN)

Esta skill contiene las pautas, patrones oficiales y ejemplos de código para la capa de **Lógica de Negocio** y exposición de **Servicios REST**, basados en `teoria/07-logic.md` y las consignas de `teoria/tp3.md`.

## 1. Principios Clave y Reglas de la Cátedra

1. **Biblioteca Estándar Exclusiva**:
   - Usar `net/http` para ruteo y handlers.
   - Usar `encoding/json` para serialización y deserialización.
   - **NO** usar routers ni frameworks de terceros (como Gin, Chi, Gorilla Mux o Echo).
2. **Lógica Pura y Desacoplada**:
   - Las reglas de negocio residen en el paquete `logic/` como funciones puras (sin dependencias de HTTP ni base de datos).
   - Firmas limpias: reciben structs o valores primitivos y devuelven el resultado o un `error`.
3. **Manejo de Respuestas HTTP**:
   - Establecer cabecera: `w.Header().Set("Content-Type", "application/json")`.
   - Códigos de estado HTTP canónicos:
     - `200 OK`: Consulta o actualización exitosa.
     - `201 Created`: Creación de recurso.
     - `204 No Content`: Eliminación exitosa (`w.WriteHeader(http.StatusNoContent)` sin body).
     - `400 Bad Request`: Datos inválidos o JSON malformado (`http.Error(w, err.Error(), http.StatusBadRequest)`).
     - `404 Not Found`: Recurso inexistente (`http.StatusNotFound`).
     - `405 Method Not Allowed`: Método no soportado (`http.StatusMethodNotAllowed`).

---

## 2. Patrones de Código de la Cátedra

### A. Lógica Pura (`logic/products.go`)
```go
package logic

import (
	"errors"
	"strings"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// ValidateProduct valida las reglas de negocio del producto.
func ValidateProduct(p Product) error {
	if strings.TrimSpace(p.Name) == "" {
		return errors.New("el nombre del producto no puede estar vacío")
	}
	if p.Price <= 0 {
		return errors.New("el precio debe ser mayor a cero")
	}
	return nil
}

// ApplyDiscount calcula y aplica un descuento porcentual retornando una copia modificada.
func ApplyDiscount(p Product, percentage float64) Product {
	if percentage > 0 && percentage <= 100 {
		p.Price = p.Price * (1 - percentage/100)
	}
	return p
}
```

### B. Router y Handlers REST con `net/http`
Patrón oficial de la cátedra para despachar según método y presencia de ID en la URL:

```go
// productsHandler maneja tanto /products como /products/{id}
func productsHandler(w http.ResponseWriter, r *http.Request) {
	// Obtenemos la parte de la ruta luego del prefijo
	path := strings.TrimPrefix(r.URL.Path, "/products")
	path = strings.Trim(path, "/")

	if path == "" {
		// Colección: /products
		switch r.Method {
		case http.MethodGet:
			getProducts(w, r)
		case http.MethodPost:
			createProduct(w, r)
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
		return
	}

	// Elemento individual: /products/{id}
	id, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, "ID de producto inválido", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		getProduct(w, r, id)
	case http.MethodPut:
		updateProduct(w, r, id)
	case http.MethodDelete:
		deleteProduct(w, r, id)
	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}
```

### C. Concurrencia con `sync.RWMutex` (Slice en memoria)
Para proteger variables globales compartidas:

```go
var (
	products = []logic.Product{}
	nextID   = 1
	mu       sync.RWMutex
)

func getProducts(w http.ResponseWriter, r *http.Request) {
	mu.RLock()
	defer mu.RUnlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	var p logic.Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	if err := logic.ValidateProduct(p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.Lock()
	p.ID = nextID
	nextID++
	products = append(products, p)
	mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}
```

### D. Middleware de Logging
Estructura clásica recomendada por Zunino y Teyseyre:

```go
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s desde %s", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}
```

---

## 3. Pruebas de la API: Hurl y cURL

La cátedra pide entregar pruebas reproducibles en `requests.hurl` o `requests.sh`.

### Ejemplo `requests.hurl`:
```hurl
# 1. Crear un producto
POST http://localhost:8080/products
Content-Type: application/json
{
  "name": "Teclado Mecánico",
  "price": 15000.0
}
HTTP 201
[Asserts]
jsonpath "$.id" exists
jsonpath "$.name" == "Teclado Mecánico"

# 2. Listar productos
GET http://localhost:8080/products
HTTP 200
[Asserts]
jsonpath "$" count >= 1

# 3. Eliminar producto
DELETE http://localhost:8080/products/1
HTTP 204
```

### Ejemplo `requests.sh` con cURL:
```bash
#!/usr/bin/env bash
BASE_URL="http://localhost:8080/products"

echo "=== Crear Producto ==="
curl -i -X POST "$BASE_URL" \
  -H "Content-Type: application/json" \
  -d '{"name": "Mouse Gamer", "price": 8500.0}'

echo -e "\n\n=== Listar Productos ==="
curl -i -X GET "$BASE_URL"

echo -e "\n\n=== Obtener Producto 1 ==="
curl -i -X GET "$BASE_URL/1"

echo -e "\n\n=== Eliminar Producto 1 ==="
curl -i -X DELETE "$BASE_URL/1"
```

