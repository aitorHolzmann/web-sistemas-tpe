# Pr

# ogramación Web

# Introducción

Dr. Alejandro Zunino & Dr. Alfredo Teyseyre alejandro.zunino@isistan.unicen.edu.ar, alfredo.teyseyre@isistan.unicen.edu.ar

  \| 1 of 73

<image redacted: 72x72px, 72x72pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

# Contenidos

1\. La Web, modelo cliente-servidor, HTML, HTTP, forms/CGI.

2\. Programación Web: Server side y client side. El modelo three tier.

3\. Contenedores y máquinas virtuales.

4\. Persistencia: relacional, NoSQL, ORM.

5\. Lógica del negocio. Escalabilidad.

6\. Presentación: ruteo, templates, DOM, SPA.

7\. Arquitecturas: onion, clean, vertical slices.

  \| 2 of 73

<image redacted: 479x479px, 479x479pt, ~72dpi, JPG, DEVICE_RGB, 32bpp>

# La Web, modelo cliente

# servidor, HTML, HTTP,

# forms/CGI

Qué es la Web y qué ofrece para desarrollar software. Cómo se estructura un sistema Web, qué es un servidor Web, qué es un cliente Web, qué es HTML, qué es HTTP, cómo se envían datos entre el cliente y el servidor, qué son los formularios y cómo funcionan.

  \| 3 of 73

# Objetivos de la materia

Conocer los conceptos del desarrollo Web:

Sistemas distribuidos, Masivamente paralelos

Conocer las tecnologías de la programación Web:

HTML, CSS, JavaScript

Servidores Web

Bases de datos

Contenedores y máquinas virtuales

Aprender haciendo: No preguntando a una IA

  \| 4 of 73

# Objetivos de la materia

Conocer los conceptos del desarrollo Web:

Sistemas distribuidos, Masivamente paralelos

Conocer las tecnologías de la programación Web:

HTML, CSS, JavaScript

Servidores Web

Bases de datos

Contenedores y máquinas virtuales

Aprender haciendo: No preguntando a una IA

Full stack: Las partes de todo sistema Web

¿Todas esas? ¿Algunas?

  \| 4 of 73

<image redacted: 28x28px, 28x28pt, ~73dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 20x21px, 19x20pt, ~75dpi, PNG, DEVICE_RGB, 32bpp>

# Sistemas distribuidos

Prácticamente todos los sistemas de software actuales son distribuidos

El procesamiento de la información se realiza entre varias computadoras en lugar de estar confinado a una sola máquina.

El software distribuido es un conjunto de computadoras independientes que se presentan al usuario como un sistema único y coherente.

La distribución de la información y el procesamiento permite :

Escalabilidad: agregar más computadoras para aumentar la capacidad de procesamiento.

Tolerancia a fallos: si una computadora falla, el sistema puede seguir funcionando.

Flexibilidad: se pueden agregar o quitar computadoras según sea necesario.

tenemos que pensar en qué lugar está cada componente .

  \| 5 of 73

# Sistema Web

Firefox Chrome Safari

Web Server

Business logic

Database File Server

  \| 6 of 73

# Ventajas y desventajas

Compartir recursos

Tecnologías abiertas

Concurrencia

Escalabilidad

Tolerancia a fallos

Transparencia

  \| 7 of 73

# Ventajas y desventajas

Compartir recursos

Tecnologías abiertas

Concurrencia

Escalabilidad

Tolerancia a fallos

Transparencia

Complejidad

Seguridad

Manejabilidad

Imprevisibilidad

Comunicación

  \| 7 of 73

¿Qué buscamos con sistemas distribuidos?

Adaptabilidad a cambios en el entorno

Cambios en la carga de trabajo

Cambios en la configuración del sistema

Cambios en los requisitos de los usuarios

Disponibilidad

Debe estar disponible en todo momento

Capaz de recuperarse ante fallos

Mantener buen tiempo de respuesta

Simple de mantener, monitorear y administrar

Fácil de actualizar y escalar

Proveer herramientas de monitoreo

y administración

  \| 8 of 73

¿Qué esperar de la materia?

# Aprenderemos

Conceptos de desarrollo Web.

Principios de diseño:

Organización general

Buenas prácticas

Arquitecturas típicas

Descripción de algunas tecnologías

Desarrollo de un sistema Web simple

  \| 9 of 73

¿Qué esperar de la materia?

# Aprenderemos

Conceptos de desarrollo Web.

Principios de diseño:

Organización general

Buenas prácticas

Arquitecturas típicas

Descripción de algunas tecnologías

Desarrollo de un sistema Web simple

# No esperes

Ser experto en Web o FullStack dev.

Seguridad, escalabilidad, rendimiento, etc.

Diseño HTML/CSS.

Mobile, SPA.

Monocultura: , , , , , etc

  \| 9 of 73

# Modalidad de la materia

Hay mucha práctica, pero no es sólo de programación.

| En la realidad                                                                   |                          | : sólo           |
|----------------------------------------------------------------------------------|--------------------------|------------------|
| no vas a encontrar servidores                                                    | con o                    |                  |
| si usas esos sistemas y algo no                                                  | es tu problema funciona: | .                |
| WSL no es Linux. Si lo utilizas, recuerda Utilizaremos servidores Linux, Docker, | . bases de datos,        | terminales, etc: |
| Amigate con la línea de comandos                                                 | .                        |                  |

Sugerencias:

Linux: Mint, Debian, Manjaro, Arch, etc.

Editores: VSCode, NeoVim, Helix, etc.

  \| 10 of 73

<image redacted: 15x6px, 14x6pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

# Modalidad de la materia

Hay mucha práctica, pero no es sólo de programación.

