Trabajo Práctico 4 \- La Capa de Presentación 

Trabajo de Cursada: Construyendo la Interfaz de Usuario 

**Objetivo**: Crear una interfaz de usuario (frontend) en HTML, CSS y JavaScript que consuma la API REST que construiste en el práctico anterior. 

En este paso, te enfocarás exclusivamente en el cliente. Tu servidor Go con la API REST debe estar corriendo en segundo plano. 

1\. **Estructura HTML**: 

Crea un nuevo index.html (puedes servirlo con un servidor de archivos estáticos simple o abrirlo directamente en el navegador). 

Diseña la estructura de la página: 

• Un título principal. 

• Un formulario para crear nuevas entidades (con campos de input y un botón de envío).   
• Un contenedor (un \<div\> o \<ul\>) donde se mostrará la lista de entidades. Asígnale un id para poder seleccionarlo con JavaScript. 

2\. **Comunicación con la API (JavaScript)**: 

Crea un archivo app.js y enlázalo a tu index.html. 

Escribe una función que, al cargar la página, haga una petición fetch a tu endpoint GET /\<entidades\>. Cuando reciba los datos, debe generar dinámicamente los elementos HTML (por ejemplo, \<li\> o \<div\>) para cada entidad y agregarlos al contenedor del DOM. Cada elemento debe mostrar la información relevante y tener un botón para eliminarlo. 

Añade un “event listener” al formulario para que, al enviarse, prevenga el comportamiento por defecto, tome los datos de los inputs, y haga una petición POST con fetch al endpoint de creación. 

Después de crear un nuevo elemento, la aplicación debe volver a pedir la lista completa y refrescar la vista para mostrar el nuevo ítem. 

3\. **Implementación del Borrado**: 

Añade un “event listener” a los botones de borrado. Como los botones se crean dinámicamente, deberás usar delegación de eventos. 

Al hacer clic, el listener debe obtener el ID del elemento y realizar una petición DELETE con fetch al endpoint correspondiente. 

Tras un borrado exitoso, refresca la lista para que el elemento desaparecido ya no se muestre. 4\. **Estilos con CSS**: 

Crea un archivo styles.css y enlázalo. 

Utiliza Flexbox o CSS Grid para organizar los elementos principales de tu página (formulario y lista). Añade estilos básicos para que la aplicación sea visualmente agradable y fácil de usar.