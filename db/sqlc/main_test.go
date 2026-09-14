// TestMain se ejecuta una sola vez antes de correr todos los Test* de este
// paquete. Acá abrimos la conexión a Postgres y la dejamos disponible en la
// variable global testQueries para que cada test la use.
package db

import (
	"database/sql"
	"log"
	"os"
	"testing"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// testQueries guarda la conexion con la db.
var testQueries *Queries

func TestMain(m *testing.M) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("falta la variable de entorno DATABASE_URL")
	}

	conn, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("no se pudo abrir la conexión: %v", err)
	}
	defer conn.Close()

	//Esperamos a que la DB este lista para recibir peticiones.
	var pingErr error
	for i := 0; i < 30; i++ {
		if pingErr = conn.Ping(); pingErr == nil {
			break
		}
		time.Sleep(1 * time.Second)
	}
	if pingErr != nil {
		log.Fatalf("no se pudo conectar a la base tras varios intentos: %v", pingErr)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
