package logic

import (
	"errors"
	"fmt"
	"strings"
)

type Producto struct {
	ID          int32   `json:"id"`
	IDCategoria int32   `json:"id_categoria"`
	Nombre      string  `json:"nombre"`
	Descripcion string  `json:"descripcion"`
	Stock       float64 `json:"stock"` // <- que sea float64
	Precio      float64 `json:"precio"`
}

func ValidateProduct(p Producto) error {
	if strings.TrimSpace(p.Nombre) == "" {
		return errors.New("el nombre no puede estar vacio")
	}
	if p.IDCategoria <= 0 {
		return fmt.Errorf("la categoria es obligatoria")
	}
	if p.Stock < 0.0 {
		return fmt.Errorf("el stock no puede ser menor a 0")
	}

	if p.Precio <= 0 {
		return fmt.Errorf("el precio debe ser mayor a cero")
	}

	return nil
}
