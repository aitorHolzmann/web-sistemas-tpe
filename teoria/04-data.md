# Pr

# ogramación Web

# El nivel de datos

Dr. Alejandro Zunino & Dr. Alfredo Teyseyre alejandro.zunino@isistan.unicen.edu.ar, alfredo.teyseyre@isistan.unicen.edu.ar

  \| 1 of 37

<image redacted: 72x72px, 72x72pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

# Contenidos

1\. Modelo Three-Tier

2\. Funciones principales

3\. Alternativas para el acceso a datos

4\. Bases de datos SQL en el nivel de datos

5\. Enfoque"SQL Directo" en la Capa de Datos

6\."SQL Directo" en Go

7\. Características de"SQL Directo"

8\. Tecnologías que siguen el enfoque"SQL Directo"

9\. Discordancia de impedancia con BD Relacionales

  \| 2 of 37

<image redacted: 479x479px, 479x479pt, ~72dpi, JPG, DEVICE_RGB, 32bpp>

# Modelo Three-Tier

Presentación : interactúa con el usuario para mostrar información y capturar entradas.

Lógica de Negocio (o capa de aplicación o capa intermedia): procesa datos, implementa reglas de negocio, realiza cálculos y coordina operaciones.

Datos : responsable de almacenar, recuperar y gestionar los datos de la aplicación. Incluye las bases de datos, sistemas de archivos y otros almacenamientos persistentes.

  Im portant Las flechas son invocaciones: llamadas a funciones locales, peticiones HTTP, mensajes asíncronos, u otros tipos de comunicación.

Navegador Web

Backend

Capa de Presentación

Capa de Lógica de Negocio

Sistemas externos

Capa de Datos BusinessServer

Bases de Datos Servidor de archivos

  \| 3 of 37

# Funciones principales

Alm acenamiento persistente de datos: Guardar la información de la aplicación de forma segura y duradera.

R ecuperación de datos: Proporcionar mecanismos para acceder a los datos almacenados.

Gestión de la base de datos: Realizar operaciones de creación, lectura, actualización y eliminación (CRUD) de los datos.

M antenimiento de la integridad de los datos: Asegurar la consistencia y validez de la información almacenada.

Optimización del acceso a los datos: Mejorar la eficiencia de las consultas y las operaciones de lectura/escritura.

Im plementación de mecanismos de seguridad: Controlar el acceso a los datos y protegerlos contra accesos no autorizados.

  \| 4 of 37

# Alternativas para el acceso a datos

Usualmente se usan bases de datos relacionales (SQL) , con algún nivel de abstracción para el acceso a datos:

SQL directo escrito por los desarrolladores:

Gap semántico objetos/tablas

Alta dependencia del RDBMS

Control, pero con alto costo de desarrollo y mantenimiento

Query Builders y mappers para reducir el código repetitivo:

los query builders construyen SQL de forma programática;

los mappers convierten filas y columnas en structs Go;

algunas herramientas combinan ambas funciones; sqlc parte de SQL escrito por el desarrollador y genera el código de acceso y mapeo.

  \| 5 of 37

# Alternativas para el acceso a datos

M apeadores objeto-relacional (ORM):

Objetos -\> Tablas (generación de esquema), Tablas -\> Objetos (ingeniería inversa), OQL (Object Query Language) en vez de SQL.

Independencia RDBMS, caching.

No relacionales (NoSQL): ideales para documentos estructurados, datos jerárquicos o grandes volúmenes de datos no estructurados.

  \| 6 of 37

# Bases de datos NoSQL

Ideales para documentos estructurados, datos jerárquicos o grandes volúmenes de datos no estructurados:

Documentales: almacenan documentos (ej: JSON). Ej: MongoDB, CouchDB.

Clave-Valor: pares clave-valor, a menudo en memoria. Ej: Redis, DynamoDB.

Grafos: nodos y relaciones. Ej: Neo4j.

Columnares: optimizadas para grandes volúmenes de datos. Ej: Cassandra, Bigtable.

  N ote

Se eligen según la naturaleza de los datos y las consultas: no existe una"mejor"en general.

  \| 7 of 37

Criterio SQL Directo

Dificultad de Alta (conocimiento Uso profundo de SQL)

Baja (SQL disperso, Mantenibilidad propenso a errores)

Potencialmente la

Performance más alta (control total)

Total (control Control del completo sobre Programador cada consulta)

