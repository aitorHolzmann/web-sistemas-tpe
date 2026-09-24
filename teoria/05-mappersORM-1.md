# Pr

# ogramación Web

# El nivel de datos (2)

Dr. Alejandro Zunino & Dr. Alfredo Teyseyre alejandro.zunino@isistan.unicen.edu.ar, alfredo.teyseyre@isistan.unicen.edu.ar

  \| 1 of 56

<image redacted: 72x72px, 72x72pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

# Contenidos

1\. Query Builders y mappers

2\. Objetivos de los Query Builders/Mappers

3\. Ventajas vs. Escribir SQL a Mano

4\. Ejemplo: sqlc - mapper y generador de código para Go

5\. Object Relational Mappers (ORM)

6\. Ejemplo de funcionamiento de un ORM

7\. Mapeo O-R y Modelo E-R

8\. ORMs Populares por Lenguaje

9\. Ejemplo: GORM - El ORM amigable para Go

10\. ¿Qué es GORM?

  \| 2 of 56

<image redacted: 479x479px, 479x479pt, ~72dpi, JPG, DEVICE_RGB, 32bpp>

# Query Builders y mappers

Alternativas para simplificar la interacción con bases de datos.

  \| 3 of 56

¿Qué son los Query Builders y mappers?

Son bibliotecas y herramientas que facilitan la interacción con bases de

datos, pero no son una única categoría.

Un Query Builder construye SQL y sus parámetros de manera programática.

Un mapper convierte filas y columnas en structs Go, y puede generar funciones tipadas para ejecutar consultas.

Un ORM agrega una abstracción mayor: mapea entidades y relaciones, y normalmente también construye SQL.

  N ote

Un Query Builder no necesariamente mapea resultados, y un mapper no necesariamente construye consultas. Conviene verificar qué funciones ofrece cada herramienta.

Aplicación

Query Builder o mapper

Mapeo a structs Construcción de Query

Structs Go SQL

Base de Datos

Resultados

  \| 4 of 56

# Objetivos de los Query Builders/Mappers

Abstracción de SQL: Reducir el código repetitivo y, en algunos builders, ocultar parte de los dialectos específicos de cada base de datos.

Seguridad: Ayudar a prevenir vulnerabilidades de inyección SQL mediante el uso de placeholders y parámetros.

Productividad: Agilizar el desarrollo al proporcionar una forma más intuitiva y programática de interactuar con la base de datos.

Mantenibilidad: Facilitar la refactorización y el mantenimiento del código relacionado con la base de datos.

Tipado: Algunas herramientas generan código tipado basado en el esquema de la base de datos, mejorando la seguridad y la detección temprana de errores.

  \| 5 of 56

# Ventajas vs. Escribir SQL a Mano

Característica Query Builders/Mappers SQL a Mano

Suele parametrizar valores, pero hay Seguridad que revisar la API

Desarrollo más rápido, sintaxis Productividad concisa

Refactorización más sencilla, código Mantenibilidad organizado

Portabilidad Abstracción de dialectos (a veces)

Generación de código tipado (a Tipado veces)

Código más legible para operaciones Legibilidad comunes

Pruebas Más fácil de aislar la construcción; Unitarias igual conviene probar contra una BD

Riesgo de inyección SQL si no se parametriza

Requiere escribir y depurar cada consulta

Difícil de refactorizar y mantener

Dependencia del dialecto específico

No hay verificación de tipos en tiempo de compilación

Consultas SQL pueden ser largas y complejas

Requiere configurar un RDBMS

  \| 6 of 56

# Ejemplo: sqlc - mapper y generador de código

# para Go

sqlc es un compilador que genera código Go fuertemente tipado a partir de SQL.

En esta clasificación lo consideramos un mapper: genera el código que ejecuta las consultas y mapea filas a structs Go.

No es un Query Builder: las consultas se escriben en SQL, en vez de construirse mediante una API de Go.

Permite a los desarrolladores:

1\. escribir consultas SQL en un archivo independiente

2\. generar el código Go correspondiente, que puede ser comprobado en tiempo de compilación para evitar errores de ejecución.

  N ote

Evita la necesidad de escribir código redundante y frágil para interactuar con la base de datos.

  \| 7 of 56

# Funcionamiento de sqlc

1\. Escribir Archivos SQL: Se definen las consultas SQL directamente en archivos .sql .

