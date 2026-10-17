package entrada

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"maze/laberinto"
	TDAPila "tdas/pila"
)

const salidaError = "ERROR"

var (
	errLineaIncompleta = errors.New("la línea del laberinto tiene menos columnas de las esperadas")
	errSinInicio       = errors.New("el laberinto no tiene celda de inicio")
	errSinFin          = errors.New("el laberinto no tiene celda de fin")
)

func ResolverLaberintosDeEntradaEstandar() {
	lector := bufio.NewReader(os.Stdin)

	for {
		lab, err := LeerLaberinto(lector)

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			break
		}

		camino, err := lab.Resolver()

		if err != nil {
			fmt.Println(salidaError)
			continue
		}

		ImprimirCamino(camino)
	}
}

func LeerLaberinto(buffer *bufio.Reader) (laberinto.Laberinto, error) {
	var filas int
	var columnas int

	_, err := fmt.Fscan(buffer, &filas, &columnas)

	if err != nil {
		return laberinto.Laberinto{}, err
	}

	matriz := make([][]laberinto.TipoCelda, filas)

	var inicio laberinto.Celda
	var fin laberinto.Celda
	var hayInicio bool
	var hayFin bool

	for i := 0; i < filas; i++ {
		var linea string

		_, err = fmt.Fscan(buffer, &linea)

		if err != nil {
			return laberinto.Laberinto{}, err
		}

		if len(linea) < columnas {
			return laberinto.Laberinto{}, errLineaIncompleta
		}

		matriz[i] = make([]laberinto.TipoCelda, columnas)

		for j := 0; j < columnas; j++ {
			tipo := laberinto.TipoDeCelda(linea[j])
			matriz[i][j] = tipo

			if tipo == laberinto.Inicio {
				inicio = laberinto.Celda{
					Fila:    i,
					Columna: j,
				}
				hayInicio = true
			}

			if tipo == laberinto.Fin {
				fin = laberinto.Celda{
					Fila:    i,
					Columna: j,
				}
				hayFin = true
			}
		}
	}

	if !hayInicio {
		return laberinto.Laberinto{}, errSinInicio
	}

	if !hayFin {
		return laberinto.Laberinto{}, errSinFin
	}

	return laberinto.CrearLaberinto(matriz, inicio, fin), nil
}

func ImprimirCamino(camino TDAPila.Pila[laberinto.Celda]) {
	var movimientos []string

	actual := camino.Desapilar()

	for !camino.EstaVacia() {
		siguiente := camino.Desapilar()

		direccion := actual.DireccionHacia(siguiente)
		movimientos = append(movimientos, string(direccion))

		actual = siguiente
	}

	fmt.Println(len(movimientos))
	fmt.Println(strings.Join(movimientos, " "))
}