Query/Struct Mapeadores O/R Mappers

Media (SQL + herramienta de Media a Baja (abstracción del SQL) generación o builder)

Alta (modelo de datos Alta en consultas abstracto, definidas y código refactorización más generado sencilla)

Generalmente Variable (puede buena (SQL generar SQL optimizado ineficiente, requiere optimización) generado)

Alto (el Medio a Bajo (el ORM programador escribe el SQL genera el SQL) base)   \| 8 of 37

Criterio SQL Directo Query/Struct Mappers Mapeadores O/R

Variable: sqlc sigue Baja (SQL Independencia el dialecto elegido; un específico de del RDBMS builder puede cada BD) abstraerlo

Empinada Curva de (necesita Moderada (entender la Aprendizaje experiencia en generación de código) SQL)

Más lento

(escribir y Medio (genera código Desarrollo mantener mucho repetitivo) SQL)

Mayor riesgo de Menor riesgo (facilita Seguridad uso de parámetros) inyecciones SQL

Alta, con límites según el ORM y sus extensiones

Moderada

(comprender los conceptos del ORM)

Más rápido (abstracción y funcionalidades

integradas)

Menor riesgo (generalmente maneja la parametrización)

  \| 9 of 37

Query/Struct Criterio SQL Directo Mappers

Portabilidad Baja (reescribir) Media (modificar) Alta (el ORM lo maneja)

Ideal op. Bueno para mapear Naturaleza de resultados a complejas y los Datos estructuras específicas

Manejo directo de Complejidad de Bueno para consultas consultas Consultas bien definidas complejas

Mapeadores O/R

Ideal para trabajar con objetos y sus relaciones

Puede simplificar consultas complejas con relaciones

  \| 10 of 37

# Bases de datos SQL en el nivel

# de datos

Funciones. Enfoque SQL directo. Query builders. Mapeadores objeto-relacional (ORM). NoSQL.

  \| 11 of 37

# Bases de datos SQL en el nivel de datos

Se accede con objetos o estructuras

Oculta que los datos están almacenados en una RDB

Abstrae el nivel de Lógica de Negocio de cómo se almacenan los datos

Debería ser independiente del RDBMS

Realiza el mapeo objeto-relacional (ORM) o estructura-tabla

Capa de Lógica del Negocio

objetos/estructuras

Capa de Datos

SQL tablas

RDBMS

  \| 12 of 37

# Bases de datos SQL en el nivel de datos

1\. Establecer y Gestionar la Conexión:

Abrir y cerrar conexiones a la base de datos de manera eficiente.

Manejar pools de conexiones para optimizar el rendimiento.

2\. Ejecutar Operaciones CRUD (Crear, Leer, Actualizar, Borrar):

Traducir las solicitudes de la lógica de negocio en consultas SQL.

Ejecutar estas consultas contra la base de datos.

3\. Mapear Resultados:

Convertir los resultados de las consultas SQL (filas y columnas) a estructuras de datos o objetos que la capa de lógica de negocio pueda entender y utilizar (structs en Go, objetos en Java/JS).

  \| 13 of 37

# Bases de datos SQL en el nivel de datos

4\. Manejar Transacciones:

Iniciar, confirmar (commit) o revertir (rollback) transacciones para asegurar la integridad de los datos en operaciones complejas.

5\. Gestionar Errores:

Capturar errores específicos de la base de datos y traducirlos a errores más genéricos o específicos del dominio de la aplicación para la capa de lógica de negocio.

6\. Implementar Lógica de Acceso a Datos Específica:

Encapsular consultas complejas, joins, filtros y ordenamientos para simplificar el acceso a datos para la lógica de negocio.

  \| 14 of 37

# "

# "

# en la

# SQL Directo

# Enfoque

# Capa de Datos

El desarrollador escribe y gestiona directamente las consultas SQL para interactuar con la base de datos.

  \| 15 of 37

# "

# en Go

# SQL Directo"

