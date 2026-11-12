# Documentación Técnica Completa: Implementación de la Parte 5 (TPE - Programación Web)
**Para uso como contexto / prompt para IA en VS Code**  
**Cátedra:** Dr. Alejandro Zunino & Dr. Alfredo Teyseyre (ISISTAN / UNICEN)  
**Materia:** Programación Web  
**Proyecto:** MarkeTandil  

---

## 🎯 Objetivo de este documento

Este documento contiene el **registro técnico exhaustivo de todos los cambios, decisiones de diseño, archivos modificados y código implementado** en la Parte 5 del Trabajo Especial. Está estructurado para que cualquier asistente de IA en VS Code (como GitHub Copilot o Claude) pueda generar una **guía paso a paso con referencias directas a archivos, funciones y líneas de código**.

---

## 📚 1. Contexto y Transición Arquitectónica (TP4 ➔ TP5)

### ¿Qué se tenía en TP4?
- **Arquitectura:** API REST desacoplada + Frontend SPA / AJAX con Vanilla JS.
- **Frontend:** `static/index.html`, `static/app.js` (con `fetch`, eventos DOM, delegación de eventos para borrado) y `static/styles.css`.
- **Backend:** `main.go` exponiendo rutas JSON (`/productos`, `/categorias`, `/clientes`) para todos los verbos HTTP (`GET`, `POST`, `PUT`, `DELETE`). El servidor servía estáticos en `/` montando `http.FileServer(http.Dir("./static"))`.

### ¿Qué exige la Parte 5 (`consigna/tp5.md`)?
1. **Abandono de la API JSON:** Prescindir de los controladores que devuelven JSON. La vista se integra y renderiza directamente en el servidor Go.
2. **Server-Side Rendering (SSR) con Templ:** Modularizar la interfaz del TP4 en componentes `.templ` dentro del directorio `views/` (`layout.templ`, `entity_form.templ`, `entity_list.templ`, `index.templ`).
3. **Estilos semánticos con Pico.css:** Vincular una hoja de estilos como Pico.css en el layout.
4. **Patrón Post/Redirect/Get (PRG):** El formulario de alta realiza un `POST /productos` tradicional (`application/x-www-form-urlencoded`). El servidor valida, persiste mediante `sqlc` y redirige al navegador a `GET /` con código `HTTP 303 See Other`.
5. **Automatización:** Actualizar el `Makefile` y Docker para compilar con `templ generate` y `sqlc generate`.
6. **Arquitectura en Tres Capas (Three-Tier):** Mantener la separación entre Presentación (handlers y Templ), Lógica de Negocio (`logic/`) y Datos (`db/sqlc/`).

---

## 📖 2. Referencias a las Filminas de la Cátedra

1. **`09-view-static.pdf`**:
   - *Filmina 12:* Limitaciones de `html/template` (type-unsafe, errores en runtime, no es Go nativo).
   - *Filminas 13-18:* **Templ**: tipado estático en tiempo de compilación, IDE-friendly, seguridad XSS automática, renderizado en streaming con `.Render(ctx, w)`.
   - *Filmina 19:* **Servicio de archivos estáticos**:
     ```go
     fs := http.FileServer(http.Dir("static"))
     http.Handle("/static/", http.StripPrefix("/static/", fs))
     ```
   - *Filmina 20:* Estructura estándar del proyecto con `views/`, `handlers/`, `db/`.
   - *Filminas 28-33:* Inyección de dependencias entre handlers, `db.Queries` y vistas Templ.
2. **`10-view-dynamic.pdf`**:
   - *Filminas 3-5:* Navegación SSR completa vs actualizaciones parciales. El estado se transmite en la URL / HTML y no en sesiones de servidor.
3. **`07-logic.pdf` y `12-slides.pdf`**:
   - *Modelo Three-Tier:* La capa de presentación no debe mezclar reglas de negocio ni lógica de base de datos. Los handlers capturan la entrada y delegan la validación en `logic.ValidateProduct`.
4. **`06-devTools.pdf`**:
   - *Filminas 13-16:* Orquestación de desarrollo con `make`. Regla `generate: @sqlc generate @templ generate`.

---

## 🛠️ 3. Detalle de Archivos Modificados y Creados

### A. Entorno y Configuración de Dependencias

#### 1. `.devcontainer/Dockerfile` (Modificado)
- **Ruta:** `tpe/.devcontainer/Dockerfile`
- **Cambio:** Se agregó la descarga e instalación del binario oficial de `templ` v0.3.833 para que esté disponible dentro del contenedor de la aplicación:
  ```dockerfile
  # templ para compilar componentes SSR (.templ -> .go)
  RUN curl -fsSL https://github.com/a-h/templ/releases/download/v0.3.833/templ_Linux_x86_64.tar.gz \
      | tar -xz -C /usr/local/bin templ
  ```
- **Propósito:** Permitir que `templ generate` corra tanto en el contenedor como en tareas automatizadas de Make.

#### 2. `go.mod` y `go.sum` (Modificados)
- **Ruta:** `tpe/go.mod`
- **Cambio:** Se ejecutó `go get github.com/a-h/templ@v0.3.833` agregando el paquete de runtime de Templ necesario para compilar y renderizar componentes.

