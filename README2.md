# Maze Solver (BFS)

🇬🇧 [Read in English](README.md)

Resolvedor de laberintos en **Go** que encuentra el **camino más corto** entre un inicio y un fin usando **búsqueda en anchura (BFS)**. Lee uno o varios laberintos por entrada estándar e imprime la cantidad de movimientos y la secuencia de direcciones a seguir.

![Go](https://img.shields.io/badge/Go-1.18+-00ADD8?logo=go&logoColor=white)

## Formato de entrada

El programa lee de `stdin` uno o más laberintos, uno a continuación del otro, hasta llegar al fin de la entrada. Cada laberinto tiene este formato:

```
<filas> <columnas>
<fila 1>
<fila 2>
...
<fila N>
```

Cada celda se representa con un carácter:

| Símbolo | Significado |
|---------|-------------|
| `#` | Pared |
| `.` | Celda libre |
| `S` | Inicio |
| `E` | Fin |

Las filas no pueden contener espacios, y cada laberinto debe tener exactamente un inicio y un fin.

## Formato de salida

Por cada laberinto se imprime:

- Si hay camino: dos líneas, la **cantidad de movimientos** y luego los **movimientos** separados por espacios (`ARRIBA`, `ABAJO`, `IZQUIERDA`, `DERECHA`).
- Si no hay camino entre el inicio y el fin: la palabra `ERROR`.

### Ejemplo

Entrada:

```
4 7
#######
#S#..E#
#...#.#
#######
4 4
####
#SE#
#..#
####
3 5
#####
#S#E#
#####
```

Salida:

```
6
ABAJO DERECHA DERECHA ARRIBA DERECHA DERECHA
1
DERECHA
ERROR
```

### Entradas inválidas

Si la entrada está mal formada (una fila con menos columnas de las indicadas, o un laberinto sin inicio o sin fin), el programa escribe el motivo por `stderr` y deja de procesar el resto de la entrada.

## Requisitos

- Go 1.18 o superior (usa genéricos).
- Los TDAs Pila y Cola (carpeta `tdas/`).

## Compilar y ejecutar

```bash
# Compilar
go build -o maze

# Ejecutar con un archivo
./maze < lab.txt

# O leyendo de la terminal
./maze
```

## Pruebas

Las pruebas se corren contra el ejecutable `maze`, por lo que hay que compilarlo antes:

```bash
go build -o maze
./pruebas.sh
```

## Cómo funciona

1. **Lectura**: el paquete `entrada` lee las dimensiones y las filas, convierte cada carácter en un tipo de celda y ubica el inicio y el fin.
2. **BFS**: `Resolver` recorre el laberinto con una **cola**, partiendo del inicio y visitando los vecinos transitables. Una matriz `visitadas` evita volver a pasar por la misma celda y una matriz `padre` guarda desde qué celda se llegó a cada una.
3. **Reconstrucción**: si se llegó al fin, se recorre `padre` desde el fin hasta el inicio apilando las celdas en una **pila**, de modo que al desapilar quedan en orden de inicio a fin.
4. **Salida**: se comparan celdas consecutivas para obtener la dirección de cada movimiento y se imprime el resultado.

Como BFS explora por niveles, el primer camino que encuentra al fin es el de **menor cantidad de movimientos**.

**Complejidad**: O(F × C) en tiempo y espacio, siendo F las filas y C las columnas.

## Decisiones de diseño

- **Separación de responsabilidades**: la entrada/salida (`entrada`) está separada de la lógica del laberinto (`laberinto`), y `main` solo delega.
- **Tipos en lugar de strings sueltos**: las celdas y las direcciones se modelan con tipos propios y constantes.
- **Errores explícitos**: los casos de entrada inválida y de laberinto sin solución se representan con errores propios.
- **TDAs propios**: se usan Pila y Cola implementadas en la materia, en lugar de slices manipulados a mano.

## Licencia

Distribuido bajo la licencia MIT. Ver [`LICENSE`](LICENSE) para más información.