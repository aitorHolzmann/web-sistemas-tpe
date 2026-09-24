# Pr

# ogramación Web

# El nivel de Presentación

Dr. Alejandro Zunino & Dr. Alfredo Teyseyre alejandro.zunino@isistan.unicen.edu.ar, alfredo.teyseyre@isistan.unicen.edu.ar

  \| 1 of 82

<image redacted: 72x72px, 72x72pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

# Contenidos

1\. Funciones de la capa de Presentación

2\. El Navegador de Web

3\. Decisiones de Diseño

4\. Tipos de Tecnologías Web

  \| 2 of 82

<image redacted: 718x479px, 718x479pt, ~72dpi, JPG, DEVICE_RGB, 32bpp>

# Modelo Three-Tier

Presentación : interactúa con el usuario para mostrar la información y capturar sus entradas.

Lógica de Negocio (o capa de aplicación o capa intermedia): procesa datos, implementa reglas de negocio, realiza cálculos y coordina operaciones.

Datos : responsable de almacenar, recuperar y gestionar los datos de la aplicación. Incluye las base de datos, sistemas de archivos y otros almacenamientos persistentes.

  Im portant Las flechas son invocaciones: llamadas a funciones locales, peticiones HTTP, mensajes asíncronos, u otros tipos de comunicación.

Navegador Web

Backend

Capa de Presentación

Capa de Lógica de Negocio

Sistemas externos

Capa de Datos BusinessServer

Bases de Datos Servidor de archivos

  \| 3 of 82

# Funciones de la capa de Presentación

Es la interfaz con la que el usuario final interactúa directamente. Su principal responsabilidad es presentar los datos al usuario y recibir sus entradas.

Es lo que comúnmente llamamos el frontend de la aplicación.

Divide sus responsabilidades en dos áreas principales:

Lógica de Presentación (UI): Cómo se muestran los datos. Incluye elementos como plantillas, componentes y widgets.

Lógica de Interfaz de Usuario (Client-Side Logic): Cómo responde la aplicación a las entradas del usuario. Esto abarca la validación de formularios, las animaciones y la gestión del estado de la interfaz.

  \| 4 of 82

# Requerimientos

Renderizado de la Interfaz: Genera el HTML, CSS y JavaScript que el navegador interpreta para mostrar la página web.

Manejo de la Interacción del Usuario: Captura eventos como clics, entradas de teclado y gestos táctiles para proporcionar una experiencia interactiva.

Comunicación con el Backend: Realiza solicitudes a la capa de negocio para obtener datos o enviar nueva información.

Gestión del Estado del Cliente: Mantiene un registro del estado de la interfaz, como qué modales están abiertos, qué pestañas están activas o el contenido de un carrito de compras.

Validación de Entradas: Proporciona retroalimentación inmediata al usuario validando los datos del lado del cliente antes de enviarlos al servidor.

  \| 5 of 82

# Requerimientos

Accesible para todos los usuarios Intuitiva y fácil de usar Tiempos de carga rápidos Usabilidad

Funcionar en diferentes Rendimiento navegadores

Capa de Compatibilidad Presentación Diseño responsivo para diversos dispositivos

Seguridad

Mantenibilidad Código limpio y modular

Manejo seguro de tokens Uso de buenas prácticas de diseño

Estabilidad visual

Responsividad a las interacciones

Protección contra Cross

Site Scripting

  \| 6 of 82

# Requerimientos

Accesible para todos los usuarios Intuitiva y fácil de usar Tiempos de carga rápidos Usabilidad

Funcionar en diferentes Rendimiento navegadores Independientemente de Capa de Compatibilidad Presentación la tecnología utilizada, la Diseño responsivo para diversos dispositivos n a v e capa de presentación usualmente involucra un g Seguridad in terprador web que Mantenibilidad Código limpio y modular eta HTML, CSS y JavaScript. Manejo seguro de tokens Uso de buenas prácticas de diseño

Estabilidad visual

Responsividad a las interacciones

Protección contra Cross

Site Scripting

  \| 6 of 82

El Entorno de Ejecución de Aplicaciones Web

1\. ¿Qué es un Navegador Web?

2\. ¿Qué es un Navegador Web?

3\. Motor de renderizado

4\. Motor JavaScript

5\. APIs Web

6\. Same-Origin Policy y CORS

7\. Content Security Policy (CSP)

# El Navegador

8\. Modelo de Objetos del Documento (DOM)

# de Web

9\. Manipulación del DOM con JS

10\. Ciclo de Eventos (Event Loop)

11\. ¿Qué es HTML?

12\. Mejorando el estilo con CSS

13\. Sólo una introducción al Navegador de Web

  \| 7 of 82

¿Qué es un Navegador Web?

Desde la perspectiva del usuario, es la herramienta para acceder a internet o a Aplicaciones Web.

Desde la perspectiva del desarrollador de aplicaciones web, el navegador es mucho más:

Es el entorno de ejecución (runtime) donde tus aplicaciones web cobran vida.

  Im portant

Piénsenlo como una máquina virtual especializada diseñada para interpretar y presentar contenido web.

  \| 8 of 82

¿Qué es un Navegador Web?

Usuario

HTML + CSS + JS + SVG Internet

HTTP / Fetch / WS

Navegador Web — VM (sobre OS/Hardware) Motores

Motor de Renderizado

DOM + CSSOM → Layout → Paint → Com positor

Motor JS Pixels en Pantalla Call Stack + JIT

APIs Web Event Loop Microtasks / Macrotasks DOM • Fetch • Storage ...

Geoloc • Notif • Storage Sandbox / Seguridad UI LocalStorage • Clipboard • Sensores Same-Origin • CORS • CSP IndexedDB...

  \| 9 of 82

# Componentes del Navegador de Web

Document

Manipulación

Canvas / WebGL

Fetch / XHR

Microtasks Macrotasks

Call Stack Nodos Event Loop

DOM

DOM

Motor de Layout y Paint Renderizado Navegador Web

CSSOM APIs Web

Motor JavaScript

Storage V8 / SpiderMonkey

JIT Single Thread

  \| 10 of 82

# Motor de renderizado

Renderizado HTML:

Interpreta HTML, CSS y SVG → construye árbol DOM.

DOM: representación en memoria con objetos JS.

Renderizado CSS:

Aplica estilos → construye CSSOM y calcula layout.

CSSOM: estilos en memoria

accesibles desde JS.

Renderizado SVG:

Gráficos vectoriales escalables, integrado a HTML/CSS.

HTML CSS SVG

Parser HTML Parser CSS

DOM CSSOM

Render Tree

Layout

Paint

Compositor GPU

Pantalla

  \| 11 of 82

# Motor JavaScript

Ejecuta JS del lado cliente; JIT optimiza en tiempo de ejecución.

Motores: V8 ( / / ), SpiderMonkey ( ), JavaScriptCore ( ).

Single thread: JS y render comparten el hilo principal → bloquear = congelar UI.

Desde JS se accede al DOM y Web APIs.

  Im portant

Permite modificar contenido/apariencia, responder a eventos y hacer fetch.

Código JS

Parser

AST

JIT

Compilación +

Optimización

Call Stack

Single Thread

bloqueo

DOM Web APIs UI Congelada

  \| 12 of 82

# APIs Web

Manipulación DOM : crear/modificar/eliminar

Networking : Fetch, XHR

Storage : Local/Session/IndexedDB

Gráficos : Canvas, WebGL

Multimedia : Audio/Video

Otros: Geoloc. , Notif. , WebSockets , Clipboard , Drag&Drop , Sensores ,W orkers

