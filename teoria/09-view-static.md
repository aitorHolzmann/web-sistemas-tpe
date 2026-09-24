# Pr

# ogramación Web

# Server-Side Rendering (SSR)

Dr. Alejandro Zunino & Dr. Alfredo Teyseyre alejandro.zunino@isistan.unicen.edu.ar, alfredo.teyseyre@isistan.unicen.edu.ar

  \| 1 of 34

<image redacted: 72x72px, 72x72pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

# Contenidos

1\. Server-Side Rendering (SSR) Tradicional

2\. Handler en Go con html/template

3\. Limitaciones de

html/template / text/template

4\. templ : Una Solución Moderna y Segura

5\. Sirviendo archivos estáticos

6\. Proyecto con templ desde cero: estructura

7\. Proyecto con templ desde cero: vista

8\. Conclusiones SSR Tradicional

  \| 2 of 34

<image redacted: 718x479px, 718x479pt, ~72dpi, JPG, DEVICE_RGB, 32bpp>

# Server-Side Rendering (SSR) Tradicional

En SSR tradicional, el servidor genera el HTML completo para cada solicitud del usuario:

El servidor procesa la solicitud, obtiene los datos necesarios (ej. base de datos) y renderiza el HTML.

Se envía un documento HTML listo para ser mostrado, que puede referenciar hojas de estilo, scripts e imágenes.

¿Cómo se genera el HTML?

1\. Imprimiendo HTML directamente en el código del servidor.

2\. A partir de plantillas (templates) que combinan HTML con datos dinámicos:

El servidor utiliza un motor de plantillas para combinar el HTML con los datos dinámicos.

Dependiendo del motor, se pueden usar sintaxis diferentes (ej. EJS, Handlebars, Jinja2, etc.) con condicionales y bucles.

  \| 3 of 34

# SSR tradicional vs SSR interactivo

Esta clase se centra en el renderizado en el servidor (SSR): el backend genera el HTML de la respuesta, normalmente a partir de datos obtenidos para esa solicitud.

En esta clase hablamos de SSR tradicional: la navegación o el envío de un formulario reemplaza el documento completo.

En la siguiente clase ( 10-view-dynamic.md ) veremos SSR interactivo con HTMX/JS, donde el cliente puede actualizar partes de la página sin recargas completas.

Ambas se apoyan en el modelo de 3 capas (presentación, lógica de negocio en 06- logic.md , acceso a datos), separando responsabilidades.

  \| 4 of 34

# SSR con html/template y text/template

Las bibliotecas estándar de Go para plantillas.

text/template : Para cualquier tipo de texto (email, configuración, etc.).

html/template : Especializado para HTML, ofrece escaping contextual automático para reducir el riesgo de XSS.

basado en cadenas Ambas usan un sistema de plantillas .

La base de Hugo: https://gohugo.io/ (The world’s fastest framework for building websites)

  \| 5 of 34

# SSR con html/template y text/template

Las bibliotecas estándar de Go para plantillas.

text/template : Para cualquier tipo de texto (email, configuración, etc.).

html/template : Especializado para HTML, ofrece escaping contextual automático para reducir el riesgo de XSS.

basado en cadenas Ambas usan un sistema de plantillas .

La base de Hugo: https://gohugo.io/ (The world’s fastest framework for building websites)

  \| 5 of 34

# Plantilla HTML con html/template

\<!DOCTYPE html\>

\<html lang="es"\> \<head\>

\<meta charset=" UTF-8"

\<meta name=" viewport" \<title\> {{.Title}}\</title\> \</head\>

\<body\> \<h1\> {{.Title}}\</h1\> \<p\>¡Bienvenidos a mi página web!\</p\> \</body\> \</html\>

/\>

content=" /\> width=device-width, initial-scale=1.0"

  \| 6 of 34

# Handler en Go con html/template

package main

// Ejemplo de renderizado de una plantilla HTML con Go. import ("html/template""net/http" )

