# Informe de Defensa - Trabajo Práctico Especial 5 (TPE 5)
## Programación Web - ISISTAN / UNICEN
**Cátedra:** Dr. Alejandro Zunino & Dr. Alfredo Teyseyre  
**Proyecto:** MarkeTandil - Renderizado del Lado del Servidor (SSR) con Templ, PRG y Pico.css  

---

## 1. Introducción y Contexto del Trabajo

El objetivo del Trabajo Práctico Especial 5 es transformar la arquitectura del sistema *MarkeTandil*, pasando del enfoque de cliente rico con API REST (TP4) a una aplicación **renderizada íntegramente en el servidor (Server-Side Rendering - SSR)** utilizando **Templ**.

Siguiendo las directivas de la cátedra:
- Se abandonan los controladores de API que retornaban JSON para servir documentos HTML completos generados dinámicamente en Go.
- Se implementa el motor de plantillas **Templ**, que a diferencia de `html/template`, ofrece **seguridad de tipos en tiempo de compilación**, generación de código Go nativo y composición de componentes reutilizables.
- Se adopta el patrón clásico **Post/Redirect/Get (PRG)** para el procesamiento de formularios, garantizando navegación fluida y evitando reenvíos accidentales ante recargas del navegador.
- Se utiliza la biblioteca de estilos semánticos **Pico.css**, complementada con los recursos locales servidos mediante `http.FileServer` y `http.StripPrefix`.
- Se respeta estrictamente la **Arquitectura en Tres Capas (Three-Tier)**: la capa de presentación (handlers y vistas Templ) delega las validaciones en la capa de lógica pura (`logic/`) y persiste a través de la capa de acceso a datos (`db/sqlc/`).

---

## 2. Matriz de Requerimientos y Justificación Técnica para la Defensa

### Requerimiento 1: Creación y Composición de Componentes Templ en `views/`
- **Consigna puntual:** Analizar el `index.html` del TP4 y dividirlo en componentes lógicos en un directorio `views`:
  - `layout.templ`: Estructura base de la página (HTML, HEAD, BODY) con enlace a Pico.css.
  - `entity_form.templ`: Componente que renderiza el formulario para dar de alta una nueva entidad.
  - `entity_list.templ`: Componente que recibe la colección de entidades (`[]db.Producto`) y renderiza la tabla.
- **¿Qué se hizo?:**
  - Se crearon los archivos `.templ` correspondientes en `tpe/views/`.
  - `layout.templ` define `templ Layout(title string)` con etiquetas semánticas de HTML5, meta-viewport responsivo, Pico.css y el bloque de contenido `{ children... }`.
  - `entity_form.templ` define `templ EntityForm(categorias []db.Categoria)` con campos de input para nombre, descripción, precio y stock, sumando un `<select name="id_categoria">` poblado dinámicamente con las categorías registradas en la base de datos.
  - `entity_list.templ` define `templ EntityList(productos []db.Producto, categoriasMap map[int32]string)` renderizando una tabla semántica (`<table class="striped">`), resolviendo el nombre legible de la categoría y ofreciendo un botón de eliminación en un formulario POST tradicional.
  - `index.templ` define `templ IndexPage(...)` que ensambla el layout, las alertas visuales de feedback (error y éxito) y organiza el formulario y la lista dentro de la grilla principal.
- **Justificación teórica (Filminas 13 a 18 de `09-view-static.pdf`):**
  - *Type Safety (Seguridad de tipos en compilación):* `html/template` procesa cadenas de texto y cualquier error de tipado o campo inexistente se detecta recién en tiempo de ejecución. Templ convierte las plantillas en código Go fuertemente tipado (`*_templ.go`); si pasamos un tipo erróneo o un struct inválido, el compilador de Go no compila el binario.
  - *Escapado contextual XSS automático:* Al igual que `html/template` pero con la seguridad de tipos de Go, Templ analiza el contexto y escapa automáticamente cadenas potencialmente maliciosas evitando ataques Cross-Site Scripting.
  - *Composición mediante `@Componente`:* Facilita la modularidad y el mantenimiento del código, permitiendo anidar componentes como funciones Go comunes.

---

