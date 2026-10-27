---
name: web-unicen-persistencia
description: >-
  Guía y patrones de la cátedra de Programación Web (UNICEN) para el nivel de persistencia de datos en Go.
  Usar para TP2, consultas con database/sql, consultas parametrizadas, repositorios, sqlc y GORM.
---

# Persistencia de Datos (Cátedra Web UNICEN)

Esta skill documenta las 3 variantes de persistencia enseñadas en la cátedra:
1. **SQL Directo** con `database/sql`.
2. **Generador y Mapper** con `sqlc`.
3. **ORM** con `GORM`.

Basado en `teoria/04-data.md`, `teoria/05-mappersORM-1.md` y `teoria/tp2-2.md`.

---

## 1. Conexión y Pool de Conexiones (PostgreSQL)

Usar el driver estándar `pgx` (`github.com/jackc/pgx/v5/stdlib`).

```go
func abrirDB() (*sql.DB, error) {
	dsn := "host=database port=5432 user=postgres password=supersecret dbname=tp2_db sslmode=disable"
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	// Configuración del pool de conexiones según lo visto en clase
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	return db, nil
}
```

---

## 2. Enfoque 1: "SQL Directo" con `database/sql` y Patrón Repositorio

### Reglas Críticas de la Cátedra:
- **SIEMPRE** consultas parametrizadas (`$1, $2, ...` en Postgres). Prohibido concatenar strings (riesgo de inyección SQL).
- Al iterar filas con `rows.Next()`:
  - Siempre `defer rows.Close()`.
  - Siempre verificar `if err := rows.Err(); err != nil { return nil, err }` al final del loop.
- En `Update` y `Delete`, verificar `result.RowsAffected()`: si es 0, devolver `sql.ErrNoRows` para permitir enviar un 404 en el handler HTTP.

### Estructura de Repositorio:
```go
type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *User) error {
	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id, created_at`
	return r.db.QueryRow(query, user.Name, user.Email).Scan(&user.ID, &user.CreatedAt)
}

func (r *UserRepository) GetUserByID(id int) (*User, error) {
	user := &User{}
	query := `SELECT id, name, email, created_at FROM users WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}
```

---

## 3. Enfoque 2: `sqlc` (Type-Safe SQL Compiler)

### Configuración (`sqlc.yaml`):
```yaml
version: "2"
sql:
  - engine: "postgresql"
    schema: "db/schema/"
    queries: "db/queries/"
    gen:
      go:
        package: "sqlc"
        out: "db/sqlc"
        sql_package: "pgx/v5"
```

### Consultas SQL anotadas (`db/queries/users.sql`):
```sql
-- name: GetUser :one
SELECT * FROM users WHERE id = $1;

-- name: ListUsers :many
SELECT * FROM users ORDER BY id;

-- name: CreateUser :one
INSERT INTO users (name, email) VALUES ($1, $2)
RETURNING *;

-- name: UpdateUser :exec
UPDATE users SET name = $2, email = $3 WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;
```

### Uso en Go:
```go
queries := db.New(conn)
usuario, err := queries.GetUser(ctx, id)
```

---

## 4. Enfoque 3: `GORM` (Object-Relational Mapping)

Usar cuando el práctico lo solicite explícitamente.

### Definición y AutoMigrate:
```go
type User struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"not null"`
	Email     string    `gorm:"unique;not null"`
	CreatedAt time.Time
}

db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
db.AutoMigrate(&User{})
```

### Operaciones CRUD:
- **Create**: `db.Create(&user)`
- **Read One**: `db.First(&user, id)` (devuelve `gorm.ErrRecordNotFound` si no existe)
- **Read All**: `db.Find(&users)`
- **Update**: `db.Model(&user).Update("Name", "Nuevo Nombre")`
- **Delete**: `db.Delete(&User{}, id)`
- **Relaciones (evitar N+1)**: Usar `.Preload("Relacion")`

