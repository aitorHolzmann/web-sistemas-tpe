# Pr

# ogramación Web

# Herramientas de Desarrollo

Dr. Alejandro Zunino & Dr. Alfredo Teyseyre alejandro.zunino@isistan.unicen.edu.ar, alfredo.teyseyre@isistan.unicen.edu.ar

  \| 1 of 16

<image redacted: 72x72px, 72x72pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

# Contenidos

1\. Instalación de Atlas y Air

2\. Atlas: Esquema deseado

3\. Air: Recompilación automática para Go

4\. Automatizando con make

5\. Flujo de trabajo integrado

  \| 2 of 16

<image redacted: 718x479px, 718x479pt, ~72dpi, JPG, DEVICE_RGB, 32bpp>

# Herramientas de Desarrollo

El desarrollo de aplicaciones web modernas no se limita a escribir código.

Necesitamos herramientas para gestionar el esquema de la base de datos y mejorar el flujo de trabajo.

En esta clase veremos dos herramientas muy útiles para el desarrollo con Go:

Atlas: para generar y aplicar migraciones de bases de datos.

Air: para recompilar y reiniciar automáticamente la aplicación.

Usaremos un Makefile para orquestar estas tareas y las herramientas vistas en clases anteriores.

  \| 3 of 16

# Instalación de Atlas y Air

Atlas (esquemas y migraciones):

curl -sSf https://atlasgo.sh \| sh # o, con Go: go install ariga.io/atlas/cmd/atlas@latest

Air (recompilación y reinicio automático):

go install github.com/air-verse/air@latest

  Ti p

Verifica la instalación con atlas version y air -v . Si instalas herramientas con go install , asegúrate de que $(go env GOPATH)/bin esté en tu PATH .

  \| 4 of 16

# Atlas: Esquema deseado

Atlas es una herramienta independiente del lenguaje de la aplicación para gestionar esquemas y migraciones.

Permite definir el esquema de la base de datos como código (en HCL, SQL o desde un ORM).

En este ejemplo usamos PostgreSQL. La sintaxis SQL y las URLs deben corresponder al motor elegido.

Atlas compara el esquema deseado con el estado actual y genera una migración versionada.

Ejemplo de flujo de trabajo:

1\. Definir el esquema deseado en db/schema/schema.sql :

CREATE TABLE users ( i d SERIAL PRIMARY KEY,

n ame TEXT NOT NULL, email TEXT UNIQUE NOT NULL );

  \| 5 of 16

# Atlas: Generar migraciones

2\. Generar la primera migración:

migrate diff genera archivos SQL; no modifica directamente la base de datos.

atlas migrate diff initial --dir"file://db/migrations""file://db/schema/schema.sql"

3\. Modificar el esquema (por ejemplo, agregar created\_at ):

El archivo de esquema representa el estado deseado completo, no una migración aislada.

CREATE TABLE users ( i d SERIAL PRIMARY KEY,

n ame TEXT NOT NULL, email TEXT UNIQUE NOT NULL, created\_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT\_TIMESTAMP );

\-- to \\

\-- dev-url"docker://postgres/15/dev?search\_path= public"

  \| 6 of 16

# Atlas: Versionar cambios

4\. Generar una nueva migración para el cambio:

atlas migrate diff add\_created\_at --dir"file://db/migrations"

\-- to"file://db/schema/schema.sql"

  N ote

