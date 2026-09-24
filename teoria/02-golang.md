# Pr

# ogramación Web

# Introducción a Go

Dr. Alejandro Zunino & Dr. Alfredo Teyseyre alejandro.zunino@isistan.unicen.edu.ar, alfredo.teyseyre@isistan.unicen.edu.ar

  \| 1 of 71

<image redacted: 72x72px, 72x72pt, ~72dpi, PNG, DEVICE_RGB, 32bpp>

# Contenidos

1\. ¿Por qué no Java ni TypeScript?

2\. Características del lenguaje Go

3\. Elementos del lenguaje Go

4\. Logging en Go

5\. Testing Unitario: introducción

6\. Concurrencia, gorutinas y canales

7\. Interfaces

8\. Genéricos

9\. Creando un Proyecto Go: inicialización

10\. Y ahora qué?

11\. Go y el desarrollo Web

  \| 2 of 71

<image redacted: 479x479px, 479x479pt, ~72dpi, JPG, DEVICE_RGB, 32bpp>

¿Por qué no Java ni TypeScript?

Java y TypeScript son lenguajes excelentes y muy usados en la industria. Para esta materia, sin embargo, elegimos Go porque facilita la transparencia: entender cómo funciona una aplicación web sin capas de magia intermedia.

Los criterios que guían esta elección:

Simplicidad: un solo modo idiomático de hacer cada cosa.

Herramientas mínimas: compilador, formateo, tests y módulos integrados.

Recursos acotados: binarios livianos y bajo consumo de memoria.

Funcionamiento visible: lo que escribimos es (casi) todo lo que hay.

  \| 3 of 71

¿Por qué no Java?

Alto consumo de recursos : la JVM agrega overhead de memoria y arranque pesado incluso en programas triviales.

Fragmentación de versiones : Java 8/11/17/21, SE vs EE: sintaxis y librerías distintas según el entorno.

Abstracción que oculta el funcionamiento : ej. Spring Boot: inyección de dependencias y anotaciones que esconden qué sucede por debajo en HTTP.

Herramientas obligatorias : Maven o Gradle para compilar y gestionar dependencias.

Crecimiento del lenguaje : múltiples formas de resolver lo trivial (loops clásicos vs streams, lambdas, clases anónimas, anotaciones) con impacto no obvio en el desempeño.

  \| 4 of 71

¿Por qué no TypeScript?

Ecosistema fragmentado : múltiples runtimes (Node/Deno/Bun), bundlers y frameworks para lo mismo.

Cambios frecuentes del lenguaje : rotación constante de configuración ( tsconfig , ESM/CommonJS).

Package managers obligatorios : npm, pnpm, yarn…, cada uno con su propio lockfile.

Problemas de seguridad : ataques a la cadena de suministro y vulnerabilidades en dependencias transitivas.

Muchísimas dependencias : hasta para desarrollos simples ( node\_modules gigante).

Múltiples formas de lo trivial : var / let / const , funciones flecha, callbacks/promesas/ async-await , decorators…

Single thread : JS corre en un único hilo (event loop); concurrencia menos natural que las gorutinas.

  \| 5 of 71

# Características del lenguaje Go

Desarrollado en 2007 por Google para sistemas, desarrollo web, herramientas de CLI, etc.

Énfasis en la simplicidad, eficiencia y legibilidad.

Compilado (sin máquina virtual), tipado estático.

Compilación y ejecución muy rápida.

Sintaxis similar a C.

Recolección de basura (garbage collection).

Concurrencia incorporada y algo de detección de deadlocks.

Sin clases, herencia, sobrecarga, anotaciones, constructores o excepciones.

Mecanismo inusual de interfaces en lugar de herencia.

  \| 6 of 71

# "

# It is intended that programs written to

# the Go 1 specification will continue to

# compile and run correctly, unchanged, over

# the lifetime of that specification. … Go

# programs that work today should continue to

# work even as future releases of Go 1

# "

arise.

\-R uss Cox, Google

  \| 7 of 71

# "

# There will not be a Go 2 that breaks Go 1

# programs. Instead, we are going to double

# down on compatibility, which is far more

# valuable than any possible break with the

# past. In fact, we believe that prioritizing

# compatibility was the most important design

# "

decision we made for Go 1.

\-R uss Cox, Google

  \| 8 of 71

Quiénes lo diseñaron?

Robert Griesemer

Dr. with Niklaus Wirth (Pascal & more)

V8, Java Hotspot, Sawzall (antes MapReduce)

Rob Pike

Unix, Inferno OS, Limbo lang, Plan9 OS, etc @Bell labs

Ken Thompson

Diseñó Unix @Bell labs

B prog lang (before C), expresiones regulares en ed , UTF-8

Premio Turing

  \| 9 of 71

ALGOL 60 (Backus et al., 1960)

Pascal (Wirth, 1970)

Modula-2 (Wirth, 1980) CSP (Hoare, 1978)

Oberon (Wirth & Squeak (Cardelli & Pike, Gutknecht, 1986) 1985)

Object Oberon (Mössenböck, Templ & Newsqueak (Pike, 1989) Griesemer, 1990)

Oberon-2 (Wirth & Alef (Winterbottom, 1992) Mössenböck, 1991)