| En la realidad                                                                   |                          | : sólo           |
|----------------------------------------------------------------------------------|--------------------------|------------------|
| no vas a encontrar servidores                                                    | con o                    |                  |
| si usas esos sistemas y algo no                                                  | es tu problema funciona: | .                |
| WSL no es Linux. Si lo utilizas, recuerda Utilizaremos servidores Linux, Docker, | . bases de datos,        | terminales, etc: |
| Amigate con la línea de comandos                                                 | .                        |                  |

Sugerencias:

Linux: Mint, Debian, Manjaro, Arch, etc.

Editores: VSCode, NeoVim, Helix, etc.

  \| 10 of 73

<image redacted: 15x6px, 14x6pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

Por qué no Windows?

| Inestabilidad | , ineficiencia    | , mal | desempeño   | , necesita GUI  | , es para escritorio |
|---------------|-------------------|-------|-------------|-----------------|----------------------|
| , dificultad  | de automatización | ,     | inseguridad | , elevado costo | , etc.               |

  \| 11 of 73

<image redacted: 554x362px, 554x362pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

Por qué no Windows?

| Inestabilidad | , ineficiencia    | , mal | desempeño   | , necesita GUI  | , es para escritorio |
|---------------|-------------------|-------|-------------|-----------------|----------------------|
| , dificultad  | de automatización | ,     | inseguridad | , elevado costo | , etc.               |

  \| 11 of 73

<image redacted: 554x362px, 554x362pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

# Evaluación

Se ofrecen dos esquemas de evaluación:

1\. Evaluación Tradicional

2\. Trabajo Práctico Integrador:

Desarrollo de una aplicación web simple.

Seguimiento continuo durante el cuatrimestre.

  N ote

Se recomienda optar por el Trabajo Práctico, ya que fomenta un aprendizaje más profundo y práctico.

Tres etapas con entregas parciales.

Cada etapa cuenta con una instancia de reentrega.

La nota del trabajo (1-10) es la nota de la cursada.

Con 7 o más, se promociona la materia.

  Im portant

La aprobación del trabajo práctico no es condición para rendir los parciales.

  \| 12 of 73

¿Qué es la Web?

  \| 13 of 73

<image redacted: 479x479px, 479x479pt, ~72dpi, JPG, DEVICE_RGB, 32bpp>

# La Web

La WWW es una colección hipervinculada de documentos y programas que residen en computadoras alrededor del mundo, conectadas a Internet

Internet: una colección de protocolos:

TCP/IP, DNS, ARP, etc.

HTTP, FTP, SMTP, IMAP, etc.

Basada en el modelo Cliente/Servidor

| Firefox    | Chrome Sistema Web | Safari     |
|------------|--------------------|------------|
| Web Server | Web Server         | Web Server |

  \| 14 of 73

¿Qué es HTML?

HTML (HyperText Markup Language): el lenguaje de marcado que define la estructura de las páginas web.

Se basa en etiquetas (tags) que envuelven el contenido:

\<h1\>Título\</h1\>

\<p\>Un párrafo con \<strong\>texto\</strong\>.\</p\> \<a href=" https://www.unicen.edu.ar"\>Un enlace\</a\> alt=" \<img src=" Logo"\> /logo.png"

El navegador interpreta las etiquetas y renderiza la página.

HTML no es un lenguaje de programación: no tiene lógica ni control de flujo .

  N ote

Profundizaremos en la estructura HTML en las clases de la capa de presentación.

  \| 15 of 73

# Cliente-Servidor

Las aplicaciones se modelan como servicios proporcionados por un servidor a un cliente.

Los clientes son:

Navegadores Web capaces de mostrar/ejecutar:

Software que consume esos servicios mediante: REST, SOAP, Protocol Buffers, etc.

La comunicación entre cliente y servidor se realiza a través del protocolo HTTP

los clientes conocen a los servidores por su URL, pero no a la inversa

Web Browser

HTTP Request Response HTTP Request Response

Web Server 1 Web Server 2

  \| 16 of 73

<image redacted: 23x23px, 22x23pt, ~74dpi, PNG, DEVICE_RGB, 32bpp>

# HTTP: Hypertext Transfer Protocol

El protocolo de la World Wide Web.

Define cómo se comunican los

navegadores (clientes) y los servidores web:

Codificación de solicitudes y respuestas (formatos de documentos).

Basado en un modelo de solicitud

respuesta.

Sin estado (stateless): Cada solicitud es independiente de las anteriores.

Funciona sobre TCP/IP.

Cliente Servidor

Petición HTTP (Ej: GET /index.html)

Respuesta HTTP (Ej: 200 OK + Contenido HTML)

Cliente Servidor

  \| 17 of 73

¿Qué son las URLs?

URL (Uniform Resource Locator) es la dirección única de un recurso en la web.

https://www.ejemplo.com:443/camino/recurso?param1=valor1&param2=valor2

https://

Protocolo www.ejemplo.com

Dominio/Host :443

/camino/recurso Puerto

Ruta/Path ?param1=valor1

?param2=valor2 Parámetros/Query

Parámetros/Query

  \| 18 of 73

# Interacción Simple

# ¿Cómo obtenemos una

# página y sus imágenes?:

1\. El navegador pide un HTML.

2\. El servidor envía el HTML.

3\. El navegador parsea el HTML y encuentra recursos (CSS, JS, imágenes).

4\. El navegador pide cada recurso al servidor.

5\. El servidor envía cada recurso.

6\. El navegador renderiza la página completa.

Navegador Servidor

GET /pagina.html

HTML de pagina.html

GET /estilo.css

CSS

GET /logo.png

Imagen PNG

Navegador Servidor

  \| 19 of 73

# Petición

Así se ve una petición GET para /index.html y la respuesta del servidor.

GET /index.html HTTP/1.1

