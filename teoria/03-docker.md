# Pr

# ogramación Web

# Docker

## Dr. Alejandro Zunino & Dr. Alfredo Teyseyre

## alejandro.zunino@isistan.unicen.edu.ar, alfredo.teyseyre@isistan.unicen.edu.ar

  \| 1 of 27

<image redacted: 72x72px, 72x72pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

# Contenidos

# Docker Fundamentos

¿Qué es Docker?

## Imágenes Docker / Dockerfile -

## Construyendo Imágenes

## Contenedores Docker /

## Volúmenes y Redes

# Docker Compose

## Docker Compose / Comandos

## Docker Compose

## Buenas Prácticas

  \| 2 of 27

# Virtual Machines vs Docker

# Virtual Machines (VM)

## Emulan hardware completo

## Incluyen sistema operativo invitado

## Consumen más recursos

## Arranque más lento

## Aislamiento fuerte

# Docker (Contenedores)

## Comparten el kernel del host

## Solo incluyen la app y dependencias

## Consumen menos recursos

## Arranque casi instantáneo

## Aislamiento a nivel de proceso

  \| 3 of 27

<image redacted: 128x153px, 128x152pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

<image redacted: 128x150px, 128x150pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

# ¿Por qué Docker en Desarrollo Web? (1/2)

# El problema:"

# En mi máquina funciona"

## En el desarrollo web, reproducir el entorno es difícil:

## Distintas versiones de lenguajes, librerías y bases de datos

## Diferencias entre desarrollo, testing y producción

## Configuraciones ocultas y dependencias no documentadas

## "Works on my machine"

## al pasar el código a otro/a o a producción

  W arning Sin aislamiento, el entorno se vuelve frágil e irreproducible.

  \| 4 of 27

# ¿Por qué Docker en Desarrollo Web? (2/2)

# La solución: entornos reproducibles

## Docker aborda estos problemas:

## Portabilidad: la app y sus dependencias viajan juntas

## Consistencia: el mismo entorno en dev, test y prod

## Aislamiento: cada servicio (web, DB, cache) en su contenedor

## Onboarding rápido: clonar repo + docker compose up

  Ti p En web, donde coexisten múltiples servicios (API, DB, frontend), Docker simplifica levantar todo el stack de forma idéntica en cualquier máquina.

  \| 5 of 27

¿Qué es Docker?

## Docker es una plataforma que permite empaquetar aplicaciones y sus dependencias en

contenedores portables y ligeros.

# Conceptos Clave

## Contenedor: Unidad estándar de software que empaqueta código y dependencias

## Imagen: Plantilla de solo lectura para crear contenedores

## Dockerfile: Archivo de texto con instrucciones para construir imágenes

## Docker Hub: Registro público de imágenes Docker

  Ti p Los contenedores comparten el kernel del sistema operativo, haciéndolos más eficientes que las máquinas virtuales tradicionales.

  \| 6 of 27

# Arquitectura Docker

# Componentes Principales

## Docker Client: Interfaz de línea de comandos

## Docker Daemon: Servicio que gestiona contenedores

## Docker Registry: Almacén de imágenes (Docker Hub)

## Docker Objects: Imágenes, contenedores, redes, volúmenes

\# Verificar instalación

docker --version

docker info

\# Listar contenedores

docker ps docker ps -a

  \| 7 of 27

# Arquitectura Docker

Docker Host

Docker Client

CLI Commands

API calls

Docker Daemon

dockerd

push/pull manages creates/runs manages manages

Docker Docker Docker Docker

Containers Volumes Networks Images

Docker Registry Docker Hub

  \| 8 of 27

# Imágenes Docker

## Las imágenes son plantillas

## inmutables que contienen:

## Sistema operativo base

## Aplicación y dependencias

## Configuración predeterminada

# Comandos Básicos

\# Descargar imagen docker pull nginx:latest # Listar imágenes docker images # Eliminar imagen docker rmi nginx:latest # Construir imagen docker build -t mi-app .

# Layers (Capas)

## Las imágenes están compuestas por capas:

Base OS Layer

Runtime Layer

Application Layer

Configuration Layer

  N ota

Las capas se reutilizan entre imágenes, ahorrando espacio de almacenamiento.

  \| 9 of 27

