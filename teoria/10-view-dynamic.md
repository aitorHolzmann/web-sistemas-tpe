# Pr

# ogramación Web

# Dinamismo en la Capa de

# Presentación

# Dr. Alejandro Zunino & Dr. Alfredo Teyseyre

# alejandro.zunino@isistan.unicen.edu.ar, alfredo.teyseyre@isistan.unicen.edu.ar

  \| 1 of 60

<image redacted: 72x72px, 72x72pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

# Contenidos

1\. Navegación SSR completa

2\. Actualizaciones parciales de HTML

3\. Actualización parcial de HTML con Vanilla JS

4\. Actualización del DOM a partir de JSON

5\. Actualización parcial de HTML con HTMX

6\. Conclusiones

  \| 2 of 60

<image redacted: 479x479px, 479x479pt, ~72dpi, JPG, DEVICE_RGB, 32bpp>

# Navegación SSR completa

# Construimos una página que

# muestra una tabla de usuarios, pero

# todavía no permite interactuar con

ella.

# Con SSR y navegación tradicional,

# la interactividad de este ejemplo se

# logra mediante la recarga del

# documento completo , lo que

puede ser más lento y costoso.

# Con JS & DOM se evita recarga,

# actualizando solo fragmentos:

# Mayor esfuerzo pero

experiencia fluida .

## JS/DOM — Partial Update

## SSR — navegación completa

## Click th

## Click th

GET /?sort=...

## JS handler

## Server render

## fetch fragment

## o JSON

## HTML completo

## Browser repaint

## DOM patch

## solo tabla

## toda la página

  \| 3 of 60

# Navegación SSR completa

Al hacer click en las cabeceras de las columnas, el servidor ordena la tabla por esa

# columna:

# Dependiendo del orden, se envía una URL al servidor ej: /?sort=email&order=asc

El servidor procesa la solicitud, ordena los datos y renderiza el HTML actualizado.

# Para lograr esto hay que navegar a /?sort=COLUMNA&order=ORDEN , siendo:

COLUMNA: id , name o email dependiendo de donde se hizo click.

ORDER: asc o desc dependiendo del orden actual.

  \| 4 of 60

<image redacted: 173x114px, 173x114pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

# Navegación SSR completa

Cliente Servidor BaseDeDatos

Primera interacción (Orden por nombre ascendente)

GET /?sort=name&order=asc

HTML con tabla ordenada (↑ en Name)

Segunda interacción (Orden por ID descendente)

GET /?sort=id&order=desc

HTML con tabla ordenada (↓ en ID)

Cliente Servidor BaseDeDatos

OM users ORDER BY name ASC Query: SELECT\*FR

Resultados ordenados

OM users ORDER BY id DESC Query: SELECT\*FR

Resultados ordenados

  \| 5 of 60

# Navegación SSR completa

| Cliente Servidor Primera interacción (Orden por nombre ascendente) GET /?sort=name&order=asc Query: En este ejemplo, y el sentido del URL y en el HTML sesión del servidor. HTML con tabla ordenada (↑ en Name) Segunda interacción (Orden por ID descendente) GET /?sort=id&order=desc Query: HTML con tabla ordenada (↓ en ID) | BaseDeDatos OM users ORDER BY name ASC SELECT el estado de la columna orden se representa en la Resultados ordenados generado, no en una OM users ORDER BY id DESC SELECT Resultados ordenados |
|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Cliente Servidor                                                                                                                                                                                                                                                                                                                 | BaseDeDatos   \| 5 of 60                                                                                                                                                                       |

# En users.sql agregamos más SELECT

\-- n ame: ListUsers :many SELECT\*FR

\-- n ame: ListUsersOrderByIdAsc :many SELECT\*FR

\-- n ame: ListUsersOrderByIdDesc :many SELECT\*FR

\-- n ame: ListUsersOrderByNameAsc :many SELECT\*FR

\-- n ame: ListUsersOrderByNameDesc :many SELECT\*FR

\-- n ame: ListUsersOrderByEmailAsc :many SELECT\*FR

\-- n ame: ListUsersOrderByEmailDesc :many SELECT\*FR

OM users;

OM users ORDER BY id ASC;

OM users ORDER BY id DESC;

OM users ORDER BY name ASC;

OM users ORDER BY name DESC;

OM users ORDER BY email ASC;

OM users ORDER BY email DESC;

  \| 6 of 60

# El handler en users.go procesa la URL

func (h\*UserHandler) ListUsers(w http.ResponseWriter, r\*http.Request) { ctx := r.Context() sortColumn := r.URL.Query().Get("sort") sortOrder := r.URL.Query().Get("order") if sortOrder !="desc" { sortOrder ="asc" // Valor por defecto } var users \[\]db.User var err error

switch sortColumn { case"id":

if sortOrder =="asc" { users, err = h.queries.ListUsersOrderByIdAsc(ctx) } else { users, err = h.queries.ListUsersOrderByIdDesc(ctx) } case"name":

if sortOrder =="asc" { users, err = h.queries.ListUsersOrderByNameAsc(ctx) } else { users, err = h.queries.ListUsersOrderByNameDesc(ctx) }

  \| 7 of 60

# El handler en users.go procesa la URL

case"email":

if sortOrder =="asc" { users, err = h.queries.ListUsersOrderByEmailAsc(ctx) } else { users, err = h.queries.ListUsersOrderByEmailDesc(ctx) } default:

// Orden por defecto users, err = h.queries.ListUsers(ctx) sortColumn ="id" // Columna por defecto para resaltar sortOrder ="asc"

}

if err != nil { http.Error(w, err.Error(), http.StatusInternalServerError) return

}

views.UserPage(users, sortColumn, sortOrder).Render(r.Context(), w) }

  \| 8 of 60

# La página HTML ahora debe hacer cosas

El HTML generado ahora debe incluir la lógica para resaltar la columna ordenada y el

# orden actual:

CSS al rescate: aplicaremos un estilo sorted-desc o sorted-asc a la cabecera de

# la columna ordenada para indicar el orden actual con o . Ej:

th.sorted-asc::after { content:" ↑";

} th.sorted-desc::after { content:" ↓";

}

\<th class=" sortable sorted-desc"\>Name\</th\>

Al hacer click en una cabecera, se actualizará la URL con los parámetros sort y order .

Lo más simple y accesible es usar un enlace \<a\> dentro de la cabecera.

# También puede usarse JavaScript; en ese caso conviene construir la query con

URLSearchParams .

  \| 9 of 60

# La página HTML ahora debe hacer cosas

Usaremos scripts templates de templ para facilitar la generación de enlaces dinámicos.

## Por ejemplo, en views/user.templ :

// navigateToSort defines a reusable JavaScript function. //\`templ\`will

script navigateToSort(column string, order string) { // Use template literals for clean URL construction in JavaScript. window.location.href =\`?sort=${column}&order=${order}\`;

}

  N ote

templ permite definir scripts reutilizables que se pueden llamar desde cualquier parte del HTML mediante el tipo script .

ensure this script is available on the page.

  \| 10 of 60

# La página HTML ahora debe hacer cosas

Usaremos scripts templates de templ para facilitar la generación de enlaces dinámicos.

## Por ejemplo, en views/user.templ :

// navigateToSort defines a reusable JavaScript function. //\`templ\`will ensure this script is available on the page. script navigateToSort(column string, order string) { // Use template literals for clean URL construction in JavaScript. window.location.href =\`?sort=${column}&order=${order}\`;

}

  N ote

