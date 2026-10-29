# Informe de Defensa - Trabajo Práctico Especial 4 (TPE 4)
## Programación Web - ISISTAN / UNICEN
**Cátedra:** Dr. Alejandro Zunino & Dr. Alfredo Teyseyre  
**Proyecto:** MarkeTandil - Capa de Presentación (Frontend con Vanilla JS, DOM, Fetch y CSS)

---

## 1. Introducción y Contexto del Trabajo

El objetivo del Trabajo Práctico Especial 4 es construir la **Capa de Presentación** para el sistema *MarkeTandil*, consumiendo la API REST desarrollada en el TP3. Siguiendo las directivas estrictas de la cátedra:
- Se prescinde por completo de frameworks y librerías externas en el frontend (nada de React, Angular, Vue, jQuery o Tailwind). Se utiliza exclusivamente **HTML5 semántico**, **CSS3 nativo** (con Flexbox / Grid) y **Vanilla JavaScript**.
- La comunicación se basa en el modelo **AJAX tradicional con JSON**: el cliente solicita datos asincrónicamente mediante la API nativa `fetch`, y manipula el **DOM (Document Object Model)** de forma dinámica sin necesidad de recargar la página completa.
- La entrega de archivos estáticos se integra en el servidor Go existente mediante `http.FileServer`, garantizando que tanto la interfaz como la API compartan el **mismo origen (Same-Origin)**, evitando problemas de CORS y facilitando el despliegue.

---

## 2. Matriz de Requerimientos y Justificación Técnica para la Defensa

### Requerimiento 1: Servir la Interfaz Web desde el Servidor Go
- **Consigna puntual:** Disponer de un archivo `index.html` accesible para el usuario (pudiendo servirse mediante un servidor de archivos estáticos o abriéndose en el navegador).
- **¿Qué se hizo?:** Se integró en `main.go` un handler de archivos estáticos utilizando la biblioteca estándar:
  ```go
  fileServer := http.FileServer(http.Dir("./static"))
  mux.Handle("/", fileServer)
  ```
  Los archivos de la interfaz (`index.html`, `styles.css`, `app.js`) residen en la carpeta `static/`.
- **Justificación teórica (¿Por qué se hizo así?):**
  - *SOP (Same-Origin Policy):* Al servir el frontend en el mismo puerto (`:8080`) y host (`localhost`) que la API REST, el navegador considera que ambas partes pertenecen al mismo origen (protocolo, dominio y puerto coincidentes). Esto elimina de raíz la necesidad de configurar cabeceras CORS (`Access-Control-Allow-Origin`), tal como se explica en la clase *08-view* y *10-view-dynamic*.
  - *Alineación con la cátedra:* En la filmina 34 de *10-view-dynamic*, la cátedra establece que el servidor debe exponer dos tipos de respuestas: `/` entrega el HTML inicial y las rutas de recursos (`/productos`, `/categorias`, `/clientes`) entregan los datos en formato JSON.

---

### Requerimiento 2: Estructura HTML5 Semántica y Accesible
- **Consigna puntual:** Crear `index.html` con un título principal, un formulario para crear nuevas entidades (con inputs y botón de envío) y un contenedor (`<div>` o `<ul>`) con un identificador único para renderizar dinámicamente el listado de entidades.
- **¿Qué se hizo?:**
  - Se estructuró el documento utilizando etiquetas semánticas de HTML5 (`<header>`, `<main>`, `<section>`, `<form>`, `<table>` o `<ul>`).
  - Se incluyó un `<form id="form-producto">` con campos etiquetados con `<label>` para nombre, descripción, precio, stock y categoría asociada, finalizando con un `<button type="submit">`.
  - Se definió un contenedor con ID unívoco (`<div id="lista-productos">` o `<tbody>` con ID) para la inyección programática de elementos.
- **Justificación teórica (¿Por qué se hizo así?):**
  - El uso de elementos semánticos favorece la accesibilidad (árbol de accesibilidad del navegador) y permite que los selectores de JavaScript (`document.getElementById`, `document.querySelector`) operen de forma limpia y predecible.
  - La presencia del atributo `id` en el contenedor es el punto de anclaje indispensable para que el script pueda localizar el nodo raíz donde se montará la vista dinámica.

---