### Requerimiento 2: Refactorización del Servidor Go y Retiro de Endpoints JSON
- **Consigna puntual:** Modificar `main.go`, prescindir de los handlers de la API REST que devuelven JSON y crear un handler principal para `GET /` que obtenga los registros mediante sqlc (`Listar...`), llame al componente principal (`IndexPage`) y lo renderice en el `http.ResponseWriter`.
- **¿Qué se hizo?:**
  - Se retiraron las rutas y handlers que exponían JSON (`/clientes`, `/categorias`, etc.).
  - Se configuró el servicio de archivos estáticos bajo el prefijo `/static/`:
    ```go
    fileServer := http.FileServer(http.Dir("./static"))
    mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
    ```
  - Se implementó `handleIndex` sobre `GET /`:
    1. Ejecuta `s.queries.ListarProductos(ctx)` y `s.queries.ListarCategorias(ctx)`.
    2. Construye un diccionario `catMap` para asociar IDs de categoría con sus nombres.
    3. Lee los parámetros de URL `error` y `success` para proveer feedback al usuario.
    4. Renderiza el componente directamente en el socket de respuesta:
       ```go
       component := views.IndexPage(productos, categorias, catMap, errorMsg, successMsg)
       err := component.Render(ctx, w)
       ```
- **Justificación teórica (Filminas 19, 24 y 29 de `09-view-static.pdf`):**
  - *Separación de rutas de assets y vistas:* En la filmina 19 se resalta que en SSR el servidor genera el HTML para la raíz, mientras que las hojas de estilo e imágenes se entregan desde una carpeta estática con `http.StripPrefix`.
  - *Renderizado en streaming y eficiencia:* El método `.Render(ctx, w)` de Templ escribe directamente en el `http.ResponseWriter`, reduciendo el uso de buffers intermedios en memoria y permitiendo una respuesta inmediata al cliente.

---

### Requerimiento 3: Manejo del Formulario con Patrón Post/Redirect/Get (PRG)
- **Consigna puntual:** El formulario debe hacer un POST a una ruta (`POST /productos`). El handler debe parsear los datos del formulario, guardar la entidad con sqlc (`CrearProducto`) y redirigir al usuario al inicio (`/`) con `http.Redirect`.
- **¿Qué se hizo?:**
  - Se implementó el handler `POST /productos`:
    1. Parsea el cuerpo codificado en `application/x-www-form-urlencoded` con `r.ParseForm()`.
    2. Convierte strings a valores numéricos (`strconv.ParseFloat`, `strconv.Atoi`).
    3. Valida las reglas de negocio con la capa pura `logic.ValidateProduct` y verifica la existencia de la categoría asociada en la base de datos.
    4. Si hay error: redirige a `/?error=...` usando `http.Redirect(w, r, ..., http.StatusSeeOther)`.
    5. Si es válido: persiste mediante `s.queries.CrearProducto(...)` y redirige a `/?success=...` usando `http.StatusSeeOther` (HTTP 303).
  - Se implementó el handler de borrado `POST /productos/{id}/eliminar` aplicando el mismo patrón PRG tras invocar `s.queries.BorrarProducto(...)`.
- **Justificación teórica (Filmina 3 de `10-view-dynamic.pdf` y estándar HTTP):**
  - *¿Por qué no responder HTML directamente en el POST?:* Si el servidor devolviera el HTML resultante en la respuesta al `POST`, el navegador mantendría la petición `POST` en el historial de navegación. Al presionar F5 (recargar), el navegador advertiría al usuario sobre el reenvío de formulario y podría insertar registros duplicados.
  - *¿Por qué HTTP 303 See Other?:* El código 303 indica explícitamente al navegador que la operación terminó con éxito y que debe realizar una petición `GET` a la URL indicada en la cabecera `Location`. De esta manera, el navegador queda en un estado limpio e idempotente en `GET /`.

---

### Requerimiento 4: Integración con la Capa de Lógica de Negocio (Three-Tier)
- **Consigna puntual:** Integrar la solución de la parte 5 respetando los conceptos y la arquitectura del trabajo especial.
- **¿Qué se hizo?:**
  - Los handlers de presentación no contienen lógica de validación dura ni queries SQL en línea.
  - Se conservó y reutilizó el paquete `logic/`: `logic.ValidateProduct` valida que el nombre no sea vacío, que la categoría sea mayor a cero, que el stock sea mayor o igual a 0 y que el precio sea estrictamente positivo.
- **Justificación teórica (Filminas 5 de `09-view-static.pdf`, `07-logic.pdf` y `12-slides.pdf`):**
  - *Modelo Three-Tier:* La capa de presentación interactúa con el usuario y captura entradas; la capa de lógica de negocio valida y ejecuta reglas independientes del medio de transporte; la capa de datos almacena y recupera información.
  - Al aislar las reglas en `logic/`, las pruebas de negocio pueden ejecutarse sin levantar servidores HTTP ni bases de datos.

---

