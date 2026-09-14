# Trabajo Práctico 2

## Requisitos
En la maquina deben estar instalados:
- Git
- Docker con Docker Compose
- Make

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

El archivo `.env` contiene la configuración local de PostgreSQL.