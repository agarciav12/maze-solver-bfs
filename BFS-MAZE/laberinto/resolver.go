package laberinto

import (
	"errors"

	TDACola "tdas/cola"
	TDAPila "tdas/pila"
)

var ErrSinCamino = errors.New("no hay camino entre el inicio y el fin")

func (l Laberinto) Resolver() (TDAPila.Pila[Celda], error) {
	cola := TDACola.CrearColaEnlazada[Celda]()

	visitadas := make([][]bool, l.filas)
	padre := make([][]Celda, l.filas)

	for i := 0; i < l.filas; i++ {
		visitadas[i] = make([]bool, l.columnas)
		padre[i] = make([]Celda, l.columnas)
	}

	cola.Encolar(l.inicio)
	visitadas[l.inicio.Fila][l.inicio.Columna] = true

	for !cola.EstaVacia() {
		actual := cola.Desencolar()

		if actual == l.fin {
			break
		}

		for _, v := range l.Vecinos(actual) {
			if !visitadas[v.Fila][v.Columna] {
				visitadas[v.Fila][v.Columna] = true
				padre[v.Fila][v.Columna] = actual
				cola.Encolar(v)
			}
		}
	}

	if !visitadas[l.fin.Fila][l.fin.Columna] {
		return nil, ErrSinCamino
	}

	return reconstruirCamino(padre, l.inicio, l.fin), nil
}

func reconstruirCamino(padre [][]Celda, inicio Celda, fin Celda) TDAPila.Pila[Celda] {
	camino := TDAPila.CrearPilaDinamica[Celda]()

	actual := fin
	camino.Apilar(actual)

	for actual != inicio {
		actual = padre[actual.Fila][actual.Columna]
		camino.Apilar(actual)
	}

	return camino
}