Pueden incluir placeholders ( ? , $1 , etc.) para parámetros.

2\. Archivo de Configuración sqlc.yaml : especifica la base de datos, la ubicación de los archivos SQL, el paquete Go donde se generará el código, y otras opciones.

3\. Ejecutar sqlc generate : analiza los archivos SQL y el archivo de configuración.

4\. sqlc genera código Go que incluye:

Funciones para ejecutar cada consulta SQL de forma segura y tipada.

Definiciones de structs Go que corresponden a las filas de los resultados de las consultas.

Una interfaz Querier opcional para facilitar la prueba y el mocking.

  \| 8 of 56

# Uso de sqlc: configuración

version:"2"

sql:

\- o"mysql", engine:"postgresql"#"sqlite" queries:"./db/ queries/" schema:"./db/schema/"

gen:

go: package:"db" out:"./db/s qlc/" emit\_json\_tags: true

engine : tipo de base de datos

queries : directorio que contiene los archivos .sql con las consultas

schema : directorio con los archivos

SQL de definición del esquema de la BD

gen.go.package : nombre del paquete Go para el código generado

gen.go.out : ruta donde se guardará el código

gen.go.emit\_json\_tags : indica si se deben generar tags json en los

structs

  \| 9 of 56

# Uso de sqlc: definición del esquema

db/schema/schema.sql

CREATE TABLE users ( i d SERIAL PRIMARY KEY,

n ame VARCHAR(255) NOT NULL, email VARCHAR(255) UNIQUE NOT NULL, created\_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT\_TIMESTAMP );

  \| 10 of 56

# Uso de sqlc: definición de consultas

db/queries/users.sql

\-- n ame: GetUser :one

SELECT id, name, email, created\_at FROM users

WHERE id = $1;

\-- n ame: ListUsers :many SELECT id, name, email, created\_at FROM users

ORDER BY name;

\-- n ame: CreateUser :one

INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id, name, email, created\_at;

\-- n ame: UpdateUser :exec UPDATE users

SET name = $2, email = $3 WHERE id = $1;

\-- n ame: DeleteUser :exec

DELETE FROM users

WHERE id = $1;

Cada consulta tiene -- name que sqlc utiliza para generar la función Go correspondiente.

Se utilizan placeholders ($1, $2, etc.).

:one espera una fila (también con INSERT / UPDATE … RETURNING ); si no hay ninguna, la función devuelve un error.

:many devuelve múltiples filas.

:exec no devuelve filas.

  \| 11 of 56

# Uso de sqlc: código generado

db/sqlc/users.go

package db

import ("context"

"time"

)

