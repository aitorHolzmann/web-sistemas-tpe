# Trabajo Práctico 2

## Requisitos

En la máquina del profesor deben estar instalados:

- Git
- Docker con Docker Compose
- Make

No es necesario instalar Go ni `sqlc` en la máquina host. Ambos se ejecutan
dentro del contenedor de desarrollo.

## Después de clonar

Desde la raíz del proyecto:

```bash
make test
```

Ese comando:

1. Genera el código de `sqlc`.
2. Descarga las dependencias de Go y compila el proyecto.
3. Crea una base PostgreSQL limpia.
4. Ejecuta todos los tests.
5. Detiene los contenedores y elimina el volumen de la base.

El archivo `.env` contiene la configuración local de PostgreSQL. Si no está
incluido en la copia del proyecto, hay que crearlo antes de ejecutar Make.

## Comandos disponibles

```bash
make              # Igual que make test
make test         # Genera, compila, crea la base y ejecuta los tests
make sqlc         # Regenera el código de sqlc
make build        # Descarga dependencias y compila
make up           # Levanta PostgreSQL y espera a que esté listo
make down         # Detiene contenedores y elimina el volumen de la base
```

Para abrir el proyecto dentro del entorno de VS Code, usar la opción
**Reopen in Container**. Al crearlo, el Dev Container descarga las
dependencias de Go automáticamente.