Revisar el SQL generado antes de aplicarlo. Versionar las migraciones junto con\`atlas.sum\`.

\\

\-- dev-url"docker://postgres/15/dev?search\_path= public"

  \| 7 of 16

# Atlas: Aplicar migraciones

Una vez generados y revisados los archivos de migración, podemos aplicarlos a la BD.

1\. Aplicar las migraciones:

Aplica todas las migraciones pendientes en db/migrations a la base de datos de desarrollo:

\\ atlas migrate apply --dir"file://db/migrations"

\-- url"postgres://user:password@localhost:5432/mydb?sslmode=disable"

2\. Verificar el estado de las migraciones:

Podemos ver qué migraciones se aplicaron y cuáles quedaron pendientes:

\\ atlas migrate status --dir"file://db/migrations"

\-- url"postgres://user:password@localhost:5432/mydb?sslmode=disable"

3\. Inspeccionar el esquema actual:

Podemos pedirle a Atlas que inspeccione la base de datos y nos muestre su esquema en formato HCL.

atlas schema inspect --url"postgres://user:password@localhost:5432/mydb?sslmode=disable"

  \| 8 of 16

# Atlas: Verificar el estado

4\. Analizar el esquema:

Atlas puede analizar el esquema en busca de problemas potenciales, como la falta de índices en claves foráneas o el uso de tipos de datos problemáticos.

atlas migrate lint --dir"file://db/migrations"

\-- dev-url"docker://postgres/15/dev?search\_path=

  Ti p

En desarrollo puede formar parte del arranque. En producción, conviene ejecutarlo como un paso separado del despliegue, evitando que varias réplicas intenten migrar simultáneamente.

\\

\-- latest 1 public"

  \| 9 of 16

# Air: Recompilación automática para Go

Air es una utilidad de línea de comandos para el ciclo de desarrollo de aplicaciones Go:

Observa los cambios en los archivos de código fuente.

Recompila automáticamente.

Reinicia la aplicación.

Air reinicia el proceso de Go; la recarga automática del navegador requiere una herramienta adicional.

Esto acelera significativamente el ciclo de desarrollo.

  \| 10 of 16

# Air: Integración con make

Podemos configurar Air para delegar la construcción en make . Así, Air observa los archivos y make coordina la construcción de la aplicación. Las herramientas de generación utilizadas por el proyecto se integran en el target correspondiente.

.air.toml de e jemplo:

root ="."

tmp\_dir ="tmp" \[build\] cmd ="make build"

entrypoint = \["./tm p/my-app"\] in clude\_ext = \["go","templ","sql"\]"vendor"\] exclude\_dir = \["tmp",

  N ote

Con esta configuración, air ejecutará make build cada vez que se modifique un archivo observado. El target build debe generar el binario en tmp/my-app .

  \| 11 of 16

# Desglosando el .air.toml

root : El directorio raíz del proyecto que Air debe observar ( . es el actual).

tmp\_dir : Directorio temporal para los archivos de Air (se recomienda ignorarlo en git).

\[build\] : Define cómo se construye la aplicación.

cmd : Comando para construir el binario. Delegamos la lógica a nuestro Makefile ( make build ).

entrypoint : Binario que air ejecuta después de una compilación exitosa.

include\_ext : Extensiones de archivo que dispararán una reconstrucción.

exclude\_dir : Directorios a ignorar.

  Ti p

La clave es la integración con make : air observa los archivos y make orquesta la construcción. No conviene omitir la construcción inicial si el binario todavía no existe.

  \| 12 of 16

# Automatizando con make

Un Makefile nos permite definir un conjunto de tareas que podemos ejecutar con un simple comando. Esto es útil para estandarizar y simplificar el flujo de trabajo de desarrollo.

APP\_NAME := my-app DB\_URL := postgres://usuario:password@localhost:5432/mydb?sslmode=disable

.PHONY: all run generate migrate apply status build test clean

all: build

run:

@air

  \| 13 of 16

# Targets del Makefile

\# Ejecuta las tareas de generación configuradas por el proyecto generate: @sqlc generate @templ generate

\# Genera una migración: make migrate name=add\_created\_at migrate: @test -n " $(name)" \|\| (echo "Uso: make migrate name=nombre" && exit 1)

\-- dir"file://db/migrations" atlas migrate diff"$(name)"

\-- dev-url"docker://postgres/15/dev?search\_path="file://db/schema/schema.sql"

\# Aplica las migraciones pendientes apply: atlas migrate apply --dir"file://db/migrations"

\# Muestra el estado de las migraciones status:

atlas migrate status --dir"file://db/migrations"

\-- to \\

public"

\-- url"$(DB\_URL)"

\-- url"$(DB\_URL)"

  \| 14 of 16

# Targets del Makefile

\# Ejecuta las tareas de generación configuradas por el proyecto El @ al inicio de una línea en el Makefile evita generate: @sqlc generate que el comando se imprima en la terminal, @templ generate mostrando solo su salida.

\# Genera una migración: make migrate name=add\_created\_at migrate: @test -n " $(name)" \|\| (echo "Uso: make migrate name=nombre" && exit 1)

\-- dir"file://db/migrations" atlas migrate diff"$(name)"

\-- dev-url"docker://postgres/15/dev?search\_path="file://db/schema/schema.sql"

\# Aplica las migraciones pendientes apply: atlas migrate apply --dir"file://db/migrations"

\# Muestra el estado de las migraciones status:

atlas migrate status --dir"file://db/migrations"

\-- to \\

public"

\-- url"$(DB\_URL)"

\-- url"$(DB\_URL)"

  \| 14 of 16

# Targets del Makefile

\# Construye el binario de la aplicación build: generate @mkdir -p tmp @go build -o tmp/$(APP\_NAME) .

test:

@go test ./...

\# Limpia los artefactos de construcción clean:

@rm -rf tmp

  \| 15 of 16

# Targets del Makefile

\# Construye el binario de la aplicación build: generate @mkdir -p tmp @go build -o tmp/$(APP\_NAME) .

test:

@go test ./...

\# Limpia los artefactos de construcción clean:

@rm -rf tmp

Las dependencias (ej: build: generate ) aseguran que las tareas se ejecuten en el orden correcto.

  \| 15 of 16

# Flujo de trabajo integrado

Modificar el esquema deseado

make migrate name=...

Revisar y versionar la migración + atlas.sum

make apply: Aplicar en desarrollo

make run Air: recompila y reinicia

  \| 16 of 16

# Flujo de trabajo integrado

Modificar el esquema deseado

En producción, aplicar las make migrate name=... migraciones como un paso separado del despliegue y revisar siempre el SQL generado. Revisar y versionar la migración + atlas.sum

make apply: Aplicar en desarrollo

make run Air: recompila y reinicia

  \| 16 of 16