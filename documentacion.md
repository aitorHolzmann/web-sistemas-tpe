# Documentacion

En esta entrega agregamos schema.sql donde estan definidas 3 tablas junto con sus relaciones:
* Producto
* Categoria
* Cliente

En el .env guardamos las credenciales para la conexion a la db. Si bien sabemos que no es una buena practica, para entrega decidimos priorizar otros puntos

Admeas tenemos el archivo queries.sql donde se definen las consultas para que sqlc genere el codigo go correspondiente. 

El archivo readme.md contiene las instrucciones para ejecutar los test.

Ademas como detalle *extra* sumamos al docker-compose.yml el contenedor de go, por lo que al levantarse al db tambien va a estar levantando go (en el Dockerfile).

Usamos un healthcheck para verificar que la db esta preparada para recibir consultas. 