templ permite definir scripts reutilizables que se pueden llamar desde cualquier parte del HTML mediante el tipo script .

  Im portant

Este script se asocia a un evento onClick en las cabeceras de la tabla para cambiar el orden de los datos.

  \| 10 of 60

# Creación de componentes

Son HTML y código que se compilan en funciones que devuelven un templ.Component :

# Esto permite crear partes reutilizables de la interfaz de usuario

# Crearemos un componente SortableHeader que aplica estilo sorted-asc o

sorted-desc si es la columna ordenada y contiene un control para cambiar el orden.

func getNextSortOrder(columnName, currentSortColumn, currentSortOrder string) string { if currentSortColumn == columnName && currentSortOrder =="asc" { return"desc"

} return"asc"

}

templ SortableHeader(columnName, displayName, currentSortColumn, currentSortOrder string) { col" class= \<th scope=" { map\[string\] bool{"sortable": true,"sorted-" + currentSortOrder: currentSortColumn == columnName, } } \>

\<a href= columnName, { fmt.Sprintf("/?sort=%s&order=%s", getNextSortOrder(columnName, currentSortColumn, currentSortOrder)) }\> { displayName } \</a\>

\</th\>

}

  \| 11 of 60

# Creación de componentes

## class=

{ map\[string\] bool{ ... } } : define las clases CSS del encabezado. templ permite usar un mapa de string a bool para definir dinámicamente las

# clases CSS que se aplicarán:

## "

## "sortable

# : true : añade la clase sortable al encabezado (highlight para

indicar que es ordenable).

## "sorted-"

## \+ currentSortOrder: currentSortColumn == columnName : si la

columna actual ( columnName ) es la misma que la columna por la que se está

# ordenando ( currentSortColumn ), entonces se añade una clase como sorted

## asc o sorted-desc (dependiendo del valor de currentSortOrder )

  \| 12 of 60

# La página users.templ