# Dockerfile - Construyendo Imágenes

  D ockerfile El Dockerfile es un archivo de texto que contiene instrucciones para construir una imagen Docker

# Ejemplo: Aplicación Go (single-stage)

\# Imagen base (incluye el toolchain de Go) FROM golang:1.27-alpine # Directorio de trabajo WORKDIR /app # Copiar dependencias primero (aprovecha la caché de capas) COPY go.mod go.sum ./ RUN go mod download # Copiar el código fuente COPY . .

\# Compilar el binario RUN go build -o /app/api ./cmd/api # Exponer puerto EXPOSE 8080

\# Comando por defecto CMD \["/app/api"\]

Cada instrucción genera una capa de la imagen. La imagen resultante incluye todo el toolchain de Go (varios cientos de MB): no es la más eficiente, pero es clara para entender los conceptos de Docker.

  \| 10 of 27

# Instrucciones Dockerfile

## FROM: Imagen base

## WORKDIR: Directorio de trabajo

## COPY/ADD: Copiar archivos

## RUN: Ejecutar comandos

## EXPOSE: Exponer puertos

## ENV: Variables de entorno

ENV PORT=8080

RUN apt- get update && apt- rm-rf /var/lib/apt/lists/\*

  N ote

El gestor de paquetes depende de la imagen base: apt- get en imágenes basadas en Debian/Ubuntu, apk en Alpine.

# Mejores Prácticas

## Usar imágenes base oficiales

## Minimizar capas combinando

## comandos RUN

## Limpiar archivos temporales

## Usar .dockerignore

## Ejecutar como usuario no

## privilegiado

get install - y curl && \\

  W arning Evitar instalar paquetes innecesarios para reducir tamaño.

  \| 11 of 27

# Dockerfile para Go

# Build multi-stage: compila y empaqueta en una imagen liviana

\# Etapa 1: compilación FROM golang:1.27-alpine AS builder WORKDIR /app COPY go.mod go.sum ./ RUN go mod download COPY . .

RUN CGO\_ENABLED=0 go build -o /app/api ./cmd/api

\# Etapa 2: imagen final (solo el binario) FROM alpine:3.24 RUN adduser -D -u 1000 app USER app COPY --from=builder /app/api /usr/local/bin/api EXPOSE 8080

CMD \["api"\]

CGO\_ENABLED=0 genera un binario estático que no depende de librerías del sistema.

La imagen final solo contiene el binario: mucho más pequeña.

USER app evita ejecutar con privilegios de root.

  Ti p

  \| 12 of 27

Las imágenes multi-stage reducen drásticamente el tamaño final (de cientos de MB a decenas).

# .dockeri

# gnore

# Go Node.js

.git bin/

.env

\*.lo g

node\_modules .git .env

\*.lo g

  Ti p Reduce el contexto de build y evita copiar secretos o archivos innecesarios a la imagen.

  \| 13 of 27

# Contenedores Docker

  Contenedores Los Contenedores son instancias ejecutables de imágenes Docker.

# Ciclo de Vida del Contenedor

\# Crear y ejecutar contenedor docker run -d --name mi-nginx - # Listar contenedores en ejecución docker ps # Ver todos los contenedores

docker ps -a # Parar contenedor

docker stop mi-nginx # Iniciar contenedor parado docker start mi-nginx # Eliminar contenedor

docker rm mi-nginx

# Opciones útiles de docker

# run

\# Ejecutar en modo interactivo docker run -it ubuntu:latest /bin/bash

\# Montar volumen

docker run -v /host/path:/container/path nginx # Variables de entorno

docker run -e ENV\_VAR=value mi-app p 8080:80 nginx:latest # Ver detalles de un contenedor

docker inspect mi-nginx # Ver logs de un contenedor docker logs mi-nginx # Copiar archivos entre host y contenedor docker cp archivo.txt mi-nginx:/ruta/destino # Reiniciar un contenedor

docker restart mi-nginx # Cambiar nombre a un contenedor

docker rename mi-nginx nginx- prod # Ejecutar un comando dentro del contenedor docker exec -it mi-nginx /bin/bash

  \| 14 of 27

# Volúmenes y Redes

  V olúmenes Los Volúmenes permiten persistir datos fuera del contenedor.

## Volume: Gestionado por Docker

## Bind Mount: Directorio del host

## tmpfs: Almacenamiento temporal

\# Crear volumen

