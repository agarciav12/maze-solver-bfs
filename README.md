# Maze Solver (BFS)

🇦🇷 [Leer en español](README2.md)

A **Go** maze solver that finds the **shortest path** between a start and an end cell using **breadth-first search (BFS)**. It reads one or more mazes from standard input and prints the number of moves and the sequence of directions to follow.

![Go](https://img.shields.io/badge/Go-1.18+-00ADD8?logo=go&logoColor=white)

## Input format

The program reads one or more mazes from `stdin`, one after the other, until the end of the input. Each maze has the following format:

```
<rows> <columns>
<row 1>
<row 2>
...
<row N>
```

Each cell is represented by a single character:

| Symbol | Meaning |
|--------|---------|
| `#` | Wall |
| `.` | Open cell |
| `S` | Start |
| `E` | End |

Rows must not contain spaces, and each maze must have exactly one start and one end.

## Output format

For each maze, the program prints:

- If a path exists: two lines, the **number of moves** followed by the **moves** separated by spaces (`ARRIBA`, `ABAJO`, `IZQUIERDA`, `DERECHA`, meaning up, down, left and right).
- If there is no path between the start and the end: the word `ERROR`.

### Example

Input:

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

Output:

```
6
ABAJO DERECHA DERECHA ARRIBA DERECHA DERECHA
1
DERECHA
ERROR
```

### Invalid input

If the input is malformed (a row with fewer columns than declared, or a maze with no start or no end), the program writes the reason to `stderr` and stops processing the rest of the input.

## Requirements

- Go 1.18 or later (uses generics).
- The Stack and Queue ADTs (`tdas/` folder).

## Build and run

```bash
# Build
go build -o maze

# Run with a file
./maze < lab.txt

# Or read from the terminal
./maze
```

## Tests

Tests run against the `maze` executable, so build it first:

```bash
go build -o maze
./pruebas.sh
```

## Project structure

## How it works

1. **Reading**: the `entrada` package reads the dimensions and rows, converts each character into a cell type, and locates the start and end cells.
2. **BFS**: `Resolver` traverses the maze using a **queue**, starting at the start cell and visiting walkable neighbors. A `visitadas` matrix prevents visiting the same cell twice, and a `padre` matrix records which cell each one was reached from.
3. **Reconstruction**: if the end was reached, `padre` is walked back from the end to the start, pushing each cell onto a **stack**, so that popping yields the cells in start-to-end order.
4. **Output**: consecutive cells are compared to derive the direction of each move, and the result is printed.

Since BFS explores level by level, the first path it finds to the end is the one with the **fewest moves**.

**Complexity**: O(R × C) time and space, where R is the number of rows and C the number of columns.

## Design decisions

- **Separation of concerns**: input/output (`entrada`) is kept apart from the maze logic (`laberinto`), and `main` only delegates.
- **Types instead of loose strings**: cells and directions are modeled with dedicated types and constants.
- **Explicit errors**: invalid input and unsolvable mazes are represented with dedicated errors.
- **Custom ADTs**: it uses the Stack and Queue implemented in the course instead of hand-managed slices.

## License

Distributed under the MIT License. See [`LICENSE`](LICENSE) for details.