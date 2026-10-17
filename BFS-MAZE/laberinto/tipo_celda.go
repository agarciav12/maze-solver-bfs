package laberinto

const (
	caracterEntradaPared  = '#'
	caracterEntradaInicio = 'S'
	caracterEntradaFin    = 'E'
	caracterEntradaLibre  = '.'
)

type TipoCelda struct {
	caracter    byte
	transitable bool
}

var (
	Libre  = TipoCelda{caracter: caracterEntradaLibre, transitable: true}
	Pared  = TipoCelda{caracter: caracterEntradaPared, transitable: false}
	Inicio = TipoCelda{caracter: caracterEntradaInicio, transitable: true}
	Fin    = TipoCelda{caracter: caracterEntradaFin, transitable: true}
)

func TipoDeCelda(caracter byte) TipoCelda {
	switch caracter {
	case caracterEntradaPared:
		return Pared
	case caracterEntradaInicio:
		return Inicio
	case caracterEntradaFin:
		return Fin
	default:
		return Libre
	}
}

func (t TipoCelda) EsTransitable() bool {
	return t.transitable
}

func (t TipoCelda) String() string {
	return string(t.caracter)
}
