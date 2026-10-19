# Trabajo de Cursada - TP3: Exponiendo la Lógica como API

API REST de **MarkeTandil** para gestionar **categorías**, **productos** y **clientes**.


## Requisitos

Solo hace falta tener instalado:

- **Docker** (con Docker Compose)
- **Make** 

Go, `sqlc`, PostgreSQL y Hurl se ejecutan dentro de los contenedores, no hace falta instalarlos.

## 1. Ejecutar el servidor

Desde la raíz del proyecto:

```bash
make run
```

Este comando:

1. Construye la imagen y levanta dos contenedores: la base de datos PostgreSQL (con las tablas de `db/schema/schema.sql`) y la app.
2. Genera el código de acceso a datos con `sqlc generate`.
3. Inicia el servidor en **http://localhost:8080**.

Cuando el servidor está listo se ve en consola:

```
Conectado a la base de datos correctamente
Servidor escuchando en http://localhost:8080
```

## 2. Probar la API

Con el servidor corriendo, en **otra terminal**:

```bash
make api-test
```

Ejecuta el archivo `requests.hurl`, que prueba todas las operaciones CRUD de las tres entidades (crear, listar, ver, actualizar y eliminar), además de los casos de error (400 por datos inválidos y 404 por recursos inexistentes). Si todo funciona, termina con:

```
requests.hurl: Success (25 request(s) in ...)
```

## 3. Detener todo

Cortar el servidor con `Ctrl+C` y luego:

```bash
make down
```

Esto baja los contenedores y borra los datos de la base.


## Endpoints

| Método | Ruta | Descripción | Respuesta |
|---|---|---|---|
| `GET` | `/categorias` | Lista todas las categorías | 200 |
| `POST` | `/categorias` | Crea una categoría | 201 / 400 |
| `GET` | `/categorias/{id}` | Obtiene una categoría | 200 / 404 |
| `PUT` | `/categorias/{id}` | Actualiza una categoría | 200 / 400 / 404 |
| `DELETE` | `/categorias/{id}` | Elimina una categoría | 204 / 404 |
| `GET` | `/productos` | Lista todos los productos | 200 |
| `POST` | `/productos` | Crea un producto | 201 / 400 |
| `GET` | `/productos/{id}` | Obtiene un producto | 200 / 404 |
| `PUT` | `/productos/{id}` | Actualiza un producto | 200 / 400 / 404 |
| `DELETE` | `/productos/{id}` | Elimina un producto | 204 / 404 |
| `GET` | `/clientes` | Lista todos los clientes | 200 |
| `POST` | `/clientes` | Crea un cliente | 201 / 400 |
| `GET` | `/clientes/{id}` | Obtiene un cliente | 200 / 404 |
| `PUT` | `/clientes/{id}` | Actualiza un cliente | 200 / 400 / 404 |
| `DELETE` | `/clientes/{id}` | Elimina un cliente | 204 / 404 |

### Ejemplos de JSON

Categoría (el `nombre` es obligatorio):

```json
{ "nombre": "Bebidas", "descripcion": "Gaseosas y jugos" }
```

Producto (`nombre` obligatorio, `precio` mayor a 0, `stock` no negativo y `id_categoria` de una categoría existente):

```json
{ "nombre": "Agua mineral", "descripcion": "Botella 2L", "stock": 20, "precio": 1500.50, "id_categoria": 1 }
```

Cliente (todos los campos obligatorios y `email` con `@`):

```json
{ "nombre": "Juan", "apellido": "Perez", "email": "juan@mail.com", "direccion": "Pinto 399, Tandil" }
```

Ejemplo con curl:

```bash
curl -X POST http://localhost:8080/categorias \
  -H "Content-Type: application/json" \
  -d '{"nombre": "Bebidas", "descripcion": "Gaseosas y jugos"}'
```

## Estructura del proyecto

```
main.go              conexión a la base de datos, rutas y handlers HTTP
logic/               validaciones de negocio (funciones puras)
db/schema/           schema.sql con las tablas
db/queries/          consultas SQL para sqlc
db/sqlc/             código generado por sqlc y tests de la capa de datos
requests.hurl        pruebas de la API
.devcontainer/       Dockerfile y docker-compose.yml
```

## Tests de la capa de datos (TP2)

```bash
make test
```

Levanta una base de datos limpia, corre los tests de `db/sqlc/` y al terminar baja los contenedores.