docker volume create mi-volumen

\# Usar volumen

docker run -v mi-volumen:/data nginx # Listar volúmenes

docker volume ls

\# Bind mount (directorio host) docker run -v /home/user/data:/app/data nginx

  R edes Las Redes permiten conectar y aislar contenedores entre sí, permitiendo una comunicación segura y controlada

\# Crear red personalizada docker network create mi-red

\# Conectar contenedor a red

docker run --network mi-red --name web nginx # Listar redes

docker network ls

\# Inspeccionar red docker network inspect mi-red

Container A

mi-red

mi-red Container B

mi-red

Container C

  \| 15 of 27

# Docker Compose

  D ocker Compose Docker Compose es una herramienta para definir y ejecutar aplicaciones Docker multi-contenedor usando archivos YAML (YAML Ain't Markup Language)

# Ventajas

## Definir servicios complejos de manera declarativa

## Gestionar múltiples contenedores como una aplicación

## Facilitar el desarrollo y testing

## Reproducir entornos de manera consistente

  \| 16 of 27

# Archivo docker-compose.yml

# Ejemplo: Aplicación Web

# con Base de Datos

services:

w eb:

build: .

ports:

\-"3000:3000"

environment:

\-N ODE\_ENV=development -DB\_HOST=database

depends\_on:

\- database

v olumes:

\-.: /app

\- /app/node\_modules

database:

im age: postgres:18 environment:

\-P OSTGRES\_DB=myapp -P OSTGRES\_USER=user

\-P OSTGRES\_PASSWORD= password

v olumes:

\- postgres\_data:/var/lib/postgresql/data ports:

\-"5432:5432"

volumes:

postgres\_data:

  Ti p: Variables de Entorno y Seguridad Definir variables sensibles (como contraseñas) en un archivo .env y referenciarlas en el docker compose.yml usando la sintaxis ${VAR} . Así se evita exponer datos sensibles en el archivo de configuración y se facilita la portabilidad del entorno.

  \| 17 of 27

# Comandos Docker Compose

# Comandos Básicos Comandos Avanzados

\# Iniciar servicios

docker compose up # Iniciar en segundo plano docker compose up -d # Construir imágenes docker compose build # Ver servicios ejecutándose docker compose ps # Ver logs docker compose logs docker compose logs web # Parar servicios

docker compose stop # Parar y eliminar docker compose down

\# Escalar servicios

docker compose up --scale web=3 # Ejecutar comando en servicio docker compose exec web bash # Ver configuración generada docker compose config # Recrear contenedores

docker compose up --force-recreate # Eliminar volúmenes

docker compose down -v # Construir sin caché

docker compose build --no-cache

  Ti p Usa\`docker compose up -d\`para desarrollo y \`docker compose logs -f\`para seguir los logs.

  \| 18 of 27

# Ejemplo Completo: API REST en Go +

# PostgreSQL

## Configuración completa para una

## aplicación API REST desarrollada en Go

que utiliza una base de datos PostgreSQL.

## Este ejemplo muestra cómo definir ambos

## servicios y su comunicación a través de

Docker Compose.

services:

api: build: ./api container\_name: go-api ports:

\-"8080:8080"

environment:

\-DB\_HOST=database

\-DB\_PORT=5432

\-DB\_USER= postgres -DB\_PASSWORD= postgres -DB\_NAME=apirest

depends\_on:

\- database

v olumes:

\-. /api:/app database:

im age: postgres:18 container\_name: postgres-db environment:

\-P OSTGRES\_DB=apirest -P OSTGRES\_USER= postgres -P OSTGRES\_PASSWORD= postgres ports:

\-"5432:5432"

v olumes:

\- pgdata:/var/lib/postgresql/data

volumes:

pgdata:

  Ti p Coloca el código fuente de tu API Go en la carpeta ./api y asegúrate de que la aplicación lea las variables de entorno para conectarse a la base de datos.

  \| 19 of 27

# Variables de Entorno

# Archivo .env

MYSQL\_ROOT\_PASSWORD=supersecret MYSQL\_DATABASE=myapp API\_KEY=abc123

NODE\_ENV= production

services:

w eb:

im age: node:26 environment:

\-N ODE\_ENV=${NODE\_ENV} -API\_KEY=${API\_KEY}

database:

im age: mysql:9 environment:

\-MY SQL\_ROOT\_PASSWORD=${MYSQL\_ROOT\_PASSWORD} -MY SQL\_DATABASE=${MYSQL\_DATABASE}

  W arning Nunca commitees archivos .env con credenciales reales al control de versiones.

  \| 20 of 27

# Redes Personalizadas en Compose

## Docker Compose puede crear redes

## personalizadas para mejorar el

aislamiento y la comunicación.

services:

fr ontend:

im age: nginx n etworks:

\-fr ontend-network

ports:

\-"80:80"

backend:

im age: node:26 n etworks:

\-fr ontend-network

\- backend-network

environment:

\-DB\_HOST=database

database:

im age: postgres:18

n etworks:

\- backend-network

environment:

\-P OSTGRES\_DB=myapp

networks:

fr ontend-network:

driver: bridge backend-network:

driver: bridge in ternal: true # Solo comunicación interna

## Ai

## slamiento de servicios

## R

## esolución DNS automática

## Control de acceso granular

  \| 21 of 27

# Profiles

## Permite activar servicios

## específicos:

services:

w eb:

im age: nginx database:

im age: postgres:18

r edis:

im age: redis:8 profiles:

\- caching m onitoring: im age: prometheus profiles: -m onitoring

\# Solo servicios principales docker compose up # Con profile de caching docker compose -- profile caching up # Múltiples profiles docker compose -- profile caching -- profile monitoring up

  \| 22 of 27

# Overrides

## Personalizar configuración para diferentes entornos:

\# docker-compose.override.yml services:

w eb:

v olumes:

\-.: /app environment:

\-DEB UG=true

database:

ports:

\-"5432:5432"# Ex poner en desarrollo

\# Usa automáticamente override

docker compose up # Archivo específico docker compose -f docker-compose.yml -f docker-compose.prod.yml up

  \| 23 of 27

# Monitoreo y Logs

## Docker y Docker Compose proporcionan

## herramientas para monitorear

aplicaciones.

# Comandos de Monitoreo

\# Ver estadísticas en tiempo real docker stats

\# Logs de contenedor específico docker logs -f container\_name

\# Logs con Docker Compose docker compose logs -f service\_name

\# Últimas 100 líneas

docker compose logs --tail=100 web

\# Logs con timestamp docker compose logs -t web

# Configuración de Logging

services:

w eb:

im age: nginx l ogging: driver: json-file options: m ax-size:"10m"

m ax-file:"3"

app: im age: node:26 l ogging: driver: syslog options: syslog-address:"tcp://log-server:514"

  \| 24 of 27

# Mejores Prácticas

# Desarrollo Producción

\# docker-compose.dev.yml services:

w eb:

build:

context: .

target: development

v olumes:

\-.: /app

\- /app/node\_modules environment:

\-N ODE\_ENV=development ports:

\-"3000:3000"

\-"9229:9229"# D ebug port

\# docker-compose.prod.yml services:

w eb:

im age: myapp:latest r estart: unless-stopped environment:

\-N ODE\_ENV= production deploy: r esources:

limi ts:

m emory: 512M

r eservations:

m emory: 256M

  \| 25 of 27

# Consejos Generales

## Usa multi-stage builds para imágenes optimizadas

## Implementa health checks para servicios críticos

## Gestiona secrets de manera segura

## Usa .dockerignore para optimizar contexto de build

## Documenta tu configuración

services:

w eb:

h ealthcheck:

"curl", test: \["CMD", in terval: 30s

timeout: 10s

r etries: 3

start\_period: 40s

"-f""

,"http://localhost:3000/health \]

  \| 26 of 27

# Resumen y Recursos

# Conceptos Clave

# Aprendidos

## Docker: Contenedorización de

## aplicaciones

## Dockerfile: Construcción de imágenes

## personalizadas

## Docker Compose: Orquestación de

## aplicaciones multi-contenedor

## Volúmenes y Redes: Persistencia y

## comunicación

## Mejores Prácticas: Desarrollo y

## producción

# Recursos Útiles

## Documentación Oficial Docker

## Docker Compose Reference

## Docker Hub - Registro de imágenes

## Dockerfile Best Practices

## Docker Labs

  Ti p La práctica es clave: experimenta con diferentes configuraciones y casos de uso para dominar Docker y Docker Compose.

  \| 27 of 27