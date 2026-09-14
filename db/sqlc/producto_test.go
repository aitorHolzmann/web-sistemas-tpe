package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"testing"
	"time"
)

// TestProductoCRUD prueba el CRUD de producto. Como producto tiene una FK
// obligatoria (id_categoria) hacia categoria, primero creamos una categoría
// "de apoyo" y la borramos al final con t.Cleanup, sin importar si el test
// falla a mitad de camino.
func TestProductoCRUD(t *testing.T) {
	ctx := context.Background()

	categoria, err := testQueries.CrearCategoria(ctx, CrearCategoriaParams{
		Nombre:      fmt.Sprintf("Categoria para producto %d", time.Now().UnixNano()),
		Descripcion: sql.NullString{Valid: false},
	})
	if err != nil {
		t.Fatalf("no se pudo crear la categoría de apoyo: %v", err)
	}
	t.Cleanup(func() {
		_ = testQueries.BorrarCategoria(ctx, categoria.ID)
	})

	nombre := fmt.Sprintf("Producto Test %d", time.Now().UnixNano())
	var producto Producto

	t.Run("crear", func(t *testing.T) {
		producto, err = testQueries.CrearProducto(ctx, CrearProductoParams{
			Nombre:      nombre,
			Descripcion: sql.NullString{String: "descripcion del producto", Valid: true},
			Stock:       "10",
			Precio:      "199.99",
			Foto:        sql.NullString{Valid: false},
			IDCategoria: categoria.ID,
		})
		if err != nil {
			t.Fatalf("CrearProducto devolvió error: %v", err)
		}
		if producto.ID == 0 {
			t.Fatal("se esperaba que el producto creado tenga un ID asignado por la base")
		}
	})

	t.Run("obtener_por_id", func(t *testing.T) {
		get, err := testQueries.GetProducto(ctx, producto.ID)
		if err != nil {
			t.Fatalf("GetProducto devolvió error: %v", err)
		}
		if get.Nombre != nombre {
			t.Errorf("nombre = %q, se esperaba %q", get.Nombre, nombre)
		}
	})

	t.Run("listar", func(t *testing.T) {
		lista, err := testQueries.ListarProductos(ctx)
		if err != nil {
			t.Fatalf("ListarProductos devolvió error: %v", err)
		}
		encontrado := false
		for _, p := range lista {
			if p.ID == producto.ID {
				encontrado = true
				break
			}
		}
		if !encontrado {
			t.Error("el producto creado no aparece en ListarProductos")
		}
	})

	t.Run("actualizar", func(t *testing.T) {
		producto.Nombre = nombre + " actualizado"
		producto.Stock = "3"
		producto.Precio = "149.50"
		producto, err = testQueries.ActualizarProducto(ctx, ActualizarProductoParams{
			ID:          producto.ID,
			Nombre:      producto.Nombre,
			Descripcion: producto.Descripcion,
			Stock:       producto.Stock,
			Precio:      producto.Precio,
			Foto:        producto.Foto,
			IDCategoria: producto.IDCategoria,
		})
		if err != nil {
			t.Fatalf("ActualizarProducto devolvio error: %v", err)
		}
		if producto.Stock != "3" || producto.Precio != "149.50" {
			t.Errorf("producto actualizado = %+v", producto)
		}
	})

	t.Run("borrar", func(t *testing.T) {
		if err := testQueries.BorrarProducto(ctx, producto.ID); err != nil {
			t.Fatalf("BorrarProducto devolvió error: %v", err)
		}

		_, err := testQueries.GetProducto(ctx, producto.ID)
		if !errors.Is(err, sql.ErrNoRows) {
			t.Errorf("GetProducto tras borrar: se esperaba sql.ErrNoRows, se obtuvo %v", err)
		}
	})
}
