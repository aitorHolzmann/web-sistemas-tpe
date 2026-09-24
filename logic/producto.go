package logic

import (
	"errors"
	"strconv"
	"strings"
)

func ValidateProducto(nombre string, stock string, precio string, idCategoria int32) error {
	if strings.TrimSpace(nombre) == "" {
		return errors.New("el nombre no puede estar vacio")
	}
	if idCategoria <= 0 {
		return errors.New("la categoria es obligatoria")
	}
	s, err := strconv.ParseFloat(stock, 64)
	if err != nil || s < 0 {
		return errors.New("el stock debe ser un numero mayor o igual a cero")
	}
	p, err := strconv.ParseFloat(precio, 64)
	if err != nil || p <= 0 {
		return errors.New("el precio debe ser mayor a cero")
	}
	return nil
}
