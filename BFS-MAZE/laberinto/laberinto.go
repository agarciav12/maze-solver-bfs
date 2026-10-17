package laberinto

type Laberinto struct {
	filas    int
	columnas int
	matriz   [][]TipoCelda
	inicio   Celda
	fin      Celda
}

func CrearLaberinto(matriz [][]TipoCelda, inicio Celda, fin Celda) Laberinto {
	return Laberinto{
		filas:    len(matriz),
		columnas: len(matriz[0]),
		matriz:   matriz,
		inicio:   inicio,
		fin:      fin,
	}
}

func (l Laberinto) estaDentroDelLaberinto(c Celda) bool {
	return c.Fila >= 0 && c.Fila < l.filas &&
		c.Columna >= 0 && c.Columna < l.columnas
}

func (l Laberinto) Vecinos(c Celda) []Celda {
	desplazamientos := []Celda{
		{Fila: -1, Columna: 0},
		{Fila: 1, Columna: 0},
		{Fila: 0, Columna: -1},
		{Fila: 0, Columna: 1},
	}

	var validos []Celda

	for _, d := range desplazamientos {
		posible := Celda{
			Fila:    c.Fila + d.Fila,
			Columna: c.Columna + d.Columna,
		}

		if !l.estaDentroDelLaberinto(posible) {
			continue
		}

		if !l.matriz[posible.Fila][posible.Columna].EsTransitable() {
			continue
		}

		validos = append(validos, posible)
	}

	return validos
}