https://developer.mozilla.org/en-US/docs/Web/API

  \| 13 of 82

# Same-Origin Policy y CORS

Origen = esquema + host + puerto . Ej: https://exa.com:443

https://exa.com ≠ http://exa.com ≠

https://api.exa.com ≠

https://exa.com:8443

SOP (bloquea): JS de un origen no puede leer DOM ni respuestas fetch/XHR de otro origen. Protege cookies/sesión.

CORS (apertura opt-in): el servidor autoriza con headers:

Access-Control-Allow-Origin: https://exa.com Vary: Origin

Simple ( GET/POST ) vs preflight ( OPTIONS para PUT/DELETE , JSON, headers custom).

API Browser (https://api.otro.com) (https://exa.com)

GET /users Origin: https://exa.com

alt \[Origin permitido\]

200 + ACAO: https://exa.com JS lee respuesta

\[Origin no permitido\]

200 sin ACAO Navegador bloquea lectura

API Browser (https://api.otro.com) (https://exa.com)

  Im portant

El servidor responde igual; es el navegador quien bloquea a JS. curl no aplica SOP.

  \| 14 of 82

# Content Security Policy (CSP)

SOP/CORS controlan de dónde se lee; CSP controla qué se ejecuta/carga en tu página. Mitiga XSS.

Lista blanca declarada por el servidor:

Content-Security-Policy: default-src'self';

script-src'self'h ttps://cdn.jsdelivr.net; object-src'none'

Bloquea \<script\> inline, eval , \<iframe\> /objetos no listados. Ver violación en DevTools (F12 → Console).

, Endurecer: object-src'none'

,r uri'self' eport-uri /csp-report .

Página + CSP

¿Origen del recurso permitido?

self / cdn listado inline / evil.com

Bloqueado + error en Ejecuta / carga Consola

  N ote base SOP + CORS + CSP = defensa en capas: aislamiento + intercambio controlado + anti XSS.

  \| 15 of 82

# Modelo de Objetos del Documento (DOM)

Es una API y una representación jerárquica del documento HTML.

Permite que JavaScript acceda y manipule la estructura, contenido y estilos de la página.

Cada elemento HTML se convierte en

un objeto nodo en el árbol DOM.

  N ote

El navegador construye esta representación en memoria para cada página que carga.

Por ejemplo, este documento HTML…

\<!DOCTYPE html\>

\<html\>

\<head\>

\<meta charset=" UTF-8" /\>

\<title\>Mi Página\</title\> \</head\>

\<body\> \<h1\>Hola Mundo\</h1\>

\<p\>Contenido del párrafo\</p\> \</body\> \</html\>

  \| 16 of 82

# Modelo de Objetos del Documento (DOM)

… es re presentado internamente por el navegador como el siguiente árbol de objetos:

Document (Root)

HTML Element

Head Element Body Element

Meta Element Title Element H1 Element P Element

Text Node:'Contenido del Text Node:'Mi Página' Text Node:'Hola Mundo' párrafo'

  \| 17 of 82

# Manipulación del DOM con JS

JavaScript document

**Acceder**

getElementById querySelector

**Modificar contenido**

textContent

innerHTML

**Modificar estilos**

Seleccionar Navegador actualiza Página visible style Árbol DOM nodos la interfaz classList

**Modificar estructura**

createElement

appendChild

remove

**Responder a eventos**

addEventListener

  \| 18 of 82

# Accediendo elementos del DOM con JS

Dado el siguiente fragmento de HTML:

container"\> \<div class="

main-title"\>Título Principal\</h1\> \<h1 id="

intro- \<p class=" paragraph"\>Este es el primer párrafo.\</p\> \<p\>Este es el segundo párrafo.\</p\> \<ul\>

\<li\>Primer item\</li\>

\<li\>Segundo item\</li\> \</ul\>

\<button class=" btn- primary"\>Botón Principal\</button\> \</div\>

container"\> \<div class="

\<!-- ... más contenido ... --\>

\</div\>

  \| 19 of 82

# Accediendo elementos del DOM con JS

Toda página Web posee un único objeto document que representa el DOM de la página.

Podemos acceder a sus elementos de varias formas:

// Por ID (único) - el más rápido const miTitulo = document.getElementById("main-title");

// Por clase (devuelve una colección) const parrafosIntro = document.getElementsByClassName("intro-

// Por nombre de etiqueta (devuelve una colección) const todosLosDivs = document.getElementsByTagName("div");

// Con selectores CSS (muy flexible) const primerBoton = document.querySelector(".btn-primary"); const itemsLista = document.querySelectorAll("ul li");

paragraph");

  \| 20 of 82

# Modificando elementos con JS

Partiendo de este HTML:

\<h1 id=" pageTitle"\>Título Original\</h1\> container"\> \<div id="

\<p\>Contenido inicial.\</p\> \</div\>

src=" \<img id=" myImage" old.jpg"

alt=" /\> Imagen Vieja"

  \| 21 of 82

# Modificando elementos con JS

Partiendo de este HTML:

\<h1 id=" pageTitle"\>Título Original\</h1\> container"\> \<div id="

\<p\>Contenido inicial.\</p\> \</div\>

src=" \<img id=" myImage" old.jpg"

Después de ejecutar el JS, el DOM se verá así:

\<h1 id=" pageTitle"\>¡Nuevo Título Dinámico!\</h1\> container"\> \<div id="

\<h2\>Un subtítulo añadido\</h2\>

\<p\>Contenido \<strong\>en negrita\</strong\>.\</p\> \</div\>

\<img i d=" myImage" src=" new\_image.png" alt=" Nueva imagen de ejemplo" 123" data-id="

/\>

alt=" /\> Imagen Vieja"

  \| 21 of 82

# Modificando elementos con JS

Podemos modificar los elementos existentes:

// Obtener un elemento por su ID const tituloPagina = document.getElementById("pageTitle");

// Cambiar el contenido de texto

tituloPagina.textContent ="¡Nuevo Título Dinámico!";

// Cambiar el contenido HTML

const contenedor = document.getElementById("container"); contenedor.innerHTML =

\+"\<h2\>Un subtítulo añadido\</h2\>"

"\< p\>Contenido \<strong\>en negrita\</strong\>.\</p\>";

// Modificar atributos

const miImagen = document.getElementById("myImage"); miImagen.src ="new\_image.png"; miImagen.alt ="Nueva imagen de ejemplo";"123"); miImagen.setAttribute("data-id",

  \| 22 of 82

# Modificando estilos con JS

Partiendo de este HTML:

main-title"\>Título Principal\</h1\> \<h1 id="

class=" btn"\>Click me\</button\> \<button id=" toggleButton"

const miTitulo = document.getElementById("main-title");

// Acceso directo a propiedades CSS miTitulo.style.color ="red"; miTitulo.style.fontSize ="32px"; miTitulo.style.backgroundColor ="#eee";

// Usar classList para añadir/quitar clases CSS const toggleButton = document.getElementById("toggleButton"); // Añade la clase'active'

toggleButton.classList.add("active"); // Quita la clase'btn' toggleButton.classList.remove("btn"); // Alterna la clase'highlight'

toggleButton.classList.toggle("highlight");

  \| 23 of 82

# Creando y eliminando elementos con JS

// Crear un nuevo elemento

const nuevoParrafo = document.createElement("p"); nuevoParrafo.textContent ="Este párrafo fue creado con JavaScript.";

nuevoParrafo.classList.add("dynamic-content"); // Añadir el nuevo elemento al final del body document.body.appendChild(nuevoParrafo); // Añadirlo a un contenedor específico (ej. \<div id=" myContainer"\>\</div\>) const miContenedor = document.getElementById("myContainer"); const nuevaLista = document.createElement("ul"); miContenedor.appendChild(nuevaLista); const itemLista = document.createElement("li"); itemLista.textContent ="Primer ítem dinámico";

nuevaLista.appendChild(itemLista); // Eliminar un elemento (ej. el nuevoParrafo que creamos) // 1. Necesitas una referencia al padre const padreNuevoParrafo = nuevoParrafo.parentNode; if (padreNuevoParrafo) { // Asegurarse de que el padre existe padreNuevoParrafo.removeChild(nuevoParrafo); } // O, si tienes la referencia directa al elemento miContenedor.remove(); // Elimina el contenedor y todo su contenido

  \| 24 of 82

# Manejo de eventos con JS

// 1. Obtener el elemento

const miBoton = document.getElementById("myButton"); // \<button id="myButton"\>Haz clic\</button\>

// 2. Añadir un"event listener"

miBoton.addEventListener("click",f unction () { alert("¡Botón presionado!"); // Aquí puedes cambiar texto, mostrar/ocultar elementos, enviar datos, etc. const mensaje = document.getElementById("message"); m ensaje.textContent ="¡Has hecho clic en el botón!"; m ensaje.style.display ="block"; });

Otros eventos comunes:

mouseover , mouseout , keydown , submit , load , resize , scroll

document.addEventListener("DOMContentLoaded", () =\> { /\* Código se ejecuta cuando el DOM está listo\*/ });

  \| 25 of 82

# El DOM y el rendimiento

Manipulación masiva: Cada manipulación del DOM puede ser costosa, ya que el navegador tiene que recalcular el renderizado (reflow y repaint).

Batching: Si vas a hacer muchas operaciones, es mejor agruparlas o hacerlas offline:

1\. Crea un fragmento de documento que no se muestra

2\. Luego adjuntar el resultado a la página

Frameworks: Librerías como React o Vue abstraen y optimizan la manipulación del DOM por ti (usando un Virtual DOM).

  \| 26 of 82

# Ciclo de Eventos (Event Loop)

El navegador es single-threaded en su ejecución de JavaScript.

  N ote

El Event Loop es el mecanismo que permite que JavaScript maneje operaciones asíncronas (como clics, temporizadores, peticiones de red) sin bloquear el hilo principal.

  \| 27 of 82

# El Problema del Hilo Único Bloqueante

User Browser JS Thread

Clic en botón (inicia JS)

Ejecuta funcionLargaSincrona()

¡3 segundos de cálculo!

Termina

Actualiza UI (después de 3s)

La UI estuvo congelada 3s!

User Browser JS Thread

  \| 28 of 82

# La Solución: Asincronía y el Event Loop

  N ote

El Event Loop es el mecanismo que permite que JS maneje operaciones asíncronas (como clics, temporizadores, peticiones de red) sin bloquear el hilo principal

Call Stack (Pila de Llamadas)

Web APIs (APIs del Navegador)

Callback Queue (Cola de Callbacks/Tareas)

Event Loop (Bucle de Eventos)

Microtask Queue (Cola de Microtareas): un agregado crucial

  \| 29 of 82

Motor JavaScript

Event Loop (Coordina flujo)

7b.Prioridad 9.Si no hay microtasks

7a.¿CallStack vacío?

8\.Ejecuta todos 10.Siguiente callback

Call Stack

1\.Ejecuta 2.Encuentra

Funciones

Navegador

Web APIs (setTimeout, DOM, Fetch, etc.)

Resolución de 4.Completa

Callback (ej: setTimeout Promesa callback)

5\.Encola en 6.Encola en

Colas de Espera

Callback Queue Microtask Queue

3\.Delega a (Macrotasks: setTimeout, (Promesas, eventos DOM) MutationObserver)

Call Stack (sincrónico) API Asíncrona (ej: setTimeout) W eb APIs (navegador) M acrotasks (menor prioridad) Mi crotasks (máxima prioridad) Ev ent Loop

  \| 30 of 82

Navegador

Web APIs (setTimeout, DOM, Fetch, etc.)

Resolución de 4.Completa

Motor JavaScript

Callback (ej: setTimeout Promesa Event Loop (Coordina flujo) callback)

5\.Encola en 6.Encola en 7b.Prioridad 9.Si no hay microtasks

Colas de Espera

Callback Queue Microtask Queue

7a.¿CallStack vacío? (Macrotasks: setTimeout, (Promesas, eventos DOM) MutationObserver)

8\.Ejecuta todos 10.Siguiente callback

Call Stack Ejecuta funciones sincrónicamente (LIFO) 1.Ejecuta 2.Encuentra

API Asíncrona (ej: Funciones setTimeout)

3\.Delega a

Call Stack (sincrónico) W eb APIs (navegador) M acrotasks (menor prioridad) Mi crotasks (máxima prioridad) Ev ent Loop

  \| 30 of 82

Motor JavaScript

Event Loop (Coordina flujo)

7b.Prioridad 9.Si no hay microtasks

7a.¿CallStack vacío?

8\.Ejecuta todos 10.Siguiente callback

Call Stack

1\.Ejecuta 2.Encuentra

Funciones

Navegador

Web APIs (setTimeout, DOM, Fetch, etc.)

Resolución de 4.Completa

Callback (ej: setTimeout Promesa callback)

5\.Encola en 6.Encola en

Colas de Espera

Callback Queue Microtask Queue

3\.Delega a (Macrotasks: setTimeout, (Promesas, eventos DOM) MutationObserver)

Cuando encuentra

operaciones asíncronas, las delega a las Web APIs

Call Stack (sincrónico) API Asíncrona (ej: setTimeout) W eb APIs (navegador) M acrotasks (menor prioridad) Mi crotasks (máxima prioridad) Ev ent Loop

  \| 30 of 82

Navegador Procesa las operaciones y, Web APIs (setTimeout, DOM, Fetch, etc.) al completarse, encola Macrotasks o Microtasks Resolución de 4.Completa

Motor JavaScript

Callback (ej: setTimeout Promesa Event Loop (Coordina flujo) callback)

5\.Encola en 6.Encola en 7b.Prioridad 9.Si no hay microtasks

Colas de Espera

Callback Queue Microtask Queue

7a.¿CallStack vacío? (Macrotasks: setTimeout, (Promesas, eventos DOM) MutationObserver)

8\.Ejecuta todos 10.Siguiente callback

Call Stack

1\.Ejecuta 2.Encuentra

API Asíncrona (ej: Funciones setTimeout)

3\.Delega a

Call Stack (sincrónico) W eb APIs (navegador) M acrotasks (menor prioridad) Mi crotasks (máxima prioridad) Ev ent Loop

  \| 30 of 82

Navegador

Web APIs (setTimeout, DOM, Fetch, etc.)

Resolución de 4.Completa

Motor JavaScript Ejecuta todo el Call Stack, Callback (ej: setTimeout Promesa luego todos los Microtask y Event Loop (Coordina flujo) callback) un Macrotask (Callback) 5.Encola en 6.Encola en 7b.Prioridad 9.Si no hay microtasks

Colas de Espera

Callback Queue Microtask Queue

7a.¿CallStack vacío? (Macrotasks: setTimeout, (Promesas, eventos DOM) MutationObserver)

8\.Ejecuta todos 10.Siguiente callback

Call Stack

1\.Ejecuta 2.Encuentra

API Asíncrona (ej: Funciones setTimeout)

3\.Delega a

Call Stack (sincrónico) W eb APIs (navegador) M acrotasks (menor prioridad) Mi crotasks (máxima prioridad) Ev ent Loop

  \| 30 of 82

Motor JavaScript

Event Loop (Coordina flujo)

7b.Prioridad 9.Si no hay microtasks

7a.¿CallStack vacío?

8\.Ejecuta todos 10.Siguiente callback

Call Stack

1\.Ejecuta 2.Encuentra

Funciones

Navegador

Web APIs (setTimeout, DOM, Fetch, etc.)

Resolución de 4.Completa

Callback (ej: setTimeout Promesa callback)

5\.Encola en 6.Encola en

Colas de Espera

Callback Queue Microtask Queue

3\.Delega a (Macrotasks: setTimeout, (Promesas, eventos DOM) MutationObserver)

Call Stack (sincrónico) API Asíncrona (ej: setTimeout) W eb APIs (navegador) M acrotasks (menor prioridad) Mi crotasks (máxima prioridad) Ev ent Loop

  \| 30 of 82

# Call stack

Es una pila donde JavaScript coloca las funciones que se están ejecutando.

Cuando una función se

invoca, se coloca en la pila.

Cuando una función

termina, se quita de la pila.

  Im portant Solo puede haber una función ejecutándose en la pila en un momento dado. Si la pila no está vacía, el Event Loop no puede actuar.

Navegador

Web APIs

JavaScript

Promesa Callback Event Loop

Colas de Espera

Microtask Callback 7a.¿CallStack vacío?

8\.Ejecuta todos 10.Siguiente callback

Call Stack

1\.Ejecuta 2.Encuentra

Funciones API Asíncrona

  \| 31 of 82

# Web APIs

Funcionalidades del navegador.

Para operaciones asíncronas:

Temporizadores: setTimeout(), setInterval()

Peticiones de red: fetch(), XMLHttpRequest

Eventos DOM como click, scroll: addEventListener()

Geolocation, Web Storage, etc.

Cuando JS invoca una Web API, la operación se"descarga" esta área del navegador, liberando el Call Stack.

Navegador

Web APIs

Resolución de 4.Completa

JavaScript

Promesa Callback Event Loop

Colas de Espera

Microtask Queue Callback Queue 3.Delega a

Call Stack

2\.Encuentra

Funciones API Asíncrona

a

  \| 32 of 82

# Callback Queue

Una vez que una operación asíncrona de una Web API ha terminado (ej. un setTimeout ha expirado, una petición fetch ha recibido una respuesta):

su función de callback

se coloca en esta

cola.

Es una cola FIFO

Navegador

Web APIs

4\.Completa

JavaScript

Promesa Callback Event Loop

9\.Si no hay microtasks 5.Encola en

Colas de Espera

Microtask Queue Callback Queue

10\.Siguiente callback

Call Stack

Funciones API Asíncrona

  \| 33 of 82

# Microtask Queue

Una cola de mayor prioridad que la Callback Queue.

Contiene callbacks de

Promesas ( .then() ,

.catch () , .finally() ) y MutationObserver .

  Im portant

El Event Loop procesará todas las microtareas de esta cola antes de pasar a la siguiente macrotarea (de la Callback Queue).

Navegador

Web APIs

JavaScript

Promesa Callback Event Loop

7b.Prioridad 6.Encola en

Colas de Espera

Microtask Queue Callback Queue

8\.Ejecuta todos

Call Stack

Funciones API Asíncrona

  \| 34 of 82

# Event Loop

Monitorea constantemente el Call Stack y las colas de callbacks/microtareas:

Verifica si el Call Stack

está vacío.

Si está vacío, vacía completamente la Microtask Queue, moviendo cada callback al

Call Stack.

Una vez que la Microtask Queue está vacía, toma la primera tarea de la Callback Queue y la mueve al Call Stack.

Repite el proceso.

Navegador

Web APIs

JavaScript

Promesa Callback Event Loop

7b.Prioridad 9.Si no hay microtasks

Colas de Espera

Microtask Queue Callback Queue 7a.¿CallStack vacío?

Call Stack

Funciones API Asíncrona

  \| 35 of 82

# Ejemplo de ejecución de un programa JS

console.log("1. Inicio"); // Síncrono

setTimeout(() =\> { console.log("3. setTimeout callback"); // MacroTask }, 0); // ¡0ms no significa inmediatamente!

Promise.resolve().then(() =\> { console.log("2. Promise callback"); // MicroTask });

console.log("4. Fin"); // Síncrono

Resultado en consola:

1\. Inicio

4\. Fin

2\. Promise callback

3\. setTimeout callback

  \| 36 of 82

# Ejemplo de ejecución de un programa JS

console.log("1. Inicio"); // Síncrono

setTimeout(() =\> { console.log("3. setTimeout callback"); // MacroTask }, 0); // ¡0ms no significa inmediatamente!

Promise.resolve().then(() =\> { console.log("2. Promise callback"); // MicroTask }); console.log('1. Inicio') se ejecuta console.log("4. Fin"); // Síncrono

Resultado en consola:

1\. Inicio

4\. Fin

2\. Promise callback

3\. setTimeout callback

  \| 36 of 82

# Ejemplo de ejecución de un programa JS

El callback () =\> {

console.log('3. setTimeout

callback'); } se envía a las

Web APIs.

El temporizador de 0ms expira casi instantáneamente.

El callback se mueve a la Callback

Queue.

console.log("1. Inicio"); // Síncrono

setTimeout(() =\> { console.log("3. setTimeout callback"); // MacroTask }, 0); // ¡0ms no significa inmediatamente!

Promise.resolve().then(() =\> { console.log("2. Promise callback"); // MicroTask });

console.log("4. Fin"); // Síncrono

Resultado en consola:

1\. Inicio

4\. Fin

2\. Promise callback

3\. setTimeout callback

  \| 36 of 82

# Ejemplo de ejecución de un programa JS

El callback () =\> { console.log('2. Promise

callback'); } se envía a la

Microtask Queue.

console.log("1. Inicio"); // Síncrono

setTimeout(() =\> { console.log("3. setTimeout callback"); // MacroTask }, 0); // ¡0ms no significa inmediatamente!

Promise.resolve().then(() =\> { console.log("2. Promise callback"); // MicroTask });

console.log("4. Fin"); // Síncrono

Resultado en consola:

1\. Inicio

4\. Fin

2\. Promise callback

3\. setTimeout callback

  \| 36 of 82

# Ejemplo de ejecución de un programa JS

console.log("1. Inicio"); // Síncrono

setTimeout(() =\> { console.log("3. setTimeout callback"); // MacroTask }, 0); // ¡0ms no significa inmediatamente!

Promise.resolve().then(() =\> { console.log("2. Promise callback"); // MicroTask }); console.log('4. Fin') se ejecuta. console.log("4. Fin"); // Síncrono

Resultado en consola:

1\. Inicio

4\. Fin

2\. Promise callback

3\. setTimeout callback

  \| 36 of 82

# Ejemplo de ejecución de un programa JS

El Call Stack está vacío y el Event Loop mira la Microtask Queue:

Encuentra () =\> {

console.log('2. Promise

callback'); }.

Lo mueve al Call Stack

console.log("1. Inicio"); // Síncrono

setTimeout(() =\> { console.log("3. setTimeout callback"); // MacroTask }, 0); // ¡0ms no significa inmediatamente!

Promise.resolve().then(() =\> { console.log("2. Promise callback"); // MicroTask });

console.log("4. Fin"); // Síncrono

Resultado en consola:

1\. Inicio

4\. Fin

2\. Promise callback

3\. setTimeout callback

  \| 36 of 82

# Ejemplo de ejecución de un programa JS

console.log("1. Inicio"); // Síncrono

setTimeout(() =\> { console.log("3. setTimeout callback"); // MacroTask }, 0); // ¡0ms no significa inmediatamente!

Promise.resolve().then(() =\> { console.log("2. Promise callback"); // MicroTask }); console.log('2. Promise

callback') se ejecuta console.log("4. Fin"); // Síncrono

Resultado en consola:

1\. Inicio

4\. Fin

2\. Promise callback

3\. setTimeout callback

  \| 36 of 82

# Ejemplo de ejecución de un programa JS

El Call Stack y la Microtask Queue están vacías

El Event Loop mira la Callback Queue:

Encuentra () =\> {

console.log('3. setTimeout

callback'); }.

Lo mueve al Call Stack

console.log("1. Inicio"); // Síncrono

setTimeout(() =\> { console.log("3. setTimeout callback"); // MacroTask }, 0); // ¡0ms no significa inmediatamente!

Promise.resolve().then(() =\> { console.log("2. Promise callback"); // MicroTask });

console.log("4. Fin"); // Síncrono

Resultado en consola:

1\. Inicio

4\. Fin

2\. Promise callback

3\. setTimeout callback

  \| 36 of 82

# Ejemplo de ejecución de un programa JS

console.log("1. Inicio"); // Síncrono

setTimeout(() =\> { console.log("3. setTimeout callback"); // MacroTask }, 0); // ¡0ms no significa inmediatamente!

Promise.resolve().then(() =\> { console.log("2. Promise callback"); // MicroTask }); console.log('3. setTimeout

callback') se ejecuta console.log("4. Fin"); // Síncrono

Resultado en consola:

1\. Inicio

4\. Fin

2\. Promise callback

3\. setTimeout callback

  \| 36 of 82

# Ejemplo de ejecución de un programa JS

console.log("1. Inicio"); // Síncrono

setTimeout(() =\> { console.log("3. setTimeout callback"); // MacroTask }, 0); // ¡0ms no significa inmediatamente!

Promise.resolve().then(() =\> { console.log("2. Promise callback"); // MicroTask }); El Call Stack está vacío. console.log("4. Fin"); // Síncrono

Resultado en consola:

1\. Inicio

4\. Fin

2\. Promise callback

3\. setTimeout callback

  \| 36 of 82

¿Qué es HTML?

Es el lenguaje de marcado que estructura el contenido con el que el usuario interactúa.

HyperText Markup Language

Es un lenguaje de marcado para estructurar contenido en la web.

No es un lenguaje de programación, pero se integra con CSS/JS

Es estático, no dinámico.

No fue pensado para aplicaciones, sino para documentos.

Define la estructura y el contenido de una página web:

Títulos

Párrafos

Imágenes

Enlaces

Listas

Formularios

¡Y mucho más!

  \| 37 of 82

# HTML para programadores

Si piensan en una aplicación:

HTML es como la definición de la interfaz de usuario (UI) :

botones, campos de texto, etiquetas.

: CSS (Cascading Style Sheets) es como el diseño o"styling"

colores, fuentes, tamaños, posiciones de los elementos.

JavaScript es la lógica del cliente (interfaz) y la interactividad :

qué sucede cuando haces clic en un botón.

  \| 38 of 82

# Los componentes de HTML

HTML se construye con elementos. Cada elemento tiene una etiqueta de apertura y, generalmente, una etiqueta de cierre:

\<p\>Este es un párrafo.\</p\>

Los elementos pueden tener atributos que proporcionan información adicional:

\<a href=" target="\_blank"\>Visitar Ejemplo\</a\> https://www.ejemplo.com"

href : El destino del enlace.

target : Cómo se abre el enlace.

  \| 39 of 82

# Estructura Básica de un Documento HTML

\<!DOCTYPE html\>

\<html lang="es"\> \<head\>

\<meta charset=" UTF-8" /\>

\<meta

n ame=" viewport" content=" width=device-width, ini tial-scale=1.0"

/\>

\<title\>Mi Primera Página HTML5\</title\> \</head\>

\<body\> \<h1\>Hola, Mundo!\</h1\> \<p\>¡Bienvenidos a mi página web!\</p\> \</body\> \</html\>

\<!DOCTYPE html\> : Declara el tipo de documento. Esencial para HTML5.

\<html lang="es"\> : El elemento raíz

es"in dica el idioma. de la página. lang="

\<head\> : Contiene metadatos de la

página (no visibles para el usuario).

\<meta charset=" UTF-8"\> :

Codificación de caracteres.

... \<meta name=" \> : Para viewport" la adaptabilidad móvil (responsive design).

\<title\> : El título que aparece en la pestaña del navegador.

\> : Contiene todo el contenido \<body visible de la página.

  \| 40 of 82

# Elementos HTML más comunes

ul

li

img

video Multimedia

audio

Texto p

span

ol a

label Listas Enlaces

input

Formularios

HTML form

button

Contenedores section

div h1...h6

article

  \| 41 of 82

# Tablas HTML

\<table\>

\<thead\>

\<tr\>

\<th\>Nombre\</th\>

\<th\>Edad\</th\>

\</tr\>

\</thead\>

\<tbody\> \<tr\>

\<td\>Pepa\</td\> \<td\>23\</td\>

\</tr\>

\<tr\>

\<td\>Juan\</td\>

\<td\>28\</td\>

\</tr\>

\<tr\>

\<td\>Rodolfo\</td\>

\<td\>26\</td\>

\</tr\>

\</tbody\> \</table\>

Nombre Edad

Pepa 23

Juan 28

Rodolfo 26

  \| 42 of 82

# Formularios HTML

\<p\>Login\</p\> \<form action=" /submit"m POST"\> ethod="

username:

text"n ame=" username"r \<input type=" equired /\>\<br /\> password: ame=" \<input type=" password"n pwd"r equired /\>\<br /\> email:

email"n ame=" email"r \<input type=" equired /\>\<br /\> contact preference: \<select name=" contact"\>

email"\>Email\</option\> \<option value=" \<option value=" phone"\>Phone\</option\>\</select \>\<br /\>

delivery options:\<br /\> checkbox"i ame=" d=" standard"n \<input type=" delivery"v \<label for=" standard"\>Standard\</label\>\<br /\>

checkbox"i ame=" d=" \<input type=" express"n delivery"v \<label for=" express"\>Express\</label\>\<br /\> submit"\>Enviar\</button\> \<button type=" \</form\>

Login username: MyName \*\*\*\* password: email: my@email.com contact: Email

\[x\] Standard Express

Enviar

alue=" /\> standard"

alue=" /\> express"

  \| 43 of 82

# Mejorando el estilo con CSS

CSS (Cascading Style Sheets) se encarga de la presentación y el estilo de los elementos HTML.

Define colores, tipografías, tamaños, espaciados, posiciones y mucho más.

Sin CSS, el HTML se vería muy simple y sin diseño.

Se compone de selectores y reglas de estilo:

selector { property: value; }

selector : especifica el elemento HTML al que se aplicará el estilo.

property : el nombre de la propiedad CSS que se desea modificar.

value : el valor que se asigna a la propiedad CSS.

  \| 44 of 82

# Mejorando el estilo con CSS

...

\<head\>

...

\<style\> p { color: blue; font-weight: bold; } \</style\> \</head\>

\<body\> \<h1\>Hola, Mundo!\</h1\> \<p\>¡Bienvenidos a mi página web!\</p\> \</body\>

Hola, Mundo!

¡Bienvenidos a mi página web!

  \| 45 of 82

# Mejorando el estilo con CSS

...

\<head\>

...

\<style\> p { color: blue; font-weight: bold; } \</style\> \</head\>

\<body\> \<h1\>Hola, Mundo!\</h1\> \<p\>¡Bienvenidos a mi página web!\</p\> \</body\>

Hola, Mundo!

¡Bienvenidos a mi página web!

Se puede aplicar CSS directamente en el HTML con style :

color: blue; font-weight: bold"\>¡Bienvenidos a mi página web!\</p\> \<p style="

O con un archivo CSS externo:

\<head\>

\<link rel=" ef=" /\> stylesheet"hr styles.css" \</head\>

  \| 45 of 82

# Selectores CSS

Se utilizan para seleccionar los elementos HTML a los que se les aplicarán los estilos.

Selector de elemento: todos los elementos de un tipo específico.

p { /\* Selecciona todos los párrafos\*/ color: green; }

Selector de clase (.): elementos con un atributo class específico. Un elemento puede tener múltiples clases

resaltado"\>Texto importante\</p\> \<p class=" \<div class=" tarjeta"\>Contenido\</div\>

.resaltado{

elementos con class=" /\*

f ont-weight: bold; }

\*/ resaltado"

  \| 46 of 82

# Selectores CSS

Selector de ID (#): elemento único con un atributo id específico. Los IDs deben ser únicos en una página.

titulo- \<h2 id=" principal"\>Mi Título\</h2\>

\#titulo- principal { \*/ titulo- /\* elem con id=" principal" text-align: center; }

Selectores combinadores:

: descendientes

\> : hijo directo

\+ : hermano adyacente

\~ : hermano general

  \| 47 of 82

# Propiedades y Valores de CSS

Una vez seleccionado un elemento, se le aplican declaraciones de estilo. Cada declaración consiste en una propiedad y un valor.

Selector Propiedad: Valor;

Propiedad: El aspecto del elemento que quieres cambiar (ej. color , font size , margin ).

Valor: El ajuste específico para esa propiedad (ej. blue , 16px ).

Se escriben como { propiedad: valor;

... }

ste es un comentario en CSS\*/ /\*E

h1 { color: blue; /\* Color del texto\*/

color: #333; /\*Ti po de fuente \*/ f ont-family: Arial, sans-serif; font-size: 16px; /\*Alin eación del texto\*/

text-align: center; } margin-top: 10px;

.tar jeta { /\* Color de fondo\*/

background-color: #f0f0f0; /\*E spacio interno\*/ padding: 20px; /\*B orde\*/

border: 1px solid #ccc; /\*M argen inferior\*/ m argin-bottom: 15px; }

  \| 48 of 82

# Modelo de caja de CSS

Content: El área donde se muestra el

Todos los elementos HTML se consideran contenido real del elemento (texto, cajas rectangulares: imágenes).

Padding: Espacio transparente entre el contenido y el borde. Empuja el borde hacia afuera.

Border: La línea que rodea el padding y el contenido.

Margin: Espacio transparente fuera del borde. Empuja otros elementos lejos de esta caja.

  \| 49 of 82

<image redacted: 225x164px, 225x164pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

# Modelo de caja de CSS

.ca ja { wi cho del contenido\*/ dth: 200px; /\*An h to del contenido\*/ eight: 100px; /\*Al padding: 15px; /\*1 5px de padding en todos los lados\*/ border: 2px solid black; /\*B orde de 2px sólido negro\*/ m argin: 10px auto; /\*1 0px de margen arriba/abajo, auto para centrar horizontalmente\*/ background-color: lightblue; }

  \| 50 of 82

<image redacted: 188x152px, 188x152pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

# Display: block e inline

La propiedad display determina cómo se comporta un elemento en el flujo del documento.

Los elementos HTML tienen una propiedad display por defecto (se puede cambiar con CSS)

inline :

Ocupa solo el ancho necesario para su contenido.

No comienza en una nueva línea.

No permite definir width ni height . Solo margin y padding horizontales.

Ejemplos: span , a , img , strong ,

em .

block :

Ocupa todo el ancho disponible.

Comienza en una nueva línea.

Permite definir width , height , margin y padding en todas las direcciones.

Ejemplos: div , p , h1 , ul , li .

  \| 51 of 82

# Display: block e inline

inline-block :

Se comporta como un elemento inline : no rompe la línea.

Permite definir width , height , margin y padding en todas las direcciones (como block ).

Útil para elementos que necesitan tamaño pero no deben ocupar toda la línea.

  \| 52 of 82

# Display: block e inline

\<span style=" \<span style="

\<div style=" \<div style="

\<a href="#"

style=" \>Botón\</a\>

\<a href="#"

style="

\>Otro Botón\</a\>

background-color: yellow"\>Soy un span (inline) \</span\> background-color: lightgreen"\>Otro span\</span\>

\</div\> background-color: lightcoral"\>Soy un div (block) background-color: lightblue"\>Otro div\</div\>

display: inline-block; width: 100px; height: 30px; background-color: orange;"

display: inline-block; width: 100px; height: 30px; background-color: purple; color: white;"

  \| 53 of 82

<image redacted: 248x85px, 248x85pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

# Layouts con CSS: Flexbox (Flexible Box

# Layout)

Un modelo de diseño unidimensional para . organizar elementos en filas o columnas

Ideal para distribuir el espacio y alinear ítems dentro de un contenedor.

  \| 54 of 82

# Layouts con CSS: Flexbox (Flexible Box

# Layout)

Un modelo de diseño unidimensional para . organizar elementos en filas o columnas

Ideal para distribuir el espacio y alinear ítems dentro de un contenedor.

  \| 54 of 82

# Layouts con CSS: Flexbox

Contenedor Flex (display: flex)

1 2 3

flex-direction: row (default)

flex- grow: los items 1 y 2 se expanden

flex-wrap: wrap (ajusta múltiples líneas)

flex-direction flex-wrap flex-

row row-reverse column column-reverse

(horizontal →) (← horizontal) (vertical ↓) (↑ vertical) (una línea)

c ol u m n

Flex Container

grow

nowrap wrap wrap-reverse Número (proporción de (múltiples líneas) (líneas invertidas) espacio libre)

  \| 55 of 82

# Layouts con CSS: Flexbox

\<div class=" contenedor-flex"\>

item-flex"\>Item 1\</div\> \<div class="

item-flex"\>Item 2\</div\> \<div class="

item-flex"\>Item 3\</div\> \<div class="

\</div\>

.contenedor-flex {

/\* Convierte el div en un contenedor flex\*/

display: flex; /\*L os ítems se organizan en fila\*/ fl ex-direction: row; /\*Di stribuye el espacio alrededor de los ítems\*/ justify-content: space-around; Centra los ítems verticalmente\*/ /\*

align-items: center; h eight: 150px; background-color: #eee; border: 1px solid #ccc; }

.item-flex {

background-color: lightblue; padding: 20px; m argin: 5px; border: 1px solid blue; }

  \| 56 of 82

<image redacted: 225x94px, 225x93pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

# Layouts con CSS: Flexbox

Propiedades del Contenedor Flex:

flex-direction : Dirección del eje principal (row, column, row-reverse, column reverse).

justify-content : Alinea ítems a lo largo del eje principal (flex-start, flex-end, center, space-between, space-

align-items : Alinea ítems a lo largo del eje cruzado (flex-start, flex-end, center, baseline).

flex-wrap : Controla si los ítems deben envolverse en múltiples líneas (nowrap,

wrap, wrap-reverse).

gap : Espacio entre ítems (atajo para row-

around).

gap y column- gap).

  \| 57 of 82

# Layouts con CSS: Flexbox

Propiedades de los Ítems Flex:

flex- grow : Define la capacidad de un ítem para crecer.

flex-shrink : Define la capacidad de un ítem para encogerse.

flex-basis : Define el tamaño inicial de un ítem antes de que crezca o se encoja.

flex : Atajo para flex- grow, flex-shrink y flex-basis.

align-self : Sobrescribe align-items para un ítem individual.

  \| 58 of 82

# Layouts con CSS: Grid

Modelo de diseño bidimensional para organizar elementos en filas y columnas.

Ideal para diseñar la estructura principal de una página o áreas complejas.

1 2 3

4 5 6

grid-template-columns: 1fr 1fr 1fr

Grid responsive con auto-fit grid-template-columns: repeat(auto-fit, minmax(120px, 1fr))

header

nav main aside

grid-template-areas

  \| 59 of 82

# CSS Grid Layout

grid-template-columns

grid-template-rows

Grid Container

grid-template-areas

gap

fr (unidad flexible) Propiedades del Contenedor Grid:

grid-template-columns : repeat() Define el número y ancho de las columnas (ej. 1fr 1fr 1fr, auto minmax() 200px 1fr).

grid-template-rows : Define el auto-fit/auto-fill

número y alto de las filas.

Nombrado de áreas gap : Espacio entre celdas (atajo para row- gap y column Posicionamiento con grid gap). area

grid-template-areas : Permite row- gap nombrar áreas de la cuadrícula

para posicionar elementos column- gap fácilmente.

  \| 60 of 82

# CSS Grid Layout

grid-template-columns

grid-template-rows

Grid Container

grid-template-areas

gap

fr (unidad flexible) Propiedades del Contenedor Grid:

grid-template-columns : repeat() Define el número y ancho de las columnas (ej. 1fr 1fr 1fr, auto minmax() 200px 1fr). La unidad\`fr\`r epresenta una fracción del grid-template-rows : Define el auto-fit/auto-fill espacio disponible en el contenedor de la número y alto de las filas. grilla. Por ejemplo, si tenemos 2 columnas

Nombrado de áreas cada una ocupará la y cada una tiene\`1fr\`, gap : Espacio entre celdas mitad del espacio disponible. (atajo para row- gap y column Posicionamiento con grid gap). area

grid-template-areas : Permite row- gap nombrar áreas de la cuadrícula

para posicionar elementos column- gap fácilmente.

  \| 60 of 82

# Ejemplo de CSS Grid Layout

\<div class=" contenedor- grid"\> item- \<div class=" grid header"\>Header\</div\> item- \<div class=" grid sidebar"\>Sidebar\</div\> item- \<div class=" grid content"\>Content\</div\> item- \<div class=" grid footer"\>Footer\</div\> \</div\>

  \| 61 of 82

<image redacted: 248x206px, 248x206pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

# Ejemplo de CSS Grid Layout

.contenedor-grid {

/\* Convierte el div en un contenedor grid\*/ display: grid; /\* Columna izquierda 1 parte, derecha 3 partes\*/ grid-template-columns: 1fr 3fr; /\*Fil a de arriba auto, medio 1 parte, abajo auto\*/ grid-template-rows: auto 1fr auto; /\*E spacio de 10px entre celdas\*/ gap: 10px; grid-template-areas:""header header

""sidebar content

efine áreas con nombres\*/ /\*D

"footer footer";

h eight: 300px; background-color: #f9f9f9; border: 1px solid #ddd; }

.item-grid {

background-color: #e0f7fa; padding: 15px; border: 1px solid #00bcd4; text-align: center; } .header{

grid-area: header; background-color: #b3e5fc; } .sidebar{

grid-area: sidebar; background-color: #c8e6c9; } .content{

grid-area: content; background-color: #ffccbc; } .footer{

grid-area: footer; background-color: #ffe0b2; }

  \| 62 of 82

# Sólo una introducción al Navegador de Web

Explora más propiedades CSS: Hay muchísimas más, como transformaciones, transiciones, animaciones, position, z-index, etc.

Diseño Responsivo: para adaptar tus diseños a diferentes tamaños de pantalla (@media queries).

Herramientas de Desarrollador: para depurar y experimentar con estilos (tecla F12 en el navegador).

JavaScript: HTML y CSS son lo básico, el siguiente paso es JavaScript para añadir interactividad.

Web APIs: APIs para acceder a funcionalidades avanzadas (ej. geolocalización, almacenamiento local, etc.).

WebAssembly: para ejecutar código de otros lenguajes en el navegador.

Web Workers: para ejecutar código en segundo plano sin bloquear la UI.

https://developer.mozilla.org/en-US/

  \| 63 of 82

# Decisiones de Diseño

La capa de presentación ejecuta en el backend y en el navegador de Web, pero…:

¿quién valida la entrada del usuario?

¿quién ordena las filas de una tabla?

¿quién genera una visualización de datos?

¿quién almacena el carrito de compras?

¿quién calcula el total de una compra?

  \| 64 of 82

# Decisiones de Diseño

La capa de presentación ejecuta en el backend y en el navegador de Web, pero…:

¿quién valida la entrada del usuario?

¿quién ordena las filas de una tabla?

¿quién genera una visualización de datos?

¿quién almacena el carrito de compras?

¿quién calcula el total de una compra?

¿Quién decide de qué color se pinta cada zona? ¿Quién pinta cada zona?

  \| 64 of 82

<image redacted: 266x275px, 266x275pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

# Línea de Tiempo de Tecnologías Web

1990

2000

2010

2015

2020

2023

  \| 65 of 82

<image redacted: 237x39px, 236x38pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 237x26px, 236x25pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 237x26px, 236x25pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 237x26px, 236x25pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 237x26px, 236x25pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 237x26px, 236x25pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

# 1990: CGI + HTML Estático

La web era un"directorio global" de solo lectura, con algo de dinamismo del lado del servidor.

Características Principales:

Páginas estáticas para consumir información, con formularios para enviar datos.

La interactividad se lograba recargando la página tras enviar un formulario.

CGI (Common Gateway Interface): Permitía a los servidores web ejecutar scripts externos (ej. en Perl) para procesar datos y generar HTML dinámicamente.

  \| 66 of 82

# 1990: CGI + HTML Estático

Avances Clave:

Nacimiento de la World Wide Web y los primeros navegadores.

Posibilidad de procesar entradas de usuario en el servidor.

Tecnologías:

HTML: Para estructurar el contenido.

HTTP: Protocolo para transferir datos.

CGI: Para la lógica del servidor, comúnmente con lenguajes como Perl o C.

  \| 67 of 82

# 2000: PHP/JSP/ASP (SSR)

El contenido se genera de forma más integrada en el servidor (Server-Side Rendering).

Características Principales:

El código de la lógica se mezcla directamente con el HTML.

Facilita la creación de sitios con contenido de bases de datos (foros, blogs, CMS).

del lado del servidor. Nace el concepto de"plantillas"

  \| 68 of 82

# 2000: PHP/JSP/ASP (SSR)

Avances Clave:

Se simplifica el desarrollo de aplicaciones web dinámicas.

CSS madura, permitiendo una mejor separación entre contenido y presentación.

JavaScript se usa para validaciones y pequeños efectos visuales.

Tecnologías:

PHP, ASP (Microsoft), JSP (Java): Lenguajes/plataformas que se ejecutan en el servidor para generar HTML.

Bases de datos: MySQL, PostgreSQL se vuelven populares.

  \| 69 of 82

# 2010: AJAX + SPA Incipientes

La web se vuelve interactiva, actualizando datos sin recargar la página.

Características Principales:

AJAX (Asynchronous JavaScript and XML): Permite a JS hacer peticiones HTTP asíncronas.

La interfaz se siente más rápida y fluida, como una aplicación de escritorio.

Nacen las primeras Single-Page Applications (SPAs).

  \| 70 of 82

# 2010: AJAX + SPA Incipientes

Avances Clave:

jQuery: Librería que simplificó drásticamente la manipulación del DOM y las llamadas AJAX.

JSON reemplaza a XML como el formato de datos preferido por su simplicidad.

Tecnologías:

JavaScript: Se convierte en el pilar de la interactividad.

jQuery, MooTools, Dojo: Librerías para facilitar el desarrollo en JS.

APIs RESTful: Se popularizan en el backend para servir datos en formato JSON.

  \| 71 of 82

# 2015: React/Vue/Angular

La era de los frameworks de JavaScript y las arquitecturas basadas en componentes.

Características Principales:

Componentes: La UI se construye como un árbol de piezas reutilizables y autocontenidas.

DOM Virtual: Los frameworks optimizan la manipulación del DOM para un rendimiento superior.

El estado de la aplicación se gestiona de forma declarativa.

  \| 72 of 82

# 2015: React/Vue/Angular

Avances Clave:

El desarrollo del frontend se profesionaliza y se vuelve más complejo.

Ecosistema masivo de herramientas: gestores de paquetes (npm), bundlers (Webpack, Vite).

Tecnologías:

Frameworks: Angular (de Google), React (de Facebook), Vue.js.

JavaScript Moderno (ES6+): Introduce clases, módulos, promesas, etc.

TypeScript: Añade un sistema de tipos a JavaScript, mejorando la robustez.

  \| 73 of 82

# 2020: JAMstack + SSR Moderno

Se busca el balance perfecto entre rendimiento, SEO y experiencia de desarrollador.

Características Principales:

JAMstack (JavaScript, APIs, Markup): Pre-renderizar sitios estáticos para máxima velocidad y servirlos desde un CDN.

SSR Moderno: El mismo código JS se ejecuta en el servidor (para la primera carga) y en el cliente (para la interactividad).

SSG (Static Site Generation) y ISR (Incremental Static Regeneration).

  \| 74 of 82

# 2020: JAMstack + SSR Moderno

Avances Clave:

Tiempos de carga casi instantáneos.

Mejora del SEO para SPAs.

Tecnologías:

Meta-Frameworks: Next.js (para React), Nuxt.js (para Vue), SvelteKit.

Headless CMS: Sistemas de gestión de contenido que solo proveen una API.

GraphQL: Lenguaje de consulta para APIs que permite al cliente pedir solo los datos que necesita.

  \| 75 of 82

# 2023: Islands Architecture + Edge SSR

Optimizaciones para enviar menos JavaScript y acercar la computación al usuario.

Características Principales:

de interactividad Arquitectura de Islas: La página es HTML estático con"islas" solo donde se necesita.

Hidratación Parcial: Solo se carga el JS necesario para los componentes interactivos, no para toda la página.

Edge Computing: El renderizado del lado del servidor (SSR) se realiza en servidores CDN en el"borde" de la red, más cerca del usuario.

  \| 76 of 82

# 2023: Islands Architecture + Edge SSR

Avances Clave:

Reducción drástica del JavaScript enviado al cliente.

Time To Interactive (TTI) mucho más rápido.

Tecnologías:

Frameworks: Astro, Qwik, Fresh (Deno).

Edge Functions: Vercel Edge Functions, Cloudflare Workers, Deno Deploy.

  \| 77 of 82

# Server-Side Rendering (SSR) Tradicional

Request Navegador Servidor Web App Server Base de Datos HTML Completo

El usuario solicita una página.

El servidor procesa la solicitud, obtiene los datos necesarios (ej. base de datos) y renderiza HTMLs.

Se envía un documento HTML, CSS y JS completamente listo para ser mostrado.

El navegador simplemente muestra el contenido. La interactividad puede ser añadida con JS.

Cada interacción requiere recarga.

Ejemplos: PHP, Ruby on Rails, Go html/template.

  \| 78 of 82

# Client-Side Rendering (CSR)

Request HTML inicial Servidor Web

HTML vacío + JS

API Calls

Backend JSON Navegador

JSON

Renderizado Client-Side

DOM

El usuario solicita la aplicación.

El servidor responde con un HTML mínimo y un bundle de JS.

El navegador descarga y ejecuta el JS.

El framework de JS toma el control, realiza llamadas a APIs para obtener datos (en formato JSON) y renderiza la interfaz dinámicamente

en el cliente

Interacciones sin recarga

Ejemplos: React, Vue, Angular

  \| 79 of 82

# Comparativa: SSR vs. CSR

Server-Side Rendering Característica (SSR)

Rápida (contenido visible al Carga inicial instante)

SEO Excelente y directo

Interactividad Lenta (requiere recargas) Rápida (sin recargas)

Alta (renderiza en cada Carga del servidor petición)

Carga del cliente Baja (sólo renderiza) Alta (estado, conversiones, DOM)

Complejidad Menor en el frontend Mayor en el frontend (estado)

Client-Side Rendering (CSR)

Lenta (pantalla en blanco inicial)

Requiere configuraciones adicionales

Baja (sirve archivo y JSON)

  \| 80 of 82

¿Hago SSR o CSR?

Intenta no ir a los extremos, usa una combinación de ambos

Depende mucho del la aplicación:

si es principalmente CRUD, SSR es más apropiado

si es muy interactiva, CSR es más adecuada

SSR no necesariamente es feo, malo o aburrido: se puede condimentar con CSR sin incrementar complejidad

Muchas modas y frameworks que aportan poco valor real

CSR requiere que el backend posea una API JSON y utilizar frameworks como React, Vue o Angular

  \| 81 of 82

¿Hago SSR o CSR?

Intenta no ir a los extremos, usa una combinación de ambos

Depende mucho del la aplicación:

si es principalmente CRUD, SSR es más apropiado

si es muy interactiva, CSR es más adecuada

SSR no necesariamente es feo, malo o aburrido: se puede condimentar con CSR sin incrementar complejidad

Muchas modas y frameworks que aportan poco valor real

CSR requiere que el backend posea una API JSON y utilizar frameworks como React, Vue o Angular

  \| 81 of 82

# El Espectro Intermedio: Soluciones Híbridas

Static Site Generation (SSG or JAMstack), HTML generado en compilación:

sitios donde el contenido no cambia frecuentemente. Ejemplos: , , .

Incremental Static Regeneration (ISR): Como SSG, pero permite regenerar páginas estáticas en segundo plano a medida que los datos cambian. Ejemplo: , .

Universal / Isomorphic Apps: El mismo código (generalmente JS) puede ejecutarse tanto en el servidor (para la carga inicial) como en el cliente (para la interactividad posterior). Ejemplos: (para ), (para ).

Hypermedia: aprovecha el poder de los enlaces y la navegación basada en hipermedios para crear aplicaciones web más interactivas y dinámicas. Ejemplo: ,Al pine-Ajax, Turbo, Unpoly.

  Im portant

Estas soluciones permiten aprovechar lo mejor de ambos mundos. Traen HTML del servidor, no JSON.

  \| 82 of 82