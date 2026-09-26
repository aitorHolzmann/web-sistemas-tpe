package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestCategoriaCRUD(t *testing.T) {
	ctx := context.Background()
	nombre := fmt.Sprintf("Categoria Test %d", time.Now().UnixNano())

	var categoria Categoria

	t.Run("crear", func(t *testing.T) {
		var err error
		categoria, err = testQueries.CrearCategoria(ctx, CrearCategoriaParams{
			Nombre:      nombre,
			Descripcion: sql.NullString{String: "descripcion inicial", Valid: true},
		})
		if err != nil {
			t.Fatalf("CrearCategoria devolvio error: %v", err)
		}
		if categoria.ID == 0 {
			t.Fatal("se esperaba que la categoria creada tenga un ID asignado por la base")
		}
	})

	t.Run("obtener_por_id", func(t *testing.T) {
		get, err := testQueries.GetCategoria(ctx, categoria.ID)
		if err != nil {
			t.Fatalf("GetCategoria devolvio error: %v", err)
		}
		if get.Nombre != nombre {
			t.Errorf("nombre = %q, se esperaba %q", get.Nombre, nombre)
		}
	})

	t.Run("listar", func(t *testing.T) {
		lista, err := testQueries.ListarCategorias(ctx)
		if err != nil {
			t.Fatalf("ListarCategorias devolvio error: %v", err)
		}
		encontrada := false
		for _, c := range lista {
			if c.ID == categoria.ID {
				encontrada = true
				break
			}
		}
		if !encontrada {
			t.Error("la categoria creada no aparece en ListarCategorias")
		}
	})

	t.Run("actualizar", func(t *testing.T) {
		nuevoNombre := nombre + " editada"
		err := testQueries.ActualizarCategoria(ctx, ActualizarCategoriaParams{
			ID:          categoria.ID,
			Nombre:      nuevoNombre,
			Descripcion: sql.NullString{String: "descripcion editada", Valid: true},
		})
		if err != nil {
			t.Fatalf("ActualizarCategoria devolvio error: %v", err)
		}

		get, err := testQueries.GetCategoria(ctx, categoria.ID)
		if err != nil {
			t.Fatalf("GetCategoria devolvió error tras actualizar: %v", err)
		}
		if get.Nombre != nuevoNombre {
			t.Errorf("nombre tras actualizar = %q, se esperaba %q", get.Nombre, nuevoNombre)
		}
	})

	t.Run("borrar", func(t *testing.T) {
		if err := testQueries.BorrarCategoria(ctx, categoria.ID); err != nil {
			t.Fatalf("BorrarCategoria devolvió error: %v", err)
		}

		_, err := testQueries.GetCategoria(ctx, categoria.ID)
		if !errors.Is(err, sql.ErrNoRows) {
			t.Errorf("GetCategoria tras borrar: se esperaba sql.ErrNoRows, se obtuvo %v", err)
		}
	})
}
