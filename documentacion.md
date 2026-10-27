# Documentación - TP Especial (TP3: Lógica de Negocio y API REST)

En esta etapa del Trabajo Especial conectamos la capa de persistencia generada por `sqlc` con una API REST construida exclusivamente con la biblioteca estándar de Go (`net/http`, `encoding/json`, `database/sql`), cumpliendo con los lineamientos de la cátedra de Programación Web.

## 1. Arquitectura y Separación en Capas

El proyecto sigue una arquitectura Three-Tier limpia y directa, sin sobreingeniería:

- **Capa de Persistencia (`db/sqlc/`)**:
  - Código generado automáticamente a partir de `schema.sql` y `querys.sql` utilizando `sqlc`.
  - Expone el tipo `*db.Queries` con métodos tipados para interactuar con PostgreSQL.
- **Capa de Lógica de Negocio Pura (`logic/`)**:
  - `logic/producto.go`, `logic/categoria.go` y `logic/cliente.go`.
  - Contiene funciones puras independientes de HTTP y de la base de datos (por ejemplo, `ValidateProducto`, `ValidateCategoria`, `ValidateCliente`).
  - Valida invariantes de dominio: nombres no vacíos, formato de emails, precios mayores a cero y stock no negativo.
- **Capa de Presentación y Servicios REST (`main.go`)**:
  - Servidor HTTP con `net/http.ServeMux`.
  - Inyección de dependencias mediante el struct `Server` que contiene `queries *db.Queries`.
  - Enrutamiento unificado por entidad siguiendo el patrón de la cátedra:
    - `/categorias/` y `/categorias/{id}`
    - `/productos/` y `/productos/{id}`
    - `/clientes/` y `/clientes/{id}`
    - `/` para archivos estáticos de `./static`.

## 2. Decisiones de Diseño y Manejo HTTP

- **Códigos de Estado Canónicos**:
  - `200 OK`: Consultas y actualizaciones exitosas.
  - `201 Created`: Creación de recursos vía POST.
  - `204 No Content`: Eliminación exitosa sin cuerpo de respuesta.
  - `400 Bad Request`: Payload JSON inválido, campos vacíos o reglas de validación fallidas.
  - `404 Not Found`: Identificador inexistente o `sql.ErrNoRows`.
  - `405 Method Not Allowed`: Métodos HTTP no soportados para la ruta.
  - `500 Internal Server Error`: Errores de conexión o fallos internos de base de datos.
- **Middleware de Logging**:
  - Implementado mediante la función canónica `loggingMiddleware(next http.Handler) http.Handler`.
  - Registra por consola cada petición con su método, ruta y dirección IP remota.
- **Resiliencia en el Arranque**:
  - Se implementó un bucle de reintento (`conn.Ping()`) de hasta 30 segundos en `main.go`, garantizando que el servidor espere a que el contenedor de PostgreSQL complete su inicialización antes de fallar.

## 3. Pruebas y Validación

- **Pruebas de Base de Datos**: Ejecutables con `make test`.
- **Pruebas de API REST**: Se incluye el script reproducible `requests.sh` con comandos `curl` y el archivo `requests.hurl` para verificar aserciones sobre respuestas y códigos HTTP.