func handler(w http.ResponseWriter, r\*http.Request) { data := map\[string\]string{"Title":"Hola, Mundo!"} // En producción, cargar una vez al iniciar. tmpl := template.Must(template.ParseFiles("index.html")) if err := tmpl.Execute(w, data); err != nil { h ttp.Error(w,"internal server error",h ttp.StatusInternalServerError) } }

  \| 7 of 34

# text/template : El Riesgo de la Inseguridad

Sintaxis simple para interpolación y lógica básica.

No tiene conocimiento del contexto de salida (HTML, JSON, etc.).

Principal peligro: No aplica escaping automáticamente a la salida.

package main import ("os"

"text/template" ) func main() { // ¡Usando text/template para demostrar el riesgo! tmpl, \_ := template.New("name").Parse("\<h1\>Hola, {{.}}!\</h1\>") // Entrada de usuario maliciosa

data :=\`\<script\>alert('XSS Atacado!');\</script\>\`

// text/template no sanitiza la salida, por lo que el script se ejecuta. tmpl.Execute(os.Stdout, data) }

  W arning Un atacante capaz de inyectar código malicioso podría ejecutar scripts en el navegador del usuario.

  \| 8 of 34

# html/template : Mejor Seguridad

consciente de HTML Basado en text/template pero .

Escaping contextual automático: Escapa según el contexto para reducir el riesgo de XSS (Cross-Site Scripting).

Escapa según el contexto del dato: texto HTML, atributos, URLs, CSS y JavaScript, no con un único esquema genérico.

El escaping convierte una entrada maliciosa en texto inofensivo cuando se inserta en la plantilla.

Esto no reemplaza la validación de datos ni vuelve seguro el uso incorrecto de tipos confiables como template.HTML .

  \| 9 of 34

# html/template : Mejor Seguridad

consciente de HTML Basado en text/template pero .

Escaping contextual automático: Escapa según el contexto para reducir el riesgo de XSS (Cross-Site Scripting).

Escapa según el contexto del dato: texto HTML, atributos, URLs, CSS y JavaScript, no con un único esquema genérico.

El escaping convierte una entrada maliciosa en texto inofensivo cuando se inserta en la plantilla.

Esto no reemplaza la validación de datos ni vuelve seguro el uso incorrecto de tipos confiables como template.HTML .

  \| 9 of 34

# html/template : Mejor Seguridad

package main

import ("html/template""os"

) func main() { tmpl, \_ := template.New("name").Parse("\<h1\>Hola, {{.}}!\</h1\>") data :=\`\<script\>alert('XSS Atacado!');\</script\>\`

tmpl.Execute(os.Stdout, data) //\`html/template\` }

Saludaría al usuario con:

\<h1\>Hola, &lt;script&gt;alert(&#39;XSS Atacado!&#39;);&lt;/script&gt;!\</h1\>

// Entrada de usuario maliciosa

sanitiza automáticamente

  \| 10 of 34

# Limitaciones de

# html/template / text/template

A pesar de la mejora de html/template, ambos tienen:

Type-unsafe (No seguro en tipos): Errores de plantilla se descubren en tiempo de ejecución.

Falta de verificación: No hay validación de estructura o datos en tiempo de compilación.

Mantenimiento: Plantillas complejas se vuelven difíciles de mantener.

No son Go nativo: Mezcla de HTML/CSS/JS con sintaxis de plantillas.

  N ote Existen alternativas más modernas y seguras como Templ (similar a Tsx), que ofrecen una mejor experiencia de desarrollo y seguridad.

  \| 11 of 34

# templ : Una Solución Moderna y Segura

templ es un lenguaje de plantillas para Go que genera código Go nativo:

Seguridad de tipos: Los parámetros y expresiones se verifican durante la compilación.

IDE-friendly: Buen soporte de IDE para refactorización y autocompletado.

Código Go: Las plantillas son archivos .templ que generan archivos .go .

Composición: Componentes reutilizables y composables.

Seguridad: Hereda la seguridad de tipos de Go.

package components templ Hello(name string) { \<div\>

\<h1\>{"Hola," + name +"!"}\</h1\>

\<p\>{"Este es un componente templ."}\</p\> \</div\>

}

  \| 12 of 34

# templ : Una Solución Moderna y Segura

package main import (""log"net/http""myproject/components" // hello.templ está en components )

func main() { http.HandleFunc("/",f unc(w http.ResponseWriter, r\*http.Request) { name := r.URL.Query().Get("name")"" if name == { name ="Mundo"

} if err := components.Hello(name).Render(r.Context(), w); err != nil { h ttp.Error(w,"internal server error",h } })

log.Println("Servidor iniciado en :8080") log.Fatal(http.ListenAndServe(":8080",nil }

ttp.StatusInternalServerError)

))

  \| 13 of 34

# templ : Una Solución Moderna y Segura

package main import (""log"net/http""myproject/components" // hello.templ está en components )

func main() { http.HandleFunc("/",f unc(w http.ResponseWriter, r\*http.Request) { name := r.URL.Query().Get("name")"" if name == { name ="Mundo"

} if err := components.Hello(name).Render(r.Context(), w); err != nil { h ttp.Error(w,"internal server error",h } })

log.Println("Servidor iniciado en :8080") log.Fatal(http.ListenAndServe(":8080",nil }

Si name no fuera un string , el compilador de Go fallaría.

ttp.StatusInternalServerError)

))

  \| 13 of 34

# templ : Una Solución Moderna y Segura

V entajas clave:

Verificación en compilación: los parámetros y los tipos declarados en el templ se verifican durante la compilación.

Sintaxis parecida a Go: Más familiar para desarrolladores Go.

Manejo de contexto: Soporte para context.Context .

Renderizado en streaming: Puede mejorar el tiempo hasta el primer contenido en respuestas adecuadas.

D esventajas:

Los archivos .templ deben compilarse a .go, lo que puede ser un paso adicional en el flujo de trabajo.

El código que se ejecuta es distinto al código del templ , lo que puede ser confuso.

  \| 14 of 34

# Funcionamiento de templ

Archivo .templ

Compilador Templ

Código Go generado

Datos Go Binario de tu app

Runtime

Renderiza HTML

  \| 15 of 34

# Funcionamiento de templ

Navegador Servidor Go Templ Handler Modelos

GET /users (HTTP Request)

Llama al Handler

Obtiene datos de DB

Devuelve structs de datos

Genera template con datos

Se ejecuta el template compilado a Go

HTML renderizado

Respuesta HTTP

HTML (HTTP Response)

Navegador Servidor Go Templ Handler Modelos

  \| 16 of 34

# templ y alternativas Go SSR

| Característica                   | text/template                                   | html/template templ                                     |
|----------------------------------|-------------------------------------------------|---------------------------------------------------------|
| Escaping HTML Seguridad de tipos |                                                 | (interpolaciones) (compilación)                         |
| Sintaxis                         | String                                          |                                                         |
| Composición Mantenimiento        | Básica Básica Depende del Depende tamaño tamaño | (componentes) del Componentes tipados Ejecuta código Go |
| Ejecución                        | Interpreta plantillas                           | Interpreta plantillas generado   \| 17 of 34            |

# Sirviendo archivos estáticos

En SSR el servidor genera el HTML, pero el CSS, las imágenes y demás recursos se entregan como archivos estáticos. Go incluye http.FileServer para esto:

Coloca los assets en un directorio (ej. static/ ) y regístralos con un handler dedicado.

Usa http.StripPrefix para que la ruta /static/ no se busque literalmente dentro de la carpeta.

fs := http.FileServer(http.Dir("static")) http.Handle("/static/",h ttp.StripPrefix("/static/",f s))

Luego se referencian normalmente desde la plantilla:

\<link rel=" ef=" /\> stylesheet"hr /static/styles.css"

Tanto en templ como en html/template los enlaces a assets son HTML común; la diferencia está en cómo se genera el documento, no en cómo se enlazan los recursos.

  \| 18 of 34

# Proyecto con templ desde cero: estructura

Primero crea un proyecto Go:

mkdir templ-example; cd templ-example; go mod init templ-example

Instala templ y sqlc (BD):

go install github.com/a-h/templ/cmd/templ@latest go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

**/**

**templ-example**

main.go

go.mod

go.sum

**db**

**queries**

Estructura del proyecto

**schema**

**sqlc**

**models**

**handlers**

user.go

**views**

user.templ

  \| 19 of 34

# Proyecto con templ desde cero: vista

En views/user.templ :

package views

import"templ-example/models"

templ UserList(users \[\]models.User) { \<ul\>

f or \_, user := range users { \<li\> { user.Name } - { user.Email }\</li\> } \</ul\>

}

templ UserPage(users \[\]models.User) { \<html\>

\<body\> \<h1\>Lista de Usuarios\</h1\>

@UserList(users) \</body\> \</html\>

}

  \| 20 of 34

# Modelo de datos en memoria

En models/user.go :

package models type User struct {

| ID N   | in     |
|--------|--------|
| ame Em | string |
| ail }  | string |

var Users = \[\]User{ ail:"alice@example.com"}, {ID: 1, Name: "Alice",Em ail:"bob@example.com"}, {ID: 2, Name: "Bob",Em

  \| 21 of 34

# Vista con datos en memoria

En views/user.templ :

package views

import"templ-example/models"

templ UserList(users \[\]models.User) { \<ul\>

f or \_, user := range users { \<li\> { user.Name } - { user.Email }\</li\> } \</ul\>

}

templ UserPage(users \[\]models.User) { \<html\>

\<body\> \<h1\>Lista de Usuarios\</h1\>

@UserList(users) \</body\> \</html\>

}

  \| 22 of 34

# Handler de usuarios

El handler que atiende la solicitud en handlers/user.go :

package handlers

import ("net/http"""templ-example/models""templ-example/views )

func UserHandler(w http.ResponseWriter, r\*http.Request) { if err := views.UserPage(models.Users).Render(r.Context(), w); err != nil { h ttp.Error(w,"internal server error",h ttp.StatusInternalServerError) } }

  \| 23 of 34

# Arranque de la aplicación

En main.go :

package main import (""log"net/http"""templ-example/handlers ) func main() { h ttp.HandleFunc("/users",h andlers.UserHandler) l og.Fatal(http.ListenAndServe(":8080",nil )) }

Finalmente, ejecuta el servidor:

templ generate go run .

  \| 24 of 34

# Arranque de la aplicación

En main.go :

package main import (""log"net/http"""templ-example/handlers ) func main() { h ttp.HandleFunc("/users",h andlers.UserHandler) l og.Fatal(http.ListenAndServe(":8080",nil )) }

Finalmente, ejecuta el servidor:

Ver el código generado en templ generate views/user.templ.go go run .

  \| 24 of 34

# Accediendo a una Base de Datos con sqlc

Este es un proyecto independiente ( templ-sqlc-example ) que retoma el ejemplo anterior agregándole acceso a datos.

Definir esquema en db/schema/ :

CREATE TABLE users ( i d SERIAL PRIMARY KEY,

n ame TEXT NOT NULL, email TEXT UNIQUE NOT NULL );

Crear queries en db/queries :

\-- n ame: ListUsers :many SELECT\*FR OM users;

Configuración de sqlc en sqlc.yaml :

version:"2"

sql:

\- engine:"postgresql" queries:"./db/ queries/" schema:"./db/schema/"

gen:

go: package:"db" out:"./db"

emit\_json\_tags: true

Generar código Go:

sqlc generate

  \| 25 of 34

# Accediendo a una Base de Datos con sqlc

Este es un proyecto independiente ( templ-sqlc-example ) que retoma el ejemplo anterior agregándole acceso a datos.

Definir esquema en db/schema/ :

CREATE TABLE users ( i d SERIAL PRIMARY KEY,

n ame TEXT NOT NULL, email TEXT UNIQUE NOT NULL );

Crear queries en db/queries :

\-- n ame: ListUsers :many SELECT\*FR OM users;

Ver el código generado en ./db/

Configuración de sqlc en sqlc.yaml :

version:"2"

sql:

\- engine:"postgresql" queries:"./db/ queries/" schema:"./db/schema/"

gen:

go: package:"db" out:"./db"

emit\_json\_tags: true

Generar código Go:

sqlc generate

  \| 25 of 34

# Accediendo a una Base de Datos con sqlc

Cliente HTTP

Handler

SQLC Queries Templ

PostgreSQL

  \| 26 of 34

# Accediendo a una Base de Datos con sqlc

package handlers

import ("net/http""templ-sqlc-example/db""templ-sqlc-example/views" )

type UserHandler struct { db.Queries queries \* }

func NewUserHandler(q\*db.Queries) \* return &UserHandler{queries: q} }

UserHandler {

  \| 27 of 34

# Accediendo a una Base de Datos con sqlc

func (h\*UserHandler) ServeHTTP(w http.ResponseWriter, r\*http.Request) { switch r.URL.Path { case"/":

h.ListUsers(w, r) default:

http.NotFound(w, r) } }

func (h\*UserHandler) ListUsers(w http.ResponseWriter, r\*http.Request) { users, err := h.queries.ListUsers(r.Context()) if err != nil { http.Error(w,"internal server error",h return

} if err := views.UserList(users).Render(r.Context(), w); err != nil { h ttp.Error(w,"internal server error",h } }

ttp.StatusInternalServerError)

ttp.StatusInternalServerError)

  \| 28 of 34 package views import"templ-sqlc-example/db"

templ UserList(users \[\]db.User) { \<!DOCTYPE html\>

\<html lang="es"\> \<head\>

\<meta charset=" UTF-8"/\>

\<meta name=" content=" viewport" \<title\>Lista de usuarios\</title\>

\</head\>

\<body\> \<h1\>Usuarios\</h1\>

\<table\>

\<thead\> \<tr\> \<th\>ID\</th\> \<th\>Name\</th\> \<th\>Email\</th\> \</tr\> \</thead\>

\<tbody\> for \_, user := range users { \<tr\>

\<td\> { user.ID }\</td\> \<td\>{ user.Name }\</td\> \<td\>{ user.Email }\</td\> \</tr\>

} \</tbody\> \</table\>

\</body\> \</html\>

}

width=device-width, initial-scale=1.0"/\>

  \| 29 of 34

# Accediendo a una Base de Datos con sqlc

package db import ("database/sql""fmt"

"os"

\_"github.com/lib/pq"

)

func ConnectDB() (\*sql.DB, error) { dsn := os.Getenv("DATABASE\_URL")"" if dsn == { return nil, fmt.Errorf("DATABASE\_URL is not set") } conn, err := sql.Open("postgres", dsn) if err != nil { return nil, err } if err := conn.Ping(); err != nil { conn.Close() return nil, err } return conn, nil }

  \| 30 of 34

# Accediendo a una Base de Datos con sqlc

package main

import (""log"net/http""templ-sqlc-example/db"""templ-sqlc-example/handlers ) func main() { conn, err := db.ConnectDB() if err != nil { log.Fatal(err) } defer conn.Close()

queries := db.New(conn) userHandler := handlers.NewUserHandler(queries)

http.Handle("/", userHandler) log.Println("Server running on :8080") log.Fatal(http.ListenAndServe(":8080",nil )) }

  \| 31 of 34

**main**

\+main()

uses (via db.New)

uses

**db.Queries**

\+CreateUser(ctx context.Context, arg CreateUserParams)(User, error) +DeleteUser(ctx context.Context, id int32) : error +GetUser(ctx context.Context, id int32)(User, error) +ListUsers(ctx context.Context)(\[\]User, error)

returns/expects

**db.User**

**db.ConnectDB**

\+ID int32

\+Name string +Email string +ConnectDB()(\*sql.DB, error)

returns

**sql.DB**

uses

**handlers.UserHandler**

\+NewUserHandler(queries\*db.Queries)

depends on renders

**views.UserList**

\+UserList(users \[\]db.User) : templ.Component

uses interface

expects

**db.DBTX**

implements

  \| 32 of 34

**main**

\+main()

Aparecen dependencias indeseables?: uses (via db.New) los handlers dependen directamente de sqlc uses

**db.Queries**

\+CreateUser(ctx context.Context, arg CreateUserParams)(User, error) +DeleteUser(ctx context.Context, id int32) : error +GetUser(ctx context.Context, id int32)(User, error) +ListUsers(ctx context.Context)(\[\]User, error)

returns/expects

**db.User**

**db.ConnectDB**

\+ID int32

\+Name string +Email string +ConnectDB()(\*sql.DB, error)

returns

**sql.DB**

uses

**handlers.UserHandler**

\+NewUserHandler(queries\*db.Queries)

depends on renders

**views.UserList**

\+UserList(users \[\]db.User) : templ.Component

uses interface

expects

**db.DBTX**

implements

  \| 32 of 34

# Conclusiones SSR Tradicional

La presentación se renderiza en el servidor, lo que simplifica el cliente y el flujo inicial.

Rápido para la carga inicial, poco JS.

Requiere menos recursos del cliente y simplicidad en el navegador.

Se reduce la dependencia de APIs específicas de JavaScript, aunque HTML y CSS siguen requiriendo compatibilidad entre navegadores.

No requiere frameworks de JavaScript pesados, transpilers ni herramientas de construcción complejas.

La UX puede ser menos fluida, ya que cada interacción puede requerir una nueva solicitud al servidor. La latencia

de red sigue siendo relevante, incluso con conexiones rápidas .

Difícil aprovechar la interactividad moderna de los navegadores.

Si no usan eso ( ), no están a la moda.

  \| 33 of 34

# Conclusiones SSR Tradicional

La presentación se renderiza en el servidor, lo que simplifica el cliente y el flujo inicial.

Rápido para la carga inicial, poco JS.

Requiere menos recursos del cliente y simplicidad en el navegador.

Se reduce la dependencia de APIs específicas de JavaScript, aunque HTML y CSS siguen requiriendo compatibilidad entre navegadores.

No requiere frameworks de JavaScript pesados, transpilers ni herramientas de construcción complejas.

La UX puede ser menos fluida, ya que cada interacción puede requerir una nueva solicitud al servidor. La latencia

de red sigue siendo relevante, incluso con conexiones rápidas . La mayoría de las aplicaciones Difícil aprovechar la interactividad web no necesitan ser una moderna de los navegadores. SPA compleja, por lo que SSR tradicional puede ser una Si no usan eso ( ), no están a la excelente opción. moda.

  \| 33 of 34

# CSS Minecraft

**There is no JavaScript on**

this page. All the logic is made 100% with pure HTML & CSS. For the best performance, please close other tabs and running programs. View on GitHub, CodePen, benjaminaster.com

Block type

Vertical Horizontal Rotation View

# Web Minecraft

https://benjaminaster.com/css

# ⨉

minecraft

  N ote

Una página web que simula el juego Minecraft utilizando solo HTML y CSS

Cero líneas de JavaScript.

  \| 34 of 34

<image redacted: 14x39px, 13x38pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 21x38px, 20x37pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 16x60px, 15x60pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 14x23px, 13x23pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 21x24px, 20x23pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 51x39px, 50x38pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 48x38px, 48x37pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 12x60px, 11x60pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 16x60px, 15x60pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 51x23px, 50x23pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 48x24px, 48x23pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 49x39px, 48x38pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 46x38px, 46x37pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 7x60px, 7x60pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 12x60px, 11x60pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 49x23px, 48x23pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 46x24px, 46x23pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 46x39px, 46x38pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 44x38px, 44x37pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 3x60px, 3x60pt, ~77dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 7x60px, 7x60pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 46x23px, 46x23pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 44x24px, 44x23pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 45x39px, 45x38pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 43x38px, 43x37pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 3x60px, 3x60pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 3x60px, 3x60pt, ~77dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 45x23px, 45x23pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 43x24px, 43x23pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 47x39px, 46x38pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 44x38px, 44x37pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 7x60px, 7x60pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 3x60px, 3x60pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 47x23px, 46x23pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 44x24px, 44x23pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 49x39px, 48x38pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 46x38px, 46x37pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 12x60px, 11x60pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 7x60px, 7x60pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 49x23px, 48x23pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 46x24px, 46x23pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 51x39px, 50x38pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 48x38px, 48x37pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 16x60px, 15x60pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 12x60px, 11x60pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 51x23px, 50x23pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 48x24px, 48x23pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 13x39px, 13x38pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 20x38px, 20x37pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 16x60px, 15x60pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 13x23px, 13x23pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 20x24px, 20x23pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 6x39px, 6x39pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 14x39px, 13x38pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 14x64px, 13x63pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 6x26px, 6x25pt, ~76dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 14x26px, 13x26pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 50x39px, 50x39pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 51x39px, 50x38pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 13x64px, 12x63pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 14x64px, 13x63pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 50x26px, 50x25pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 51x26px, 50x26pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 51x39px, 51x39pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 49x39px, 48x38pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 8x64px, 8x63pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 13x64px, 12x63pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 51x26px, 51x25pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 49x26px, 48x26pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 49x39px, 49x39pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 46x39px, 46x38pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 3x64px, 3x63pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 8x64px, 8x63pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 49x26px, 49x25pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 46x26px, 46x26pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 48x39px, 47x39pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 45x39px, 45x38pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 4x64px, 3x63pt, ~83dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 3x64px, 3x63pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 48x26px, 47x25pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 45x26px, 45x26pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 49x39px, 48x39pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 47x39px, 46x38pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 8x64px, 8x63pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 4x64px, 3x63pt, ~83dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 49x26px, 48x25pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 47x26px, 46x26pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 51x39px, 51x39pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 49x39px, 48x38pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 13x64px, 12x63pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 8x64px, 8x63pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 51x26px, 51x25pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 49x26px, 48x26pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 50x39px, 49x39pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 51x39px, 50x38pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 13x64px, 13x63pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 13x64px, 12x63pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 50x26px, 49x25pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 51x26px, 50x26pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 6x39px, 5x39pt, ~77dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 13x39px, 13x38pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 13x64px, 13x63pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 6x26px, 5x25pt, ~78dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 13x26px, 13x26pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 6x39px, 6x39pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 6x67px, 6x67pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 6x29px, 6x29pt, ~76dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 44x40px, 44x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 50x39px, 50x39pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 14x67px, 13x67pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 6x67px, 6x67pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 44x28px, 44x28pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 50x29px, 50x29pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 54x40px, 54x40pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 51x39px, 51x39pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 9x67px, 8x67pt, ~76dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 14x67px, 13x67pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 54x28px, 54x28pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 51x29px, 51x29pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 52x40px, 51x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 49x39px, 49x39pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 4x67px, 3x67pt, ~79dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 9x67px, 8x67pt, ~76dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 52x28px, 51x28pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 49x29px, 49x29pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 51x40px, 50x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 48x39px, 47x39pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 4x67px, 3x67pt, ~79dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 4x67px, 3x67pt, ~79dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 51x28px, 50x28pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 48x29px, 47x29pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 52x40px, 51x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 49x39px, 48x39pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 9x67px, 8x67pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 4x67px, 3x67pt, ~79dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 52x28px, 51x28pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 49x29px, 48x29pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 55x40px, 54x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 51x39px, 51x39pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 14x67px, 14x67pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 9x67px, 8x67pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 55x28px, 54x28pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 51x29px, 51x29pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 44x40px, 43x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 50x39px, 49x39pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 6x67px, 5x67pt, ~77dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 14x67px, 14x67pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 44x28px, 43x28pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 50x29px, 49x29pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 6x39px, 5x39pt, ~77dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 6x67px, 5x67pt, ~77dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 6x29px, 5x29pt, ~77dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 37x40px, 37x40pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 44x40px, 44x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 16x71px, 15x71pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 37x32px, 37x31pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 44x32px, 44x32pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 58x40px, 58x40pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 54x40px, 54x40pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 10x71px, 9x71pt, ~76dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 16x71px, 15x71pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 58x32px, 58x31pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 54x32px, 54x32pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 55x40px, 54x40pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 52x40px, 51x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 4x71px, 4x71pt, ~76dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 10x71px, 9x71pt, ~76dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 55x32px, 54x31pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 52x32px, 51x32pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 54x40px, 53x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 51x40px, 50x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 4x71px, 4x71pt, ~76dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 4x71px, 4x71pt, ~76dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 54x32px, 53x31pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 51x32px, 50x32pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 55x40px, 54x40pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 52x40px, 51x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 10x71px, 9x71pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 4x71px, 4x71pt, ~76dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 55x32px, 54x31pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 52x32px, 51x32pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 58x40px, 58x40pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 55x40px, 54x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 16x71px, 15x71pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 10x71px, 9x71pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 58x32px, 58x31pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 55x32px, 54x32pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 37x40px, 36x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 44x40px, 43x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 16x71px, 15x71pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 37x32px, 36x31pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 44x32px, 43x32pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 30x41px, 29x40pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 37x40px, 37x40pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 17x75px, 17x75pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 30x36px, 29x35pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 37x36px, 37x36pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 62x41px, 62x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 58x40px, 58x40pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 11x75px, 11x75pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 17x75px, 17x75pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 62x36px, 62x35pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 58x36px, 58x36pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 59x41px, 58x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 55x40px, 54x40pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 4x75px, 4x75pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 11x75px, 11x75pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 59x36px, 58x35pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 55x36px, 54x36pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 57x41px, 56x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 54x40px, 53x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 4x75px, 4x75pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 4x75px, 4x75pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 57x36px, 56x35pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 54x36px, 53x36pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 59x41px, 58x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 55x40px, 54x40pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 11x75px, 11x75pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 4x75px, 4x75pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 59x36px, 58x35pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 55x36px, 54x36pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 62x41px, 62x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 58x40px, 58x40pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 18x75px, 17x75pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 11x75px, 11x75pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 62x36px, 62x35pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 58x36px, 58x36pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 29x41px, 29x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 37x40px, 36x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 18x75px, 17x75pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 29x36px, 29x35pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 37x36px, 36x36pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 21x41px, 21x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 30x41px, 29x40pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 19x80px, 19x80pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 21x41px, 21x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 30x41px, 29x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 67x41px, 66x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 62x41px, 62x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 12x80px, 12x80pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 19x80px, 19x80pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 67x41px, 66x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 62x41px, 62x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 63x41px, 62x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 59x41px, 58x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 5x80px, 4x80pt, ~78dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 12x80px, 12x80pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 63x41px, 62x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 59x41px, 58x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 61x41px, 60x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 57x41px, 56x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 5x80px, 5x80pt, ~76dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 5x80px, 4x80pt, ~78dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 61x41px, 60x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 57x41px, 56x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 63x41px, 63x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 59x41px, 58x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 12x80px, 12x80pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 5x80px, 5x80pt, ~76dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 63x41px, 63x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 59x41px, 58x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 67x41px, 66x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 62x41px, 62x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 20x80px, 19x80pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 12x80px, 12x80pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 67x41px, 66x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 62x41px, 62x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 21x41px, 20x40pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 29x41px, 29x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 20x80px, 19x80pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 21x41px, 20x40pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 29x41px, 29x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 12x40px, 11x40pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 21x41px, 21x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 21x86px, 21x85pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 12x47px, 11x46pt, ~76dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 21x46px, 21x46pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 71x40px, 71x40pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 67x41px, 66x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 14x86px, 13x85pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 21x86px, 21x85pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 71x47px, 71x46pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 67x46px, 66x46pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 68x40px, 67x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 63x41px, 62x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 6x86px, 5x85pt, ~79dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 14x86px, 13x85pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 68x47px, 67x46pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 63x46px, 62x46pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 66x40px, 65x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 61x41px, 60x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 6x86px, 5x85pt, ~79dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 6x86px, 5x85pt, ~79dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 66x47px, 65x46pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 61x46px, 60x46pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 68x40px, 67x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 63x41px, 63x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 14x86px, 13x85pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 6x86px, 5x85pt, ~79dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 68x47px, 67x46pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 63x46px, 63x46pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 71x40px, 70x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 67x41px, 66x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 21x86px, 20x85pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 14x86px, 13x85pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 71x47px, 70x46pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 67x46px, 66x46pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 11x40px, 10x40pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 21x41px, 20x40pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 21x86px, 20x85pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 11x47px, 10x46pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 21x46px, 20x46pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 12x40px, 11x40pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 12x92px, 11x92pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 12x53px, 11x53pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 64x39px, 64x39pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 71x40px, 71x40pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 16x92px, 15x92pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 12x92px, 11x92pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 64x54px, 64x54pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 71x53px, 71x53pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 73x39px, 73x39pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 68x40px, 67x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 6x92px, 6x92pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 16x92px, 15x92pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 73x54px, 73x54pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 68x53px, 67x53pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 71x39px, 70x39pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 66x40px, 65x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 6x92px, 6x92pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 6x92px, 6x92pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 71x54px, 70x54pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 66x53px, 65x53pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 73x39px, 73x39pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 68x40px, 67x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 16x92px, 15x92pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 6x92px, 6x92pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 73x54px, 73x54pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 68x53px, 67x53pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 64x39px, 63x39pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 71x40px, 70x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 11x92px, 10x92pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 16x92px, 15x92pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 64x54px, 63x54pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 71x53px, 70x53pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 11x40px, 10x40pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 11x92px, 10x92pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 11x53px, 10x53pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 57x37px, 56x37pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 64x39px, 64x39pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 18x100px, 18x99pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 57x64px, 56x63pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 64x61px, 64x61pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 80x37px, 79x37pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 73x39px, 73x39pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 7x100px, 6x99pt, ~77dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 18x100px, 18x99pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 80x64px, 79x63pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 73x61px, 73x61pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 77x37px, 76x37pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 71x39px, 70x39pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 7x100px, 6x99pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 7x100px, 6x99pt, ~77dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 77x64px, 76x63pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 71x61px, 70x61pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 80x37px, 80x37pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 73x39px, 73x39pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 18x100px, 18x99pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 7x100px, 6x99pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 80x64px, 80x63pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 73x61px, 73x61pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 56x37px, 55x37pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 64x39px, 63x39pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 18x100px, 18x99pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 56x64px, 55x63pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 64x61px, 63x61pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

# CSS Minecraft

**There is no JavaScript on**

this page. All the logic is made 100% with pure HTML & CSS. For the best performance, please close other tabs and running programs. View on GitHub, CodePen, benjaminaster.com

Block type

Vertical Horizontal Rotation View

# Web Minecraft

https://benjaminaster.com/css

# ⨉

minecraft

  N ote

Una página web que simula el juego Minecraft utilizando solo HTML y CSS

Cero líneas de JavaScript.

  W arning

React es una herramienta válida, pero no reemplaza el conocimiento de HTML, CSS y JavaScript ni siempre justifica su complejidad.

  \| 34 of 34

<image redacted: 14x39px, 13x38pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 21x38px, 20x37pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 16x60px, 15x60pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 14x23px, 13x23pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 21x24px, 20x23pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 51x39px, 50x38pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 48x38px, 48x37pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 12x60px, 11x60pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 16x60px, 15x60pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 51x23px, 50x23pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 48x24px, 48x23pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 49x39px, 48x38pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 46x38px, 46x37pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 7x60px, 7x60pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 12x60px, 11x60pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 49x23px, 48x23pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 46x24px, 46x23pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 46x39px, 46x38pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 44x38px, 44x37pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 3x60px, 3x60pt, ~77dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 7x60px, 7x60pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 46x23px, 46x23pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 44x24px, 44x23pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 45x39px, 45x38pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 43x38px, 43x37pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 3x60px, 3x60pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 3x60px, 3x60pt, ~77dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 45x23px, 45x23pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 43x24px, 43x23pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 47x39px, 46x38pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 44x38px, 44x37pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 7x60px, 7x60pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 3x60px, 3x60pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 47x23px, 46x23pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 44x24px, 44x23pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 49x39px, 48x38pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 46x38px, 46x37pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 12x60px, 11x60pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 7x60px, 7x60pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 49x23px, 48x23pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 46x24px, 46x23pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 51x39px, 50x38pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 48x38px, 48x37pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 16x60px, 15x60pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 12x60px, 11x60pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 51x23px, 50x23pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 48x24px, 48x23pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 13x39px, 13x38pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 20x38px, 20x37pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 16x60px, 15x60pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 13x23px, 13x23pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 20x24px, 20x23pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 6x39px, 6x39pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 14x39px, 13x38pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 14x64px, 13x63pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 6x26px, 6x25pt, ~76dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 14x26px, 13x26pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 50x39px, 50x39pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 51x39px, 50x38pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 13x64px, 12x63pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 14x64px, 13x63pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 50x26px, 50x25pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 51x26px, 50x26pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 51x39px, 51x39pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 49x39px, 48x38pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 8x64px, 8x63pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 13x64px, 12x63pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 51x26px, 51x25pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 49x26px, 48x26pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 49x39px, 49x39pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 46x39px, 46x38pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 3x64px, 3x63pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 8x64px, 8x63pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 49x26px, 49x25pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 46x26px, 46x26pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 48x39px, 47x39pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 45x39px, 45x38pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 4x64px, 3x63pt, ~83dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 3x64px, 3x63pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 48x26px, 47x25pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 45x26px, 45x26pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 49x39px, 48x39pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 47x39px, 46x38pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 8x64px, 8x63pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 4x64px, 3x63pt, ~83dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 49x26px, 48x25pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 47x26px, 46x26pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 51x39px, 51x39pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 49x39px, 48x38pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 13x64px, 12x63pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 8x64px, 8x63pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 51x26px, 51x25pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 49x26px, 48x26pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 50x39px, 49x39pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 51x39px, 50x38pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 13x64px, 13x63pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 13x64px, 12x63pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 50x26px, 49x25pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 51x26px, 50x26pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 6x39px, 5x39pt, ~77dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 13x39px, 13x38pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 13x64px, 13x63pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 6x26px, 5x25pt, ~78dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 13x26px, 13x26pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 6x39px, 6x39pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 6x67px, 6x67pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 6x29px, 6x29pt, ~76dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 44x40px, 44x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 50x39px, 50x39pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 14x67px, 13x67pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 6x67px, 6x67pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 44x28px, 44x28pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 50x29px, 50x29pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 54x40px, 54x40pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 51x39px, 51x39pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 9x67px, 8x67pt, ~76dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 14x67px, 13x67pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 54x28px, 54x28pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 51x29px, 51x29pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 52x40px, 51x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 49x39px, 49x39pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 4x67px, 3x67pt, ~79dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 9x67px, 8x67pt, ~76dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 52x28px, 51x28pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 49x29px, 49x29pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 51x40px, 50x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 48x39px, 47x39pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 4x67px, 3x67pt, ~79dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 4x67px, 3x67pt, ~79dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 51x28px, 50x28pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 48x29px, 47x29pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 52x40px, 51x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 49x39px, 48x39pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 9x67px, 8x67pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 4x67px, 3x67pt, ~79dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 52x28px, 51x28pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 49x29px, 48x29pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 55x40px, 54x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 51x39px, 51x39pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 14x67px, 14x67pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 9x67px, 8x67pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 55x28px, 54x28pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 51x29px, 51x29pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 44x40px, 43x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 50x39px, 49x39pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 6x67px, 5x67pt, ~77dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 14x67px, 14x67pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 44x28px, 43x28pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 50x29px, 49x29pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 6x39px, 5x39pt, ~77dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 6x67px, 5x67pt, ~77dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 6x29px, 5x29pt, ~77dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 37x40px, 37x40pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 44x40px, 44x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 16x71px, 15x71pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 37x32px, 37x31pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 44x32px, 44x32pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 58x40px, 58x40pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 54x40px, 54x40pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 10x71px, 9x71pt, ~76dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 16x71px, 15x71pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 58x32px, 58x31pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 54x32px, 54x32pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 55x40px, 54x40pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 52x40px, 51x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 4x71px, 4x71pt, ~76dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 10x71px, 9x71pt, ~76dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 55x32px, 54x31pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 52x32px, 51x32pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 54x40px, 53x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 51x40px, 50x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 4x71px, 4x71pt, ~76dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 4x71px, 4x71pt, ~76dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 54x32px, 53x31pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 51x32px, 50x32pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 55x40px, 54x40pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 52x40px, 51x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 10x71px, 9x71pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 4x71px, 4x71pt, ~76dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 55x32px, 54x31pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 52x32px, 51x32pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 58x40px, 58x40pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 55x40px, 54x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 16x71px, 15x71pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 10x71px, 9x71pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 58x32px, 58x31pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 55x32px, 54x32pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 37x40px, 36x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 44x40px, 43x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 16x71px, 15x71pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 37x32px, 36x31pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 44x32px, 43x32pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 30x41px, 29x40pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 37x40px, 37x40pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 17x75px, 17x75pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 30x36px, 29x35pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 37x36px, 37x36pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 62x41px, 62x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 58x40px, 58x40pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 11x75px, 11x75pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 17x75px, 17x75pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 62x36px, 62x35pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 58x36px, 58x36pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 59x41px, 58x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 55x40px, 54x40pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 4x75px, 4x75pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 11x75px, 11x75pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 59x36px, 58x35pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 55x36px, 54x36pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 57x41px, 56x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 54x40px, 53x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 4x75px, 4x75pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 4x75px, 4x75pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 57x36px, 56x35pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 54x36px, 53x36pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 59x41px, 58x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 55x40px, 54x40pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 11x75px, 11x75pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 4x75px, 4x75pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 59x36px, 58x35pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 55x36px, 54x36pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 62x41px, 62x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 58x40px, 58x40pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 18x75px, 17x75pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 11x75px, 11x75pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 62x36px, 62x35pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 58x36px, 58x36pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 29x41px, 29x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 37x40px, 36x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 18x75px, 17x75pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 29x36px, 29x35pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 37x36px, 36x36pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 21x41px, 21x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 30x41px, 29x40pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 19x80px, 19x80pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 21x41px, 21x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 30x41px, 29x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 67x41px, 66x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 62x41px, 62x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 12x80px, 12x80pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 19x80px, 19x80pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 67x41px, 66x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 62x41px, 62x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 63x41px, 62x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 59x41px, 58x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 5x80px, 4x80pt, ~78dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 12x80px, 12x80pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 63x41px, 62x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 59x41px, 58x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 61x41px, 60x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 57x41px, 56x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 5x80px, 5x80pt, ~76dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 5x80px, 4x80pt, ~78dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 61x41px, 60x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 57x41px, 56x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 63x41px, 63x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 59x41px, 58x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 12x80px, 12x80pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 5x80px, 5x80pt, ~76dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 63x41px, 63x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 59x41px, 58x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 67x41px, 66x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 62x41px, 62x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 20x80px, 19x80pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 12x80px, 12x80pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 67x41px, 66x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 62x41px, 62x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 21x41px, 20x40pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 29x41px, 29x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 20x80px, 19x80pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 21x41px, 20x40pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 29x41px, 29x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 12x40px, 11x40pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 21x41px, 21x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 21x86px, 21x85pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 12x47px, 11x46pt, ~76dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 21x46px, 21x46pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 71x40px, 71x40pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 67x41px, 66x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 14x86px, 13x85pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 21x86px, 21x85pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 71x47px, 71x46pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 67x46px, 66x46pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 68x40px, 67x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 63x41px, 62x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 6x86px, 5x85pt, ~79dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 14x86px, 13x85pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 68x47px, 67x46pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 63x46px, 62x46pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 66x40px, 65x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 61x41px, 60x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 6x86px, 5x85pt, ~79dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 6x86px, 5x85pt, ~79dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 66x47px, 65x46pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 61x46px, 60x46pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 68x40px, 67x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 63x41px, 63x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 14x86px, 13x85pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 6x86px, 5x85pt, ~79dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 68x47px, 67x46pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 63x46px, 63x46pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 71x40px, 70x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 67x41px, 66x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 21x86px, 20x85pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 14x86px, 13x85pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 71x47px, 70x46pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 67x46px, 66x46pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 11x40px, 10x40pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 21x41px, 20x40pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 21x86px, 20x85pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 11x47px, 10x46pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 21x46px, 20x46pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 12x40px, 11x40pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 12x92px, 11x92pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 12x53px, 11x53pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 64x39px, 64x39pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 71x40px, 71x40pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 16x92px, 15x92pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 12x92px, 11x92pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 64x54px, 64x54pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 71x53px, 71x53pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 73x39px, 73x39pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 68x40px, 67x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 6x92px, 6x92pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 16x92px, 15x92pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 73x54px, 73x54pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 68x53px, 67x53pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 71x39px, 70x39pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 66x40px, 65x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 6x92px, 6x92pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 6x92px, 6x92pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 71x54px, 70x54pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 66x53px, 65x53pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 73x39px, 73x39pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 68x40px, 67x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 16x92px, 15x92pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 6x92px, 6x92pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 73x54px, 73x54pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 68x53px, 67x53pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 64x39px, 63x39pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 71x40px, 70x40pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 11x92px, 10x92pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 16x92px, 15x92pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 64x54px, 63x54pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 71x53px, 70x53pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 11x40px, 10x40pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 11x92px, 10x92pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 11x53px, 10x53pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 57x37px, 56x37pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 64x39px, 64x39pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 18x100px, 18x99pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 57x64px, 56x63pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 64x61px, 64x61pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 80x37px, 79x37pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 73x39px, 73x39pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 7x100px, 6x99pt, ~77dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 18x100px, 18x99pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 80x64px, 79x63pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 73x61px, 73x61pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 77x37px, 76x37pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 71x39px, 70x39pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 7x100px, 6x99pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 7x100px, 6x99pt, ~77dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 77x64px, 76x63pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 71x61px, 70x61pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 80x37px, 80x37pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 73x39px, 73x39pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 18x100px, 18x99pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 7x100px, 6x99pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 80x64px, 80x63pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 73x61px, 73x61pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 56x37px, 55x37pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 64x39px, 63x39pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 18x100px, 18x99pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 56x64px, 55x63pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 64x61px, 63x61pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>