---

### B. Capa de Presentación: Componentes Templ (`views/`)

Se creó el directorio `tpe/views/` con 4 componentes `.templ` interconectados:

#### 3. `views/layout.templ` (Nuevo)
- **Ruta:** `tpe/views/layout.templ`
- **Firma del componente:** `templ Layout(title string)`
- **Contenido clave:**
  - Estructura semántica base: `<!DOCTYPE html>`, `<html lang="es">`, `<head>`, `<body>`.
  - Enlace local a `/static/pico.min.css` (para disponibilidad offline) y enlace a `/static/styles.css`.
  - Encabezado general de *MarkeTandil*.
  - Bloque `{ children... }` dentro de `<main class="container">` para inyección de componentes hijos.

#### 4. `views/entity_form.templ` (Nuevo)
- **Ruta:** `tpe/views/entity_form.templ`
- **Firma del componente:** `templ EntityForm(categorias []db.Categoria)`
- **Contenido clave:**
  - Renderiza `<form method="POST" action="/productos" id="form-producto">`.
  - Campo `nombre` (input text, requerido).
  - Campo `id_categoria`: selector dinámico `<select id="id_categoria" name="id_categoria">` que itera sobre `categorias []db.Categoria` y genera `<option value="{ id }">{ nombre }</option>`.
  - Campo `descripcion` (input text).
  - Campo `precio` (input number, `step="0.01"`, `min="0.01"`).
  - Campo `stock` (input number, `step="1"`, `min="0"`).
  - Botón de envío `<button type="submit" class="btn-primario">Crear Producto</button>`.

#### 5. `views/entity_list.templ` (Nuevo)
- **Ruta:** `tpe/views/entity_list.templ`
- **Firma del componente:** `templ EntityList(productos []db.Producto, categoriasMap map[int32]string)`
- **Contenido clave:**
  - Encabezado con contador dinámico: `fmt.Sprintf("%d producto(s) en total", len(productos))`.
  - Condicional: si no hay productos, renderiza `<p class="vacio">No hay productos registrados actualmente.</p>`.
  - Si hay productos, renderiza una tabla con clases Pico.css `<table class="striped" id="tabla-productos">`:
    - Columnas: ID, Nombre, Categoría (mapeada a su nombre legible usando `categoriasMap`), Descripción (con manejo de `sql.NullString`), Precio formateado y Stock.
    - Columna de Acciones: formulario tradicional `<form method="POST" action="/productos/{id}/eliminar">` con botón de eliminación y confirmación nativa `onclick="return confirm(...)";`.

#### 6. `views/index.templ` (Nuevo)
- **Ruta:** `tpe/views/index.templ`
- **Firma del componente:** `templ IndexPage(productos []db.Producto, categorias []db.Categoria, categoriasMap map[int32]string, errorMsg string, successMsg string)`
- **Contenido clave:**
  - Invoca a `@Layout("MarkeTandil - Gestión de Productos")`.
  - Bloques de notificación condicionales:
    - `if errorMsg != ""` ➔ `<div class="mensaje error">...</div>`.
    - `if successMsg != ""` ➔ `<div class="mensaje exito">...</div>`.
  - Contenedor en cuadrícula `<div class="contenedor-principal">` que compone `@EntityForm(categorias)` a la izquierda y `@EntityList(productos, categoriasMap)` a la derecha.

#### 7. Código Go Generado (`*_templ.go`)
- Al ejecutar `templ generate`, se producen automáticamente:
  - `views/layout_templ.go`
  - `views/entity_form_templ.go`
  - `views/entity_list_templ.go`
  - `views/index_templ.go`

---

### C. Refactorización del Servidor HTTP y Handlers

#### 8. `main.go` (Refactorizado)
- **Ruta:** `tpe/main.go`
- **Estructura principal:**
  ```go
  type Server struct {
      queries *db.Queries
  }
  ```
- **Rutas configuradas en `mux`:**
  1. `mux.Handle("/static/", http.StripPrefix("/static/", fileServer))`: Servidor de archivos estáticos (filmina 19 de `09-view-static`).
  2. `mux.HandleFunc("/", server.handleIndex)`: Handler principal de Server-Side Rendering.
  3. `mux.HandleFunc("/productos", server.handleProductos)`: Manejo de alta mediante formulario (POST).
  4. `mux.HandleFunc("/productos/", server.handleProductoAcciones)`: Manejo de acciones individuales (POST para eliminar).