Host: \[www.ejemplo.com\](https://www.ejemplo.com) User-Agent: Mozilla/5.0 ... Accept: text/html,\*/\* Accept-Language: es-ES Connection: keep-alive

# Respuesta

## Cabecera

HTTP/1.1 200 OK

Date: Fri, 27 Oct 2024 10:00:00 GMT Server: Apache Content-Type: text/html Content-Length: 150

## Contenido

\<!DOCTYPE html\>

\<html\>

\<head\>

\<title\>Página\</title\> \</head\>

\<body\> \<h1\>Hola!\</h1\>

\</body\> \</html\>

  \| 20 of 73

# Enviando Parámetros: GET y POST

Los clientes a menudo necesitan enviar datos al servidor (búsquedas, formularios, etc.). Los métodos más comunes son GET y POST.

GET: Envía datos en la URL.

POST: Envía datos en el cuerpo de la petición.

Login username: MyName \*\*\*\* password:

| email:                | my@email.com |
|-----------------------|--------------|
| contact:              | Email        |
| \[x\] Standard Enviar | Express      |

  \| 21 of 73

Los datos (query string) van en la URL: https://search.brave.com/search?q=unicen

Limi tado en tamaño.

# Método

Vi sible en el historial del navegador y logs del servidor.

# GET:

Cacheable y marcable como favorito.

# usado

i dempotente (repetirlo no cambia el resultado).

# para

# solicitar

Petición GET

# datos

GET /buscar?q=unicen Datos: q=unicen Cliente Servidor Respuesta

  \| 22 of 73

# Método POST:

# enviar datos

# para ser

# procesados

Los datos van en el cuerpo de la petición: https://www.ejemplo.com/login

Sin límite práctico de tamaño.

N o visible directamente en la URL/historial.

N o cacheable por defecto.

N o idempotente (repetirlo puede tener efectos secundarios).

Petición POST

POST /guardar {nombre:'Ana',edad:30} Cliente Servidor

Respuesta

  \| 23 of 73

# Formularios HTML para obtener datos

Usan los métodos GET o POST para comunicarse con el servidor.

formUsuario" action=" \<form id=" /procesar"m \<label for=" nombre"\>Nombre:\</label\>

text"i nombre"n ame=" d=" \<input type=" submit"\>Enviar\</button\> \<button type=" \</form\>

Nombre:

GET"\> ethod="

usuario"\>

Enviar

  \| 24 of 73

# Formularios HTML para obtener datos

Usan los métodos GET o POST para comunicarse con el servidor.

formUsuario" action=" \<form id=" /procesar"m \<label for=" nombre"\>Nombre:\</label\>

text"i nombre"n ame=" d=" \<input type=" submit"\>Enviar\</button\> \<button type=" \</form\>

Nombre:

GET"\> ethod="

usuario"\>

Enviar

  \| 24 of 73

# Formularios HTML para obtener datos

Usan los métodos GET o POST para comunicarse con el servidor.

formUsuario" action=" \<form id=" ethod=" /procesar"m \<label for=" nombre"\>Nombre:\</label\>

text"i nombre"n ame=" usuario"\> d=" \<input type=" submit"\>Enviar\</button\> \<button type=" \</form\>

Nombre:

Enviar

Al presionar Enviar , el navegador realiza una petición GET al servidor:

Usuario Navegador Servidor

Rellena formulario y envía

Muestra resultados

Usuario Navegador Servidor

GET"\>

GET /procesar?usuario=Alejandro

Respuesta (Resultados)

  \| 24 of 73

# Cabeceras HTTP

Metadatos enviados en cada petición y respuesta que proporcionan información crucial sobre la transacción:

Generales: Aplicables a ambos (Ej: Date , Connection ).

De Petición: Info sobre el cliente o el recurso solicitado (Ej: User-Agent , Accept , Host ).

De Respuesta: Info sobre el servidor o el recurso enviado (Ej: Server , Content Type , Content-Length ).

De Entidad: Info sobre el cuerpo del mensaje (Ej: Content-Type , Content-Length ).

  \| 25 of 73

# Cabeceras Comunes

# Petición :

Host: Dominio del servidor.

User-Agent: Identificación del navegador/cliente.

Accept: Tipos de contenido que el cliente acepta.

Accept-Language: Idiomas preferidos.

Cookie: Datos de sesión enviados por el cliente.

Authorization: Credenciales de

autenticación.

# Respuesta :

Content-Type: Tipo de dato del cuerpo (HTML, JSON, PNG…).

Content-Length: Tamaño del cuerpo.

Server: Software del servidor web.

Set-Cookie: Pide al navegador guardar información.

Location: Redirección (con 3xx).

Cache-Control: Directivas de caché.

  \| 26 of 73

# Códigos de Estado HTTP

Cada respuesta HTTP incluye un código de estado de 3 dígitos que indica el resultado de la petición.

1xx - Informativas

100 Continue : el servidor espera el resto de la petición.

2xx - Éxito

200 OK : petición exitosa (como vimos con curl ).

201 Created : recurso creado (ej: POST).

204 No Content : exitosa, sin contenido.

3xx - Redirecciones

301 Moved Permanently / 302 Found : la URL cambió de ubicación.

304 Not Modified : usa la copia en caché.

  \| 27 of 73

# Códigos de Estado HTTP

4xx - Errores del cliente

400 Bad Request : petición mal formada.

401 Unauthorized : se requiere autenticación.

403 Forbidden : sin permiso para el recurso.

404 Not Found : el recurso no existe.

5xx - Errores del servidor

500 Internal Server Error : error inesperado del servidor.

502 Bad Gateway / 503 Service Unavailable : problemas con el servicio.

  Ti p

Los códigos se ven en las herramientas de desarrollo del navegador (pestaña Network).

  \| 28 of 73

# El servidor de Web

  \| 29 of 73

<image redacted: 319x479px, 319x479pt, ~72dpi, JPG, DEVICE_RGB, 32bpp>

¿Qué es un Servidor Web?

Es un software que se ejecuta en una computadora y cuya función principal es escuchar y responder a peticiones realizadas por clientes a través del protocolo HTTP.

Es la columna vertebral de cualquier aplicación o sitio web.

Su trabajo es"servir" contenido a los navegadores o a otros programas cliente.

Puede servir desde un simple archivo de texto hasta complejas aplicaciones dinámicas.

  \| 30 of 73

# Funciones Principales

## 1\. Escuchar Peticiones

Se"enlaza" a una dirección IP y un puerto de red (ej: 80 para HTTP, 443 para HTTPS).

Espera a que los clientes envíen una petición HTTP.

## 2\. Procesar la Petición

Analiza la petición: el método (GET, POST), la URL, las cabeceras, el cuerpo y decide qué hacer:

Servir un archivo: Si la URL corresponde a un archivo (HTML, CSS, imagen), lo lee del disco y lo envía.

Invocar lógica dinámica: Si la URL corresponde a una ruta de la aplicación, pasa la petición a un"manejador" (handler) para que genere una respuesta dinámica.

  \| 31 of 73

# Funciones Principales

## 3\. Enviar la Respuesta

Construye una respuesta HTTP:

Línea de estado: Versión del

protocolo y código de estado (ej: HTTP/1.1 200 OK ).

Cabeceras: Metadatos como

Content-Type , Content

Length , etc.

Cuerpo (Body): El contenido real (el HTML, la imagen, datos JSON, etc.).

Envía la respuesta al cliente a través de la conexión de red.

Cliente Servidor

GET /index.html HTTP/1.1

Procesa: Lee /static/index.html

HTTP/1.1 200 OK Content-Type: text/html ...HTML...

Cliente Servidor

  \| 32 of 73

# Tipos de Servidores Web

Existen servidores de propósito general y bibliotecas que nos permiten construir los nuestros.

Servidores Dedicados:

Apache HTTP Server: Muy configurable, basado en módulos. Fue el más popular durante años.

Nginx: Conocido por su alto rendimiento y bajo consumo de memoria. Excelente como servidor de archivos estáticos y proxy inverso.

Bibliotecas de Programación:

net/http: El paquete estándar de Go, potente y usado en producción.

Express.js / Fastify: Frameworks populares para Node.js.

Jetty / Tomcat: Servidores de aplicaciones para Java.

  \| 33 of 73

# Tipos de Servidores Web

Existen servidores de propósito general y bibliotecas que nos permiten construir los nuestros.

Servidores Dedicados:

Apache HTTP Server: Muy configurable, basado en módulos. Fue el más popular durante años.   N ote

En desarrollo, a menudo usamos las bibliotecas del Nginx: Conocido por su alto rendimiento y bajo consumo de memoria. lenguaje. En producción, es común poner un Excelente como servidor de archivos estáticos y proxy inverso. servidor dedicado como Nginx delante de nuestra aplicación (como proxy inverso) para mejorar el rendimiento y la seguridad. Bibliotecas de Programación:

net/http: El paquete estándar de Go, potente y usado en producción.

Express.js / Fastify: Frameworks populares para Node.js.

Jetty / Tomcat: Servidores de aplicaciones para Java.

  \| 33 of 73

# Objetivo

Crear un servidor web en Go que:

1\. Escuche peticiones HTTP en un puerto específico (ej: 8080).

2\. Responda a todas las peticiones con una página HTML simple.

3\. Utilice únicamente las bibliotecas estándar de Go.

La Página HTML a Servir:

\<!DOCTYPE html\>

\<html\>

\<head\>

\<title\>Página\</title\> \</head\>

\<body\> \<h1\>Hola!\</h1\>

\</body\> \</html\>

  \| 34 of 73

# Paso 1: Crear el módulo Go

  Ti p Go utiliza módulos para gestionar las dependencias y la organización del proyecto.

1\. Crea un directorio para tu proyecto:

mkdir servidor- go cd servidor- go

2\. Inicializa el módulo Go:

go mod init \[ejemplo.com/servidor- go\]

Puedes reemplazar ejemplo.com/servidor- go con tu propio nombre de módulo

Esto creará un archivo go.mod en tu directorio.

  \| 35 of 73

# Mini servidor HTTP (1): main.go

package main

import ("fmt"

"net/http" )

func main() { // 1. Define el contenido HTML

htmlContent :=\`\<!DOCTYPE html\>\<html\>\<head\>

\<title\>Página\</title\>\</head\>\<body\> \>\</html\>\` \<h1\>Hola!\</h1\>\</body

// 2. Registra un manejador (handler) para la ruta raíz"/" http.HandleFunc("/",f unc(w http.ResponseWriter, r\*http.Request) { // 3. Establece la cabecera Content-Type"text/html; charset=utf-8") w.Header().Set("Content-Type", // 4. Escribe el HTML en la respuesta fmt.Fprint(w, htmlContent) })

  \| 36 of 73

# Mini servidor HTTP (2): main.go

// 5. Define el puerto y muestra un mensaje port :=":8080" fmt.Printf("Servidor escuchando en http://localhost%s\\n", port)

// 6. Inicia el servidor HTTP

err := http.ListenAndServe(port, nil) if err != nil { fmt.Printf("Error al iniciar el servidor: %s\\n", err) } }

  \| 37 of 73

# Ejecutar el servidor

  N ote

Ahora que tenemos el módulo y el código, podemos ejecutarlo o compilarlo:

Desde tu terminal, en el directorio servidor- go, ejecuta:

go run .

Si todo va bien, deberías ver un mensaje como:

Servidor escuchando en http://localhost:8080

  \| 38 of 73

# Probar con curl

curl es una herramienta de línea de comandos para transferir datos con URLs.

  Ti p Es excelente para probar y automatizar pruebas de servidores Web.

Abre otra terminal (deja el servidor corriendo en la primera) y ejecuta:

curl -v http://localhost:8080

\*H ost localhost:8080 was resolved.

\*IPv 6: ::1

\*IPv4: 127. 0.0.1

\*Tr ying \[::1\]:8080... \* Connected to localhost (::1) port 8080 \* using HTTP/1.x \> GET / HTTP/1.1

\> Host: localhost:8080

\> User-Agent: curl/8.13.0 \> Accept:\*/\* \>

\*R equest completely sent off \< HTTP/1.1 200 OK

\< Content-Type: text/html; charset=utf-8 \< Date: Sat, 24 May 2025 23:40:06 GMT \< Content-Length: 96 \<

\<!DOCTYPE html\>\<html\>\<head\>

\<title\>Página\</title\>\</head\>\<body\> \* Connection #0 to host localhost left intact

\>\</html\>⏎ \<h1\>Hola!\</h1\>\</body

  \| 39 of 73

# Extensión - Servir Archivos Estáticos (1)

1\. Crea un directorio para el contenido HTML:

mkdir static

2\. Crea static/index.html :

\<!DOCTYPE html\>

\<html\>

\<head\>\<title\>Inicio\</title\>\</head\>

\<body\> \<h1\>Bienvenido al Servidor Estático\</h1\>

\<p\>Esta es la página principal.\</p\> \<a href=" /otra.html"\>Ir a otra página\</a\>

\</body\> \</html\>

  \| 40 of 73

# Extensión - Servir Archivos Estáticos (2)

3\. Crea static/otra.html :

\<!DOCTYPE html\>

\<html\>

\<head\>\<title\>Otra Página\</title\>\</head\> \<body\> \<h1\>Esta es Otra Página\</h1\> \<a href="/"\>Volver al inicio\</a\>

\</body\> \</html\>

  \| 41 of 73

# Extensión - Servir Archivos Estáticos (3)

4\. Modifica main.go para usar http.FileServer :

package main

import ("fmt"

"net/http" )

func main() { // 1. Define el directorio que contiene los archivos estáticos. staticDir :="./static"

// 2. Crea un manejador (handler) de servidor de archivos. // http.Dir convierte la ruta del directorio en un sistema de archivos HTTP. // http.FileServer crea un manejador que sirve archivos desde ese sistema. // ¡Automáticamente sirve index.html para directorios! fil eServer := http.FileServer(http.Dir(staticDir))

  \| 42 of 73

# Extensión - Servir Archivos Estáticos (4)

// 3. Registra el manejador para que atienda todas las peticiones ("/"). // Usamos http.Handle porque fileServer es un http.Handler. h ttp.Handle("/",fil eServer)

// 4. Define el puerto y muestra un mensaje. port :=":8080" fm t.Printf("Servidor ESTÁTICO escuchando en http://localhost%s\\n", fm t.Printf("Sirviendo archivos desde: %s\\n",

// 5. Inicia el servidor.

err := http.ListenAndServe(port, nil) if err != nil { fm t.Printf("Error al iniciar el servidor: %s\\n", } }

port) staticDir)

err)

  \| 43 of 73

# HTTP en Go: La Magia de ListenAndServe

Es la forma más común de iniciar un servidor HTTP en Go. Su firma es:

func ListenAndServe(addr string, handler http.Handler) error

addr: La dirección y puerto donde escuchar (ej: :8080 o 127.0.0.1:8080 ).

handler: El objeto que maneja las peticiones. Si es nil , usa http.DefaultServeMux .

Resumiendo:

1\. Se enlaza a la dirección TCP especificada.

2\. Entra en un bucle infinito para aceptar conexiones entrantes.

3\. Sirve las peticiones HTTP que llegan por esas conexiones.

4\. Es bloqueante (normalmente se ejecuta en su propia gorutina o como la última acción de main).

  \| 44 of 73

# El bucle de Escucha

Internamente, ListenAndServe crea un listener TCP y entra en un bucle for:

ListenAndServe

Invocado

Listener.Accept() espera y bloquea hasta que un cliente se conecta.

Cuando llega una conexión, Accept() la devuelve y el bucle continúa para esperar la siguiente.

Manejar Conexión

Nueva Conexión Crea Listener Bucle Infinito TCP en addr Listener.Accept() Error Retorna Error

  \| 45 of 73

# El bucle de Escucha

Internamente, ListenAndServe crea un listener TCP y entra en un bucle for:

ListenAndServe

Invocado

Listener.Accept() espera y bloquea hasta que un cliente se conecta.

Cuando llega una conexión, Accept() la devuelve y el bucle continúa para esperar la siguiente.

  W arning

¿cómo maneja múltiples conexiones al mismo tiempo si Accept() bloquea?

Manejar Conexión

Nueva Conexión Crea Listener Bucle Infinito TCP en addr Listener.Accept() Error Retorna Error

  \| 45 of 73

¡Gorutinas al Rescate!

Aquí es donde brilla la concurrencia de Go:

Por cada conexión aceptada, ListenAndServe inicia una nueva gorutina.

La creación de una gorutina es muy barata en Go.

El bucle principal no espera a que una petición termine:

lanza la gorutina y vuelve a esperar ( Accept ).

Esto permite aceptar miles de conexiones y manejarlas concurrentemente .

  \| 46 of 73

Bucle (Listen) Conexión 1 Conexión 2 Gorutina 1 Gorutina 2

Accept()

go serve(Conn1)

Accept()

go serve(Conn2)

El bucle vuelve a Accept() INMEDIATAMENTE después de lanzar la gorutina.

Bucle (Listen) Conexión 1 Conexión 2 Gorutina 1 Gorutina 2

par \[Goroutine1:\]

Lee Petición

Llama al Handler

Escribe Respuesta

\[Goroutine2:\]

Lee Petición

Llama al Handler

Escribe Respuesta

  \| 47 of 73

# Gorutina (serve) maneja la conexión TCP

Lee la petición HTTP del cliente.

Crea los objetos http.Request y http.ResponseWriter .

Determina qué http.Handler usar (el que pasaste a ListenAndServe o DefaultServeMux ).

Llama al método ServeHTTP del handler, pasándole Request y ResponseWriter .

Espera a que tu handler escriba la respuesta y la Envía al cliente.

Si la conexión es keep-alive, vuelve al paso 1. Si no, cierra y la gorutina termina.

Llamar Escribir Crear Request/ Buscar Handler ¿Tu Código? Handler.ServeHTTP Response Respuesta

Cerrar Conexión Leer Petición No ¿Keep-Alive? y Terminar

Sí

  \| 48 of 73

# El Modelo de Concurrencia

Modelo:"Una gorutina por conexión", o más precisamente, por serve loop .

Simpleza: No necesitas gestionar hilos o pools manualmente. Go lo hace por ti .

Escalabilidad: Go puede manejar cientos de miles de gorutinas eficientemente, lo que permite a un servidor net/http manejar muchas conexiones concurrentes con bajo overhead.

Paralelismo: Si ejecuta en múltiples núcleos de CPU, Go distribuirá estas gorutinas entre ellos, logrando paralelismo real .

  \| 49 of 73

# La concurrencia es poderosa, pero

Estado Compartido : Si tus handlers acceden a datos compartidos (variables globales, bases de datos, cachés), debes usar mecanismos de sincronización:

sync.Mutex o canales para evitar race conditions

¡Cada handler se ejecuta en su propia gorutina!

Manejo de Pánicos : Si un handler entra en pánico (panic), por defecto sólo esa gorutina morirá:

La conexión se cerrará abruptamente, pero el servidor seguirá funcionando:

Es buena práctica usar middleware para recuperar pánicos y loguearlos o devolver un error 500.

Límites de Recursos : Aunque las gorutinas son baratas, no son gratuitas:

Un número extremadamente alto puede agotar la memoria o los descriptores de archivo.

  \| 50 of 73

Implementar límites si es necesario (aunque raramente para servidores web Resutmípicoes)n.

http.ListenAndServe es una forma simple y potente de iniciar servidores HTTP en Go.

Su magia reside en el uso masivo y eficiente de gorutinas.

Lanza una nueva gorutina por cada conexión entrante .

Esto proporciona alta concurrencia de forma natural y escalable.

Recuerda proteger el estado compartido cuando escribas tus handlers :

¿Te suena de alguna materia del pasado cuatrimestre?

  \| 51 of 73

# Agregando formularios HTML

Queremos construir un servidor Go que:

1\. Muestre un formulario HTML al usuario.

2\. Reciba los datos enviados (vía POST ) cuando el usuario envíe el formulario.

3\. Genere una nueva página HTML mostrando un mensaje personalizado con los datos recibidos.

El Formulario:

\<form action=" POST"\> ethod=" /login"m \<label for=" user"\>Usuario:\</label\>

text"i user"n ame=" user"\>\<br\>\<br\> d=" \<input type=" \<label for=" pass"\>Clave:\</label\> ame=" d=" \<input type=" password"i pass"n pass"\>\<br\>\<br\> submit"\>Login\</button\> \<button type=" \</form\>

  \| 52 of 73

# El Flujo de Trabajo

Necesitamos manejar dos interacciones principales:

Petición GET / : El

usuario pide la página inicial. El servidor

debe responder con el HTML del

formulario.

Petición POST

/login : El navegador envía los datos del

formulario. El servidor

debe leerlos y responder con la página de bienvenida.

Usuario Navegador ServidorGo

Abrir http://localhost:8080/

GET /

HTML del Formulario

Muestra Formulario

Rellena y Envía Formulario

POST /login (Datos: user=X, pass=Y)

HTML de Bienvenida (¡Hola X!)

Muestra Bienvenida

Usuario Navegador ServidorGo

  \| 53 of 73 package main

import ("fmt"

"net/http" )

// HTML del formulario

const loginForm =\`\<!DOCTYPE html\>\<html\> \<head\>\<title\>Login\</title\>\</head\> \>\<h2\>Login\</h2\>\<form action=" /login"m \<body \<label\>Usuario:\</label\>

text"n ame=" user"\>\<br\> \<input type=" \<label\>Clave:\</label\>

ame=" \<input type=" password"n pass"\>\<br\> submit"\>Login\</button\>\</form\>\</body \<button type="

func main() { // Ruta para mostrar el formulario http.HandleFunc("/", serveForm) // Ruta para procesar el formulario http.HandleFunc("/login",h andleLogin)

Usamos http.HandleFunc para asociar rutas (paths) con funciones específicas ( handlers ).

POST"\> ethod="

\>\</html\>\`

  \| 54 of 73 port :=":8080" fmt.Printf("Servidor con formulario escuchando en http://localhost%s\\n", err := http.ListenAndServe(port, nil) // Inicia el servidor if err != nil { fmt.Printf("Error: %s\\n", } }

Dentro de cada handler, verificamos el método HTTP ( r.Method ) y la ruta ( r.URL.Path ) para asegurarnos de que estamos manejando la petición correcta:

// serveForm: Maneja GET / para mostrar el formulario func serveForm(w http.ResponseWriter, r\*http.Request) { if r.URL.Path !="/"

http.NotFound(w, r) return

} w.Header().Set("Content-Type", fmt.Fprint(w, loginForm) }

port)

err)

\|\| r.Method != http.MethodGet {

"text/html; charset=utf-8")

  \| 55 of 73

// handleLogin: Maneja POST /login para procesar datos func handleLogin(w http.ResponseWriter, r\*http.Request) { if r.Method != http.MethodPost { http.Error(w,"Método no permitido",h return

}

// 1. Parsear los datos del formulario (¡Crucial!) if err := r.ParseForm(); err != nil { http.Error(w,"Error al parsear",h return

}

// 2. Obtener el valor del campo'user' username := r.FormValue("user")

// 3. Generar y enviar la respuesta HTML w.Header().Set("Content-Type", fmt.Fprintf(w,\`\<!DOCTYPE html\>\<html\>\<head\> \<title\>Bienvenido\</title\>\</head\> \<body \<p\>Recibimos tus datos.\</p\> \<a href=" }

ttp.StatusMethodNotAllowed)

ttp.StatusBadRequest)

"text/html; charset=utf-8")

\>\<h1\>¡Hola, %s!\</h1\> \>\</html\>\`, /"\>Volver\</a\>\</body username)

  \| 56 of 73

Manos a la obra! 1. Utiliza las herramientas de

depuración del navegador para inspeccionar el formulario y ver cómo se envían los datos.

2\. Utiliza curl para enviar una petición POST al servidor con los datos del formulario y verifica la respuesta.

3\. Modifica el programa para utilizar el método GET y repite el punto 1.

4\. Modifica el programa para que valide los datos del formulario y devuelva error 400 si el campo user o pass está vacío.

5\. Modifica el programa para que use http.FileServer y reduzcas la cantidad de HTML en el código Go.   \| 57 of 73

<image redacted: 319x480px, 319x480pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

# Entonces nos ponemos a

programar Web apps?

  \| 58 of 73

# Entonces nos ponemos a

programar Web apps?   Ti p

Primero tenemos qué pensar un poco qué es una Web app, qué partes tiene y cómo se comunican esas partes entre sí.

  \| 58 of 73

# Modelo Three-Tier

Presentación : interactúa con el usuario para mostrar información y capturar entradas.

Lógica del Negocio (o capa de aplicación o capa intermedia): procesa datos, implementa reglas de negocio, realiza cálculos y coordina operaciones.

Datos : responsable de almacenar, recuperar y gestionar los datos de la aplicación. Incluye las base de datos, sistemas de archivos y otros almacenamientos persistentes.

  Im portant Las flechas son invocaciones: llamadas a funciones locales, peticiones HTTP, mensajes asíncronos, u otros tipos de comunicación.

Navegador Web

Backend

Capa de Presentación

Capa de Lógica de Negocio

Sistemas externos

Capa de Datos BusinessServer

Bases de Datos Servidor de archivos

  \| 59 of 73

# Ventajas del Modelo Three-Tier

Modularidad: Cada capa es independiente y puede desarrollarse, modificarse y escalarse por separado.

Reusabilidad: Los componentes de la capa de lógica de negocio pueden ser reutilizados por diferentes interfaces de presentación.

Mantenibilidad: Los cambios en una capa tienen un impacto mínimo en las otras capas, facilitando el mantenimiento y las actualizaciones.

Escalabilidad: Cada capa puede escalarse horizontal o verticalmente según las necesidades. Por ejemplo, se pueden añadir más servidores para la capa de lógica de negocio si aumenta la carga de procesamiento.

Seguridad: La separación de capas permite implementar medidas de seguridad específicas para cada una. La capa de datos, por ejemplo, puede estar protegida por firewalls y controles de acceso.

  \| 60 of 73

# Server-side vs Client-side

Server-side

El servidor genera el HTML final y lo envía al navegador.

El navegador sólo lo muestra.

Ejemplos: Go + templates (templ), PHP, JSP.

Más simple de depurar y mejor para SEO .

Client-side

El servidor envía datos (ej: JSON) y el navegador (JavaScript) construye la interfaz.

Ejemplos: React, Vue, SPA.

Mayor interactividad , pero más complejo .

Client-side

Pide datos Navegador Servidor API JS renderiza Datos JSON

Server-side

Pide HTML Servidor Navegador genera HTML HTML listo

  \| 61 of 73

# Presentación

R enderizado de la interfaz de usuario: Mostrar la información al usuario en un

formato visualmente atractivo y fácil de entender (páginas web, interfaces gráficas, etc.).

Captura de la entrada del usuario: Recibir y validar la información que el usuario introduce a través de formularios, botones, etc.

Gestión de la sesión del usuario: Mantener el estado de la interacción del usuario

con la aplicación (por ejemplo, el carrito de compras, el usuario autenticado).

R edirección y navegación: Guiar al usuario a través de las diferentes secciones y funcionalidades de la aplicación.

Pr esentación de datos: Formatear y mostrar los datos recibidos de la capa de lógica de negocio.

  \| 62 of 73

Cómo se hace la presentación?

La base de la presentación es el HTML. Pero no es lo único:

HTML : Estructura el contenido de las páginas web.

CSS : Define el estilo y la presentación visual de las páginas web.

JavaScript : Añade interactividad y dinamismo al lado del cliente (navegador).

  \| 63 of 73

Cómo se hace la presentación?

La base de la presentación es el HTML. Pero no es lo único:

HTML : Estructura el contenido de las páginas web.

CSS : Define el estilo y la presentación visual de las páginas web.

JavaScript : Añade interactividad y dinamismo al lado del cliente (navegador).

  N ote

Veamos un ejemplo de cómo se puede cambiar el color de un texto en una página web usando HTML, CSS y JavaScript:

  \| 63 of 73

\<!DOCTYPE html\>

\<html\>

\<head\>

\<title\>Cambiar Color\</title\>

\<style\> #miT exto { color: black; } \</style\> \</head\>

\<body\>

miTexto"\>Este es un texto.\</p\> \<p id=" \<button onclick=" cambiarColor()"\>Cambiar Color\</button\>

\<script\> f unction cambiarColor() {

v ar texto = document.getElementById("miTexto"); texto.style.background ="yellow"; } \</script\>

\</body\> \</html\>

  \| 64 of 73

# Difícil lograr sólo con HTML, CSS y JS

Bibliotecas y frameworks JavaScript:

: biblioteca minimalista para crear aplicaciones web interactivas sin JS complejo.

: para construir interfaces de usuario dinámicas, reactivas y complejas .

: framework progresivo para construir interfaces de usuario.

Plantillas del lado del servidor (Server-Side Templating Engines):

: Un framework para construir sitios web rápidos y optimizados.

: lenguaje de scripting del lado del servidor con capacidades de templating.

: framework de React para aplicaciones web y sitios estáticos.

T emplating engines como Templ (Go), Jinja2 (Python), Handlebars (NodeJS), etc.

  \| 65 of 73

# Capa de Lógica del Negocio

Es la que se encarga de procesar la información y aplicar las reglas de negocio de la aplicación.

Actúa como intermediaria entre la capa de presentación y la capa de datos.

Tiene las siguientes responsabilidades principales:

Validación de datos : asegurar que los datos recibidos son correctos y completos.

Procesamiento de datos : cálculos, transformaciones y manipulaciones.

Implementación de reglas de negocio : calcular precios, aplicar descuentos, verificar inventario.

Orquestación de servicios : interacción con otras capas (presentación y datos) y posiblemente con otros servicios externos.

  \| 66 of 73

# Capa de Lógica del Negocio

Orquestación de servicios : interacción con otras capas (presentación y datos) y posiblemente con otros servicios externos.

Gestión de la seguridad y autorización : asegurar que sólo usuarios autorizados puedan realizar ciertas acciones.

Manejo de transacciones : asegurar que las operaciones de datos se completen correctamente o se reviertan en caso de error.

  Ti p Generalmente stateless en memoria, statefull en la persistencia: no guarda estado en RAM entre peticiones.

  \| 67 of 73

# Casi cualquier lenguaje de programación

Mantenibilidad : código limpio, modular y fácil de entender.

Rendimiento : uso eficiente de recursos para escalar si la aplicación crece (tamaño y usuarios).

Seguridad : protección contra ataques comunes (inyección SQL, XSS, CSRF, etc.).

Monitoreo y registro : capacidad de rastrear errores y eventos importantes.

Posibilidad de pruebas unitarias y de integración : para asegurar que la lógica funciona correctamente y se pueda probar de forma aislada.

Capacidad de escalar horizontalmente : permitir que la aplicación maneje más tráfico y usuarios agregando más instancias del servicio.

  \| 68 of 73

# Casi cualquier lenguaje de programación

Mantenibilidad : código limpio, modular y fácil de entender.

Rendimiento : uso eficiente de recursos para escalar si la aplicación crece (tamaño y usuarios).

Seguridad : protección contra ataques comunes (inyección SQL, XSS, CSRF, etc.).

Monitoreo y registro : capacidad de rastrear errores y eventos importantes.

Posibilidad de pruebas unitarias y de integración : para asegurar que la lógica funciona correctamente y se pueda probar de forma aislada.

Capacidad de escalar horizontalmente : permitir que la aplicación maneje más tráfico y usuarios agregando más instancias del servicio.

  Caution

Algunas tecnologías tienden a consumir muchos recursos, lo que puede elevar los costos en servidores y limitar la escalabilidad.

  \| 68 of 73

# Exponer la lógica

A veces, la lógica del negocio se expone a otros sistemas/tecnologías:

REST (Representational State Transfer): Un estilo arquitectónico para construir servicios web escalable.

GraphQL: Un lenguaje de consulta para APIs. Útil para obtener datos de forma flexible.

Mensajes asíncronos: Usar colas de mensajes (ej: RabbitMQ, Kafka) para comunicar eventos entre sistemas.

  W arning El tener varias UIs para el mismo backend podría motivar a exponer la lógica de negocio como un servicio REST o GraphQL.

  \| 69 of 73

# Capa de datos

Almacenamiento persistente de datos : Guardar la información de la aplicación de forma segura y duradera.

Recuperación de datos : Proporcionar mecanismos para acceder a los datos almacenados.

Gestión de la base de datos : Realizar operaciones de creación, lectura, actualización y eliminación (CRUD) de los datos.

Mantenimiento de la integridad de los datos : Asegurar la consistencia y validez de la información almacenada.

Optimización del acceso a los datos : Mejorar la eficiencia de las consultas y las operaciones de lectura/escritura.

Implementación de mecanismos de seguridad : Controlar el acceso a los datos y protegerlos contra accesos no autorizados.

  \| 70 of 73

# Casi cualquier motor de bases de datos

Usualmente se usan bases de datos relacionales (SQL):

SQL directo escrito por los desarrolladores:

Gap semántico objetos/tablas

Alta dependencia del RDBMS

Control, pero con alto costo de desarrollo y mantenimiento

Query Builders para generar SQL automáticamente y Structs mapeados a tablas:

ayudan en la generación de SQL usando estructuras de datos en el lenguaje de programación.

M apeadores objeto-relacional:

Objetos -\> Tablas (generación de esquema), Tablas -\> Objetos (ingeniería reversa), OQL (Object Query Language) en vez de SQL.

Independencia RDBMS, caching.

  \| 71 of 73

# Casi cualquier motor de bases de datos

No relacionales (NoSQL): ideales para documentos estructurados, datos jerárquicos o grandes volúmenes de datos no estructurados:

: almacena documentos JSON.

: almacena documentos JSON con un enfoque en la replicación y sincronización.

r edis: almacena pares clave-valor en memoria, ideal para cachés y datos temporales.

: similar a redis, pero con características adicionales como replicación y persistencia mejorada.

  \| 72 of 73

Qué tiene de diferente el desarrollo Web?

Plataformas heterogéneas en los clientes y servidores

Escalabilidad y variabilidad en número de usuarios

Seguridad

Tiempos de desarrollo cortos y muchos cambios

Tecnologías en evolución constante

Separación de roles diseñador, programador, administrador de sistemas

Calidad de servicio y experiencia de usuario

Observabilidad y monitoreo

Costos de operación y mantenimiento

Tendencias: automatización, operaciones desconectadas, …

  \| 73 of 73