Go (Griesemer, Pike & Thompson, 2009)

C (Ritchie, 1972)

  \| 10 of 71

# "

# When the three of us got started, it was

# pure research. The three of us got together

# and decided that we hated C++. We started

# off with the idea that all three of us had

# to be talked into every feature in the

# language, so there was no extraneous

# garbage put into the language for any

# "

reason.

\-K en Thompson, Distinguished Engineer, Google

  \| 11 of 71

# Elementos del lenguaje Go

tipos básicos:

bool string int8 int16 int32 int64 uint8 … int uint

float32 float64 complex64 complex128

(sin escapes) Quotes: 'a' (rune), "texto UTF-8" (string),\`raw string\`

variables (infiere los tipos de las variables a partir de la asignación)

declaraciones de variables: var x, y, z = 0, 1.23, false

declaraciones cortas: x := 0; y := 1.23; z := false

La asignación entre diferentes tipos requiere una conversión como int(float\_expression)

los operadores son como en C, pero sin ?:

  \| 12 of 71

# Punteros y los operadores \* y & : conceptos

# básicos

a la ubicación donde se encuentra En lugar de contener un valor, un puntero"apunta" ese valor.

Permite la modificación directa del valor original a través del puntero.

El operador & se utiliza para obtener la dirección de memoria de una variable.

Se lee como"dirección de".

Para pasar argumentos por referencia a funciones.

Trabajar con estructuras de datos complejas evitando copias de datos.

edad := 30

direccionMemoria := &edad

fmt.Println("Valor de edad:", edad) // Valor de edad: 30 fmt.Println("Dirección de memoria de edad:", direccionMemoria) // Dirección de memoria de edad: 0xc0000101e0

fmt.Printf("Tipo de direccionMemoria: %T\\n", direccionMemoria) // Tipo de direccionMemoria:\*int

  \| 13 of 71

# Punteros y los operadores \* y & :

# dereferenciación

El operador \* tiene dos usos principales relacionados con punteros:

Declaración de un tipo puntero: Se utiliza para indicar que una variable es un puntero a un tipo específico.

Dereferenciación: Cuando se aplica a una variable puntero, accede al valor almacenado en la dirección de memoria a la que apunta el puntero. Se lee como el valor apuntado por.

contador := 5

ptrContador := &contador

| fmt.Println( fmt.Println( // Dirección de contador: fmt.Println( // Valor apuntado por // Modificando el valor ptrContador = 15 | contador) // Valor de contador: 5 ptrContador) 0xc0000101e8 (puede variar) ptrContador) ptrContador: 5 a través del puntero |
|---------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------|
| fmt.Println(                                                                                                                    | contador) // Nuevo valor de contador: 15   \| 14 of 71                                                                      |

# Punteros y los operadores \* y & : resumen

Op Propósito Ejemplo Resultado

Obtiene la dirección de una & variable

\* 1 . Declara un ti po puntero.

2\. Dereferencia un puntero para \*

acceder al valor apuntado

ptr = ptr contiene la dirección de &variable variable

ptr puede almacenar una var ptr dirección de memoria de un \*int int

valor = valor contiene el valor al que \*ptr apunta ptr

  \| 15 of 71

# Arrays

Colección de elementos del mismo tipo con una longitud fija.

La longitud se define al declarar el array y no se puede cambiar.

Los elementos se acceden mediante un índice (base 0).

| var a \[3\]int b := \[2\]string{ fmt.Println(a\[0\]) // // b\[1\] = | // Declara un array de 3 enteros, inicializados a 0 Mundo Accede al primer elemento (0) Modifica el segundo elemento |
|---------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------|
| fmt.Println(b) //                                                   | Salida: \[Hola Go\]                                                                                                  |

  \| 16 of 71

# Slices

Proporcionan una vista dinámica (segmento) de un array subyacente.

Su longitud puede variar (dentro de la capacidad del array subyacente).

numeros := \[5\]int{1, 2, 3, 4, 5} slice1 := numeros\[1:4\] // desde 1 (inclusive) hasta 4 (exclusivo) -\> \[2 3 4\] slice2 := numeros\[:3\] // desde 0 hasta el índice 3 (exclusivo) -\> \[1 2 3\] slice3 := numeros\[2:\] // desde 2 (inclusive) hasta el final -\> \[3 4 5\] slice4 := \[\]int{10, 20} // Crea un slice directamente (el array subyacente se gestiona internamente)

fmt.Println(slice1) fmt.Println(len(slice1)) // Longitud del slice (número de elementos) fmt.Println(cap(slice1)) // Capacidad del slice (tamaño del array subyacente desde el inicio del slice)

slice1 = append(slice1, 6) // Agrega un elemento al slice (puede crear un nuevo array subyacente si es necesario) fmt.Println(slice1)

  \| 17 of 71

# slice1 := numeros\[1:4\]

1

apunta a

Slice slice1

2

3

4

Array numeros

2

3

4 5 apunta a

apunta a

  \| 18 of 71

# Funciones

Funciones regulares:

func Sin(x float64) float64 func AddScale(x, y int, f float64) int

Funciones con múltiples valores de retorno:

func Write(data \[\]byte) (written int, err error)

Funciones variádicas:

func sumarTodo(numeros ...int) int { total := 0

f or \_, num := range numeros { total += num

} r eturn total

}

  \| 19 of 71 func sumar(a int, b int) int { return a + b

}

func dividir(a float64, b float64) (float64, error) { if b == 0 { return 0, fmt.Errorf("división por cero") } return a / b, nil }

func main() { resultado := sumar(5, 3) fmt.Println("La suma es:",r esultado)

cociente, err := dividir(10, 2) if err != nil { fmt.Println("Error:", err) } else { fmt.Println("El cociente es:", cociente) } }

  \| 20 of 71

# Funciones como valores

Las funciones pueden ser asignadas a variables, pasadas como argumentos a otras funciones y devueltas por otras funciones:

func multiplicar(a int, b int) int { r eturn a\* b

}

func aplicarOperacion(a int, b int, operacion func(int, int) int) int { r eturn operacion(a, b) }

func main() { miF uncion := multiplicar // Asigna la función'multiplicar' r esultado1 := miFuncion(4, 5) fm t.Println("Resultado 1:",r esultado1) // Imprime: Resultado 1: 20

r esultado2 := aplicarOperacion(10, 3, multiplicar) // Pasa'multiplicar' fm t.Println("Resultado 2:",r esultado2) // Imprime: Resultado 2: 30 }

a la variable'miFuncion'

como argumento

  \| 21 of 71

# Mapas

Colecciones de pares clave-valor.

Las claves deben ser de un tipo comparable (ej., int , string , pero no slices).

Los valores pueden ser de cualquier tipo.

edades := map\[string\]int{"Juan": 30,

"María": 25,

"Pedro": 35,

} fmt.Println(edades\["Juan"\]) // Accede al valor asociado a la clave "Juan" edades\["Ana"\] = 28 // Agrega un nuevo par clave-valor delete(edades,"María") // Elimina el par clave-valor con la clave "María" valor, existe := edades\["Carlos"\] // Comprueba si una clave existe if existe { fm t.Println("La edad de Carlos es:",v alor) } fmt.Println(edades)

  \| 22 of 71

# Estructuras de Control de Flujo: if-else y for

Go utiliza un conjunto simple de estructuras de control de flujo.

if-else : Para la toma de decisiones. Permite una declaración de inicialización.

if x := 10; x \> 5 { fm t.Println("x es mayor que 5") } else { fm t.Println("x no es mayor que 5") }

for : Es la única estructura de bucle en Go.

// Bucle clásico

for i := 0; i \< 5; i++ { ... } // Como un"while"

sum := 1

for sum \< 1000 { sum += sum } // Bucle infinito

for { ... }

  \| 23 of 71

# Estructuras de Control de Flujo: switch

switch : Una forma más limpia de escribir cadenas de if-else .

switch os := runtime.GOOS; os { case"darwin":

fm t.Println("macOS") case"linux":

fm t.Println("Linux") default:

fm t.Printf("%s.\\n", os) }

). El ejemplo usa el paquete estándar runtime (requiere import"runtime"

  \| 24 of 71

# range : iterar sobre colecciones

for range permite iterar sobre slices, arrays, mapas, strings y canales.

Sobre slices/arrays: devuelve índice, valor .

"uva"} frutas := \[\]string{"manzana","pera", for i, fruta := range frutas { fm t.Printf("%d: %s\\n",i,fr uta) } for \_, fruta := range frutas { // Si no necesitás el índice, usá'\_' fm t.Println(fruta) }

Sobre mapas: devuelve clave, valor (en orden arbitrario).

for nombre, edad := range edades { fm t.Printf("%s tiene %d años\\n",n ombre, edad) }

  Ti p

\_ se usa para ignorar un valor que no necesitás (índice, clave o el propio valor).

  \| 25 of 71

# Manejo de Errores

Go tiene un enfoque explícito para el manejo de errores.

Las funciones que pueden fallar devuelven un valor de tipo error como su último valor de retorno.

El código que llama a la función debe verificar si el error es nil o no.

El código es robusto y fácil de entender: el flujo de error es explícito.

func main() { fil e, err := os.Open("archivo\_inexistente.txt") if err != nil { fm t.Println("Error al abrir el archivo:", err) r eturn // Termina la ejecución si hay un error } defer file.Close()

// ... hacer algo con el archivo fm t.Println("Archivo abierto exitosamente:",fil e.Name()) }

  \| 26 of 71

# Errores: envolver y comparar

Envolver un error agrega contexto, preservando el error original con %w :

func abrirConfig(nombre string) error { f, err := os.Open(nombre) if err != nil { r eturn fmt.Errorf("no se pudo abrir %s: %w",n } defer f.Close() r eturn nil

}

Comparar con errors.Is para verificar errores específicos, incluso si vienen envueltos:

err := abrirConfig("config.json") if errors.Is(err, os.ErrNotExist) { l og.Println("El archivo no existe:", }

  Ti p

Usá %w al envolver y errors.Is / errors.As para inspeccionar la cadena de errores.

ombre, err)

err)

  \| 27 of 71

# Logging en Go

El logging es crucial para entender el comportamiento de una aplicación, depurar problemas y monitorear su estado en producción.

Go proporciona un paquete estándar log que es simple y fácil de usar para necesidades básicas de logging.

Por defecto, el paquete log escribe mensajes en la salida de error estándar ( stderr ) e incluye una marca de tiempo.

import (""log )

func main() { l og.Println("Este es un mensaje de log simple.") // Salida: 2025/07/11 10:00:00 Este es un mensaje de log simple. }

  \| 28 of 71

# Funciones del paquete log : niveles de

# severidad

El paquete log ofrece varias funciones para diferentes niveles de severidad:

log.Print() / log.Println() / log.Printf() :

Escriben un mensaje de log estándar. Printf permite formatear la cadena.

log.Fatal() / log.Fatalf() :

Escriben el mensaje de log y luego llaman a os.Exit(1) , terminando el programa inmediatamente. Se usan para errores irrecuperables.

log.Panic() / log.Panicf() :

Escriben el mensaje de log y luego llaman a panic() . Se usan cuando un error es tan grave que el programa no puede continuar de forma segura.

  \| 29 of 71

# Funciones del paquete log : ejemplo

l og.Printf("Iniciando proceso con ID: %d",12 345)

\_, err := os.Open("archivo\_que\_no\_existe.txt") if err != nil { l og.Fatalf("Error fatal al abrir el archivo: %v", err) }

  \| 30 of 71

# Configurando el Logger Estándar: opciones

Es posible configurar el comportamiento del logger por defecto:

log.SetOutput(w io.Writer) : Cambia el destino de la salida del log (por ejemplo, a un archivo).

log.SetPrefix(prefix string) : Añade un prefijo a cada línea de log.

log.SetFlags(flag int) : Modifica la información que se incluye (fecha, hora, archivo, etc.).

  \| 31 of 71

# Configurando el Logger Estándar: ejemplo

func main() { l og.SetPrefix("mi-app:") l og.SetFlags(log.LstdFlags \| log.Lshortfile) // Flags estándar + nombre de archivo y línea

l og.Println("Log con prefijo y flags.") // Salida: mi-app: 2025/07/11 10:00:00 main.go:12: Log con prefijo y flags.

// Log a un archivo fil e, \_ := os.OpenFile("app.log", os.O\_CREATE\|os.O\_WRONLY\|os.O\_APPEND, 0666) defer file.Close() l og.SetOutput(file) l og.Println("Este mensaje va al archivo.") }

  \| 32 of 71

# Logging Estructurado y de Terceros

Para aplicaciones complejas o en producción, el paquete log puede ser limitado (no tiene niveles de log como DEBUG , INFO , WARN , etc.).

Logging Estructurado:

Escribe logs en formatos como JSON, con pares clave-valor.

Facilita el procesamiento automático, la búsqueda y el análisis de logs.

Bibliotecas Populares:

slog : Nuevo paquete de logging estructurado, añadido en Go 1.21. Es la opción recomendada actualmente.

zerolog : Muy rápido, con cero asignaciones de memoria y una API fluida.

  \| 33 of 71

# El Paquete slog : Logging Estructurado Nativo

Introducido en Go 1.21, slog es el logger estructurado oficial de Go:

Niveles de Log: Soporte nativo para niveles como Debug , Info , Warn , y Error .

Salida Estructurada: Emite logs en formatos como JSON, facilitando el análisis por máquinas.

Atributos Clave-Valor: Permite añadir contexto a los logs de forma estructurada.

Handlers Personalizables: TextHandler y JSONHandler , o manejadores propios.

  \| 34 of 71

# El Paquete slog : ejemplo de código

import (""log/slog"os"

)

func main() { l ogger := slog.New(slog.NewTextHandler(os.Stdout, nil)) slog.SetDefault(logger) // Establecer como logger global slog.Debug("Iniciando la aplicación...", slog.Info("Petición recibida", slog.Warn("Contraseña a punto de expirar", slog.Error("No se pudo conectar a la base de datos", }

time=2025-07-26T21:00:14.711-03:00 level=INFO msg="Petición recibida"m time=2025-07-26T21:00:14.711-03:00 level=WARN msg="Contraseña a punto de expirar" time=2025-07-26T21:00:14.711-03:00 level=ERROR msg="No se pudo conectar a la base de datos"

"version","1.2.3")

"GET","ruta","metodo","/api/usuarios")" ,12"usuario\_id 3)"error","conexión rechazada")

etodo=GET ruta=/api/usuarios usuario\_id=123 error=" con

  \| 35 of 71

# Handlers y Niveles en slog

slog.NewTextHandler : Crea un logger que emite logs en formato de texto plano ( clave=valor ).

slog.NewJSONHandler : Crea un logger que emite logs en formato JSON.

slog.HandlerOptions : Permite configurar el nivel de log mínimo que se registrará.

Niveles: LevelDebug \< LevelInfo \< LevelWarn \< LevelError .

func main() { opts := &slog.HandlerOptions{ L evel: slog.LevelDebug, // Mostrar logs desde el nivel Debug hacia arriba } h andler := slog.NewTextHandler(os.Stdout, opts) l ogger := slog.New(handler)

l ogger.Debug("Este es un mensaje de depuración.") l ogger.Info("Este es un mensaje informativo.") // Si Level fuera slog.LevelInfo, el mensaje Debug no se mostraría. }

  \| 36 of 71

# Agregando Atributos a los Logs con slog : en

# el momento

Atributos en el momento del log:

slog.Info("Usuario ha iniciado sesión",

Logger con atributos predefinidos:

requestLogger := slog.With( slog.String("componente", slog.String("request\_id", ) requestLogger.Info("Petición entrante") requestLogger.Warn("Validación fallida") // Todos los logs de requestLogger tendrán"componente"

"usuario","ana","rol","admin")

"peticion\_http"),"xyz-123"),

y"request\_id"

  \| 37 of 71

# Agregando Atributos a los Logs con slog :

# grupos

Grupos de atributos:

slog.Info("Petición procesada", slog.Group("http", slog.String("metodo", slog.Int("status",2 ), slog.Int("bytes\_escritos", ) // Salida JSON: {"http": {"metodo":"POST",

"POST"),

01\),

512\),

"status": 201},"bytes\_escritos": 512}

  \| 38 of 71

# Paquetes y Visibilidad

El código en Go se organiza en paquetes.

package main : Es un paquete que define un programa ejecutable. main() es el punto de entrada.

Otros paquetes: Son bibliotecas de código reutilizable.

La visibilidad de los identificadores (variables, funciones, tipos, etc.) fuera de un paquete se controla:

Identificadores Exportados (Públicos): Si comienza con una mayúscula, es visible y accesible desde otros paquetes que lo importen.

var Version ="1.0" // Exportado func Saludar() { ... } // Exportado

Identificadores no Exportados (Privados): Si comienza con una minúscula, es privado al paquete.

var versionInterna ="beta" // No exportado func calcular() { ... } // No exportado

  \| 39 of 71

# La Sentencia defer

defer pospone la ejecución de una llamada a función hasta que la función que la contiene está a punto de retornar.

Las llamadas a funciones diferidas se ejecutan en orden LIFO (Last-In, First-Out).

Útil para tareas de limpieza, como cerrar archivos, desbloquear mutexes o cerrar conexiones de red.

Garantiza que la limpieza se realice sí o sí (retorno normal, panic , etc.).

func leerArchivo(nombre string) error { fil e, err := os.Open(nombre) if err != nil { r eturn err

} // file.Close() se ejecutará justo antes de que leerArchivo retorne. defer file.Close() // ... lógica para leer el archivo ... fm t.Println("Procesando el archivo.") r eturn nil

}

  \| 40 of 71

# Structs: Tipos de Datos Agregados

Un struct es una colección de campos con nombre que permite agrupar datos relacionados.

Se definen con la palabra clave type seguida del nombre del struct y la palabra clave struct .

Los campos dentro de un struct pueden ser de diferentes tipos.

type Persona struct { N ombre string E dad int

}

func main() { // Creación de una instancia de Persona

p1 := Persona{ N ombre:"Ana",

E dad: 30, } fm t.Println("Nombre:", p1.Nombre) // Acceso a los campos con el operador . p1.Edad = 31 fm t.Printf("Persona: %+v\\n", p1) // Imprime la struct con nombres de campo }

  \| 41 of 71

# Testing Unitario: introducción

Go tiene soporte para pruebas unitarias integrado en su conjunto de herramientas estándar:

Las pruebas se escriben en archivos que terminan en \_test.go .

Cada función de prueba debe comenzar con el prefijo Test y recibir un argumento \*testing.T .

Se utiliza el comando go test para ejecutar todas las pruebas en el paquete actual.

Ejemplo:

Supongamos que tenemos un archivo matematicas.go :

// matematicas.go package matematicas func Suma(a, b int) int { r eturn a + b

}

  \| 42 of 71

# Testing Unitario: ejemplo

El archivo de prueba sería matematicas\_test.go :

// matematicas\_test.go package matematicas

import"testing"

func TestSuma(t\*testing.T) { r esultado := Suma(2, 3) esperado := 5 if r esultado != esperado { t.Errorf("Suma(2, 3) = %d; se esperaba %d",r esultado, esperado) } }

Para ejecutar las pruebas: go test -v

  \| 43 of 71

# Testing Unitario: Table-Driven Tests

Una práctica muy común en Go: una tabla de casos, donde cada fila define una entrada y el resultado esperado.

func TestSuma(t\*testing.T) { casos := \[\]struct { n ombre string a, b int esperado int }{ 3, 5}, {"positivos",2, 0, 0, 0}, {"cero", -1, -2, -3}, {"negativos", }

f or \_, tc := range casos { t.Run(tc.nombre, func(t\*testing.T) { if r es := Suma(tc.a, tc.b); res != tc.esperado { t.Errorf("Suma(%d, %d) = %d; esperado %d", } }) } }

t.Run agrupa cada caso como un subtest.

tc.a, tc.b, res, tc.esperado)

  \| 44 of 71

Agregar un caso nuevo sólo requiere una fila más.

# Concurrencia, gorutinas y canales

gorutina : una función que se ejecuta de manera concurrente con otras del mismo programa (proceso)

Similar a un hilo de ejecución, pero más ligeras y fáciles de usar.

canales : una generalización type-safe de los pipes de Unix para comunicación entre gorutinas

Permiten el envío y recepción de datos entre gorutinas, realizando sincronización implícita entre ellas.

  N ote

También se pueden usar locks, semáforos, mutexes, etc., pero son menos comunes en Go.

  \| 45 of 71

# Gorutinas

Una función que puede ejecutarse concurrentemente con otras funciones.

El costo de creación y gestión de gorutinas es significativamente menor que el de los threads.

El runtime de Go gestiona la planificación y ejecución de las gorutinas.

func imprimir(mensaje string) { for i := 0; i \< 5; i++ { fmt.Println(mensaje, i) time.Sleep(time.Millisecond\*1 00) } }

func main() { go imprimir("Hola desde la gorutina 1") // Inicia una nueva gorutina go imprimir("Hola desde la gorutina 2") // Inicia otra nueva gorutina // La función main también se ejecuta como una gorutina (la gorutina principal) fmt.Println("Hola desde la gorutina principal") // Esperar a que las otras gorutinas terminen (no es la mejor práctica) time.Sleep(time.Second\*1) }

  \| 46 of 71

# Canales

Mecanismos para la comunicación segura entre gorutinas.

Permiten enviar y recibir valores entre gorutinas de forma sincronizada.

Ayudan a evitar condiciones de carrera y otros problemas comunes en la programación concurrente con memoria compartida.

// Declaración de un canal que puede enviar y recibir valores de tipo int var miCanal chan int

// Creación de un canal (sin buffer) miCanal = make(chan int) // Creación de un canal con buffer de tamaño 10

canalConBuffer := make(chan int, 10)

// Enviar 50 al canal

canalConBuffer \<- 50

// Recibir un valor del canal

valor := \<-canalConBuffer

  \| 47 of 71

# Ejemplo de canales

func enviar(c chan string) { c \<-"Mensaje 1"

time.Sleep(time.Millisecond\* 500) c \<-"Mensaje 2"

close(c) // Cerrar el canal indica que no se enviarán más valores }

func recibir(c chan string) { for msg := range c { // Bucle para recibir valores hasta que el canal se cierre fmt.Println("Recibido:",m sg) } fmt.Println("Terminó de recibir") }

func main() { miCanal := make(chan string) go enviar(miCanal) go recibir(miCanal)

time.Sleep(time.Second\*1) }

  \| 48 of 71

# Go no es orientado a objetos, pero…

type Vertex struct { X,Y fl oat64

}

func (v\*Vertex) Scale(f float64) { v.X=v.X\*f

v.Y=v.Y\*f

}

func (v Vertex) Abs() float64 { r eturn math.Sqrt(v.X\*v.X + v.Y\*v.Y) }

func main() {

v := &Vertex{3, 4}

v. Scale(5) fm t.Println(v, v.Abs()) }

  \| 49 of 71

# Interfaces

Una colección de firmas de métodos. Como una clase abstracta en otros lenguajes.

Define un contrato que un tipo debe implementar para ser considerado de esa interfaz.

Especifica el comportamiento que un tipo debe tener.

No define la implementación concreta de los métodos, solo su firma (nombre, parámetros y tipo de retorno).

Proporcionan una forma de lograr polimorfismo y abstracción.

Ej: cualquier tipo puede implementar la interfaz fmt.Stringer (definiendo String() string ) para permitir su impresión con fmt.Println :

type Stringer interface { String() string }

  \| 50 of 71

# Ejemplo de uso de interfaces

type Animal interface { H acerSonido() string }

// Tipo'Perro'im plementa la interfaz'Animal' type Perro struct{}

func (p Perro) HacerSonido() string { r eturn"Guau!"

}

// Tipo'Gato'im plementa la interfaz'Animal' type Gato struct{}

func (g Gato) HacerSonido() string { r eturn"Miau!"

}

func main() {

v ar a Animal

perro := Perro{} gato := Gato{} a = perro fm t.Println("El animal hace:", a.HacerSonido()) a = gato fm t.Println("El animal hace:", a.HacerSonido()) }

«interface»

**Animal**

HacerSonido() : string

**Perro**

**Gato**

HacerSonido() : string HacerSonido() : string

La flecha punteada indica implementación de la interfaz, no herencia.

  \| 51 of 71

# Implementación implícita

En Go, la implementación de una interfaz es implícita.

Un tipo implementa una interfaz simplemente implementando todos los métodos definidos en la interfaz.

No hay una declaración explícita:

como implements en otros lenguajes.

type Printer interface { Prin t(s string) }

type ConsolePrinter struct{}

func (cp ConsolePrinter) Print(s string) { fm t.Println(s) }

func main() {

v ar p Printer console := ConsolePrinter{} //'ConsolePrinter'im plementa'Printer' p = console p.Print("¡Hola desde la consola!") }

  \| 52 of 71

# Interfaces vacías: interface{} o any

La interfaz vacía no tiene ningún método definido.

Todos los tipos implementan la interfaz vacía.

Se utiliza para representar un valor de tipo desconocido o arbitrario.

package main import"fmt"

// equivalente a func hacerAlgo(i any) { func hacerAlgo(i interface{}) { fm t.Printf("El valor es: %v, y su tipo es: %T\\n",i,i)

func main() {

| acerAlgo(10) // h acerAlgo( h              | El valor es: 10, y su tipo es: int                                                |
|--------------------------------------------|-----------------------------------------------------------------------------------|
| acerAlgo(true) // h acerAlgo(struct{}{}) } | El valor es: true, y su tipo es: bool // El valor es: {}, y su tipo es: struct {} |

  \| 53 of 71

# Uso de interfaces: beneficios

Las interfaces son poderosas y se utilizan en muchos escenarios:

Polimorfismo: Permitir que el código funcione con diferentes tipos de datos de manera uniforme si implementan la misma interfaz.

Abstracción: Ocultar los detalles de implementación y enfocarse en el comportamiento.

Desacoplamiento: Reducir la dependencia entre diferentes partes del código:

Un componente puede depender de una interfaz en lugar de una implementación concreta: cambio de implementaciones sin afectar al código que las utiliza.

Pruebas unitarias: Facilitar la creación de mocks (objetos simulados) para probar componentes de forma aislada.

Bibliotecas estándar: Muchas partes de la biblioteca estándar de Go utilizan interfaces (por ejemplo, io.Reader , io.Writer , error ).

  \| 54 of 71

# Uso de interfaces: ejemplo con io.Reader

package main import ("fmt"

"io"

"strings" )

func leerContenido(r io.Reader) error { buffer := make(\[\]byte, 10) for { n, err := r.Read(buffer) if err == io.EOF { break

} if err != nil { return fmt.Errorf("error al leer: %w", err) } fmt.Print(string(buffer\[:n\])) } fmt.Println() return nil

}

func main() { texto :="Este es un texto de ejemplo."

lectorString := strings.NewReader(texto)

err := leerContenido(lectorString) if err != nil { fmt.Println("Error:", err) }

// También podríamos pasar un archivo abierto // (que implementa io.Reader) // file, err := os.Open("miarchivo.txt") // if err == nil { // defer file.Close() // leerContenido(file) // } }

  \| 55 of 71

# Genéricos

Permiten escribir código que funciona con múltiples tipos sin sacrificar la seguridad de tipos.

El objetivo principal es reducir la duplicación de código.

Se basan en dos conceptos clave:

Parámetros de Tipo: placeholders para tipos desconocidos.

Restricciones de Tipo: definen qué tipos son válidos para un parámetro de tipo.

  \| 56 of 71

# Funciones Genéricas

Son funciones que pueden operar sobre argumentos de cualquier tipo que satisfaga una restricción.

Los parámetros de tipo se declaran entre corchetes \[\] antes de los parámetros de la función.

any es una restricción predefinida que permite cualquier tipo ( any es un alias de interface{} ).

//'T' es un parámetro de tipo, restringido por'any'. func Imprimir\[T any\](valor T) { fm t.Println(valor)

func main() { primir("Hola, mundo") // T es string

| primir(42) Im  | // T es int  |
|----------------|--------------|
| primir(true) } | // T es bool |

  \| 57 of 71

# Restricciones de Tipo (Constraints): concepto

Definen el"contrato" que un tipo debe cumplir para ser usado como argumento de tipo.

Se definen usando interfaces.

Una interfaz como constraint especifica los métodos o tipos que son requeridos.

  \| 58 of 71

# Restricciones de Tipo (Constraints): ejemplo

//'Number' es una constraint que permite cualquier tipo que sea int o float64. // El símbolo'\|' se usa para unir tipos en una constraint. type Number interface { in t \| float64 }

// Esta función suma un slice de'Number'.

func Suma\[T Number\](numeros \[\]T) T {

v ar total T

f or \_, n := range numeros { total += n

} r eturn total

}

func main() { fm t.Println("Suma de enteros:", Suma(\[\]int{1, 2, 3})) fm t.Println("Suma de flotantes:", Suma(\[\]float64{1.5, 2.5, 3.0})) }

  \| 59 of 71

# La Constraint comparable

Go proporciona una constraint predefinida útil: comparable .

Permite cualquier tipo que pueda ser comparado usando los operadores == y != .

Incluye tipos como int , string , bool , punteros, canales, y structs si todos sus campos son comparables.

No incluye slices, maps, o funciones.

// Esta función busca un elemento en un slice.

//'T' debe ser comparable para poder usar'==' func Buscar\[T comparable\](slice \[\]T, valor T) int { f or i, v := range slice {

== if v v alor { r eturn i

} } r eturn -1

}

  \| 60 of 71

# Tipos Genéricos

Los structs y otros tipos también pueden ser genéricos.

Útil para crear estructuras de datos de propósito general, como stacks, listas enlazadas o árboles.

type Stack\[T any\] struct { i tems \[\]T }

func (s\*Stack\[T\]) Push(item T) { s.items = append(s.items, item) }

func (s\*Stack\[T\]) Pop() (T, bool) { if l en(s.items) == 0 {

v ar zero T // Crea el valor cero del tipo T r eturn zero, false } i tem := s.items\[len(s.items)-1\] s.items = s.items\[:len(s.items)-1\] r eturn item, true }

  \| 61 of 71

# Herramientas: go fmt y go vet

Go incluye herramientas de línea de comandos para mantener la calidad y consistencia del código:

go fmt : Aplica el formato oficial de Go ( gofmt ) a todos los archivos. No hay discusión de estilo: un único formato para todo el código.

go fmt ./...

go vet : Analiza el código en busca de constructos sospechosos que el compilador no detecta (cadenas de formato incorrectas, struct tags mal formados, código inalcanzable, etc.).

go vet ./...

  Ti p

En la materia ejecutamos go fmt ./... y go vet ./... como parte del flujo estándar antes de compilar y probar.

  \| 62 of 71

# Creando un Proyecto Go: inicialización

Crea un proyecto con go mod init : crea un archivo go.mod que define el módulo (su nombre y la versión de Go). Las dependencias se agregan después, cuando importes paquetes externos.

1\. Crea un nuevo directorio para tu proyecto:

mkdir mi- proyecto- cd mi- go proyecto-

2\. Inicializa el módulo:

go mod init ejemplo.com/mi-

ejemplo.com/mi-

para el módulo y previene conflictos de importación. Generalmente se usa una URL de un repositorio de código.

go

go proyecto-

go es la ruta del módulo. Sirve como un nombre único proyecto-

  \| 63 of 71

# Creando un Proyecto Go: archivo main.go

3\. Crea un archivo main.go con el siguiente contenido:

package main

import"fmt"

func main() { fm t.Println("¡Hola, Go!") }

  \| 64 of 71

# Compilando y Ejecutando

Go es un lenguaje compilado. Tienes dos opciones principales para ejecutar tu código:

go run : Compila y ejecuta el programa en un solo paso. Es útil para el desarrollo y pruebas rápidas.

go run .

El . le indica a Go que compile y ejecute el paquete en el directorio actual.

go build : Compila el código y crea un archivo ejecutable. Este ejecutable se puede distribuir y ejecutar de forma independiente.

go build -o mi-app

Esto crea un ejecutable llamado mi-app .

Para ejecutarlo:

./mi-app

  \| 65 of 71

# Módulos en Go: go.mod : concepto

Define la ruta del módulo.

Especifica la versión de Go con la que se construyó el proyecto.

Lista todas las dependencias directas e indirectas requeridas, junto con sus versiones.

  \| 66 of 71

# Módulos en Go: go.mod : ejemplo

Un go.mod típico se ve así:

go module ejemplo.com/mi- proyecto-

go 1.22.0

require ( github.com/google/uuid v1.6.0 // Dependencia directa )

require ( golang.org/x/sys v0.18.0 // indirect )

La sección require lista las dependencias.

// indirect indica que la dependencia no es importada directamente por tu código, sino por una de tus dependencias directas.

  \| 67 of 71

# Agregando un Módulo Externo

Para usar una biblioteca externa, simplemente impórtala en tu código. Las herramientas de Go se encargarán del resto.

1\. Modifica tu main.go para usar un paquete externo, como github.com/google/uuid :

package main import ("fmt"

"github.com/google/uuid" )

func main() { i d := uuid.New() fm t.Println("UUID generado:",i d) }

2\. La próxima vez que compiles o ejecutes ( go build o go run ), Go notará la nueva importación, la descargará automáticamente y la agregará a tu go.mod y go.sum . El archivo go.sum contiene los hashes criptográficos de las dependencias para garantizar la integridad.

  \| 68 of 71

# Limpiando Dependencias con go mod tidy

El comando go mod tidy es esencial para mantener tu archivo go.mod limpio y preciso.

Actualiza go.mod y go.sum para que coincidan con las importaciones en tus archivos .go .

Agrega las dependencias que faltan para las importaciones que has agregado a tu código.

Elimina las dependencias que ya no se utilizan porque has eliminado importaciones.

Asegura que el archivo go.mod refleje con precisión el estado actual de tu código.

Es una buena práctica ejecutar go mod tidy antes de confirmar tus cambios en el control de versiones.

  \| 69 of 71

Y ahora qué?

Orígenes y evolución:

https://cacm.acm.org/research/the go- programming-language-and environment/

Hacer el Tour de Go:

https://go.dev/tour/welcome/

Mirar Go by Example:

https://gobyexample.com/

Mirar la documentación:

https://go.dev/doc/

Evitar el uso de bibliotecas externas

  \| 70 of 71

<image redacted: 479x479px, 479x479pt, ~72dpi, JPG, DEVICE_RGB, 32bpp>

# Go y el desarrollo Web

Todo lo visto en esta introducción es la base para lo que viene en la materia:

net/http : crear servidores y manejar peticiones HTTP (el punto de entrada a la web).

html/template y templ: generar páginas HTML de forma segura (capa de presentación).

database/sql y sqlc: acceder a bases de datos relacionales (capa de datos).

Concurrencia (gorutinas y canales): base para servidores que manejan miles de conexiones.

Interfaces: dependencias flexibles, fáciles de testear y de reemplazar.

  Ti p El resto de la materia aprovecha estas herramientas sobre servidores Linux con Docker.

  \| 71 of 71