### Requerimiento 3: Carga Inicial de Datos mediante `fetch` (GET) y Manipulación del DOM
- **Consigna puntual:** Crear `app.js`, enlazarlo al HTML, y programar una función que al iniciar la página realice una petición `fetch` al endpoint `GET` de la entidad, reciba el JSON, genere dinámicamente los elementos en el DOM con sus datos relevantes y un botón de eliminación.
- **¿Qué se hizo?:**
  - Se configuró el disparo inicial al cargarse el documento mediante el evento `DOMContentLoaded`:
    ```javascript
    document.addEventListener("DOMContentLoaded", () => {
        cargarProductos();
    });
    ```
  - La función `cargarProductos` utiliza `fetch('/productos')`, valida la respuesta con `response.ok`, transforma el cuerpo con `response.json()`, y genera las filas/tarjetas iterando la colección para insertarlas en el contenedor mediante manipulación de DOM nativo (`innerHTML` controlado o `document.createElement`).
  - Cada elemento renderizado incluye un botón `<button class="btn-eliminar" data-id="${producto.id}">Eliminar</button>`.
- **Justificación teórica (¿Por qué se hizo así?):**
  - *Evitar recargas innecesarias:* En lugar de renderizar la página completa en el servidor (SSR tradicional) cada vez que el usuario ingresa, el cliente descarga el esqueleto estático una sola vez y solicita únicamente los datos que necesita en formato liviano (JSON). Esto reduce el ancho de banda y la latencia percibida por el usuario.
  - *Atributos de datos (`data-*`):* Se utiliza el atributo estándar HTML5 `data-id` para asociar la clave primaria de la base de datos directamente al elemento del DOM sin contaminar clases ni IDs globales, permitiendo su lectura limpia mediante `e.target.dataset.id`.

---

### Requerimiento 4: Alta de Entidades vía `fetch` (POST) y Prevención de Comportamiento por Defecto
- **Consigna puntual:** Escuchar el evento de envío del formulario (`submit`), evitar el comportamiento por defecto (`e.preventDefault()`), leer los valores de los inputs, enviar la petición `POST` con `fetch` a la API y refrescar la vista tras la confirmación exitosa.
- **¿Qué se hizo?:**
  - Se agregó el listener sobre el evento `"submit"` del formulario:
    ```javascript
    form.addEventListener("submit", (e) => {
        e.preventDefault();
        // recolección de valores y conversión de tipos
        // fetch POST con headers Content-Type: application/json
        // tras status 201 Created: reset del form y recarga de la lista
    });
    ```
  - Se configuró la cabecera obligatoria `'Content-Type': 'application/json'` y se serializó el objeto JavaScript con `JSON.stringify()`.
- **Justificación teórica (¿Por qué se hizo así?):**
  - *¿Por qué `e.preventDefault()`?:* Los formularios HTML por defecto realizan una navegación síncrona enviando los datos codificados como `application/x-www-form-urlencoded` o `multipart/form-data`, lo que provocaría un parpadeo y la recarga total de la página perdiendo el estado del cliente. Al cancelar el evento, JavaScript toma el control total del flujo.
  - *Reutilización de `cargarProductos()`:* En vez de intentar insertar manualmente el nuevo elemento en el DOM recalculando estados parciales, volver a consultar el listado garantiza sincronización estricta con el estado real de la base de datos (incluyendo IDs autonuméricos generados y consistencia de datos).

---

### Requerimiento 5: Borrado con Delegación de Eventos (DELETE)
- **Consigna puntual:** Implementar el borrado de elementos mediante `fetch` (DELETE). Dado que los botones se generan dinámicamente en tiempo de ejecución, el requerimiento exige utilizar **delegación de eventos**. Al eliminar con éxito, refrescar la lista.
- **¿Qué se hizo?:**
  - Se asignó un único `addEventListener("click", ...)` sobre el contenedor padre existente (`#lista-productos`), en lugar de atar un listener a cada botón individual:
    ```javascript
    contenedor.addEventListener("click", async (e) => {
        const botonEliminar = e.target.closest(".btn-eliminar");
        if (!botonEliminar) return;

        const id = botonEliminar.dataset.id;
        // confirmación y petición fetch DELETE /productos/{id}
    });
    ```
