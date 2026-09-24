# Trabajo Práctico Especial - TP3: API REST

## Requisitos
En la máquina deben estar instalados:
- Git
- Docker con Docker Compose
- Make
- cURL (o Hurl) para ejecutar las pruebas

## Ejecución del Servidor

Desde el directorio `tpespecial/tp3`:

```bash
make run
```

Este comando:
1. Genera el código tipado con `sqlc`.
2. Compila el binario y resuelve dependencias.
3. Levanta el contenedor de PostgreSQL con las migraciones aplicadas.
4. Inicia el servidor HTTP escuchando en `http://localhost:8080`.

## Pruebas de la API REST

Con el servidor en ejecución, en otra terminal ejecutar:

```bash
./requests.sh
```

O utilizando Hurl:

```bash
hurl --test requests.hurl
```

## Pruebas Unitarias y de Integración con Base de Datos

Para correr la suite de tests automatizados de persistencia:

```bash
make test
```

## Estructura de Endpoints

- **Categorías**:
  - `GET /categorias`: Listar todas las categorías.
  - `POST /categorias`: Crear una nueva categoría.
  - `GET /categorias/{id}`: Obtener categoría por ID.
  - `PUT /categorias/{id}`: Actualizar categoría por ID.
  - `DELETE /categorias/{id}`: Eliminar categoría por ID.

- **Productos**:
  - `GET /productos`: Listar todos los productos.
  - `POST /productos`: Crear un producto (valida campos obligatorios y categoría existente).
  - `GET /productos/{id}`: Obtener producto por ID.
  - `PUT /productos/{id}`: Actualizar producto por ID.
  - `DELETE /productos/{id}`: Eliminar producto por ID.

- **Clientes**:
  - `GET /clientes`: Listar clientes.
  - `POST /clientes`: Crear un cliente.
  - `GET /clientes/{id}`: Obtener cliente por ID.
  - `PUT /clientes/{id}`: Actualizar cliente por ID.
  - `DELETE /clientes/{id}`: Eliminar cliente por ID.