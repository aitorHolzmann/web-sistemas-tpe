package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestClienteCRUD(t *testing.T) {
	ctx := context.Background()
	sufijo := time.Now().UnixNano()
	email := fmt.Sprintf("cliente.test.%d@example.com", sufijo)

	var cliente Cliente

	t.Run("crear", func(t *testing.T) {
		var err error
		cliente, err = testQueries.CrearCliente(ctx, CrearClienteParams{
			Nombre:    "Nicolas",
			Apellido:  "Benito",
			Email:     email,
			Direccion: "Arroyo Seco 2500",
		})
		if err != nil {
			t.Fatalf("CrearCliente devolvió error: %v", err)
		}
		if cliente.ID == 0 {
			t.Fatal("se esperaba que el cliente creado tenga un ID asignado por la base")
		}
	})

	t.Run("obtener_por_id", func(t *testing.T) {
		get, err := testQueries.GetCliente(ctx, cliente.ID)
		if err != nil {
			t.Fatalf("GetCliente devolvió error: %v", err)
		}
		if get.Email != email {
			t.Errorf("email = %q, se esperaba %q", get.Email, email)
		}
	})

	t.Run("listar", func(t *testing.T) {
		lista, err := testQueries.ListarClientes(ctx)
		if err != nil {
			t.Fatalf("ListarClientes devolvió error: %v", err)
		}
		encontrado := false
		for _, c := range lista {
			if c.ID == cliente.ID {
				encontrado = true
				break
			}
		}
		if !encontrado {
			t.Error("el cliente creado no aparece en ListarClientes")
		}
	})

	t.Run("actualizar", func(t *testing.T) {
		nuevoNombre := "Nicolas Actualizado"
		nuevaDireccion := "Reforma Universitaria"
		err := testQueries.ActualizarCliente(ctx, ActualizarClienteParams{
			ID:        cliente.ID,
			Nombre:    nuevoNombre,
			Apellido:  cliente.Apellido,
			Email:     cliente.Email,
			Direccion: nuevaDireccion,
		})
		if err != nil {
			t.Fatalf("ActualizarCliente devolvió error: %v", err)
		}

		get, err := testQueries.GetCliente(ctx, cliente.ID)
		if err != nil {
			t.Fatalf("GetCliente tras cambiar dirección devolvió error: %v", err)
		}
		if get.Nombre != nuevoNombre || get.Direccion != nuevaDireccion {
			t.Errorf("direccion = %q, se esperaba %q", get.Direccion, nuevaDireccion)
		}
	})

	t.Run("borrar", func(t *testing.T) {
		if err := testQueries.BorrarCliente(ctx, cliente.ID); err != nil {
			t.Fatalf("BorrarCliente devolvió error: %v", err)
		}

		_, err := testQueries.GetCliente(ctx, cliente.ID)
		if !errors.Is(err, sql.ErrNoRows) {
			t.Errorf("GetCliente tras borrar: se esperaba sql.ErrNoRows, se obtuvo %v", err)
		}
	})
}