- **Lógica de los Handlers:**
  - **`handleIndex` (`GET /`):**
    - Obtiene datos de la BD: `s.queries.ListarProductos(ctx)` y `s.queries.ListarCategorias(ctx)`.
    - Arma `catMap[id] = nombre`.
    - Captura parámetros de feedback: `r.URL.Query().Get("error")` y `r.URL.Query().Get("success")`.
    - Renderiza la vista: `views.IndexPage(...).Render(ctx, w)`.
  - **`crearProductoForm` (`POST /productos` - Patrón PRG):**
    - Parsea formulario: `r.ParseForm()`.
    - Extrae `nombre`, `descripcion`, `precio`, `stock`, `id_categoria`.
    - Valida números con `strconv.ParseFloat` y `strconv.Atoi`.
    - Delega reglas de negocio en la capa pura: `logic.ValidateProduct(prodLogic)` (capa `logic/`).
    - Valida existencia de categoría con `s.queries.GetCategoria(ctx, idCategoria)`.
    - Si falla: `http.Redirect(w, r, "/?error="+url.QueryEscape(err), http.StatusSeeOther)`.
    - Si es válido: persiste con `s.queries.CrearProducto(...)`.
    - Redirige al inicio con confirmación: `http.Redirect(w, r, "/?success="+url.QueryEscape(msg), http.StatusSeeOther)` (**HTTP 303 See Other**).
  - **`handleProductoAcciones` (`POST /productos/{id}/eliminar` - Patrón PRG):**
    - Extrae el ID numérico del path o del campo oculto.
    - Borra con `s.queries.BorrarProducto(ctx, id)`.
    - Redirige a `GET /` con `HTTP 303 See Other` y mensaje de confirmación.
  - **`asegurarCategoriasIniciales`:**
    - Verifica si la tabla `categoria` está vacía. Si es así, crea automáticamente categorías iniciales ("Almacén", "Bebidas", "Lácteos", "Limpieza") para garantizar la clave foránea en bases limpias.

---

### D. Automatización, Pruebas y Estilos

#### 9. `Makefile` (Actualizado)
- **Ruta:** `tpe/Makefile`
- **Nuevos targets y modificaciones:**
  - `run`: levanta contenedores con `docker compose`, ejecuta `sqlc generate`, luego `templ generate` e inicia `go run main.go`.
  - `generate`: ejecuta `@$(COMPOSE) exec app sqlc generate` y `@$(COMPOSE) exec app templ generate` (filmina 14 de `06-devTools`).
  - `ssr-test`: ejecuta pruebas Hurl sobre `ssr_requests.hurl`.
  - `test`: levanta base limpia, corre `sqlc generate`, `templ generate` y `go test ./... -v`.

#### 10. `ssr_requests.hurl` (Nuevo)
- **Ruta:** `tpe/ssr_requests.hurl`
- **Casos de prueba automatizados con Hurl:**
  1. `GET http://localhost:8080/`: verifica `HTTP 200`, `Content-Type: text/html`, título "MarkeTandil" y presencia de `#form-producto`.
  2. `POST http://localhost:8080/productos` con `[FormParams]`: verifica respuesta `HTTP 303` con encabezado `Location: /`.
  3. `GET http://localhost:8080/`: comprueba que el nuevo producto aparece renderizado en la tabla `#tabla-productos`.
  4. `POST /productos` con precio negativo (`-10.00`): comprueba redirección `HTTP 303` hacia `Location` conteniendo `error=`.
  5. `POST /productos` con nombre vacío: comprueba redirección `HTTP 303` hacia `Location` conteniendo `error=`.

#### 11. `static/pico.min.css` (Nuevo)
- **Ruta:** `tpe/static/pico.min.css`
- **Propósito:** Hoja de estilos de Pico.css v2 guardada localmente para garantizar que la interfaz funcione de manera idéntica aun si el evaluador o contenedor no posee salida a internet.

#### 12. `DEFENSA_TP5.md` y `README.md` (Nuevos / Actualizados)
- **`README.md`:** Documentación general del repositorio actualizada para el TP5.
- **`DEFENSA_TP5.md`:** Guía teórica exhaustiva con las respuestas fundamentadas para la defensa oral ante los profesores (Zunino y Teyseyre), cubriendo:
  - Ventajas y comparativa de Templ vs `html/template` estándar de Go.
  - Funcionamiento y necesidad del patrón Post/Redirect/Get (PRG) con código HTTP 303.
  - Justificación de `http.StripPrefix` en `http.FileServer`.
  - Interacción entre la capa de presentación y la capa de lógica pura (`logic/`).

---

## 💡 4. Prompt Recomendado para usar con la IA en VS Code

Si deseas que la IA de VS Code te redacte una guía paso a paso, puedes copiar y pegar el siguiente prompt:

```text
Actúa como profesor adjunto de Programación Web de UNICEN.
Lee atentamente el archivo DOCUMENTACION_TP5_PARA_IA.md del proyecto.
En base a este documento y a los archivos del repositorio (main.go, views/*.templ, logic/*.go, Makefile, etc.), genera una guía de estudio paso a paso detallada para mí, que explique:
1. Cómo fue la transición de TP4 a TP5.
2. Cada paso que se siguió para implementar los componentes Templ en views/.
3. Cómo se refactorizó main.go para abandonar la API JSON e implementar Server-Side Rendering con el patrón PRG (Post/Redirect/Get).
4. Cómo se vincula con la capa de lógica (logic/) y la capa de datos (sqlc).
5. Incluye fragmentos de código exactos con sus nombres de archivo y referencias a las filminas de la cátedra (09-view-static, 10-view-dynamic, etc.).
```
