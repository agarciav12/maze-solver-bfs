package laberinto

type Celda struct {
	Fila    int
	Columna int
}

type Direccion string

const (
	Arriba    Direccion = "ARRIBA"
	Abajo     Direccion = "ABAJO"
	Izquierda Direccion = "IZQUIERDA"
	Derecha   Direccion = "DERECHA"
)

func (c Celda) DireccionHacia(destino Celda) Direccion {
	if destino.Fila > c.Fila {
		return Abajo
	}

	if destino.Fila < c.Fila {
		return Arriba
	}

	if destino.Columna > c.Columna {
		return Derecha
	}

	return Izquierda
}
