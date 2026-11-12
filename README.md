# Trabajo Especial - TP5: Renderizado del Lado del Servidor con Templ

Aplicación web de **MarkeTandil** para gestión de **productos** y **categorías**, implementando **Server-Side Rendering (SSR)** con **Templ**, estilos semánticos con **Pico.css** y el patrón **Post/Redirect/Get (PRG)** en Go.

> [!NOTE]
> **Sin JavaScript**: Cumpliendo estrictamente con las consignas de la cátedra para el TP5, el proyecto no contiene archivos `.js`. Toda la lógica de presentación se ejecuta del lado del servidor (SSR) mediante componentes compilados de Templ.

---

## 📋 Requisitos Previos

Solo se requiere tener instalado en el host:

- **Docker** (con plugin Docker Compose)
- **Make** o **Bash**

Go, `sqlc`, `templ`, PostgreSQL y Hurl se ejecutan íntegramente dentro de los contenedores Docker configurados en `.devcontainer`.

---

## 🚀 Ejecución de la Aplicación

Para clonar y poner en marcha el proyecto con todas las tareas requeridas por la cátedra, ejecute desde la raíz de `tpe`:

```bash
make run-app
```

*(O de forma equivalente ejecutando directamente el script bash)*:

```bash
./run-app.sh
```

### Ciclo Completo Automatizado por `make run-app` / `./run-app.sh`:

1. **Tareas Previas**:
   - **Limpieza inicial**: Detiene y elimina contenedores y volúmenes previos para garantizar un estado limpio.
   - **Levantamiento**: Construye las imágenes y levanta los servicios (`db` y `app`).
   - **Espera activa de la BD**: Espera con `pg_isready` a que PostgreSQL esté listo y aceptando conexiones.
   - **Generación sqlc**: Ejecuta `sqlc generate` para la capa de acceso a datos.
   - **Compilación de vistas**: Ejecuta `templ generate` para compilar las plantillas `.templ` a código Go nativo.
   - **Compilación de la aplicación**: Compila el binario ejecutable del servidor Go.

2. **Ejecución de la Aplicación**:
   - Inicia el servidor HTTP en **http://localhost:8080**.
   - Acceda desde su navegador a: 👉 **http://localhost:8080**

3. **Tareas Posteriores**:
   - Al detener la ejecución (presionando `Ctrl+C` en la terminal), un manejador `trap` ejecuta automáticamente `docker compose down -v --remove-orphans`, liberando los puertos y borrando los contenedores y volúmenes creados.

---

## 🧪 Pruebas y Verificación

### 1. Pruebas de Integración SSR con Hurl
Con el servidor corriendo mediante `make run-app`, en **otra terminal** puede verificar la integración ejecutando:

```bash
make ssr-test
```

Valida:
- Renderizado HTML inicial con componentes Templ (`GET /`, HTTP 200).
- Envío de formulario tradicional (`POST /productos` con `application/x-www-form-urlencoded`).
- Redirección HTTP 303 (See Other) del patrón PRG.
- Carga de la lista con el producto recién creado.
- Validaciones de negocio del servidor (precio negativo, nombre vacío).

### 2. Pruebas Unitarias y de Base de Datos
Para correr las pruebas de CRUD sobre PostgreSQL en un contenedor efímero:

```bash
make test
```

### 3. Detención manual de contenedores
Si en cualquier momento desea forzar la detención y borrado de contenedores y volúmenes:

```bash
make down
```

---

## 🏗️ Arquitectura y Componentes Templ

Siguiendo las filminas de la cátedra (`09-view-static.pdf`, `10-view-dynamic.pdf`, `07-logic.pdf`, `12-slides.pdf`):

### Directorio `views/`
- **`layout.templ`**: Estructura base del documento (`<!DOCTYPE html>`, `<head>`, `<body>`, encabezado) con enlace local a **Pico.css** y estilos complementarios `/static/styles.css`. Aloja el contenido dinámico mediante `{ children... }`.
- **`entity_form.templ`**: Formulario tradicional (`<form method="POST" action="/productos">`) con campos para nombre, descripción, precio, stock y un `<select name="id_categoria">` poblado dinámicamente con las categorías de la base de datos.
- **`entity_list.templ`**: Componente que recibe `[]db.Producto` y un mapa de categorías, renderizando una tabla semántica (`<table class="striped">`) con los datos formateados y botones tradicionales de eliminación por producto (`POST /productos/{id}/eliminar`).
- **`index.templ`**: Componente `IndexPage` que ensambla el `Layout`, bloques de alerta para feedback (errores y éxitos) y organiza el formulario y el listado.

### Flujo Post/Redirect/Get (PRG)
1. **Envío de formulario**: El usuario completa el formulario y envía `POST /productos` con `application/x-www-form-urlencoded`.
2. **Validación en tres capas**: El servidor Go parsea los campos y delega la validación de reglas de negocio en `logic.ValidateProduct`.
3. **Persistencia**: Si es válido, inserta con `s.queries.CrearProducto`.
4. **Redirección**: En lugar de renderizar directamente en el POST, devuelve `HTTP 303 See Other` redirigiendo a `/?success=...`.
5. **Carga limpia**: El navegador ejecuta `GET /`, evitando reenvíos accidentales si el usuario actualiza la página (F5).

---

## 📁 Estructura del Proyecto

```
main.go              Servidor HTTP, rutas SSR y handlers del patrón PRG
run-app.sh           Script de tareas previas, ejecución y limpieza posterior
Makefile             Automatización (run-app, generate, ssr-test, test, down)
views/               Componentes Templ (.templ) y código Go compilado (*_templ.go)
  layout.templ       Plantilla base HTML5 con Pico.css
  entity_form.templ  Componente formulario tradicional de alta
  entity_list.templ  Componente tabla de productos y formulario de eliminación
  index.templ        Componente de página principal (IndexPage)
logic/               Capa de lógica pura del negocio (Three-Tier, validaciones)
db/schema/           schema.sql con la definición de tablas PostgreSQL
db/queries/          Consultas SQL tipadas para sqlc
db/sqlc/             Código generado por sqlc y tests unitarios
static/              Archivos estáticos servidos en /static/ (pico.min.css, styles.css)
ssr_requests.hurl    Pruebas de integración SSR y PRG con Hurl
DEFENSA_TP5.md       Guía teórica completa y preguntas para la defensa oral
.devcontainer/       Dockerfile (con sqlc, templ y hurl) y docker-compose.yml
```