- **Justificación teórica (¿Por qué se hizo así? - Pregunta clásica de examen):**
  - *¿Por qué no poner el listener directamente en el botón?:* Cuando la página se carga inicialmente, los botones de eliminar **todavía no existen en el DOM** (se crean recién cuando el servidor responde la petición JSON). Si intentáramos buscarlos al inicio, la lista estaría vacía.
  - *Eficiencia de memoria y ciclo de vida:* Si tuviéramos cientos de elementos, crear cientos de funciones listener consumiría memoria y generaría fugas de memoria (*memory leaks*) cada vez que los nodos se reemplacen.
  - *Mecanismo de propagación (Event Bubbling):* Los eventos de clic se originan en el botón y suben por el árbol del DOM hasta llegar al contenedor padre. Gracias a `e.target` o `e.target.closest()`, el padre puede identificar exactamente qué botón fue presionado, leer su `data-id` y ejecutar la llamada a la API.

---

### Requerimiento 6: Diseño Responsivo y Estilos con CSS (Flexbox / Grid)
- **Consigna puntual:** Crear `styles.css`, vincularlo al HTML y organizar la interfaz utilizando **Flexbox** o **CSS Grid**, ofreciendo una interfaz clara, agradable y fácil de usar.
- **¿Qué se hizo?:**
  - Se utilizó **CSS Grid** en el contenedor principal (`<main>`) para distribuir en dos columnas o áreas bien delimitadas: la sección del formulario de alta y la sección de listado/gestión.
  - Se utilizó **Flexbox** en los componentes internos (alineación de campos del formulario, grupos de botones, tarjetas o encabezados de tabla) para garantizar alineación fluida y espaciado consistente (`gap`, `align-items: center`).
  - Se agregaron estilos para estados interactivos (`:hover`, `:focus`) y retroalimentación visual clara ante acciones y mensajes de error.
- **Justificación teórica (¿Por qué se hizo así?):**
  - Siguiendo la clase *08-view*, CSS Grid es el modelo bidimensional óptimo para la estructura macro de la página (filas y columnas generales), mientras que Flexbox es el modelo unidimensional adecuado para la distribución de componentes internos en fila o columna.

---

## 3. Posibles Preguntas de los Profesores en la Defensa

1. **¿Por qué decidieron servir el HTML desde Go en vez de abrir el archivo `file:///` con doble clic?**  
   *Respuesta:* Al abrirlo con `file:///`, el navegador aplica la política de mismo origen (SOP) de forma restrictiva y bloquea las llamadas `fetch` a `http://localhost:8080` por ser orígenes distintos (CORS). Al servirlo directamente con `http.FileServer` en Go, la aplicación web y la API comparten protocolo, dominio y puerto (`http://localhost:8080`), permitiendo llamadas seguras sin depender de proxies ni de cabeceras permisivas.

2. **¿Qué ventaja tiene actualizar el DOM con JSON en lugar del enfoque SSR que usamos en el TP1?**  
   *Respuesta:* En el enfoque SSR tradicional, cualquier modificación obliga al servidor a regenerar el documento completo y al navegador a descartar el árbol DOM, parpadear y descargar nuevamente todos los recursos. Con el enfoque dinámico (AJAX/Fetch), el servidor solo transmite los datos en formato JSON (mucho más liviano) y el cliente actualiza quirúrgicamente únicamente la porción de la pantalla que cambió, brindando una experiencia mucho más rápida e interactiva.

3. **¿Cómo funciona la delegación de eventos y por qué es indispensable acá?**  
   *Respuesta:* La delegación de eventos aprovecha la fase de burbujeo (*bubbling*) del modelo de eventos del navegador: los eventos disparados en un nodo hijo suben hacia sus ancestros. Como los botones de eliminación se agregan dinámicamente en el DOM tras recibir la respuesta de la API, no existen al momento de cargar el script. Colocando un único escuchador en el contenedor padre estático y evaluando `e.target.closest()`, podemos atender eventos de elementos que fueron creados dinámicamente sin saturar la memoria con múltiples escuchadores.

4. **¿Por qué no usamos un framework como React o Vue?**  
   *Respuesta:* Porque la cátedra busca que comprendamos los fundamentos que hacen funcionar a cualquier framework moderno: el DOM nativo, el ciclo de eventos de JavaScript (Event Loop), la propagación de eventos, la API Fetch y el estándar CSS. Una vez dominados estos conceptos nativos del navegador, cualquier herramienta de nivel superior resulta intuitiva y predecible.