n const createUser =\`-- ame: CreateUser :one

INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id, name, email, created\_at\`

type CreateUserParams struct { Name string\`json:"name"\` Email string\`json:"email"\` }

type CreateUserRow struct {

| ID              | int64            | json:"id"\`    |
|-----------------|------------------|----------------|
| Name            | string           | json:"name"\`  |
| Email CreatedAt | string time.Time | json:"email"\` |

}

  \| 12 of 56

# Uso de sqlc: código generado

func (q\*Queries) CreateUser(ctx context.Context, arg CreateUserParams) (CreateUserRow, error) { row := q.db.QueryRowContext(ctx, createUser, arg.Name, arg.Email) var i CreateUserRow

err := row.Scan(&i.ID, &i.Name, &i.Email, &i.CreatedAt) return i, err } // otras funciones generadas para GetUser, ListUsers...

Los parámetros de las funciones se definen en structs (ej., CreateUserParams ).

Los resultados de las consultas se mapean a structs Go (ej., CreateUserRow ).

Todo el acceso a la base de datos se realiza a través de la instancia de Queries .

  \| 13 of 56

# Uso de sqlc: juntando las partes

package main import (

...

\_"github.com/jackc/pgx/v5/stdlib"

sqlc"your\_project/db/sqlc" // generado por sqlc )

func main() { connStr :="user=u password= p dbname=d" db, err := sql.Open("pgx", connStr) if err != nil { log.Fatalf("failed to connect to DB: %v", err) } defer db.Close() queries := sqlc.New(db) ctx := context.Background()

createdUser, err := queries.CreateUser(ctx, // Create sqlc.CreateUserParams{ Name:"John Doe",

Email:"john.doe@example.com",

})

  \| 14 of 56

# Uso de sqlc: juntando las partes

if err != nil { log.Fatalf("failed to create user: %v", err) } fmt.Printf("Created user: %+v\\n", createdUser)

user, err := queries.GetUser(ctx, createdUser.ID) // Read One if err != nil { log.Fatalf("failed to get user: %v", err) } fmt.Printf("Retrieved user: %+v\\n", user)

users, err := queries.ListUsers(ctx) // Read Many if err != nil { log.Fatalf("failed to list users: %v", err) } fmt.Printf("All users: %+v\\n", users)

err = queries.UpdateUser(ctx, sqlc.UpdateUserParams{ // Update ID: createdUser.ID, Name:"Johnny Doe",

Email:"johnny.doe@example.com",

})

  \| 15 of 56

# Uso de sqlc: juntando las partes

if err != nil { log.Fatalf("failed to update user: %v", err) }

fmt.Println("User updated successfully")

updatedUser, err := queries.GetUser(ctx, createdUser.ID) if err != nil { log.Fatalf("failed to get updated user: %v", err) }

fmt.Printf("Updated user: %+v\\n", updatedUser)

  \| 16 of 56

# Uso de sqlc: juntando las partes

err = queries.DeleteUser(ctx, createdUser.ID) // Delete if err != nil { log.Fatalf("failed to delete user: %v", err) }

fmt.Println("User deleted successfully")

\_, err = queries.GetUser(ctx, createdUser.ID) if err == sql.ErrNoRows { fmt.Println("User not found after deletion") } else if err != nil { log.Fatalf("failed to get user after deletion: %v", err) } }

  \| 17 of 56

# Uso de sqlc: Instalación y Compilación

1\. Configurar la Base de Datos ( XYZ es la contraseña que usarás para el servidor de BD):

docker run --name some-

2\. Conéctate a la base de datos y crea una base de datos/usuario (si es necesario):

docker exec -it some-

\# Dentro de psql CREATE DATABASE \<yourdb\>; CREATE USER youruser WITH PASSWORD'\<yourpassword\>'; GRANT ALL PRIVILEGES ON DATABASE \<yourdb\> TO \<youruser\>; \\c \<yourdb\> GRANT ALL ON SCHEMA public TO \<youruser\>; \\q

3\. Asegúrate de que la cadena de conexión en el código Go (connStr) coincida con tus credenciales.

4\. Prueba: docker exec -it some-

\<yourdb\>

postgres -e POSTGRES\_PASSWORD=XYZ - p 5432:5432 -d docker.io/postgres

postgres psql -h localhost -U postgres

postgres psql -h localhost -U youruser -d

  \| 18 of 56

# Uso de sqlc: Instalación y Compilación

5\. Crear la Estructura de Directorios y Archivos:

my\_sqlc\_project/ ├── db/

| │ ├── | queries/               |
|-------|------------------------|
| │ │   | └── users.sql          |
| │ └── | schema/                |
| │ ├── | └── schema.sql main.go |

└── sqlc.yaml

  \| 19 of 56

# Uso de sqlc: Instalación y Compilación

6\. Inicializa el módulo Go:

cd my\_sqlc\_project go mod init my\_sqlc\_project

7\. Instalar sqlc :

go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

8\. Generar el Código Go:

sqlc generate

Esto creará el directorio ./db/sqlc y los archivos Go generados (como users.go , models.go , querier.go ).

  \| 20 of 56

# Uso de sqlc: Instalación y Compilación

9\. Instalar Dependencias de Go:

go get github.com/jackc/pgx/v5/stdlib

10\. Ejecutar el Programa:

go run main.go

  N ote

Debes ver la salida de las operaciones CRUD realizadas en la base de datos.

  \| 21 of 56

# Ejemplo: Query Builder con goqu

Un Query Builder construye SQL y parámetros, pero no necesariamente mapea el resultado a structs.

import"github.com/doug-martin/goqu/v9"

type User struct {

| ID } ds := query, // | int64 Name string Email string goqu.From( name email Select( Where(goqu.Ex{ Order(goqu.I( Limit(10) args, err := ds.Prepared(true).ToSQL() name email OM query: SELECT = SC LIMIT 10           |
|----------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| //                   | WHERE ( $1) ORDER BY                                                                                                                                                                           |
| // Luego confiable.  | args: \[true\] se ejecuta query y se escanean las filas con database/sql o sqlx . El builder parametriza los valores, pero nombres de tablas y columnas deben provenir de código   \| 22 of 56 |

# Conclusiones Query Builders/Mappers

Los query builders y mappers pueden ofrecer ventajas significativas sobre la escritura manual de acceso a datos:

seguridad

productividad

mantenibilidad

tipado

sqlc , con su enfoque en la generación de código a partir de SQL, proporciona una forma robusta y segura de interactuar con BDs:

robusto ante cambios de esquema: sqlc verify analiza esquema y consultas para detectar inconsistencias.

  \| 23 of 56

# Object Relational Mappers

# (ORM)

Alternativas para abstraer la interacción con bases de datos.

  \| 24 of 56

¿Qué es un Object Relational Mapper?

Abstraen buena parte de la base de datos, permitiendo interactuar con ella utilizando entidades y sus relaciones. El ORM se encarga de generar el SQL subyacente, aunque normalmente permite usar SQL cuando la abstracción no alcanza.

Características Principales:

Abstracción Alta: No es necesario escribir SQL para la mayoría de las operaciones CRUD.

Manejo de Relaciones: Facilitan la gestión de relaciones entre tablas (uno-a-uno, uno a-muchos, muchos-a-muchos).

Mayor Nivel de Abstracción: El desarrollador interactúa con objetos Go, el ORM traduce estas interacciones a SQL.

Características Adicionales: Suelen incluir funcionalidades como migraciones de esquemas, transacciones, caché, etc.

Menos Control Directo sobre el SQL: Puede ser más difícil optimizar consultas complejas o utilizar características específicas de la base de datos.

  \| 25 of 56

# Query Builders y mappers vs. ORMs

Ambas son tecnologías que facilitan la interacción con bases de datos, pero difieren en su nivel de abstracción y funcionalidades.

Aplicación Go

Query Builder o mapper

SQL Generado (más Mapeo básico a structs directo)

ORM Base de Datos

SQL Generado (más Mapeo complejo a objetos abstracto) (relaciones)

  \| 26 of 56

Característica Query Builders y mappers ORMs

Abstracción Menor, más cercano a SQL Mayor, interacción a través de entidades

Control SQL Mayor, permite escribir SQL específico Menor, el ORM genera el SQL

Mapeo básico a structs, gestión manual de Relaciones relaciones

Curva de Aprendizaje Puede requerir mayor conocimiento de SQL Aprender la API del ORM

Mejor control para optimizaciones Rendimiento específicas

Flexibilidad Mayor flexibilidad para consultas complejas Menos flexibilidad para consultas muy específicas

Mantenibilidad Mantenimiento del SQL y Go por separado

Prevención de inyección SQL con Seguridad placeholders

Características Generalmente menos características Adicionales integradas

Ejemplos en Go goqu, sqlf, jet, sqlc gorm, ent

Gestión automática de relaciones entre entidades

Puede generar SQL ineficiente en casos complejos

Mantenimiento del modelo de objetos, el ORM gestiona SQL

Generalmente seguros, pero depende de configuración

Suelen incluir migraciones, transacciones, etc.

  \| 27 of 56

Query Builders, mappers u ORMs?

# Query Builders/mappers

Necesitas control sobre el SQL.

Estás trabajando con consultas complejas o específicas de la base de datos.

El rendimiento es crítico y necesitas optimizar el SQL.

Prefieres una abstracción ligera sobre la BD.

Estás integrando con una BD existente.

En proyectos donde la complejidad de un ORM completo no está justificada.

# Object Relational Mapper

# (ORM)

Se busca productividad y rapidez.

Necesitas una gestión sencilla de las relaciones entre entidades.

Prefieres objetos en lugar de tablas.

La aplicación tiene muchas operaciones CRUD.

Equipo con menos experiencia en SQL.

Las características del ORM

(migraciones, transacciones, caching) son importantes.

  \| 28 of 56

# El ORM ideal

Mapeo Declarativo: Definición clara y concisa de cómo las clases se mapean a tablas.

Soporte para Relaciones: Manejo eficiente de relaciones uno a uno, uno a muchos y muchos a muchos.

Consultas Expresivas: Un lenguaje de consulta que sea fácil de usar y traducir a SQL.

Gestión de Transacciones:

Mecanismos robustos para asegurar la integridad de los datos.

Lazy Loading y Eager Loading: Estrategias para optimizar la carga de datos relacionados.

Manejo de Concurrencia: Mecanismos para prevenir problemas cuando múltiples usuarios acceden a los mismos datos.

Migraciones: Herramientas para gestionar la evolución del esquema de la base de datos.

Extensibilidad: Arquitectura que permita la personalización y la adición de nuevas funcionalidades.

Rendimiento: Diseño que minimice la sobrecarga y genere SQL eficiente.

Soporte para Múltiples BDs: Abstracción de las diferencias entre

distintas SBDRs.

  \| 29 of 56

# Ejemplo de funcionamiento de un ORM

**User**

\+ID int

\+Name string +Email string +Profile Profile

\+Posts \[\]Post

1

1 has has

\*

1

**Post**

**Profile**

\+ID int +ID int +UserID int +UserID int +Title string +Bio string +Body string

  \| 30 of 56

# Mapeo O-R y Modelo E-R

# Relación 1:1

USER

int id PK

string name

string email

has

PROFILE

int id PK

int user\_id FK

string biography

// Modelo de Objeto: Usuario type User struct { ID int \`orm:"primaryKey"\` Name string\`orm:"column:name"\` Email string\`orm:"unique"\` Profile Profile\`orm:"one2one"\` // Relación 1:1 con Profile

\`orm: Posts \[\]Post"one2many"\` }

// Modelo de Objeto: Perfil type Profile struct { ID int \`orm:"primaryKey"\` UserID int \`orm: // FK a tabla Users"unique,foreignKey:Users"\` Bio string\`orm:"column:biography"\` }

  N ote

son ilustrativos: cada ORM define los suyos (por Los tags orm:"..." ejemplo, GORM usa gorm:"...").

  \| 31 of 56

# Mapeo O-R y Modelo E-R

# Relación 1:N

USER

int id PK

string name

string email

has

POST

int id PK

int user\_id FK

string title

string content

// Modelo de Objeto: Usuario type User struct { ID int \`orm:"primaryKey"\` Name string\`orm:"column:name"\` Email string\`orm:"unique"\` Profile Profile\`orm:"one2one"\`

\`orm: // Relación 1:N con Post Posts \[\]Post"one2many"\` }

// Modelo de Objeto: Publicación type Post struct { ID int \`orm:"primaryKey"\` UserID int \`orm: // FK a tabla Users"foreignKey:Users"\` Title string\`orm:"column:title"\` Body string\`orm:"column:content"\` }

  \| 32 of 56

# Relationship Traversal

El ORM ideal permite navegar fácilmente entre objetos relacionados:

es una instancia de GORM // Asumiendo que'db' var user models.User

err := db.First(&user, 1).Error if err != nil { // Manejar error }

// Acceder al perfil del usuario profile := user.Profile fmt.Println(profile.Bio)

// Acceder a las publicaciones del usuario for \_, post := range user.Posts { fmt.Println(post.Title) }

  \| 33 of 56

# Carga explícita de relaciones

No todos los ORM implementan carga perezosa. En GORM, el acceso a un campo de relación no ejecuta automáticamente otra consulta: la relación debe cargarse de forma explícita.

// Cargar el usuario y luego sus posts explícitamente user, err := db.FindUserByID(1) if err != nil { // Manejar error } fmt.Println(user.Name)

err = db.Model(&user).Association("Posts").Find(&user.Posts) if err != nil { // Manejar error } for \_, post := range user.Posts { fmt.Println(post.Title) }

  Ti p La carga explícita evita traer relaciones que la operación no necesita

  W arning Si se hace dentro de un recorrido, puede llevar al problema de N+1 lecturas.

  \| 34 of 56

# Problema de N+1 Lecturas

Ocurre cuando se realiza:

1\. una consulta inicial para obtener una lista de N objetos,

2\. N consultas adicionales para cargar los datos relacionados de cada objeto.

// Una consulta para usuarios y otra para todos sus posts var users \[\]models.User err := db.Find(&users).Error // 1 consulta: N usuarios if err != nil { // Manejar error }

for \_, user := range users { fmt.Println(user.Name) // Si cada usuario carga aquí su relación, se agrega 1 consulta por usuario for \_, post := range user.Posts { fmt.Println(post.Title) } }

  Ti p El problema no es exclusivo de los ORM: también puede ocurrir con SQL directo. Se evita planificando la carga, por ejemplo con un\`JOIN\`o con una consulta que recupere todas las relaciones.

  \| 35 of 56

# Solución: Eager Loading

GORM permite especificar qué relaciones cargar junto con la consulta principal mediante Preload , evitando las N consultas adicionales.

var users \[\]models.User err := db.Preload("Posts").Find(&users).Error // 1 consulta de usuarios + 1 consulta de posts if err != nil { // Manejar error }

for \_, user := range users { fmt.Println(user.Name) for \_, post := range user.Posts { // Acceso directo a los datos cargados fmt.Println(post.Title) } }

  \| 36 of 56

# Consultas en GORM

Un ORM debe proporcionar una forma intuitiva y segura de realizar consultas.

// Consultar por ID var user User

err := db.First(&user, 1).Error if err != nil { // Manejar error }

// Consultar por nombre var users \[\]User"Alice").Find(&users).Error err = db.Where("name = ?", if err != nil { // Manejar error }

// Consultas más complejas 5,"New York"). err = db.Where("age \> ? AND city = ?",2 Order("name ASC").Limit(10).Find(&users).Error if err != nil { // Manejar error }

  \| 37 of 56

# Transacciones

El ORM debe facilitar la gestión de transacciones en el lenguaje:

err := db.Transaction(func(tx\*gorm.DB) error { ail:"bob@example.com"} user := models.User{Name:"Bob",Em if err := tx.Create(&user).Error; err != nil { r eturn err

}

profile := models.Profile{UserID: user.ID, Bio:"Backend developer"} if err := tx.Create(&profile).Error; err != nil { return err

}

// Si todo va bien, la transacción se commite automáticamente return nil

})

if err != nil { fmt.Println("Error durante la transacción:", err) } else { fmt.Println("Transacción exitosa") }

  \| 38 of 56

# Updates/Deletes en Cascada

La propagación depende de una restricción de clave foránea configurada en el esquema, no solamente del struct de Go. En GORM puede declararse con constraint:OnDelete:CASCADE y verificarse en la migración.

type User struct { gorm.Model N ame string // La FK se configura para eliminar los posts asociados. P osts \[\]Post\`gorm:"constraint:OnDelete:CASCADE;"\` }

// AutoMigrate debe crear la restricción en la BD. db.AutoMigrate(&User{}, &Post{}) result := db.Delete(&User{}, 1) if result.Error != nil { // Manejar error } // Los posts asociados al usuario con ID 1 también se eliminarán automáticamente

  \| 39 of 56

# Desafíos para un ORM Ideal

Complejidad del Mapeo: Traducir la riqueza de los modelos de objetos a las limitaciones del modelo relacional puede ser complejo, especialmente con relaciones avanzadas y herencia (que no se mapea directamente).

Rendimiento: Generar SQL eficiente para todas las posibles consultas y relaciones es un desafío constante. La abstracción del ORM puede introducir sobrecarga si no se diseña cuidadosamente.

Diferencias entre Bases de Datos: Cada sistema de base de datos tiene sus propias peculiaridades en cuanto a sintaxis SQL, tipos de datos y funcionalidades. Abstraer estas diferencias sin perder rendimiento ni funcionalidad es difícil.

Manejo de Casos Edge: Lidiar con escenarios complejos como concurrencia optimista/pesimista, herencia de tablas, y tipos de datos específicos de la base de datos requiere una implementación sofisticada.

  \| 40 of 56

# Desafíos para un ORM Ideal

Facilidad de Uso vs. Flexibilidad: Encontrar el equilibrio entre una API fácil de usar para casos comunes y la flexibilidad necesaria para consultas y operaciones más avanzadas es crucial.

Mantenimiento y Evolución: A medida que Go evoluciona y surgen nuevas funcionalidades en las bases de datos, mantener el ORM actualizado y eficiente requiere un esfuerzo continuo.

Curva de Aprendizaje: Un ORM potente puede tener una curva de aprendizaje pronunciada debido a la cantidad de conceptos y configuraciones involucradas. La documentación clara y los ejemplos son esenciales.

Pruebas Exhaustivas: Asegurar que el ORM funcione correctamente con diferentes bases de datos y en diversos escenarios requiere pruebas rigurosas.

  \| 41 of 56

# ORMs Populares por Lenguaje

Go

GORM: amigable para desarrolladores, que apunta a ser rico en características y fácil de usar.

Ent: impulsado por esquemas para Go con generación de código estática, seguridad de tipos y navegación de grafos sencilla.

Bob: generador de código ORM para Go que produce código idiomático adaptado a tu esquema de base de datos.

TypeScript

TypeORM: soporta Active Record y Data Mapper, con un enfoque en la experiencia del desarrollador.

Prisma: ORM de próxima generación que facilita el acceso a la base de datos con seguridad de tipos y migración automática.

  \| 42 of 56

# ORMs Populares por Lenguaje

Python

SQLAlchemy: El kit de herramientas SQL y ORM de Python que proporciona potencia y flexibilidad para interactuar con bases de datos.

Django ORM: El ORM integrado con el framework Django, conocido por su facilidad de uso y su estrecha integración con el ecosistema Django.

Peewee: Un ORM pequeño y expresivo para Python, fácil de aprender y usar, ideal para proyectos más pequeños o cuando se prefiere un enfoque minimalista.

Java

Hibernate: maduro y ampliamente utilizado para la plataforma Java, que proporciona una potente capa de persistencia para aplicaciones empresariales.

Spring Data JPA: Parte del ecosistema Spring, simplifica el acceso a datos con repositorios y abstracciones sobre JPA (Java Persistence API).

  \| 43 of 56

# Ejemplo: GORM - El ORM

# amigable para Go

Interactuando con la base de datos a través de objetos Go.

  \| 44 of 56

¿Qué es GORM?

GORM (The Go ORM) es un ORM completo para Go.

Abstrae la base de datos, permitiendo trabajar con objetos Go en lugar de SQL:

Mapeo Objeto-Relacional: Convierte structs de Go a tablas de base de datos.

Consultas Flexibles: API fluida para construir consultas complejas.

Migraciones Automáticas: Puede crear y actualizar el esquema de la BD automáticamente.

Hooks: Permite ejecutar código antes o después de operaciones (Create, Save, Update, Delete).

Soporte para Relaciones: Maneja relaciones has one , has many , belongs to ,

many to many .

  N ote

GORM busca maximizar la productividad del desarrollador al reducir la necesidad de escribir SQL.

  \| 45 of 56

# Uso de GORM: Definición del Modelo

En lugar de un esquema SQL, se define un struct en Go.

GORM utiliza tags en los campos del struct para configurar el mapeo a la tabla.

gorm.Model es un struct que GORM provee, e incluye campos comunes como ID ,

CreatedAt , UpdatedAt , DeletedAt .

models/user.go

package models

import"gorm.io/gorm"

type User struct { gorm.Model // Incluye ID, CreatedAt, UpdatedAt, DeletedAt

| Name  | string                                         |
|-------|------------------------------------------------|
| Email | // Crea una restricción UNIQUE en la BD string |

}

  \| 46 of 56

# Uso de GORM: Conexión y Migración

para el motor de BD GORM necesita una cadena de conexión y un"dialecto" específico.

db.AutoMigrate() es una de las características más potentes:

Analiza el struct del modelo ( User ).

Crea la tabla correspondiente si no existe.

Agrega columnas, índices o restricciones faltantes si el struct cambia.

  \| 47 of 56

# Uso de GORM: Conexión y Migración

package main

import ("gorm.io/gorm""gorm.io/driver/postgres""your\_project/models" )

func main() { dsn :="host=localhost user=

db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}) if err != nil { panic("failed to connect database") }

// Migrar el esquema (en un proyecto real, revisar el error). if err := db.AutoMigrate(&models.User{}); err != nil { panic(err) }

// ... aquí va el código para las operaciones CRUD }

youruser password= yourpassword dbname= yourdb port=5432 sslmode=disable

  \| 48 of 56

# Uso de GORM: Operaciones CRUD

# Create

Cea una instancia del struct y se pasa a db.Create() . GORM genera el INSERT :

// Crear un usuario

ail:"jinzhu@cn.com"} user := models.User{Name:"Jinzhu",Em result := db.Create(&user) // Pasa el puntero del objeto

if result.Error != nil { l og.Fatalf("failed to create user: %v",r esult.Error) } fmt.Printf("Usuario creado con ID: %d", user.ID)

# Read (One)

Se usa db.First() para recuperar un registro por su clave primaria.

var readUser models.User

// Obtener el primer registro ordenado por clave primaria if err := db.First(&readUser, user.ID).Error; err != nil { l og.Fatalf("failed to read user: %v", err) } fmt.Printf("Usuario leído: %+v",r eadUser)

  \| 49 of 56

# Uso de GORM: Operaciones CRUD

# Read (Many)

db.Find() recupera múltiples registros que coinciden con una condición.

var users \[\]models.User if err := db.Find(&users).Error; err != nil { l og.Fatalf("failed to list users: %v", err) } fmt.Printf("Todos los usuarios: %+v", users)

  \| 50 of 56

# Uso de GORM: Operaciones CRUD

# Update

Se puede actualizar un registro con db.Model().Updates() .

if err := db.Model(&readUser).Update("Name", l og.Fatalf("failed to update user: %v", } fmt.Printf("Usuario actualizado: %+v",r

# Delete

db.Delete() elimina un registro. GORM por defecto usa borrado lógico (soft delete) si el modelo incluye gorm.DeletedAt .

if err := db.Delete(&models.User{}, readUser.ID).Error; err != nil { l og.Fatalf("failed to delete user: %v", } fmt.Println("Usuario eliminado")

"Jinzhu 2").Error; err != nil {

err)

eadUser)

err)

  \| 51 of 56

# Uso de GORM: Relaciones

Las relaciones se declaran con tags en el struct:

type User struct { gorm.Model Name string Email string\`gorm:"unique"\` Posts \[\]Post\`gorm:"foreignKey:UserID"\` // 1:N (has many) }

type Post struct { gorm.Model UserID uint

Title string Body string }

db.AutoMigrate(&models.User{}, &models.Post{}) crea ambas tablas y la clave foránea.

  \| 52 of 56

# Uso de GORM: Relaciones

# Evitando N+1: Preload

var users \[\]models.User db.Preload("Posts").Find(&users) // 1 consulta + 1 por la relación (NO N consultas)

for \_, user := range users { fmt.Println(user.Name) for \_, post := range user.Posts { // Ya están cargados fmt.Println(post.Title) } }

  Ti p

Preload es la carga ansiosa (eager loading) de GORM: resuelve el problema de N+1 visto antes.

  \| 53 of 56

# Uso de GORM: Juntando las partes

package main import ("fmt"

""log"gorm.io/driver/postgres""gorm.io/gorm""your\_project/models" ) func main() { dsn :="host=localhost user=

db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}) if err != nil { log.Fatal("failed to connect database") }

if err := db.AutoMigrate(&models.User{}); err != nil { log.Fatal(err) }

// Create

user := models.User{Name:"Gopher",Em if err := db.Create(&user).Error; err != nil { log.Fatal(err) } fmt.Printf("Created user: %+v",

youruser password= yourpassword dbname= yourdb port=5432 sslmode=disable

ail:"gopher@example.com"}

user)

  \| 54 of 56

# Uso de GORM: Juntando las partes

// Read

var readUser models.User

if err := db.First(&readUser, user.ID).Error; err != nil { log.Fatal(err) } fmt.Printf("Retrieved user: %+v",r

// Update if err := db.Model(&readUser).Update("Name", log.Fatal(err) } fmt.Printf("Updated user: %+v",r

// Delete

if err := db.Delete(&models.User{}, readUser.ID).Error; err != nil { log.Fatal(err) } fmt.Println("User deleted successfully") }

eadUser)

"Go Gopher").Error; err != nil {

eadUser)

  \| 55 of 56

# Conclusiones sobre GORM

GORM representa un alto nivel de abstracción, lo que trae consigo ventajas y desventajas claras:

Productividad Alta: Ideal para prototipos y aplicaciones con operaciones CRUD. La migración automática y la API fluida aceleran el desarrollo.

Menor Necesidad de SQL: Los programadores pueden centrarse en la lógica de negocio (no SQL).

Curva de Aprendizaje: Requiere aprender la API y los conceptos de GORM, que pueden ser complejos (ej. Preload , asociaciones, transacciones y hooks).

Menos Control y Rendimiento Potencialmente Menor: El SQL generado por el ORM puede no ser tan optimizado como el SQL escrito a mano. Para consultas muy complejas, puede ser un obstáculo.

  Im portant

GORM es una excelente opción cuando la velocidad de desarrollo es prioritaria sobre el control granular del SQL. Para sistemas de alto rendimiento con consultas complejas, sqlc o SQL directo pueden ser más adecuados.

  \| 56 of 56