import ("database/sql"

\_"github.com/go-sql-driver/mysql" // Importa el driver MySQL )

type userRepository struct { db \* sql.DB }

type User struct { ID in t64

N ame string Em ail string }

  \| 16 of 37

# "

# SQL Directo"

func (r\*userRepository) GetUserByID(id int64) (\*User, error) { r ow := r.db.QueryRow("SELECT id, name, email FROM users WHERE id = ?",i user := &User{} err := row.Scan(&user.ID, &user.Name, &user.Email) if err != nil { r eturn nil, err } r eturn user, nil }

# en Go

d)

  \| 17 of 37

# "

# SQL Directo"

func (r\*userRepository) GetUserByID(id int64) (\*User, error) { r ow := r.db.QueryRow("SELECT id, name, email FROM users WHERE id = ?",i user := &User{} err := row.Scan(&user.ID, &user.Name, &user.Email) if err != nil { r eturn nil, err } r eturn user, nil }

  Caution

La forma en que se pasan los parámetros a la consulta es crucial para la seguridad: posibles inyecciones SQL

# en Go

d)

  \| 17 of 37

# Conectando a la base de datos

import ("database/sql"

\_"github.com/go-sql-driver/mysql" // Importa el driver MySQL ) func abrirDB() (\*sql.DB, error) { db, err := sql.Open("mysql","usuario:clave@tcp(localhost:3306)/mi\_app") if err != nil { r eturn nil, err } if err := db.Ping(); err != nil { // Verifica que la BD responde r eturn nil, err } db.SetMaxOpenConns(25) // Configura el tamaño del pool r eturn db, nil }

sql.Open no conecta todavía: crea el pool y valida la configuración.

db.Ping() verifica que la conexión funcione.

El pool se reutiliza entre consultas; se cierra con defer db.Close() al final.

  Ti p El pool de conexiones es clave: abrir y cerrar una conexión por consulta es lento.

  \| 18 of 37

# Previniendo Inyección de SQL

## Consultas Parametrizadas (La forma segura)

El ? en la consulta es un marcador de posición, no un simple reemplazo de texto.

// El driver de la BD y Go se encargan de la seguridad row := r.db.QueryRow("SELECT id, name, email FROM users WHERE id = ?",i d)

1\. Código y Datos por Separado: La consulta SQL ( SELECT... ) se envía a la base de datos separada del valor ( id ).

2\. Pre-compilación: La base de datos analiza la consulta y entiende su estructura antes de recibir los datos.

3\. Ejecución Segura: El valor del parámetro id es tratado siempre como un dato, nunca como código ejecutable. Si id contuviera 1; DROP TABLE users; , la base de datos buscaría un usuario con ese ID literal, sin ejecutar el DROP TABLE .

  \| 19 of 37

# Previniendo Inyección de SQL

## Concatenación de Strings:

## forma insegura y mala práctica

// ¡PELIGRO! ¡NO HACER ESTO! idStr :="1; DROP TABLE users; --" // Input de un atacante + idStr query :="SELECT ... FROM users WHERE id =" r.db.Query(query) // Se ejecuta: SELECT ... FROM users WHERE id = 1; DROP TABLE users; --

En este caso, el código malicioso se inyecta en la consulta y la base de datos lo ejecuta, porque no puede distinguir entre el código original y los datos del atacante.

  Caution

NUNCA construyas consultas concatenando strings con datos del usuario.

  \| 20 of 37

# Previniendo Inyección de SQL

## Concatenación de Strings:

## forma insegura y mala práctica

// ¡PELIGRO! ¡NO HACER ESTO! idStr :="1; DROP TABLE users; --" // Input de un atacante + idStr query :="SELECT ... FROM users WHERE id =" r.db.Query(query) // Se ejecuta: SELECT ... FROM users WHERE id = 1; DROP TABLE users; --

En este caso, el código malicioso se inyecta en la consulta y la base de datos lo ejecuta, porque no puede distinguir entre el código original y los datos del atacante.

  Caution

NUNCA construyas consultas concatenando strings con datos del usuario.

  \| 20 of 37

# Características de"

# SQL Directo"

# Dificultad de Uso: Alta

Requiere un conocimiento profundo y detallado del lenguaje SQL específico de la base de datos que se esté utilizando (MySQL, PostgreSQL, SQL Server, etc.).

Implica entender la sintaxis, las funciones, los tipos de datos y las características avanzadas de cada motor de base de datos.

El desarrollador es responsable de construir cada consulta manualmente, incluyendo la gestión de parámetros y la concatenación segura de cadenas (para evitar inyecciones SQL).

No hay abstracción sobre la estructura de la base de datos:

cualquier cambio en el esquema puede requerir la modificación de numerosas consultas en el código.

  \| 21 of 37

# Características de"

# SQL Directo"

# Mantenibilidad: Baja

El SQL disperso por toda la base de código puede volverse difícil de rastrear y mantener, especialmente en aplicaciones grandes.

Los cambios en el esquema de la base de datos a menudo requieren la revisión y modificación manual de múltiples consultas SQL.

La falta de una estructura coherente para el acceso a datos puede llevar a inconsistencias y errores difíciles de depurar.

La refactorización del código relacionado con la base de datos puede ser compleja y propensa a introducir errores.

  \| 22 of 37

# Características de"

# SQL Directo"

# Performance: Potencialmente la Más Alta

Ofrece el control más granular sobre las consultas SQL, permitiendo a los desarrolladores optimizar cada consulta para casos de uso específicos.

Se pueden escribir consultas altamente especializadas que aprovechen las características únicas del motor de base de datos para obtener el máximo rendimiento.

Evita la sobrecarga potencial que pueden introducir las capas de abstracción de los ORM.

Sin embargo, requiere un conocimiento experto en optimización de SQL por parte del desarrollador para lograr este potencial. Consultas mal escritas pueden resultar en un rendimiento deficiente.

  \| 23 of 37

# Características de"

# SQL Directo"

# Control del Programador: Total

El desarrollador tiene control completo sobre cada aspecto de la interacción con la BD, desde la construcción de la consulta hasta la gestión de la conexión y los resultados.

Permite la implementación de lógica de acceso a datos muy específica y optimizada para las necesidades exactas de la aplicación.

No está limitado por las abstracciones o las decisiones de diseño de un ORM o un mapper.

  \| 24 of 37

# Características de"

# SQL Directo"

# Independencia del RDBMS: Baja

El SQL directo suele ser altamente específico del motor de base de datos que se esté utilizando. Las diferencias en sintaxis, funciones y tipos de datos hacen que cambiar de base de datos sea un proceso complejo.

La migración a otra base de datos puede requerir la reescritura significativa de las consultas SQL en toda la aplicación.

La portabilidad del código entre diferentes sistemas de gestión de bases de datos es limitada.

  \| 25 of 37

# Características de"

# SQL Directo"

# Desarrollo: Más Lento

Escribir y mantener manualmente una gran cantidad de código SQL puede ser un proceso tedioso y propenso a errores, lo que ralentiza el desarrollo.

La falta de herramientas y abstracciones para tareas comunes de acceso a datos (CRUD, relaciones) requiere más código manual.

La depuración de problemas relacionados con la base de datos puede ser más difícil sin las herramientas que ofrecen los ORM.

  \| 26 of 37

# Características de"

# SQL Directo"

# Seguridad: Mayor Riesgo (si no se toman precauciones)

La construcción manual de consultas SQL aumenta el riesgo de vulnerabilidades, especialmente inyecciones SQL, si no se implementan medidas de seguridad adecuadas (como consultas parametrizadas o prepared statements).

El desarrollador es responsable de garantizar la seguridad de las interacciones con la BD.

# Curva de Aprendizaje: Empinada

Requiere una inversión significativa de tiempo y esfuerzo para adquirir un dominio profundo del lenguaje SQL y las mejores prácticas de acceso a datos.

El desarrollador debe entender los conceptos relacionales, la normalización, la optimización de consultas y la seguridad de la base de datos.

  \| 27 of 37

# Características de"

# SQL Directo"

# Portabilidad: Baja

Como se mencionó anteriormente, la dependencia del SQL específico de cada base de datos dificulta la migración y el despliegue de la aplicación en diferentes entornos con diferentes sistemas de gestión de bases de datos.

# Naturaleza de los Datos: Ideal para Operaciones

# Complejas y Específicas

Es el enfoque más adecuado cuando se requieren consultas muy complejas, optimizaciones finas o el uso de características específicas de una base de datos en particular.

Permite un control total sobre cómo se manipulan y recuperan los datos, lo que puede ser crucial para ciertos tipos de aplicaciones o análisis.

  \| 28 of 37

# Características de"

# SQL Directo"

# Complejidad de Consultas: Manejo de Consultas

# Complejas

Permite escribir y gestionar consultas SQL de cualquier nivel de complejidad, sin las limitaciones que a veces pueden imponer las abstracciones de los ORM.

El desarrollador puede aprovechar al máximo las capacidades avanzadas del lenguaje SQL para realizar operaciones sofisticadas.

# Madurez de la Tecnología: Muy Alta y Bien Establecida

El lenguaje SQL es una tecnología fundamental y extremadamente madura con una amplia documentación, herramientas y una gran comunidad de desarrolladores.

La interacción directa con BD a través de SQL es una práctica popular en el desarrollo.

  \| 29 of 37

# SQL

# Tecnologías que siguen el enfoque"

# Directo"

Go:

database/sql : El paquete estándar de Go para interactuar con bases de datos SQL. Requiere escribir SQL directamente.

Bibliotecas de terceros que facilitan ejecución de consultas (ej: github.com/jmoiron/sqlx ).

Java:

JDBC (Java Database Connectivity): La API estándar de Java para interactuar con BD. Implica escribir SQL utilizando PreparedStatement para seguridad.

Bibliotecas ligeras como jOOQ (es más fácil, ya que genera y ejecuta SQL directo).

  \| 30 of 37

# SQL

# Tecnologías que siguen el enfoque"

# Directo"

JavaScript/TypeScript:

pg (para PostgreSQL), mysql (para MySQL), sqlite3 (para SQLite), mssql (para SQL Server): Drivers específicos para cada base de datos que permiten ejecutar consultas SQL directamente desde Node.js.

Bibliotecas como knex.js (un"query builder" que ayuda a construir SQL de forma programática pero aún genera y ejecuta SQL directo).

  \| 31 of 37

# Discordancia de impedancia con BD

# Relacionales

La capa de lógica se implementa en lenguajes como Go, Java o TS, mientras que la capa de persistencia con RDBMS.

Existe discordancia de impedancia (impedance mismatch ) entre estos dos mundos debido a sus diferentes paradigmas y modelos de datos.

E sta discordancia puede llevar a:

complejidad

código repetitivo

posibles errores

mal desempeño

  \| 32 of 37

# Discordancia de impedancia

ORDER LINE\_ITEM CUSTOMER int id int order\_id int id places contains contains datetime order\_date int product\_id string name int customer\_id int quantity

**Order**

**OrderItem**

**Customer**

\+ID int +OrderID int \* \*1 +ID int has contains refersTo +OrderDate datetime +ProductID int \*1 1 +Name string +CustomerID int +Quantity int +Orders Order\[\] +Product Product +OrderItems OrderItem\[\]

PRODUCT

int id

string name

decimal price

**Product**

\+ID int

\+Name string +Price decimal

  \| 33 of 37

# Discordancia de impedancia

Lenguajes como Java, TypeScript o Go favorecen el paradigma orientado a objetos (OOP):

aunque Go no es puramente OO, utiliza conceptos similares de estructuración de datos

Los datos se representan como objetos o estructuras con atributos (campos) y comportamientos (métodos).

Las relaciones entre objetos se expresan mediante referencias o composiciones.

  \| 34 of 37

# Discordancia de impedancia

Modelo Modelo de Característica Relacional Objetos

Tablas, filas, Representación Objetos, campos columnas

Referencias, Relaciones Claves foráneas composición

Métodos de Manipulación Consultas SQL Datos objetos

Tipos SQL Tipos Go Tipos de Datos específicos específicos

Claves Identidad de Identidad primarias objetos

Herencia No inherente Jerarquía de tipos

Problema

Mapeo entre estructuras diferentes

Navegación y gestión de relaciones

Traducción entre paradigmas de manipulación

Conversión de tipos de datos

Gestión de la identidad a través de capas

Representación de jerarquías

  \| 35 of 37

# DBA versus Programador

# DBA

Cumple con las reglas de la BD (integridad referencial, procedimientos almacenados, secuencias, etc.)

El esquema no se toca

El esquema podría cambiar

Opinaré sobre las consultas/SQL para optimizarlas

Necesito poder perfilar todas las llamadas SQL

Aprovechar las funciones de la base de datos (joins externos, subconsultas, funciones especializadas…)

# Programador

El modelo de datos no debe restringir el modelo de objetos

No deseo código de base de datos en el código de objetos/componentes

El acceso a datos debe ser rápido

Debo minimizar las llamadas a la base

de datos

Necesito consultas OO, no SQL

Voy a proteger la aplicación web de cambios en el esquema

  \| 36 of 37

# Todavía queremos

# hacer"

Es hora de considerar alternativas que nos ayuden a lidiar con la discordancia de impedancia

# SQL Directo"

?

  \| 37 of 37

<image redacted: 718x479px, 718x479pt, ~72dpi, JPG, DEVICE_RGB, 32bpp>