templ UserPage(users \[\]db.User, sortColumn string, sortOrder string) { \<!DOCTYPE html\> ...

\<style\> table {width:100%;border-collapse:collapse;} th, td { padding: 8px 12px; border: 1px solid #ddd; text-align: left; } th.sortable { cursor: pointer; } th.sortable:hover {background-color:#f2f2f2;} th.sorted-asc::after { content:" th.sorted-desc::after { content:" \</style\> \</head\>

\<body\> \<h1\>Users\</h1\>

\<table\>

\<thead\>

\<tr\>

@SortableHeader("id", @SortableHeader("name", @SortableHeader("email",

↑" ; } ↓" ; }

"ID", sortColumn, sortOrder)"Name", sortColumn, sortOrder)"Email", sortColumn, sortOrder)

  \| 13 of 60

# La página users.templ

templ UserPage(users \[\]db.User, sortColumn string, sortOrder string) { \<!DOCTYPE html\> ...

\<style\> table {width:100%;border-collapse:collapse;} th, td { padding: 8px 12px; border: 1px solid #ddd; text-align: left; } th.sortable { cursor: pointer; } th.sortable:hover {background-color:#f2f2f2;} th.sorted-asc::after { content:" th.sorted-desc::after { content:" \</style\> \</head\>

\<body\> \<h1\>Users\</h1\>

\<table\>

\<thead\>

\<tr\>

@SortableHeader("id", @SortableHeader("name", @SortableHeader("email",

↑" ; } ↓" ; }

"ID", sortColumn, sortOrder)"Name", sortColumn, sortOrder)"Email", sortColumn, sortOrder)

  \| 13 of 60

# La página users.templ

templ UserPage(users \[\]db.User, sortColumn string, sortOrder string) { \<!DOCTYPE html\> ...

\<style\> table {width:100%;border-collapse:collapse;} th, td { padding: 8px 12px; border: 1px solid #ddd; text-align: left; } th.sortable { cursor: pointer; } th.sortable:hover {background-color:#f2f2f2;} th.sorted-asc::after { content:" th.sorted-desc::after { content:" \</style\> \</head\>

\<body\> \<h1\>Users\</h1\>

\<table\>

\<thead\>

\<tr\>

@SortableHeader("id", @SortableHeader("name", @SortableHeader("email",

↑" ; } ↓" ; }

"ID", sortColumn, sortOrder)"Name", sortColumn, sortOrder)"Email", sortColumn, sortOrder)

  \| 13 of 60

# La página users.templ

\</tr\>

\</thead\>

\<tbody\> if len(users) \> 0 { for \_, user := range users { \<tr\>

\<td\> {fmt.Sprint(user.ID)}\</td\> \<td\> {user.Name}\</td\> \<td\> {user.Email}\</td\> \</tr\>

} } else { \<tr\>

\<td colspan="3"\>No users found.\</td\> \</tr\>

} \</tbody\> \</table\>

\</body\> \</html\>

}

  \| 14 of 60

# Observaciones de la navegación tradicional

# El handler coordina parámetros,

# consulta y renderizado; es simple,

# aunque acopla el flujo HTTP con la

presentación.

# Carga inicial rápida, poco JS, pocos

recursos en el cliente.

# Compatibilidad con todos los

navegadores.

# No requiere frameworks de JavaScript

# pesados, transpilers ni herramientas de

construcción complejas.

# SEO-friendly: las URLs son amigables

para los motores de búsqueda.

# Cada interacción requiere una nueva

solicitud al servidor.

# Difícil aprovechar la interactividad

moderna de los navegadores.

  \| 15 of 60

# Actualizaciones parciales de HTML

La navegación tradicional de este ejemplo tiene una limitación: cada click recarga el documento completo.

# Veremos alternativas que actualizan solo un fragmento de la página, aprovechando:

# JavaScript para solicitar datos al servidor ( fetch o HTMX)

# DOM para reemplazar solo la parte afectada, sin recarga

  \| 16 of 60

# Actualizaciones parciales de HTML

La navegación tradicional de este ejemplo tiene una limitación: cada click recarga el documento completo.

# Veremos alternativas que actualizan solo un fragmento de la página, aprovechando:

# JavaScript para solicitar datos al servidor ( fetch o HTMX)

# DOM para reemplazar solo la parte afectada, sin recarga

# Alternativas:

# Vanilla JS:

# fragmentos HTML

# JSON → HTML en el cliente

# JSON + JavaScript:

# HTMX:

# fragmentos HTML declarativo

  \| 16 of 60

# Actualización parcial de HTML con Vanilla JS

# AJAX (Asynchronous JavaScript and XML) es el nombre histórico del patrón de

# realizar solicitudes asíncronas y actualizar partes de una página sin recargarla

completamente.

fetch es una API de JavaScript que permite realizar solicitudes HTTP de manera

# asíncrona:

la idea es utilizar onClick en las cabeceras de la tabla para hacer una solicitud

# GET mediante fetch con los parámetros sort y order al servidor

# el servidor responderá con un fragmento de HTML

con JS/DOM se reemplazará el contenido de la tabla actual con la nueva

AJAX

\<table\> Server ...\</table\>

  \| 17 of 60

# El handler en users.go procesa una solicitud

# con fetch

func (h\*UserHandler) ListUsers(w http.ResponseWriter, r\*http.Request) { ctx := r.Context()

// Obtener parámetros de ordenamiento sortColumn := r.URL.Query().Get("sort") sortOrder := r.URL.Query().Get("order") if sortOrder !="desc" { sortOrder ="asc" // Valor por defecto } var users \[\]db.User var err error

switch sortColumn { case"id":

if sortOrder =="asc" { users, err = h.queries.ListUsersOrderByIdAsc(ctx) } else { users, err = h.queries.ListUsersOrderByIdDesc(ctx) }

  \| 18 of 60

# El handler en users.go procesa una solicitud

# con fetch

case"name":

if sortOrder =="asc" { users, err = h.queries.ListUsersOrderByNameAsc(ctx) } else { users, err = h.queries.ListUsersOrderByNameDesc(ctx) } case"email":

if sortOrder =="asc" { users, err = h.queries.ListUsersOrderByEmailAsc(ctx) } else { users, err = h.queries.ListUsersOrderByEmailDesc(ctx) } default:

users, err = h.queries.ListUsers(ctx)"" if sortColumn == { sortColumn ="id" // Columna por defecto para resaltar } }

  \| 19 of 60

# El handler en users.go procesa una solicitud

# con fetch

if err != nil { http.Error(w, err.Error(), http.StatusInternalServerError) return

}

// Check for AJAX request if r.Header.Get("X-Requested-With") == "XMLHttpRequest" { // If it's an AJAX request, only render the table

views.UserTable(users, sortColumn, sortOrder).Render(r.Context(), w) return

}

// For initial page load, render the full page views.UserPage(users, sortColumn, sortOrder).Render(r.Context(), w) }

  \| 20 of 60

# El handler en users.go procesa una solicitud

# con fetch

if err != nil { http.Error(w, err.Error(), http.StatusInternalServerError) return

}

// Check for AJAX request if r.Header.Get("X-Requested-With") == "XMLHttpRequest" { // If it's an AJAX request, only render the table

views.UserTable(users, sortColumn, sortOrder).Render(r.Context(), w) return

}

// For initial page load, render the full page views.UserPage(users, sortColumn, sortOrder).Render(r.Context(), w) }

  \| 20 of 60

# La página users.templ

templ UserPage(users \[\]db.User, sortColumn string, sortOrder string) { \<!DOCTYPE html\>

\<html lang="es"\> \<head\>

\<meta charset=" UTF-8"/\>

\<meta name=" content=" width=device-width, initial-scale=1.0"/\> viewport" \<title\>Lista de usuarios\</title\>

\<style\> body { font-family: sans-serif; } table { width: 100%; border-collapse: collapse; } th, td { padding: 8px 12px; border: 1px solid #ddd; text-align: left; } th.sortable button { cursor: pointer; } th.sortable button:hover { background-color: #f2f2f2; } ↑" th.sorted-asc::after { content:" ; } ↓" th.sorted-desc::after { content:" ; } \</style\> \</head\>

\<body\> \<h1\>Usuarios\</h1\>

users-table-container"\> @UserTable(users, sortColumn, sortOrder) \</div\> \<div id="

@sortScript() \</body\> \</html\>

}

  \| 21 of 60

# La página users.templ

templ UserPage(users \[\]db.User, sortColumn string, sortOrder string) { \<!DOCTYPE html\>

\<html lang="es"\> \<head\>

\<meta charset=" UTF-8"/\>

\<meta name=" content=" width=device-width, initial-scale=1.0"/\> viewport" \<title\>Lista de usuarios\</title\>

\<style\> body { font-family: sans-serif; } table { width: 100%; border-collapse: collapse; } th, td { padding: 8px 12px; border: 1px solid #ddd; text-align: left; } th.sortable button { cursor: pointer; } th.sortable button:hover { background-color: #f2f2f2; } ↑" th.sorted-asc::after { content:" ; } ↓" th.sorted-desc::after { content:" ; } \</style\> \</head\>

\<body\> \<h1\>Usuarios\</h1\>

users-table-container"\> @UserTable(users, sortColumn, sortOrder) \</div\> \<div id="

@sortScript() \</body\> \</html\>

}

  \| 21 of 60

# La página users.templ

templ UserPage(users \[\]db.User, sortColumn string, sortOrder string) { \<!DOCTYPE html\>

\<html lang="es"\> \<head\>

\<meta charset=" UTF-8"/\>

\<meta name=" content=" width=device-width, initial-scale=1.0"/\> viewport" \<title\>Lista de usuarios\</title\>

\<style\> body { font-family: sans-serif; } table { width: 100%; border-collapse: collapse; } th, td { padding: 8px 12px; border: 1px solid #ddd; text-align: left; } th.sortable button { cursor: pointer; } th.sortable button:hover { background-color: #f2f2f2; } ↑" th.sorted-asc::after { content:" ; } ↓" th.sorted-desc::after { content:" ; } \</style\> \</head\>

\<body\> \<h1\>Usuarios\</h1\>

users-table-container"\> @UserTable(users, sortColumn, sortOrder) \</div\> \<div id="

@sortScript() \</body\> \</html\>

}

  \| 21 of 60

# La página users.templ

templ UserTable(users \[\]db.User, sortColumn, sortOrder string) { users-table"\> \<table id="

\<thead\>

\<tr\>

"ID", @SortableHeader("id", sortColumn, sortOrder)"Name", @SortableHeader("name", sortColumn, sortOrder)"Email", @SortableHeader("email", sortColumn, sortOrder) \</tr\>

\</thead\>

\<tbody\> for \_, user := range users { \<tr\>

\<td\> { fmt.Sprint(user.ID) }\</td\> \<td\> { user.Name }\</td\> \<td\> { user.Email }\</td\> \</tr\>

} if len(users) == 0 { \<tr\> \<td colspan="3"\>No users found.\</td\> \</tr\> } \</tbody\> \</table\>

}

  \| 22 of 60

# La página users.templ

// SortableHeader is a reusable component for table headers. templ SortableHeader(columnName, displayName, currentSortColumn, currentSortOrder string) { col" \<th scope=" class= { map\[string\] bool{"sortable": true,

"sorted-" + currentSortOrder: currentSortColumn == columnName, } } data-column= { columnName } data-order= { getNextSortOrder(columnName, currentSortColumn, currentSortOrder) } \>

button"\>{ displayName }\</button\> \<button type=" \</th\>

}

  \| 23 of 60

# La página users.templ

templ sortScript() { \<script\> document.addEventListener('DOMContentLoaded', const tableContainer = document.getElementById('users-table-container'); if (!tableContainer) return; tableContainer.addEventListener('click', const header = e.target.closest('th.sortable'); if (!header) return;

const column = header.dataset.column; const order = header.dataset.order;

const params = new URLSearchParams({ sort: column, order }); fetch(\`/?${params}\`, { headers: {'X-Requested-With':'XMLHttpRequest', }) .then (response =\> { if (!response.ok) { throw new Error('Network error'); } return response.text(); })

() =\> {

(e) =\> {

},

  \| 24 of 60

# La página users.templ

then(html =\> { tableContainer.innerHTML = html;

// Update URL const url = new URL(window.location); url.searchParams.set('sort', url.searchParams.set('order', window.history.pushState({},'', }) .catch (error =\> { console.error('Error fetching and updating table:', }); }); }); \</script\> }

column); order); url);

error);

  \| 25 of 60

# Flujo Vanilla JS: fetch + fragmento HTML

## Click en th.sortable →

## closest +

## dataset.column/order

## fetch /?sort=&order con

## X-Requested-With:

## XMLHttpRequest

# Servidor responde el

# fragmento HTML UserTable

## tableContainer.innerHTML

## =h

## tml

## history.pushState

# actualiza URL sin recarga

# catch loguea error de red

Usuario Browser Servidor

click th.sortable

column = dataset.column order = dataset.order

fetch GET /?sort=&order X-Requested-With: XMLHttpRequest

Fragmento HTML UserTable

tableContainer.innerHTML = html

history.pushState URL

sin recarga

tabla reordenada

Usuario Browser Servidor

  \| 26 of 60

# Actualización del DOM a partir de JSON

Ahora usaremos JS para ordenar los datos en el navegador.

El servidor devolverá los datos en formato JSON, y el navegador generará HTML con JavaScript y DOM.

# Ahora, el servidor tendrá dos endpoints:

/ : devuelve el HTML de la página.

/users : devuelve los datos de los usuarios en formato JSON.

3\. Tabla HTML insertada

ID Name

1 a

# AJAX: transformación JSON a HTML

{id:"1", JS

"a Server ...name:"} 2. JavaScript procesa los datos

1\. Servidor envía JSON

  \| 27 of 60

# El handler users.go devuelve HTML y JSON

func (h\*UserHandler) ServeHTTP(w http.ResponseWriter, r\*http.Request) { switch r.URL.Path { case"/":

h.ListUsers(w, r) case"/users":

h.GetUsersJSON(w, r) default:

http.NotFound(w, r) } }

  \| 28 of 60

# El handler users.go devuelve HTML y JSON

func (h\*UserHandler) ServeHTTP(w http.ResponseWriter, r\*http.Request) { switch r.URL.Path { case"/":

h.ListUsers(w, r) case"/users":

h.GetUsersJSON(w, r) default:

http.NotFound(w, r) } }

  \| 28 of 60

# El handler users.go devuelve HTML y JSON

func (h\*UserHandler) GetUsersJSON(w http.ResponseWriter, r\*http.Request) { ctx := r.Context() users, err := h.queries.ListUsers(ctx)

if err != nil { http.Error(w, err.Error(), http.StatusInternalServerError) return

}

w.H eader().Set("Content-Type","application/json") json.NewEncoder(w).Encode(users) }

func (h\*UserHandler) ListUsers(w http.ResponseWriter, r\*http.Request) { // For initial page load, render the full page vi ews.UserPage().Render(r.Context(), w) }

  \| 29 of 60

# El handler users.go devuelve HTML y JSON

func (h\*UserHandler) GetUsersJSON(w http.ResponseWriter, r\*http.Request) { ctx := r.Context() users, err := h.queries.ListUsers(ctx)

if err != nil { http.Error(w, err.Error(), http.StatusInternalServerError) return

}

w.H eader().Set("Content-Type","application/json") json.NewEncoder(w).Encode(users) }

func (h\*UserHandler) ListUsers(w http.ResponseWriter, r\*http.Request) { // For initial page load, render the full page vi ews.UserPage().Render(r.Context(), w) }

  \| 29 of 60

# Actualización del DOM a partir de JSON

## Miremos json.NewEncoder(w).Encode(users) :

## db/models.go define User con tags

## json:"..." :

type User struct { ID in t32 \`json:"id"\` N ame string\`json:"name"\` Em ail string\`json:"email"\` }

# Generado por sqlc desde

## db/schema/users.sql :

CREATE TABLE users ( i d SERIAL PRIMARY KEY,

n ame TEXT NOT NULL, email TEXT UNIQUE NOT NULL );

## SQL schema users.sql

## sqlc

## Go struct - User + json tag

## json.Encode

## JSON bytes

## fetch /users

## JS object usersData

## sort + render DOM

  \| 30 of 60

# La página users.templ

templ UserTable(users \[\]db.User) { users-table"\> \<table id="

\<thead\>

\<tr\>

col" data-column=" id"\>\<button type=" \<th scope=" col" name"\>\<button type=" data-column=" \<th scope=" col" email"\>\<button type=" data-column=" \<th scope=" \</tr\>

\</thead\>

\<tbody\> for \_, user := range users { \<tr\>

\<td\> { fmt.Sprint(user.ID) }\</td\> \<td\> { user.Name }\</td\> \<td\> { user.Email }\</td\> \</tr\> } if len(users) == 0 { \<tr\> \<td colspan="3"\>No users found.\</td\> \</tr\> } \</tbody\> \</table\>

}

button"\>ID\</button\>\</th\>

button"\>Name\</button\>\</th\>

button"\>Email\</button\>\</th\>

  \| 31 of 60

# La página users.templ

templ UserPage() { \<!DOCTYPE html\>

\<html lang="es"\> \<head\>

\<meta charset=" UTF-8"/\>

\<meta name=" content=" width=device-width, initial-scale=1.0"/\> viewport" \<title\>Lista de usuarios\</title\>

\<style\> body { font-family: sans-serif; } table { width: 100%; border-collapse: collapse; } th, td { padding: 8px 12px; border: 1px solid #ddd; text-align: left; } th.sortable { cursor: pointer; } th.sortable:hover { background-color: #f2f2f2; } th.sorted-asc::after { content:" \</style\> \</head\>

\<body\> \<h1\>Usuarios\</h1\>

@UserTable(\[\]db.User{}) @sortScript() \</body\> \</html\>

}

↑" ↓" ; } th.sorted-desc::after { content:" ; }

  \| 32 of 60

# La página users.templ

templ sortScript() { \<script\> ; let currentSortOrder =''; let usersData = \[\]; let currentSortColumn =''

function renderTable() { const tbody = document.querySelector('#users-table tbody'); tbody.innerHTML ='';

if (usersData.length === 0) { const row = document.createElement('tr'); row.innerHTML ='\<td colspan="3"\>No users found.\</td\>';

tbody.appendChild(row); return; }

const sortedUsers = \[...usersData\].sort((a, b) =\> { if (!currentSortColumn) return 0; const aValue = a\[currentSortColumn\]; const bValue = b\[currentSortColumn\]; if (aValue \< bValue) return currentSortOrder ==='asc'?-1 : 1; if (aValue \> bValue) return currentSortOrder ==='asc'? 1 :-1; return 0; });

  \| 33 of 60

# La página users.templ

Esta parte del script se encarga de renderizar la tabla con los usuarios ordenados:

sortedUsers.forEach(user =\> { const row = document.createElement('tr'); for (const value of \[user.id, user.name, user.email\]) { const cell = document.createElement('td'); cell.textContent = value; row.appendChild(cell); } tbody.appendChild(row); });

// Update headers document.querySelectorAll('#users-table th').forEach(th =\> {'sorted-desc'); th.classList.remove('sorted-asc', if (th.dataset.column === currentSortColumn) { th.classList.add('sorted-' + currentSortOrder); } }); }

  \| 34 of 60

# La página users.templ

document.addEventListener('DOMContentLoaded', () =\> { fetch('/users') .then (response =\> response.json()) .then (data =\> { usersData = data; renderTable(); }) .catch (error =\> console.error('Error fetching users:',

const table = document.getElementById('users-table'); if (!table) return; table.addEventListener('click', (e) =\> { const header = e.target.closest('th\[data-column\]'); if (!header) return; const column = header.dataset.column; if (currentSortColumn === column) { currentSortOrder = currentSortOrder ==='asc'?'desc':'asc';

} else { currentSortColumn = column; currentSortOrder ='asc'; } renderTable(); }); }); \</script\> }

error));

  \| 35 of 60

# Actualización del DOM a partir de JSON

El servidor se simplifica en cuanto a funcionalidad, pero ahora expone datos JSON:

# tipos como las fechas requieren acordar una representación y convertirla en el

cliente.

en este ejemplo el ordenamiento se traslada al cliente, aunque el servidor podría seguir resolviendo filtros y paginación.

Se reduce la cantidad de HTML generado, pero se aumenta la complejidad del JS.

Se realizan menos interacciones con el servidor.

El cliente es responsable de ordenar y renderizar, lo que puede ser más conveniente, pero habrá que tener en cuenta el rendimiento si la cantidad de datos es grande:

# paginación

# carga perezosa (lazy loading)

Movimos la lógica al cliente: el cliente hace lo que antes hacía el servidor.

  \| 36 of 60

# Actualización parcial de HTML con HTMX

HTMX es una biblioteca que permite realizar solicitudes HTTP y actualizar el DOM

# mediante atributos HTML. Las transiciones CSS forman parte de su modelo;

WebSockets y Server-Sent Events se incorporan mediante extensiones.

# Busca completar HTML como un hipertexto, permitiendo que cualquier elemento, no

solo \<a\> y \<form\> , realice solicitudes HTTP.

para la interactividad del frontend.

# Reduce la necesidad de escribir JavaScript

# Filosofía"HTML-first": la lógica de la aplicación reside en el backend, y el frontend

solo se encarga de mostrar la UI y reaccionar a las acciones del usuario.

AJAX

\<table\> Server ...\</table\>

  \| 37 of 60

# Actualización parcial de HTML con HTMX

HTMX es una biblioteca que permite realizar solicitudes HTTP y actualizar el DOM

# mediante atributos HTML. Las transiciones CSS forman parte de su modelo;

WebSockets y Server-Sent Events se incorporan mediante extensiones.

# Busca completar HTML como un hipertexto, permitiendo que cualquier elemento, no

solo \<a\> y \<form\> , realice solicitudes HTTP.

para la interactividad del frontend.

# Reduce la necesidad de escribir JavaScript

# Filosofía"HTML-first": la lógica de la aplicación reside en el backend, y el frontend

solo se encarga de mostrar la UI y reaccionar a las acciones del usuario.

AJAX

\<table\> Server ...\</table\>

  \| 37 of 60

# Instalación de HTMX 4.0

HTMX es un solo archivo JavaScript sin dependencias. No se necesita un sistema de construcción.

# CDN (recomendado para empezar)

\<script src=" https://cdn.jsdelivr.net/npm/htmx.org@4.0.0/dist/htmx.min.js" in tegrity="sha384-BvJpBiO8Kh31EqtJe5DRIeWrHWnCGkwytKs9NKFi86Hhw96dEqdEMzZDeK9iEGTc" crossorigin=" anonymous"\>\</script\>

# Descarga directa

\<script src=" /js/htmx.min.js"\>\</script\>

# NPM

npm install htmx.org@next

  N ote

En producción, se recomienda descargar el archivo y servirlo localmente en lugar de usar un CDN.

  \| 38 of 60

¿Qué es HTMX?

## load

## hx-target

## every

## hx-trigger

## Triggers

## Atributos

## click

## HTMX

## beforebegin

## hx-

## post

## innerHTML

## Partials

## Intercambio

## outerMorph

## outerHTML

## afterend

## innerMorph

## hx-

## get

## hx-swap

## hx-

## partial

  \| 39 of 60

# hx-

# get

# El atributo hx-

# get le dice a un

# elemento que realice una petición GET

# a una URL específica cuando se

# dispara un evento (por defecto, el

evento click ).

# El contenido devuelto por el servidor

# (generalmente un fragmento de HTML)

# reemplazará el contenido del propio

elemento.

\<button hx- /users"\> get=" Cargar Usuarios\</button\>

User Browser Server

Clic en el botón

GET /users

Fragmento HTML

Reemplaza el botón con el HTML

User Browser Server

  N ote

Cuando se hace clic en el botón, se hará una petición GET a /users . El HTML devuelto por el servidor reemplazará al propio botón.

  \| 40 of 60

# hx-target

# El atributo hx-target permite

# especificar qué elemento de la página

# será actualizado con la respuesta del

servidor.

# Puede ser un selector CSS (como un ID

o una clase).

# Esto desacopla el elemento que dispara

el evento del elemento que se actualiza.

\<button hx- /users"hx-target=" #user-list"\> get=" Cargar Usuarios\</button\>

user-list"\> \<div id="

\<!-- Aquí se cargarán los usuarios --\> \</div\>

User Browser Server

Clic en el botón

GET /users

Fragmento HTML

Reemplaza el contenido de

User Browser Server

  N ote

Ahora, el contenido de #user-list será reemplazado, en lugar del botón.

  \| 41 of 60

# hx-boost

# Mejora los enlaces y formularios

# para que usen AJAX en lugar de

recarga de página.

# Cuando se aplica a un elemento

# (generalmente el \<body\> ), los

# enlaces y formularios compatibles

# de sus descendientes se mejoran

para usar AJAX.

# En HTMX 4.0, se usa :inherited

# para hacer explícita la herencia

del comportamiento.

# Esto proporciona una experiencia

# de usuario más fluida, similar a

# una SPA, pero sin la complejidad

de un framework JS.

true"\> \<body hx-boost:inherited=" \<a href=" /about"\>Acerca de\</a\>

\<a href=" /contact"\>Contacto\</a\>

\</body\>

User Browser Server

Clic en el enlace"Acerca de"

GET /about (petición AJAX)

HTML del body de /about

Reemplaza el body actual con el nuevo HTML

User Browser Server

  N ote

Al hacer clic en los enlaces, HTMX solicita la página de destino y reemplaza el contenido del \<body\> sin una recarga completa. Sin JavaScript, los enlaces siguen funcionando como navegación normal.

  \| 42 of 60

# Herencia de atributos: :inherited

# En HTMX 4.0, la herencia implícita está desactivada por defecto

( implicitInheritance: false ).

# Para heredar un atributo concreto, se usa el modificador :inherited :

\<!-- Sin :inherited → en la configuración predeterminada, no heredan hx-confirm --\> \<div hx-confirm=" ¿Estás seguro?"\> /item/1"\>Eliminar\</button\> \<!-- NO muestra confirm --\> \<button hx-delete="

\</div\>

\<!-- Con :inherited → los hijos sí heredan hx-confirm --\> \<div hx-confirm:inherited=" ¿Estás seguro?"\> /item/1"\>Eliminar\</button\> \<!-- SÍ muestra confirm --\> \<button hx-delete="

\</div\>

  N ote

Esto hace explícita la herencia en cada atributo y mejora la previsibilidad.

  \| 43 of 60

# hx-trigger

# El atributo hx-trigger especifica qué

evento causa la petición.

# Por defecto, es click para botones y

enlaces, y submit para formularios.

# Se puede modificar para reaccionar a

# cualquier evento, como keyup para

búsquedas en tiempo real.

  \| 44 of 60

# hx-trigger

# El atributo hx-trigger especifica qué

evento causa la petición.

# Por defecto, es click para botones y

enlaces, y submit para formularios.

# Se puede modificar para reaccionar a

# cualquier evento, como keyup para

búsquedas en tiempo real.

# Modificadores:

# changed : dispara si el valor del

elemento ha cambiado.

# delay:\<tiempo\> : espera un tiempo

# antes de la petición (ej:

delay:500ms ).

## throttle:\<tiempo\> : ignora eventos

# durante un tiempo después de una

petición.

once : dispara la petición una vez.

# Para sondeos (polling), se puede usar

every \<tiempo\> (ej: every 2s ).

  \| 44 of 60

# hx-trigger

text"n ame="q"hx-get=" \<input type=" hx-target=" #search-results"

search-results"\>\</div\> \<div id="

  N ote

La petición se envía 500ms después de que el usuario deja de escribir, y solo si el texto cambió.

/search"hx-trigger=" keyup changed delay:500ms" /\> Buscando mientras escribes..." placeholder="

User Browser Server

Escribe en el input

Espera 500ms

GET /search?q=...

HTML con resultados

Actualiza

User Browser Server

  \| 45 of 60

# hx-indicator

# El atributo hx-indicator permite mostrar un indicador de carga mientras una

petición AJAX está en curso.

Se puede apuntar a un elemento (usando un selector CSS) que se hará visible durante la petición.

# Por defecto, HTMX busca un elemento con la clase htmx-indicator dentro del

elemento que hace la petición.

Esto mejora la experiencia de usuario al proporcionar feedback visual.

  \| 46 of 60

# hx-indicator

\<button hx- /users" get=" #user-list"\> hx-target=" Cargar Usuarios htmx-indicator" \<img class=" src=" ./s 20" dth=" pinner.gif"wi \</button\>

user-list"\>\</div\> \<div id="

  N ote

Para que funcione, el indicador (la imagen del spinner en este caso) debe estar oculto por CSS y mostrarse cuando HTMX añade la clase htmx-request al elemento padre.

  Im portant

El indicador se muestra u oculta manipulando su opacidad por defecto, pero se puede personalizar con CSS.

User Browser Server

Clic en /\>"Cargar Usuarios"

Muestra el indicador

GET /users

HTML de respuesta

Oculta el indicador

Actualiza

User Browser Server

  \| 47 of 60

# \-

# oob :"

# Out of Band"

# hx-swap

OOB swap permite actualizar múltiples elementos en la página con una sola petición.

# El servidor puede devolver fragmentos de HTML extra junto con la respuesta

principal.

.

## true"

# Cada fragmento OOB debe tener un id y el atributo hx-swap-oob="

HTMX buscará elementos en la página con los mismos id y los actualizará.

  \| 48 of 60

# \-

# oob

# hx-swap

# HTML del Cliente:

\<button hx- /add-to-cart" post=" #status"\> hx-target=" Añ adir al Carrito\</button\>

status"\>\</div\> \<div id="

\<div\>

N otificaciones:

notifications"\>\</div\> \<div id="

\</div\>

\<p\>Carrito: cart-count"\>4\</span\> items \<span id=" \</p\>

  N ote

Al hacer clic, se actualiza #status con la respuesta principal, #notifications con el div de notificación y #cart-count con el nuevo número.

# Respuesta del Servidor:

\<!-- Respuesta principal para el hx-target --\> \<p\>Acción completada.\</p\>

\<!-- Fragmentos OOB --\> notifications"hx-swap-oob=" true"\> \<div id="

\<p\>¡Tienes una nueva notificación!\</p\> \</div\>

cart-count"hx-swap-oob=" true"\>5\</span\> \<span id="

  \| 49 of 60

# hx-

# partial

# En HTMX 4.0, los \<hx-

# partial\> son la alternativa más clara a OOB swaps para

actualizaciones múltiples.

# Cada \<hx-

# partial\> puede especificar su propio hx-target y hx-swap ; es una

# alternativa declarativa a algunos usos de OOB:

\<!-- Respuesta del servidor --\> \<hx- #notifications"hx-swap="beforeend"\> partial hx-target=" \<div\>Nueva notificación\</div\>

\</hx- partial\> \<hx- #cart-count"\> partial hx-target=" \<span\>5\</span\> \</hx- partial\> \<form id=" my-form"\> \<!-- Contenido principal del formulario --\> \</form\>

  N ote

Los \<hx- partial\> son más expresivos que hx-swap-oob y permiten controlar explícitamente el destino y la estrategia de swap.

  \| 50 of 60

# hx-on

# El atributo hx-on permite ejecutar JavaScript en respuesta a eventos estándar o

eventos de HTMX.

Es útil para pequeñas tareas de scripting sin necesidad de un bloque \<script\> separado.

.

## La sintaxis es hx-on:\<evento\>="\<script\>"

# HTMX dispara muchos eventos útiles, como:

htmx:confirm : para personalizar la confirmación.

htmx:before:request : antes de enviar la petición.

htmx:after:swap : después de que el nuevo contenido se ha insertado en el DOM.

htmx:response:error : si hay un error en la respuesta.

  \| 51 of 60

# hx-on

\<!-- Limpiar form después de envío ok --\> \<form

hx-post=" /register" hx-target=" #response" hx-on:htmx:after:swap="this.reset()"

\>

text"n ame=" name" /\> \<input type=" submit"\>Enviar\</button\> \<button type=" \</form\>

\<div id=" response"\>\</div\>

  N ote

Después de que la respuesta se intercambia correctamente, el evento htmx:after:swap se dispara y el código en hx-on resetea el formulario. User Browser Server

User Browser Server

Rellena y envía el formulario

POST /register

Respuesta HTML

Actualiza

Dispara htmx:after:swap

Ejecuta this.reset()

  \| 52 of 60

# hx-on: scripting async

# HTMX 4.0 soporta async/await en hx-on , permitiendo scripts asíncronos:

\<button hx- /like" post=" hx-on:htmx:after:swap="await timeout('3s'); this.remove()"\>

Lik e, esperar 3 segundos y eliminar \</button\>

## await timeout('3s') : espera 3 segundos antes de continuar.

this.remove() : elimina el botón después del intercambio.

  N ote

HTMX 4.0 proporciona helpers async como timeout() que hacen más fácil el scripting asíncrono inline.

  \| 53 of 60

# Actualización parcial de HTML con HTMX

# A diferencia del ejemplo de JSON/JavaScript, el servidor ahora debe responder con

fragmentos HTML.

El código del servidor no cambia mucho, pero ahora el handler ListUsers debe devolver un fragmento de HTML en lugar de una página completa.

La lógica de ordenamiento y renderizado de la tabla se mantiene en el servidor, pero ahora se usa HTMX para actualizar la tabla sin recargar la página.

# Se reduce mucho el código JavaScript necesario para manejar la interactividad, ya

que HTMX se encarga de gran parte de la lógica.

No es necesario tener un endpoint que exponga los datos en formato JSON, ya que HTMX puede manejar la actualización de HTML directamente.

  \| 54 of 60

# El handler devuelve un documento completo o

# un fragmento HTML

func (h\*UserHandler) ListUsers(w http.ResponseWriter, r\*http.Request) { ctx := r.Context() sortColumn := r.URL.Query().Get("sort") sortOrder := r.URL.Query().Get("order")

...

var users \[\]db.User

switch sortColumn { case"id":

if sortOrder =="asc" { users, err = h.queries.ListUsersOrderByIdAsc(ctx)

...

// Check for HTMX request header if r.Header.Get("HX-Request") == "true" { // If it's an HTMX request, only render the table

views.UserTable(users, sortColumn, sortOrder).Render(r.Context(), w) return

}

// For initial page load, render the full page views.UserPage(users, sortColumn, sortOrder).Render(r.Context(), w)

  \| 55 of 60

# El handler devuelve un documento completo o

# un fragmento HTML

func (h\*UserHandler) ListUsers(w http.ResponseWriter, r\*http.Request) { ctx := r.Context() sortColumn := r.URL.Query().Get("sort") sortOrder := r.URL.Query().Get("order")

...

var users \[\]db.User

switch sortColumn { case"id":

if sortOrder =="asc" { users, err = h.queries.ListUsersOrderByIdAsc(ctx)

...

// Check for HTMX request header if r.Header.Get("HX-Request") == "true" { // If it's an HTMX request, only render the table

views.UserTable(users, sortColumn, sortOrder).Render(r.Context(), w) return

}

// For initial page load, render the full page views.UserPage(users, sortColumn, sortOrder).Render(r.Context(), w)

  \| 55 of 60

# El handler devuelve un documento completo o

# un fragmento HTML

func (h\*UserHandler) ListUsers(w http.ResponseWriter, r\*http.Request) { ctx := r.Context() sortColumn := r.URL.Query().Get("sort") sortOrder := r.URL.Query().Get("order")

...

var users \[\]db.User

switch sortColumn { case"id":

if sortOrder =="asc" { users, err = h.queries.ListUsersOrderByIdAsc(ctx)

...

// Check for HTMX request header if r.Header.Get("HX-Request") == "true" { // If it's an HTMX request, only render the table

views.UserTable(users, sortColumn, sortOrder).Render(r.Context(), w) return

}

// For initial page load, render the full page views.UserPage(users, sortColumn, sortOrder).Render(r.Context(), w)

  \| 55 of 60

# La página users.templ

templ UserPage(users \[\]db.User, sortColumn string, sortOrder string) { \<!DOCTYPE html\>

...

\<script src=" https://cdn.jsdelivr.net/npm/htmx.org@4.0.0/dist/htmx.min.js" in tegrity="sha384-BvJpBiO8Kh31EqtJe5DRIeWrHWnCGkwytKs9NKFi86Hhw96dEqdEMzZDeK9iEGTc" crossorigin=" anonymous"\>\</script\> \<style\> body { font-family: sans-serif; } table { width: 100%; border-collapse: collapse; } th, td { padding: 8px 12px; border: 1px solid #ddd; text-align: left; } th.sortable { cursor: pointer; } th.sortable:hover { background-color: #f2f2f2; } ↑" th.sorted-asc::after { content:" ; } ↓" th.sorted-desc::after { content:" ; } \</style\> \</head\>

\<body\> \<h1\>Usuarios\</h1\>

@UserTable(users, sortColumn, sortOrder) \</body\> \</html\>

}

  \| 56 of 60

# La página users.templ

templ UserPage(users \[\]db.User, sortColumn string, sortOrder string) { \<!DOCTYPE html\>

...

\<script src=" https://cdn.jsdelivr.net/npm/htmx.org@4.0.0/dist/htmx.min.js" in tegrity="sha384-BvJpBiO8Kh31EqtJe5DRIeWrHWnCGkwytKs9NKFi86Hhw96dEqdEMzZDeK9iEGTc" crossorigin=" anonymous"\>\</script\> \<style\> body { font-family: sans-serif; } table { width: 100%; border-collapse: collapse; } th, td { padding: 8px 12px; border: 1px solid #ddd; text-align: left; } th.sortable { cursor: pointer; } th.sortable:hover { background-color: #f2f2f2; } ↑" th.sorted-asc::after { content:" ; } ↓" th.sorted-desc::after { content:" ; } \</style\> \</head\>

\<body\> \<h1\>Usuarios\</h1\>

@UserTable(users, sortColumn, sortOrder) \</body\> \</html\>

}

  \| 56 of 60

# La página users.templ

templ UserPage(users \[\]db.User, sortColumn string, sortOrder string) { \<!DOCTYPE html\>

...

\<script src=" https://cdn.jsdelivr.net/npm/htmx.org@4.0.0/dist/htmx.min.js" in tegrity="sha384-BvJpBiO8Kh31EqtJe5DRIeWrHWnCGkwytKs9NKFi86Hhw96dEqdEMzZDeK9iEGTc" crossorigin=" anonymous"\>\</script\> \<style\> body { font-family: sans-serif; } table { width: 100%; border-collapse: collapse; } th, td { padding: 8px 12px; border: 1px solid #ddd; text-align: left; } th.sortable { cursor: pointer; } th.sortable:hover { background-color: #f2f2f2; } ↑" th.sorted-asc::after { content:" ; } ↓" th.sorted-desc::after { content:" ; } \</style\> \</head\>

\<body\> \<h1\>Usuarios\</h1\>

@UserTable(users, sortColumn, sortOrder) \</body\> \</html\>

}

  \| 56 of 60

# La página users.templ

templ UserTable(users \[\]db.User, sortColumn, sortOrder string) { users-table"\> \<table id="

\<thead\>

\<tr\>

"ID", @SortableHeader("id", sortColumn, sortOrder)"Name", @SortableHeader("name", sortColumn, sortOrder)"Email", @SortableHeader("email", sortColumn, sortOrder) \</tr\>

\</thead\>

\<tbody\> for \_, user := range users { \<tr\> \<td\> { fmt.Sprint(user.ID) }\</td\> \<td\> { user.Name }\</td\> \<td\> { user.Email }\</td\> \</tr\> } if len(users) == 0 { \<tr\> \<td colspan="3"\>No users found.\</td\> \</tr\> } \</tbody\> \</table\>

}

  \| 57 of 60

# La página users.templ

templ UserTable(users \[\]db.User, sortColumn, sortOrder string) { users-table"\> \<table id="

\<thead\>

\<tr\>

"ID", @SortableHeader("id", sortColumn, sortOrder)"Name", @SortableHeader("name", sortColumn, sortOrder)"Email", @SortableHeader("email", sortColumn, sortOrder) \</tr\>

\</thead\>

\<tbody\> for \_, user := range users { \<tr\> \<td\> { fmt.Sprint(user.ID) }\</td\> \<td\> { user.Name }\</td\> \<td\> { user.Email }\</td\> \</tr\> } if len(users) == 0 { \<tr\> \<td colspan="3"\>No users found.\</td\> \</tr\> } \</tbody\> \</table\>

}

  \| 57 of 60

# La página users.templ

templ SortableHeader(columnName, displayName, currentSortColumn, currentSortOrder string) { col" \<th scope=" class= { map\[string\] bool{"sortable": true,

"sorted-" + currentSortOrder: currentSortColumn == columnName, } } hx- columnName, get={ fmt.Sprintf("/?sort=%s&order=%s", getNextSortOrder(columnName, currentSortColumn, currentSortOrder)) } #users-table" hx-target=" hx-swap="outerHTML" hx- true" push-url=" \>

button"\>{ displayName }\</button\> \<button type=" \</th\>

}

  \| 58 of 60

# La página users.templ

templ SortableHeader(columnName, displayName, currentSortColumn, currentSortOrder string) { col" \<th scope=" class= { map\[string\] bool{"sortable": true,

"sorted-" + currentSortOrder: currentSortColumn == columnName, } } hx- columnName, get={ fmt.Sprintf("/?sort=%s&order=%s", getNextSortOrder(columnName, currentSortColumn, currentSortOrder)) } #users-table" hx-target=" hx-swap="outerHTML" hx- true" push-url="

# hace hx-

# get a la URL con los parámetros

\>

# sort y order , actualiza el contenido del

button"\>{ displayName }\</button\> \<button type="

# elemento con ID users-table ,

\</th\>

}

# reemplazando el HTML completo de la

tabla, y actualiza la URL del navegador.

  \| 58 of 60

# Conclusiones

# \+interactividad +complejidad en el cliente y en la coordinación con el servidor:

# SSR con navegación completa: HTML generado en el servidor y documento

completo en cada navegación.

Actualización parcial de HTML con Vanilla JS o HTMX: HTML generado en el servidor y reemplazo de una parte del DOM.

JSON + JavaScript: datos entregados por el servidor; el cliente los procesa y genera el DOM/HTML.

# HTMX: simplifica la interactividad del frontend, permitiendo que el servidor maneje

la lógica de la aplicación y el cliente solo se encargue de mostrar la UI.

Soluciones con renderizado y estado en el cliente, como React, Vue.js, Angular y Svelte, también son opciones válidas.

  Caution Peden ser innecesarias si la aplicación se resuelve adecuadamente con HTML, JS puntual o HTMX.

  \| 59 of 60

# Reflexiones

# En los ejemplos, el estado está en:

# SSR y HTMX: URL ( sort , order )

# \+ HTML del server

Guardo marcador → guardo estado

## Bookmarkable

## URL

## HTML

## compartible

## ?sort=&order

# AJAX JSON: variables JS en cliente

Requiere programar más

## Efímero

## JS vars

## se pierde al refrescar

## usersData

  \| 60 of 60

# Reflexiones

# En los ejemplos, el estado está en:

# SSR y HTMX: URL ( sort , order )

# \+ HTML del server

Guardo marcador → guardo estado

## Bookmarkable

## URL

## HTML

## compartible

## ?sort=&order

# AJAX JSON: variables JS en cliente

Requiere programar más

## Efímero

## JS vars

## se pierde al refrescar

## usersData

# Datos sufren transformaciones:

# SSR / HTMX: server ordena → HTML

SQL → Go → HTML (3 pasos)

## Go struct HTML

## SQL

# AJAX JSON: server JSON → cliente ordena

# → HTML

Múltiples transformaciones

## Go JSON JS sort HTML

## SQL

  \| 60 of 60