### Requerimiento 5: Automatización con Makefile y Docker
- **Consigna puntual:** Escribir un Makefile o script para automatizar la generación de código y la ejecución del servidor, incluyendo el manejo del motor de base de datos con Docker.
- **¿Qué se hizo?:**
  - Se actualizó [.devcontainer/Dockerfile](.devcontainer/Dockerfile) instalando el compilador `templ` binario.
  - Se actualizó [Makefile](Makefile):
    - `make run`: levanta la BD y el contenedor app, ejecuta `sqlc generate`, compila vistas con `templ generate` e inicia el servidor.
    - `make generate`: corre ambas herramientas de generación de código (`sqlc` y `templ`).
    - `make ssr-test`: corre las pruebas de integración SSR con Hurl sobre `ssr_requests.hurl`.
    - `make test`: ejecuta las pruebas unitarias y de base de datos con entorno aislado.
    - `make down`: detiene los contenedores y destruye volúmenes.
- **Justificación teórica (Filminas 13 a 16 de `06-devTools.pdf`):**
  - La filmina 14 establece explícitamente la tarea `generate` en el Makefile combinando `@sqlc generate` y `@templ generate`, asegurando que las dependencias generadas estén al día antes de la compilación de Go.

---

## 3. Preguntas Típicas de Examen en la Defensa Oral

### 1. ¿Cuáles son las diferencias y ventajas de Templ respecto a `html/template` estándar de Go?
*Respuesta:*
`html/template` utiliza plantillas basadas en cadenas de texto que se interpretan en tiempo de ejecución. Esto implica que errores tipográficos en nombres de campos, tipos incompatibles o fallas de sintaxis solo se descubren cuando una petición web activa esa plantilla (en runtime).  
**Templ**, en cambio, es un lenguaje que se compila a código Go nativo (`*_templ.go`). Posee **tipado estático en tiempo de compilación**, soporte completo de autocompletado y refactorización en el IDE, y el compilador de Go rechaza la compilación si los tipos o parámetros no coinciden. Además, Templ hereda la seguridad de escapado automático de HTML contra XSS y permite componer componentes como funciones Go nativas.

### 2. ¿Qué es el patrón Post/Redirect/Get (PRG) y qué problema resuelve?
*Respuesta:*
En el modelo web tradicional sin JavaScript, los formularios se envían mediante una petición HTTP `POST`. Si el servidor responde con el documento HTML final directamente sobre esa respuesta `POST`, la URL en el navegador permanece asociada al verbo POST. Si el usuario actualiza la página con F5 o retrocede en el historial, el navegador muestra el cartel de advertencia *"¿Desea reenviar los datos del formulario?"*, pudiendo generar una segunda inserción duplicada en la base de datos.  
El patrón **PRG** resuelve esto:
1. El cliente envía `POST /productos` con los datos del formulario.
2. El servidor procesa la creación en la base de datos.
3. El servidor responde con una **redirección HTTP 303 (See Other)** con la cabecera `Location: /`.
4. El navegador del cliente emite automáticamente una petición `GET /`, renderizando la lista actualizada. Ahora la última acción en el historial es un `GET` idempotente, y actualizar la página es completamente seguro.

### 3. ¿Por qué usamos `http.StripPrefix` al servir los archivos estáticos?
*Respuesta:*
Cuando registramos un handler de archivos con `http.FileServer(http.Dir("./static"))` en la ruta `/static/`, el navegador solicita, por ejemplo, `/static/styles.css`. Si no usáramos `http.StripPrefix("/static/", ...)`, el `FileServer` intentaría buscar el archivo en la ruta `./static/static/styles.css`, lo que provocaría un error 404 Not Found. `http.StripPrefix` remueve el prefijo de la URL antes de pasárselo al sistema de archivos local, buscando correctamente `./static/styles.css`.

### 4. ¿Cómo convive la capa de lógica del negocio (`logic`) con las vistas Templ?
*Respuesta:*
Las vistas Templ pertenecen exclusivamente a la **capa de presentación**: su única función es renderizar datos en HTML y estructurar el formulario de captura. Cuando el usuario envía el formulario, el handler HTTP (`crearProductoForm`) extrae los valores y construye el modelo de negocio `logic.Producto`, invocando `logic.ValidateProduct`. Si la validación falla, la presentación interrumpe el flujo y redirige con el mensaje de error correspondiente sin tocar la base de datos. Si es válida, delega la persistencia en `queries.CrearProducto` de sqlc. De este modo, la capa de presentación nunca asume reglas de negocio y la capa de lógica permanece desacoplada de la interfaz gráfica y de la base de datos.
